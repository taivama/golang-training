package services

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/taivama/golang-training/entities"
	"github.com/taivama/golang-training/utils"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	userSvc *UserService
)

func getCollection(db string, name string, t *testing.T) *mongo.Collection {
	c, err := utils.ConnectDB(context.Background())
	if err != nil {
		t.Errorf("db connection failed: %s", err.Error())
	}
	return c.Database(db).Collection(name)
}

func TestRegisterUser(t *testing.T) {
	c := getCollection("Matti", "TestUsers", t)
	userSvc = InitUserService(c)
	u := entities.User{
		FirstName:       "John",
		LastName:        "Doe",
		Email:           "john.doe@abc.com",
		Age:             30,
		Password:        "john",
		ConfirmPassword: "john",
	}
	response, err := userSvc.Register(&u)
	if err != nil {
		t.Errorf("register user failed: %s", err.Error())
	}
	assert.NotNil(t, response.Id)
}

func TestLogin(t *testing.T) {
	TestRegisterUser(t)
	u := entities.User{
		FirstName:       "John",
		LastName:        "Doe",
		Email:           "john.doe@abc.com",
		Age:             30,
		Password:        "john",
		ConfirmPassword: "john",
	}
	response := userSvc.Login(&u)
	assert.NotEmpty(t, response.TokenId)
	assert.Empty(t, response.Error)
}

func TestLogout(t *testing.T) {
	TestLogin(t)
	err := userSvc.Logout()
	assert.Nil(t, err)
}
