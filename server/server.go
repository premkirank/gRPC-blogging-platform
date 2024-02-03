package main

import (
	"context"
	"fmt"
	"google.golang.org/protobuf/types/known/emptypb"
	"log"
	"net"

	pb "gRPC-blogging-platform/blog"
	"github.com/google/uuid"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedBlogServiceServer
}

var posts = make(map[string]*pb.PostResponse)

func (s *server) CreatePost(ctx context.Context, req *pb.CreatePostRequest) (*pb.PostResponse, error) {
	//postID := fmt.Sprintf("%d", len(posts)+1)
	postID := fmt.Sprintf("%v", uuid.New())
	post := &pb.PostResponse{
		PostId:          postID,
		Title:           req.Title,
		Content:         req.Content,
		Author:          req.Author,
		PublicationDate: req.PublicationDate,
		Tags:            req.Tags,
	}

	posts[postID] = post
	return post, nil
}

func (s *server) ReadPost(ctx context.Context, req *pb.ReadPostRequest) (*pb.PostResponse, error) {
	post, ok := posts[req.PostId]
	if !ok {
		return nil, fmt.Errorf("post with ID %s not found", req.PostId)
	}

	return post, nil
}

func (s *server) ReadAllPosts(ctx context.Context, req *emptypb.Empty) (*pb.AllPostsResponse, error) {
	var allPostsResponse = pb.AllPostsResponse{}
	allPosts := make([]*pb.PostResponse, 0)

	for _, post := range posts {
		if allPosts != nil {
			allPosts = append(allPosts, post)
		}
	}

	if posts == nil {
		return nil, fmt.Errorf("no posts found")
	}

	allPostsResponse.Posts = allPosts
	return &allPostsResponse, nil
}

func (s *server) UpdatePost(ctx context.Context, req *pb.UpdatePostRequest) (*pb.PostResponse, error) {
	post, ok := posts[req.PostId]
	if !ok {
		return nil, fmt.Errorf("post with ID %s not found", req.PostId)
	}
	post.Title = req.Title
	post.Content = req.Content
	post.Author = req.Author
	post.Tags = req.Tags

	return post, nil
}

func (s *server) DeletePost(ctx context.Context, req *pb.DeletePostRequest) (*pb.DeleteResponse, error) {
	_, ok := posts[req.PostId]
	if !ok {
		return &pb.DeleteResponse{Success: false, Message: fmt.Sprintf("post with ID %s not found", req.PostId)}, nil
	}

	delete(posts, req.PostId)
	return &pb.DeleteResponse{Success: true, Message: "Post deleted successfully"}, nil
}

func main() {
	posts = make(map[string]*pb.PostResponse)

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterBlogServiceServer(s, &server{})

	log.Println("Server started on port 50051...")

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
