package controllers

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/taivama/golang-training/proto"
)

type ProductController struct {
	Product proto.ProductServiceClient
}

func InitProductController(p proto.ProductServiceClient) *ProductController {
	return &ProductController{Product: p}
}

func (pc *ProductController) AddProduct(c *gin.Context) {
	var product proto.Product
	if err := c.BindJSON(&product); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	result, err := pc.Product.AddProduct(context.Background(), &product)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	if result.Error != "" {
		c.String(http.StatusBadRequest, result.Error)
		return
	}
	c.IndentedJSON(http.StatusCreated, "Added")
}

func (pc *ProductController) GetProductById(c *gin.Context) {
	request := proto.GetRequest{Id: c.Param("id")}
	result, err := pc.Product.GetProductById(context.Background(), &request)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	if result.Error != "" {
		c.String(http.StatusNotFound, result.Error)
		return
	}
	c.IndentedJSON(http.StatusFound, result.Product)
}

func (pc *ProductController) SearchProducts(c *gin.Context) {
	request := proto.SearchRequest{Name: c.Param("name")}
	result, err := pc.Product.SearchProducts(context.Background(), &request)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	if result.Error != "" {
		c.String(http.StatusNotFound, result.Error)
		return
	}
	c.IndentedJSON(http.StatusFound, result.Products)
}
