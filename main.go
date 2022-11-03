package main

import (
	"context"
	"log"

	// "math/rand"
	"net"

	pb "example.com/usergrpc/proto"

	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

type UserServer struct {
	pb.UnimplementedUserServiceServer
}

func (s *UserServer) CreateUser(ctx context.Context, in *pb.NewUser) (*pb.User, error) {
	log.Printf("RECEIVED...")
	var userId uint32 = 101 //uint32(rand.Intn(1000))
	return &pb.User{Name: in.Name, Age: in.Age, Id: userId}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterUserServiceServer(s, &UserServer{})
	log.Printf("Server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatal("Failed to serve: %v", err)
	}
}

// protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative user.proto
