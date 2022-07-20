package main

import (
	"context"
	"log"
	"net"

	pb "github.com/Lazaro-Barros/gRPC/sumAPI/proto"
	"google.golang.org/grpc"
)

var address = "127.0.0.1:5200"

type Server struct {
	pb.SumServiceServer
}

func main() {
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("Fail to listen on %s, error: %v\n", address, err)
	}

	log.Printf("Listening on %s\n", address)

	server := grpc.NewServer()
	pb.RegisterSumServiceServer(server, &Server{})
	if err = server.Serve(listener); err != nil {
		log.Fatalf("Fail to listen on %s, error: %v\n", address, err)
	}
}

func (s Server) Sum(ctx context.Context, in *pb.SumRequest) (response *pb.SumResponse, err error) {
	log.Printf("Function Sum invoked with the parameters: %v\n", in)
	var res = pb.SumResponse{
		Result: in.Number1 + in.Number2,
	}
	return &res, nil
}
