package main

import (
	pb "grpc-demo/proto"
	"io"
	"log"
)

func (h *helloServer) SayHelloClientStreaming(stream pb.GreetService_SayHelloClientStreamingServer) error {
	var names []string
	for {
		req, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				res := &pb.MessagesList{}
				for _, name := range names {
					res.Messages = append(res.Messages, "Hello "+name)
				}
				return stream.SendAndClose(res)
			}
			log.Fatalf("Failed to receive a request: %v", err)
		}

		names = append(names, req.Name)
	}
}
