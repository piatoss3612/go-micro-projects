package main

import (
	"context"
	"log"

	pb "grpc-go/blog/proto"
)

func updateBlog(c pb.BlogServiceClient, id string) {
	log.Println("---updateBlog was invoked---")

	req := &pb.Blog{
		Id:       id,
		AuthorId: "SSOTAIP",
		Title:    "A new title",
		Content:  "author name is reversed!",
	}

	_, err := c.UpdateBlog(context.Background(), req)
	if err != nil {
		log.Fatalf("Error happened while updating: %v\n", err)
	}

	log.Println("Blog was updated!")
}
