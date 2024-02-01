package main

import (
	"context"
	pb "gRPC-blogging-platform/blog"
	"testing"
)

func TestServer_CreatePost(t *testing.T) {
	// Initialize server
	srv := &server{}

	// Prepare a sample request
	req := &pb.CreatePostRequest{
		Title:           "Sample Post",
		Content:         "This is a sample post content.",
		Author:          "John Doe",
		PublicationDate: "2024-02-01",
		Tags:            []string{"golang", "grpc", "blog"},
	}

	// Call the CreatePost method
	resp, err := srv.CreatePost(context.Background(), req)
	if err != nil {
		t.Errorf("CreatePost failed: %v", err)
	}

	// Validate response
	if resp == nil {
		t.Error("CreatePost returned nil response")
	}
}

func TestServer_ReadPost(t *testing.T) {
	// Initialize server
	srv := &server{}

	// Prepare a sample request
	req := &pb.ReadPostRequest{
		PostId: "123",
	}

	// Call the ReadPost method
	_, err := srv.ReadPost(context.Background(), req)
	if err != nil {
		t.Errorf("ReadPost failed: %v", err)
	}
}
