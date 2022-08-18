package main

import (
	"log"
	"net"

	"github.com/Sotnasjeff/gRPC-api-studying/pb"
	"github.com/Sotnasjeff/gRPC-api-studying/services"

	"google.golang.org/grpc"
)

func main() {

	lis, err := net.Listen("tcp", "localhost:5051")
	if err != nil {
		log.Fatalf("Couldn't connect %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterUserServiceServer(grpcServer, services.NewUserService)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Couldn't connect server: %v", err)
	}

}
