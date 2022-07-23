package main

import (
	"context"
	"log"
	"net"

	"com.higortavares.grpc-go/pb"
	"google.golang.org/grpc"
)

type server struct {
}

func (*server) Echo(ctx context.Context, request *pb.EchoRequest) (*pb.EchoResponse, error) {
	result := "Hello " + request.GetName()
	res := &pb.EchoResponse{
		Echo: result,
	}
	return res, nil
}

func main() {
	list, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatal("There is a problem", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterEchoServiceServer(grpcServer, &server{})
	if err := grpcServer.Serve(list); err != nil {
		log.Fatal("There is another problem", err)
	}
}
