package controller

import (
	userapiv1 "Identity/internal/proto/identity/userapi/v1"
	"context"
)

type UserServiceServer struct {
	userapiv1.UnimplementedUserServiceServer
}

var _ userapiv1.UserServiceServer = (*UserServiceServer)(nil)

func (s *UserServiceServer) CreateUser(context.Context, *userapiv1.CreateUserRequest) (*userapiv1.CreateUserResponse, error) {

	return &userapiv1.CreateUserResponse{}, nil
}

func (s *UserServiceServer) UserByID(context.Context, *userapiv1.UserByIDRequest) (*userapiv1.UserByIDResponse, error) {

	return &userapiv1.UserByIDResponse{}, nil
}
