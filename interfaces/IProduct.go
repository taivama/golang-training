package interfaces

import (
	"github.com/taivama/golang-training/entities"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type IProduct interface {
	AddProduct(p *entities.Product) error
	GetProductById(id primitive.ObjectID) (*entities.Product, error)
	SearchProducts(name string) ([]*entities.Product, error)
}
