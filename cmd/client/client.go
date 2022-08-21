package main

import (
	"context"
	"log"

	"github.com/Sotnasjeff/gRPC-api-studying/pb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	connection, err := grpc.Dial("localhost:5051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("couldn't connect to gRPC Server: %v", err)
	}

	client := pb.NewUserServiceClient(connection)

	AddUser(client)

}

func AddUser(client pb.UserServiceClient) {
	request := &pb.User{
		Id:    "0",
		Name:  "Jeff",
		Email: "jeff@jeff.com",
	}

	response, err := client.AddUser(context.Background(), request)
	if err != nil {
		log.Fatalf("couldn't make gRPC Request: %v", err)
	}

	log.Println(response)
}
