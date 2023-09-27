package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/taivama/golang-training/entities"
	"github.com/taivama/golang-training/interfaces"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ProductController struct {
	Product interfaces.IProduct
}

func InitProductController(p interfaces.IProduct) *ProductController {
	return &ProductController{Product: p}
}

func (pc *ProductController) AddProduct(c *gin.Context) {
	var product entities.Product
	if err := c.BindJSON(&product); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	if err := pc.Product.AddProduct(&product); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	c.IndentedJSON(http.StatusCreated, "Added")
}

func (pc *ProductController) GetProductById(c *gin.Context) {
	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	p, err := pc.Product.GetProductById(id)
	if err != nil {
		c.String(http.StatusNotFound, err.Error())
		return
	}
	c.IndentedJSON(http.StatusFound, p)
}

func (pc *ProductController) SearchProducts(c *gin.Context) {
	name := c.Param("name")
	products, err := pc.Product.SearchProducts(name)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	c.IndentedJSON(http.StatusFound, products)
}
