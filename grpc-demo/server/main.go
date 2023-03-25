package main

import (
	"log"
	"net"

	pb "grpc-demo/proto"

	"google.golang.org/grpc"
)

const (
	port = ":8080"
)

type helloServer struct {
	pb.GreetServiceServer
}

func main() {
	l, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to start the server: %v", err)
	}

	grpcSrv := grpc.NewServer()
	pb.RegisterGreetServiceServer(grpcSrv, &helloServer{})

	if err := grpcSrv.Serve(l); err != nil {
		log.Fatalf("Failed to start the server: %v", err)
	}
}
