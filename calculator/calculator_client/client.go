package main

import (
	"context"
	"fmt"
	"github.com/yohang88/learn-golang-grpc/calculator/calculatorpb"
	"google.golang.org/grpc"
	"io"
	"log"
)

func main()  {
	fmt.Printf("Client ready to connect")

	conn, err := grpc.Dial("localhost:50002", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}

	defer conn.Close()

	client := calculatorpb.NewCalculatorServiceClient(conn)

	doServerStreaming(client)
}

func doUnary(client calculatorpb.CalculatorServiceClient) {
	fmt.Println("Starting to do a unary RPC...")

	request := &calculatorpb.SumRequest{
		FirstInteger:  40,
		SecondInteger: 20,
	}

	response, err := client.Sum(context.Background(), request)

	if err != nil {
		log.Fatalf("Error while calling Sum RPC: %v", err)
	}

	log.Printf("Response from Sum: %v", response)
}

func doServerStreaming(client calculatorpb.CalculatorServiceClient) {
	fmt.Println("Starting to do a server streaming RPC...")

	request := &calculatorpb.PrimeNumberDecompositionRequest{
		InputInteger: 123456789,
	}

	stream, err := client.PrimeNumberDecomposition(context.Background(), request)

	if err != nil {
		log.Fatalf("Error while calling PrimeNumberDecomposition RPC: %v", err)
	}

	for {
		message, err := stream.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Error while reading stream: %v", err)
		}

		fmt.Printf("Response from RPC: %v\n", message.GetResult())
	}
}