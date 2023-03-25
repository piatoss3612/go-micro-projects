package main

import (
	"context"
	pb "grpc-demo/proto"
	"io"
	"log"
	"time"
)

func callSayHelloServerStream(client pb.GreetServiceClient, names *pb.NameList) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	stream, err := client.SayHelloServerStreaming(ctx, names)
	if err != nil {
		log.Fatalf("Failed to call SayHelloServerStreaming: %v", err)
	}

	for {
		resp, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatalf("Failed to receive a response: %v", err)
		}
		log.Printf("Response from SayHelloServerStreaming: %v\n", resp)
	}
}
