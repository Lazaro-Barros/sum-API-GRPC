package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/Lazaro-Barros/gRPC/sumAPI/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var address = "127.0.0.1:5200"

func main() {
	conecction, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Fail to connect to %v, error:%v\n", address, err)
	}
	defer conecction.Close()

	client := pb.NewSumServiceClient(conecction)
	var request = pb.SumRequest{
		Number1: 100,
		Number2: 100,
	}
	response, err := client.Sum(context.Background(), &request)
	if err != nil {
		log.Fatalf("Fail on request sum service. error:%v\n", err)
	}
	fmt.Println(response.Result)
}
