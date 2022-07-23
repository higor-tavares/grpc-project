package main

import (
	"context"
	"log"

	"com.higortavares.grpc-go/pb"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatal("Error while connecting to server", err)
	}
	defer conn.Close()
	client := pb.NewEchoServiceClient(conn)
	req := &pb.EchoRequest{
		Name: "Higor",
	}
	res, err := client.Echo(context.Background(), req)
	if err != nil {
		log.Fatal("Error", err)
	}
	log.Println(res.Echo)
}
