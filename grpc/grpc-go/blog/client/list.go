package main

import (
	"context"
	"io"
	"log"

	pb "grpc-go/blog/proto"
)

func listBlogs(c pb.BlogServiceClient) {
	log.Println("---listBlogs was invoked---")

	stream, err := c.ListBlogs(context.Background(), &pb.Empty{})
	if err != nil {
		log.Fatalf("Error while calling ListBlogs: %v\n", err)
	}

	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Something went wrong: %v\n", err)
		}

		log.Println(res)
	}
}
