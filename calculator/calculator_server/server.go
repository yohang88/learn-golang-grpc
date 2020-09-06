package main

import (
	"context"
	"fmt"
	"github.com/yohang88/learn-golang-grpc/calculator/calculatorpb"
	"google.golang.org/grpc"
	"io"
	"log"
	"net"
)

type server struct {}

func (s server) FindMaximum(stream calculatorpb.CalculatorService_FindMaximumServer) error {
	fmt.Printf("Find Maximum was invoked.\n")

	maxInteger := int32(0)

	for {
		request, err := stream.Recv()

		if err == io.EOF {
			return nil
		}

		if err != nil {
			log.Fatalf("Error while reading client stream: %v", err)
		}

		inputInteger := request.GetInputInteger()

		if inputInteger > maxInteger {
			maxInteger = inputInteger

			errSend := stream.Send(&calculatorpb.FindMaximumResponse{
				MaxInteger: maxInteger,
			})

			if errSend != nil {
				log.Fatalf("Error while sending data to client: %v", errSend)
				return errSend
			}
		}
	}
}

func (s server) ComputeAverage(stream calculatorpb.CalculatorService_ComputeAverageServer) error {
	fmt.Printf("Compute Average was invoked.\n")

	sumInput := int32(0)
	countInput := int32(0)

	for {
		request, err := stream.Recv()

		if err == io.EOF {
			resultAverage := float64(sumInput) / float64(countInput)

			return stream.SendAndClose(&calculatorpb.ComputeAverageResponse{
				ResultAverage: resultAverage,
			})
		}

		if err != nil {
			log.Fatalf("Error while reading client stream %v", err)
		}

		countInput++
		sumInput += request.GetInputInteger()
	}
}

func (s server) PrimeNumberDecomposition(request *calculatorpb.PrimeNumberDecompositionRequest, stream calculatorpb.CalculatorService_PrimeNumberDecompositionServer) error {
	fmt.Printf("Prime number decomposition was invoked with %v\n", request)

	inputInteger := request.GetInputInteger()
	k := int32(2)

	for inputInteger > 1 {
		if inputInteger % k == 0 {
			err := stream.Send(&calculatorpb.PrimeNumberDecompositionResponse{
				Result: k,
			})

			if err != nil {
				log.Fatalf("Error send response stream")
			}

			inputInteger = inputInteger / k
		} else {
			k = k + 1
			fmt.Printf("Divisor has increased to %v\n", k)
		}
	}

	return nil
}

func (s server) Sum(ctx context.Context, request *calculatorpb.SumRequest) (*calculatorpb.SumResponse, error) {
	fmt.Printf("Sum function was invoked with %v", request)

	firstInteger := request.GetFirstInteger()
	secondInteger := request.GetSecondInteger()

	result := firstInteger + secondInteger

	response := &calculatorpb.SumResponse{
		SumResult: result,
	}

	return response, nil
}

func main() {
	fmt.Println("CalculatorService start to listening...")

	listener, err := net.Listen("tcp", "0.0.0.0:50002")

	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	rpcServer := grpc.NewServer()
	calculatorpb.RegisterCalculatorServiceServer(rpcServer, &server{})

	err = rpcServer.Serve(listener)

	if err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
