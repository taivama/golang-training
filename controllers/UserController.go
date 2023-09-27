package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/taivama/golang-training/entities"
	"github.com/taivama/golang-training/interfaces"
)

type UserController struct {
	User interfaces.IUser
}

func InitUserController(u interfaces.IUser) *UserController {
	return &UserController{User: u}
}

func (uc *UserController) RegisterUser(c *gin.Context) {
	var user entities.User
	if err := c.BindJSON(&user); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	result, err := uc.User.Register(&user)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	c.IndentedJSON(http.StatusCreated, result)
}

func (uc *UserController) Login(c *gin.Context) {
	var user entities.User
	if err := c.BindJSON(&user); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	result := uc.User.Login(&user)
	c.IndentedJSON(http.StatusOK, result)
}

func (uc *UserController) Logout(c *gin.Context) {
	if err := uc.User.Logout(); err != nil {
		c.String(http.StatusNotFound, err.Error())
		return
	}
	c.IndentedJSON(http.StatusOK, "OK")
}
