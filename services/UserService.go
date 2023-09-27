package services

import (
	"context"
	"fmt"
	"time"

	"github.com/taivama/golang-training/entities"
	"github.com/taivama/golang-training/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserService struct {
	Users   *mongo.Collection
	Current *entities.User
}

func InitUserService(users *mongo.Collection) *UserService {
	return &UserService{Users: users}
}

func (us *UserService) Register(u *entities.User) (*entities.UserResponse, error) {
	if u.Password != u.ConfirmPassword {
		return nil, fmt.Errorf("given password and confirmed password do not match")
	}
	u.Id = primitive.NewObjectID()
	u.Password = utils.HashPassword(u.Password)
	u.ConfirmPassword = utils.HashPassword(u.ConfirmPassword)
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
	result, err := us.Users.InsertOne(context.Background(), u)
	if err != nil {
		return nil, err
	}
	id := result.InsertedID.(primitive.ObjectID)
	return &entities.UserResponse{Id: id, CreatedAt: u.CreatedAt, UpdatedAt: u.UpdatedAt}, nil
}

func (us *UserService) Login(u *entities.User) *entities.LoginResponse {
	result := us.Users.FindOne(context.Background(), bson.M{"email": u.Email})
	if result.Err() != nil {
		return &entities.LoginResponse{TokenId: "", Error: "login failed, user not found"}
	}
	var user entities.User
	if err := result.Decode(&user); err != nil {
		return &entities.LoginResponse{TokenId: "", Error: "login failed, user decoding failed"}
	}
	if err := utils.VerifyPassword(u.Password, user.Password); err != nil {
		return &entities.LoginResponse{TokenId: "", Error: "login failed, password mismatch"}
	}
	token, err := utils.GenerateToken(user.Email, user.FirstName, user.LastName, user.Id.Hex())
	if err != nil {
		return &entities.LoginResponse{TokenId: "", Error: "login failed, token generation failed"}
	}
	us.Current = &user
	return &entities.LoginResponse{TokenId: token, Error: ""}
}

func (us *UserService) Logout() error {
	if us.Current == nil {
		return fmt.Errorf("no users logged in")
	}
	us.Current = nil
	return nil
}
