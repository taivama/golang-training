package services

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/taivama/golang-training/entities"
)

var (
	prodSvc *ProductService
	p       entities.Product
)

func TestAddProduct(t *testing.T) {
	c := getCollection("Matti", "TestProducts", t)
	prodSvc = InitProductService(c)
	p = entities.Product{
		Name:     "item-1",
		Category: "general",
		Quantity: 10,
	}
	err := prodSvc.AddProduct(&p)
	assert.Nil(t, err)
}
func TestGetProductById(t *testing.T) {
	result, err := prodSvc.GetProductById(p.Id)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestSearchProducts(t *testing.T) {
	result, err := prodSvc.SearchProducts(p.Name)
	assert.NotNil(t, result)
	assert.Nil(t, err)
	assert.GreaterOrEqual(t, len(result), 1)
}
