package main

import (
	"context"
	"log"
	"time"

	mainapipb "simplegrpcclient/proto/gen"
	farewellpb "simplegrpcclient/proto/gen/farewell"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func main() {
	cert := "cert.pem"

	creds, err := credentials.NewClientTLSFromFile(cert, "")
	if err != nil {
		log.Fatalln("failed to load certificates:", err)
	}

	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatalln("did not connect:", err)
	}
	defer conn.Close()

	calcClient := mainapipb.NewCalculateClient(conn)
	greeterClient := mainapipb.NewGreeterClient(conn)
	farewellClient := mainapipb.NewBidFarewellClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	addReq := &mainapipb.AddRequest{
		A: 10,
		B: 20,
	}

	addRes, err := calcClient.Add(ctx, addReq)
	if err != nil {
		log.Fatalln("could not add:", err)
	}

	greetReq := &mainapipb.HelloRequest{
		Name: "Eddy",
	}

	greetRes, err := greeterClient.Greet(ctx, greetReq)
	if err != nil {
		log.Fatalln("could not greet:", err)
	}

	goodByeReq := &farewellpb.GoodByeRequest{
		Name: "Grace",
	}

	goodByeRes, err := farewellClient.BidGoodBye(ctx, goodByeReq)
	if err != nil {
		log.Fatalln("could not bid goodbye:", err)
	}

	log.Println("Sum:", addRes.Sum)
	log.Println("******************Greeter Message from the Add function in proto file:", greetRes.Message)
	log.Println("Goodbye Message:", goodByeRes.Message)

	state := conn.GetState()
	log.Println("Connection State:", state)
}

// package main
//
// import (
// 	"context"
// 	"log"
// 	mainapipb "simplegrpcclient/proto/gen"
// 	farewellpb "simplegrpcclient/proto/gen/farewell"
// 	"time"
//
// 	"google.golang.org/grpc"
// 	"google.golang.org/grpc/credentials"
// )
//
// func main() {
// 	cert := "cert.pem"
//
// 	creds, err := credentials.NewClientTLSFromFile(cert, "")
// 	if err != nil {
// 		log.Fatalln("failed to load certificates", err)
// 	}
//
// 	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(creds))
// 	if err != nil {
// 		log.Fatalln("Did not connect:", err)
// 	}
// 	defer conn.Close()
//
// 	client := mainapipb.NewCalculateClient(conn)
// 	client2 := mainapipb.NewGreeterClient(conn)
// 	//fwClient := farewellpb.NewAufWiedersehenClient(conn)
// 	client3 := mainapipb.NewBidFarewellClient(conn)
//
// 	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// 	defer cancel()
//
// 	req := mainapipb.AddRequest{
// 		A: 10,
// 		B: 20,
// 	}
//
// 	res, err := client.Add(ctx, &req)
// 	if err != nil {
// 		log.Fatalln("Could not add", err)
// 	}
//
// 	reqGreet := &mainapipb.HelloRequest{
// 		Name: "Eddy",
// 	}
//
// 	res1, err := client2.Greet(ctx, reqGreet)
// 	if err != nil {
// 		log.Fatalln("Could not greet", err)
// 	}
//
// 	resAddFromGreeter, err := client2.Add(ctx, reqGreet)
// 	if err != nil {
// 		log.Fatalln("Could not greet", err)
// 	}
//
// 	reqGoodBye := &farewellpb.GoodByeRequest{
// 		Name: "Grace",
// 	}
//
// 	// resFw, err := fwClient.BidGoodBye(ctx, reqGoodBye)
// 	// if err != nil {
// 	// 	log.Fatalln("Could not bid goodbye", err)
// 	// }
// 	res3, err := client3.BidGoodBye(ctx, reqGoodBye)
// 	if err != nil {
// 		log.Fatalln("Could not bid goodbye", err)
// 	}
//
// 	log.Println("Sum:", res.Sum)
// 	log.Println("Greeter Message:", res1.Message)
// 	log.Println("********Greeter Message From the second Add function in proto file:", resAddFromGreeter.Message)
// 	//log.Println("Goodbye message:", resFw.Message)
// 	log.Println("Goodbye message:", res3.Message)
//
// 	state := conn.GetState()
// 	log.Println("Connection State:", state)
// }
