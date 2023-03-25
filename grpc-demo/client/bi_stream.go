package main

import (
	"context"
	pb "grpc-demo/proto"
	"io"
	"log"
	"time"
)

func callSayHelloBidirectionalStream(client pb.GreetServiceClient, names *pb.NameList) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	stream, err := client.SayHelloBidirectionalStreaming(ctx)
	if err != nil {
		log.Fatalf("Failed to call SayHelloBidirectionalStreaming: %v", err)
	}

	wait := make(chan struct{})

	go func() {
		for _, name := range names.Names {
			req := &pb.HelloRequest{Name: name}
			if err := stream.Send(req); err != nil {
				log.Fatalf("Failed to send a request: %v", err)
			}
		}
		stream.CloseSend()
	}()

	go func() {
		for {
			resp, err := stream.Recv()
			if err != nil {
				if err == io.EOF {
					close(wait)
					return
				}
				log.Fatalf("Failed to receive a response: %v", err)
			}
			log.Printf("Response from SayHelloBidirectionalStreaming: %v\n", resp)
		}
	}()

	<-wait

	log.Println("SayHelloBidirectionalStreaming finished")
}
