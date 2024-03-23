package grpc

import (
	"context"

	pb "github.com/antrein/proto-repository/pb/bc"
)

type helloServer struct {
	pb.UnimplementedGreeterServer
}

func (s *helloServer) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{Message: "Hello " + in.Name}, nil
}
