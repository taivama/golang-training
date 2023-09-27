package entities

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Id              primitive.ObjectID `json:"id" bson:"_id"`
	FirstName       string             `json:"firstname" bson:"firstname,required"`
	LastName        string             `json:"lastname" bson:"lastname,required"`
	Age             int                `json:"age" bson:"age,required"`
	Email           string             `json:"email" bson:"email,required"`
	Password        string             `json:"password" bson:"password,required"`
	ConfirmPassword string             `json:"confirm" bson:"confirm,required"`
	CreatedAt       time.Time          `json:"createdAt" bson:"createdAt"`
	UpdatedAt       time.Time          `json:"updatedAt" bson:"updatedAt"`
}

type LoginResponse struct {
	TokenId string `json:"tokenId"`
	Error   string `json:"error"`
}

type UserResponse struct {
	Id        primitive.ObjectID `json:"id" bson:"_id"`
	CreatedAt time.Time          `json:"createdAt" bson:"createdAt"`
	UpdatedAt time.Time          `json:"updatedAt" bson:"updatedAt"`
}
