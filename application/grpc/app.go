package grpc

import (
	"antrein/bc-dashboard/internal/pb"
	"context"
)

type helloServer struct {
	pb.UnimplementedGreeterServer
}

func (s *helloServer) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{Message: "Hello " + in.Name}, nil
}
