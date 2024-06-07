package analytic

import (
	"antrein/bc-dashboard/model/config"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	pb "github.com/antrein/proto-repository/pb/bc"
	"github.com/gofiber/adaptor/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"

	"google.golang.org/grpc"
)

type Client struct {
	cfg        *config.Config
	grpcClient *grpc.ClientConn
}

func New(cfg *config.Config, gc *grpc.ClientConn) *Client {
	return &Client{
		cfg:        cfg,
		grpcClient: gc,
	}
}

func (c *Client) RegisterRoute(app *fiber.App) {
	app.Get("/bc/dashboard/analytic", adaptor.HTTPHandler(http.HandlerFunc(c.StreamAnalyticData)))
}

func handler(f http.HandlerFunc) http.Handler {
	return http.HandlerFunc(f)
}

func (c *Client) StreamAnalyticData(w http.ResponseWriter, r *http.Request) {
	// Set headers for SSE
	ctx := context.Background()
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	projectID := r.URL.Query().Get("project_id")

	client := pb.NewAnalyticServiceClient(c.grpcClient)

	stream, err := client.StreamRealtimeData(ctx, &pb.AnalyticRequest{
		ProjectId: projectID,
	})
	if err != nil {
		log.Error(err)
		http.Error(w, "Error connecting to gRPC stream", http.StatusInternalServerError)
		return
	}

	for ctx.Err() == nil {
		select {
		case <-ctx.Done():
			http.Error(w, "Stream done", http.StatusInternalServerError)
			return
		default:
			analyticData, err := stream.Recv()
			if err != nil {
				log.Error(err)
				http.Error(w, "Error receiving data from gRPC stream", http.StatusInternalServerError)
				return
			}
			jsonData, err := json.Marshal(analyticData)
			if err != nil {
				log.Error(err)
				http.Error(w, "Error marshaling data to JSON", http.StatusInternalServerError)
				return
			}

			_, err = fmt.Fprintf(w, "data: %s\n\n", jsonData)
			if err != nil {
				log.Error(err)
				http.Error(w, "Error writing to response", http.StatusInternalServerError)
				return
			}

			if flusher, ok := w.(http.Flusher); ok {
				flusher.Flush()
			} else {
				log.Error(err)
				http.Error(w, "Streaming unsupported", http.StatusInternalServerError)
				return
			}

			time.Sleep(1 * time.Second)
		}
	}

	http.Error(w, "Stream done", http.StatusInternalServerError)
}
