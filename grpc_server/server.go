package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "simplegrpcserver/proto/gen"
	farewellpb "simplegrpcserver/proto/gen/farewell"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type server struct {
	pb.UnimplementedCalculateServer
	pb.UnimplementedBidFarewellServer
}

type serverGreeter struct {
	pb.UnimplementedGreeterServer
}

func (s *server) Add(ctx context.Context, req *pb.AddRequest) (*pb.AddResponse, error) {
	sum := req.A + req.B
	log.Println("Sum:", sum)

	return &pb.AddResponse{
		Sum: sum,
	}, nil
}

func (s *serverGreeter) Add(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{
		Message: fmt.Sprintf("Hello %s. Nice to receive request from you", req.Name),
	}, nil
}

func (g *server) Greet(ctx context.Context, req *pb.AddRequest) (*pb.AddResponse, error) {
	sum := req.A + req.B
	log.Println("Sum:", sum)

	return &pb.AddResponse{
		Sum: sum,
	}, nil
}

func (s *server) BidGoodBye(
	ctx context.Context,
	req *farewellpb.GoodByeRequest,
) (*farewellpb.GoodByeResponse, error) {
	return &farewellpb.GoodByeResponse{
		Message: fmt.Sprintf("Goodbye %s! Nice to receive request from you. Farewell my friend!", req.Name),
	}, nil
}

func main() {
	cert := "cert.pem"
	key := "key.pem"

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal("Failed to listen:", err)
	}

	creds, err := credentials.NewServerTLSFromFile(cert, key)
	if err != nil {
		log.Fatal("Failed to load credentials:", err)
	}

	grpcServer := grpc.NewServer(grpc.Creds(creds))

	s := &server{}
	g := &serverGreeter{}

	pb.RegisterCalculateServer(grpcServer, s)
	pb.RegisterGreeterServer(grpcServer, g)
	pb.RegisterBidFarewellServer(grpcServer, s)

	log.Println("Server is running on port :50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal("Failed to serve:", err)
	}
}

// package main

//
// import (
// 	"context"
// 	"fmt"
// 	"log"
// 	"net"
//
// 	pb "simplegrpcserver/proto/gen"
// 	farewellpb "simplegrpcserver/proto/gen/farewell"
//
// 	"google.golang.org/grpc"
// 	"google.golang.org/grpc/credentials"
// )
//
// type server struct {
// 	pb.UnimplementedCalculateServer
// 	pb.BidFarewellServer
// 	//farewellpb.UnimplementedAufWiedersehenServer
// }
//
//
// type serverGreeter struct{
// pb.UnimplementedGreeterServer
// }
//
// func (s *server) Add(ctx context.Context, req *pb.AddRequest) (*pb.AddResponse, error) {
// 	sum := req.A + req.B
// 	log.Println("Sum:", sum)
//
// 	return &pb.AddResponse{
// 		Sum: sum,
// 	}, nil
// }
//
// func (s *server) Greet(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
// 	return &pb.HelloResponse{
// 		Message: fmt.Sprintf("Hello %s. Nice to receive request from you", req.Name),
// 	}, nil
// }
//
// func (s *server) BidGoodBye(
// 	ctx context.Context,
// 	req *farewellpb.GoodByeRequest,
// ) (*farewellpb.GoodByeResponse, error) {
//
// 	return &farewellpb.GoodByeResponse{
// 		Message: fmt.Sprintf(
// 			"Goodbye %s! Nice to receive request from you. Farewell my friend!",
// 			req.Name,
// 		),
// 	}, nil
// }
//
// func main() {
// 	cert := "cert.pem"
// 	key := "key.pem"
//
// 	lis, err := net.Listen("tcp", ":50051")
// 	if err != nil {
// 		log.Fatal("Failed to listen:", err)
// 	}
//
// 	creds, err := credentials.NewServerTLSFromFile(cert, key)
// 	if err != nil {
// 		log.Fatal("Failed to load credentials:", err)
// 	}
//
// 	grpcServer := grpc.NewServer(grpc.Creds(creds))
//
// 	s := &server{}
// 	pb.RegisterCalculateServer(grpcServer, s)
// 	pb.RegisterGreeterServer(grpcServer, serverGreeter{})
// 	pb.RegisterBidFarewellServer(grpcServer, s)
// 	//farewellpb.RegisterAufWiedersehenServer(grpcServer, s)
//
// 	log.Println("Server is running on port :50051")
// 	if err := grpcServer.Serve(lis); err != nil {
// 		log.Fatal("Failed to serve:", err)
// 	}
// }
