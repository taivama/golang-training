package services

import (
	"context"
	"fmt"
	"time"

	"github.com/taivama/golang-training/entities"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProductService struct {
	Products *mongo.Collection
}

func InitProductService(products *mongo.Collection) *ProductService {
	return &ProductService{Products: products}
}

func (ps *ProductService) AddProduct(p *entities.Product) error {
	p.Id = primitive.NewObjectID()
	p.CreatedAt = time.Now()
	p.UpdatedAt = time.Now()
	_, err := ps.Products.InsertOne(context.Background(), p)
	if err != nil {
		return err
	}
	return nil
}

func (ps *ProductService) GetProductById(id primitive.ObjectID) (*entities.Product, error) {
	result := ps.Products.FindOne(context.Background(), bson.M{"_id": id})
	if result.Err() != nil {
		return nil, fmt.Errorf("get product with id failed: %w", result.Err())
	}
	var product entities.Product
	if err := result.Decode(&product); err != nil {
		return nil, fmt.Errorf("decoding product failed: %w", result.Err())
	}
	return &product, nil
}

func (ps *ProductService) SearchProducts(name string) ([]*entities.Product, error) {
	result, err := ps.Products.Find(context.Background(), bson.M{"name": name})
	if err != nil {
		return nil, fmt.Errorf("search products with names: %w", err)
	}
	defer result.Close(context.Background())

	products := []*entities.Product{}
	for result.Next(context.Background()) {
		var p entities.Product
		if err := result.Decode(&p); err != nil {
			return nil, fmt.Errorf("decoding searhced products failed: %w", result.Err())
		}
		products = append(products, &p)
	}
	return products, nil
}
