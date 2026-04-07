package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "github.com/joss12/grpc-echo/proto"
	"google.golang.org/grpc"
)

type echoServer struct {
	pb.UnimplementedEchoServiceServer
}

func (s *echoServer) Echo(ctx context.Context, req *pb.EchoRequest) (*pb.EchoResponse, error) {
	log.Printf("Received: %s", req.Message)
	return &pb.EchoResponse{Message: "ECHO: " + req.Message}, nil
}

func main() {

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()

	pb.RegisterEchoServiceServer(s, &echoServer{})

	fmt.Println("Server listening on :50051")

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
