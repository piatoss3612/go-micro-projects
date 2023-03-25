package main

import (
	"context"
	pb "grpc-demo/proto"
	"log"
	"time"
)

func callSayHelloClientStream(client pb.GreetServiceClient, names *pb.NameList) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	stream, err := client.SayHelloClientStreaming(ctx)
	if err != nil {
		log.Fatalf("Failed to call SayHelloClientStreaming: %v", err)
	}

	for _, name := range names.Names {
		req := &pb.HelloRequest{Name: name}
		if err := stream.Send(req); err != nil {
			log.Fatalf("Failed to send a request: %v", err)
		}
	}

	resp, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Failed to receive a response: %v", err)
	}

	log.Printf("Response from SayHelloClientStreaming: %v\n", resp)
}
