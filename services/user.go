package services

import (
	"context"
	"time"

	"github.com/Sotnasjeff/gRPC-api-studying/pb"
)

//type UserServiceServer interface {
//	AddUser(context.Context, *User) (*User, error)
//	mustEmbedUnimplementedUserServiceServer()
//  AddUserVerbose(ctx context.Context, in *User, opts ...grpc.CallOption) (UserService_AddUserVerboseClient, error)
//}

type UserService struct {
	pb.UnimplementedUserServiceServer
}

func NewUserService() *UserService {
	return &UserService{}
}

func (*UserService) AddUser(ctx context.Context, request *pb.User) (*pb.User, error) {

	return &pb.User{
		Id:    request.GetId(),
		Name:  request.GetName(),
		Email: request.GetEmail(),
	}, nil
}

func (*UserService) AddUserVerbose(request *pb.User, stream pb.UserService_AddUserVerboseServer) error {

	stream.Send(&pb.UserResultStream{
		Status: "Init",
		User:   &pb.User{},
	})

	time.Sleep(time.Second * 3)

	stream.Send(&pb.UserResultStream{
		Status: "Inserting",
		User:   &pb.User{},
	})

	time.Sleep(time.Second * 3)

	stream.Send(&pb.UserResultStream{
		Status: "User has been inserted",
		User:   &pb.User{
			Id: request.GetId(),
			Name: request.GetName(),
			Email: request.GetEmail(),
		},
	})

	time.Sleep(time.Second * 3)

	return nil
}
