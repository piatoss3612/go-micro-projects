package main

import (
	"context"
	pb "grpc-demo/proto"
)

func (s *helloServer) SayHello(ctx context.Context, in *pb.NoParam) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{Message: "Hello"}, nil
}
