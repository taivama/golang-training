package interfaces

import "github.com/taivama/golang-training/entities"

type IUser interface {
	Register(u *entities.User) (*entities.UserResponse, error)
	Login(u *entities.User) *entities.LoginResponse
	Logout() error
}
