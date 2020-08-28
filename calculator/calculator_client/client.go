package main

import (
	"context"
	"fmt"
	"github.com/yohang88/learn-golang-grpc/calculator/calculatorpb"
	"google.golang.org/grpc"
	"log"
)

func main()  {
	fmt.Printf("Client ready to connect")

	conn, err := grpc.Dial("localhost:50002", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}

	defer conn.Close()

	rpcClient := calculatorpb.NewCalculatorServiceClient(conn)

	request := &calculatorpb.SumRequest{
		FirstInteger:  40,
		SecondInteger: 20,
	}

	response, err := rpcClient.Sum(context.Background(), request)

	if err != nil {
		log.Fatalf("Error while calling Sum RPC: %v", err)
	}

	log.Printf("Response from Sum: %v", response)
}