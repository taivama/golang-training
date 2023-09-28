package controllers

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/taivama/golang-training/proto"
)

type UserController struct {
	User proto.UserServiceClient
}

func InitUserController(u proto.UserServiceClient) *UserController {
	return &UserController{User: u}
}

func (uc *UserController) RegisterUser(c *gin.Context) {
	var user proto.User
	if err := c.BindJSON(&user); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	result, err := uc.User.Register(context.Background(), &user)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	c.IndentedJSON(http.StatusCreated, result.Response)
}

func (uc *UserController) Login(c *gin.Context) {
	var user proto.User
	if err := c.BindJSON(&user); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	result, err := uc.User.Login(context.Background(), &user)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	c.IndentedJSON(http.StatusOK, result)
}

func (uc *UserController) Logout(c *gin.Context) {
	result, err := uc.User.Logout(context.Background(), &proto.Empty{})
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	if result.Error != "" {
		c.String(http.StatusNotFound, result.Error)
		return
	}
	c.IndentedJSON(http.StatusOK, "OK")
}
