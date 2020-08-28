package main

import (
	"context"
	"fmt"
	"github.com/yohang88/learn-golang-grpc/greet/greetpb"
	"log"
	"net"

	"google.golang.org/grpc"
)

type server struct {}

func (*server) Greet(ctx context.Context, req *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {
	fmt.Printf("Greet function was invoked with %v", req)

	firstName := req.GetFirstName()

	result := "Hello " + firstName

	response := &greetpb.GreetResponse{
		Result: result,
	}

	return response, nil
}

func main()  {
	fmt.Println("Hello World")

	listener, err := net.Listen("tcp", "0.0.0.0:50001")

	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	greetpb.RegisterGreetServiceServer(s, &server{})

	err = s.Serve(listener)

	if err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}