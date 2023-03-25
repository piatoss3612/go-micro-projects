package main

import (
	pb "grpc-demo/proto"
	"io"
	"log"
)

func (h *helloServer) SayHelloBidirectionalStreaming(stream pb.GreetService_SayHelloBidirectionalStreamingServer) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				return nil
			}
			log.Fatalf("Failed to receive a request: %v", err)
		}

		res := &pb.HelloResponse{
			Message: "Hello " + req.Name,
		}

		if err := stream.Send(res); err != nil {
			log.Fatalf("Failed to send a response: %v", err)
		}
	}
}
