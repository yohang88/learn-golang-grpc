package main

import (
	"context"
	"fmt"
	"github.com/yohang88/learn-golang-grpc/greet/greetpb"
	"google.golang.org/grpc"
	"log"
)

func main()  {
	fmt.Println("Hello I'm a client")

	conn, err := grpc.Dial("localhost:50001", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}

	defer conn.Close()

	c := greetpb.NewGreetServiceClient(conn)

	request := &greetpb.GreetRequest{
		FirstName: "Yoga",
		LastName:  "Hanggara",
	}

	response, err := c.Greet(context.Background(), request)

	if err != nil {
		log.Fatalf("Error while calling Greet RPC: %v", err)
	}

	log.Printf("Response from Greet: %v", response.Result)
}