package main

import (
	"context"
	"google.golang.org/protobuf/types/known/emptypb"
	"log"

	pb "gRPC-blogging-platform/blog"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()

	c := pb.NewBlogServiceClient(conn)

	// Create Post
	createPostReq := &pb.CreatePostRequest{
		Title:           "Sample Post",
		Content:         "This is a sample post content.",
		Author:          "John Doe",
		PublicationDate: "2024-02-01",
		Tags:            []string{"golang", "grpc", "blog"},
	}

	createPostRes, err := c.CreatePost(context.Background(), createPostReq)
	if err != nil {
		log.Fatalf("could not create post: %v", err)
	}

	log.Printf("Post created: %v", createPostRes)

	// Read Post
	readPostReq := &pb.ReadPostRequest{
		PostId: createPostRes.PostId,
	}

	readPostRes, err := c.ReadPost(context.Background(), readPostReq)
	if err != nil {
		log.Fatalf("could not read post: %v", err)
	}

	log.Printf("Post read: %v", readPostRes)

	// Read All Posts
	readAllPostsRes, err := c.ReadAllPosts(context.Background(), &emptypb.Empty{})
	if err != nil {
		log.Fatalf("could not read post: %v", err)
	}

	log.Printf("Posts read: %v", readAllPostsRes)

	// Update Post
	updatePostReq := &pb.UpdatePostRequest{
		PostId:  createPostRes.PostId,
		Title:   "Updated Sample Post",
		Content: "This is the updated content of the sample post.",
		Author:  "Jane Doe",
		Tags:    []string{"golang", "grpc", "blog", "update"},
	}

	updatePostRes, err := c.UpdatePost(context.Background(), updatePostReq)
	if err != nil {
		log.Fatalf("could not update post: %v", err)
	}

	log.Printf("Post updated: %v", updatePostRes)

	// Delete Post
	deletePostReq := &pb.DeletePostRequest{
		PostId: createPostRes.PostId,
	}

	deletePostRes, err := c.DeletePost(context.Background(), deletePostReq)
	if err != nil {
		log.Fatalf("could not delete post: %v", err)
	}

	log.Printf("Post deleted: %v", deletePostRes)
}
