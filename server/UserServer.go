package server

import (
	"context"

	"github.com/taivama/golang-training/entities"
	"github.com/taivama/golang-training/interfaces"
	"github.com/taivama/golang-training/proto"
)

type UserServer struct {
	proto.UnimplementedUserServiceServer
	User interfaces.IUser
}

func InitUserServer(u interfaces.IUser) *UserServer {
	return &UserServer{User: u}
}

func (us *UserServer) Register(ctx context.Context, u *proto.User) (*proto.RegisterResponse, error) {
	user := entities.User{
		FirstName:       u.FirstName,
		LastName:        u.LastName,
		Age:             int(u.Age),
		Email:           u.Email,
		Password:        u.Password,
		ConfirmPassword: u.ConfirmPassword,
	}
	response, err := us.User.Register(&user)
	if err != nil {
		return &proto.RegisterResponse{Response: nil, Error: err.Error()}, nil
	}
	r := proto.UserResponse{
		Id:        response.Id.Hex(),
		CreatedAt: response.CreatedAt.String(),
		UpdatedAt: response.UpdatedAt.String(),
	}
	return &proto.RegisterResponse{Response: &r, Error: ""}, nil
}

func (us *UserServer) Login(ctx context.Context, u *proto.User) (*proto.LoginResponse, error) {
	user := entities.User{
		FirstName:       u.FirstName,
		LastName:        u.LastName,
		Age:             int(u.Age),
		Email:           u.Email,
		Password:        u.Password,
		ConfirmPassword: u.ConfirmPassword,
	}
	response := us.User.Login(&user)
	return &proto.LoginResponse{
		TokenId: response.TokenId,
		Error:   response.Error,
	}, nil
}

func (us *UserServer) Logout(ctx context.Context, e *proto.Empty) (*proto.LogoutResponse, error) {
	err := us.User.Logout()
	if err != nil {
		return &proto.LogoutResponse{Error: err.Error()}, nil
	}
	return &proto.LogoutResponse{}, nil
}
