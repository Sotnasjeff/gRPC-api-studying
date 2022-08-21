package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

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

	//AddUser(client)
	//AddUserVerbose(client)
	AddUsers(client)
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

func AddUserVerbose(client pb.UserServiceClient) {
	request := &pb.User{
		Id:    "0",
		Name:  "Jeff",
		Email: "jeff@jeff.com",
	}

	responseStream, err := client.AddUserVerbose(context.Background(), request)
	if err != nil {
		log.Fatalf("couldn't make gRPC Request: %v", err)
	}

	for {
		stream, err := responseStream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Couldn't receive msg %v", err)
		}
		fmt.Println("Status:", stream.Status, stream.GetUser())
	}
}

func AddUsers(client pb.UserServiceClient) {
	request := []*pb.User{
		&pb.User{
			Id:    "1",
			Name:  "Jefferson",
			Email: "jef@jef.com",
		},
		&pb.User{
			Id:    "2",
			Name:  "Andre",
			Email: "and@and.com",
		},
		&pb.User{
			Id:    "3",
			Name:  "Adriana",
			Email: "adri@adri.com",
		},
		&pb.User{
			Id:    "4",
			Name:  "Jackson",
			Email: "jack@jack.com",
		},
		&pb.User{
			Id:    "5",
			Name:  "Jessica",
			Email: "jess@jess.com",
		},
		&pb.User{
			Id:    "6",
			Name:  "Sabrina",
			Email: "sab@sab.com",
		},
	}

	stream, err := client.AddUsers(context.Background())
	if err != nil {
		log.Fatalf("Error creating request: %v", err)
	}

	for _, req := range request {
		stream.Send(req)
		time.Sleep(time.Second * 3)
	}

	responser, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error receiving response: %v", err)
	}

	fmt.Println(responser)

}
