package services

import (
	"context"

	"github.com/Sotnasjeff/gRPC-api-studying/pb"
)

//type UserServiceServer interface {
//	AddUser(context.Context, *User) (*User, error)
//	mustEmbedUnimplementedUserServiceServer()
//}

type UserService struct {
	pb.UnimplementedUserServiceServer
}

func NewUserService() *UserService {
	return &UserService{}
}

func (*UserService) AddUser(ctx context.Context, request *pb.User) (*pb.User, error) {

	return &pb.User{
		Id:    "2918928",
		Name:  request.GetName(),
		Email: request.GetEmail(),
	}, nil
}
