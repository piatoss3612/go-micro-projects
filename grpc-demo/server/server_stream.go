package main

import (
	pb "grpc-demo/proto"
	"log"
)

func (s *helloServer) SayHelloServerStreaming(in *pb.NameList, stream pb.GreetService_SayHelloServerStreamingServer) error {
	log.Printf("Received: %v\n", in.Names)

	for _, name := range in.Names {
		res := &pb.HelloResponse{Message: "Hello " + name}

		if err := stream.SendMsg(res); err != nil {
			log.Fatalf("Failed to send a response: %v", err)
		}
	}
	return nil
}
