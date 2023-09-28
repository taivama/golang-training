package server

import (
	"context"
	"fmt"

	"github.com/taivama/golang-training/entities"
	"github.com/taivama/golang-training/interfaces"
	"github.com/taivama/golang-training/proto"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ProductServer struct {
	proto.UnimplementedProductServiceServer
	Product interfaces.IProduct
}

func InitProductServer(p interfaces.IProduct) *ProductServer {
	return &ProductServer{Product: p}
}

func (ps *ProductServer) AddProduct(ctx context.Context, p *proto.Product) (*proto.AddResponse, error) {
	prod := entities.Product{
		Name:     p.Name,
		Category: p.Category,
		Quantity: int(p.Quantity),
	}
	err := ps.Product.AddProduct(&prod)
	if err != nil {
		return &proto.AddResponse{Error: err.Error()}, nil
	}
	return &proto.AddResponse{}, nil
}

func (ps *ProductServer) GetProductById(ctx context.Context, r *proto.GetRequest) (*proto.GetResponse, error) {
	id, err := primitive.ObjectIDFromHex(r.Id)
	if err != nil {
		return nil, fmt.Errorf("invalid object id: %w", err)
	}
	prod, err := ps.Product.GetProductById(id)
	if err != nil {
		return &proto.GetResponse{Error: err.Error()}, nil
	}
	p := proto.Product{
		Id:        prod.Id.Hex(),
		Name:      prod.Name,
		Category:  prod.Category,
		Quantity:  int32(prod.Quantity),
		CreatedAt: prod.CreatedAt.String(),
		UpdatedAt: prod.UpdatedAt.String(),
	}
	return &proto.GetResponse{Product: &p}, nil
}

func (ps *ProductServer) SearchProducts(ctx context.Context, r *proto.SearchRequest) (*proto.SearchResponse, error) {
	products, err := ps.Product.SearchProducts(r.Name)
	if err != nil {
		return &proto.SearchResponse{Error: err.Error()}, nil
	}
	prods := []*proto.Product{}
	for _, product := range products {
		p := proto.Product{
			Id:        product.Id.Hex(),
			Name:      product.Name,
			Category:  product.Category,
			Quantity:  int32(product.Quantity),
			CreatedAt: product.CreatedAt.String(),
			UpdatedAt: product.UpdatedAt.String(),
		}
		prods = append(prods, &p)
	}
	return &proto.SearchResponse{Products: prods}, nil
}
