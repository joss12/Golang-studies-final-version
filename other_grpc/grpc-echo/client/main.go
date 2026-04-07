package main

import (
	"context"
	"log"
	"time"

	pb "github.com/joss12/grpc-echo/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient("localhost:50051",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewEchoServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	resp, err := client.Echo(ctx, &pb.EchoRequest{Message: "Hello gRPC!"})
	if err != nil {
		log.Fatalf("could not echo: %v", err)
	}

	log.Printf("Response: %s", resp.Message)
}
