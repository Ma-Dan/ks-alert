package handler

import (
	"context"
	"github.com/carmanzhang/ks-alert/pkg/dispatcher/pb"
)

type ProductHandler struct{}

// product
func (server ProductHandler) CreateProduct(ctx context.Context, product *pb.Product) (*pb.ProductResponse, error) {
	return nil, nil
}

func (server ProductHandler) DeleteProduct(ctx context.Context, productID *pb.ProductID) (*pb.ProductResponse, error) {
	return nil, nil
}

func (server ProductHandler) UpdateProduct(ctx context.Context, product *pb.Product) (*pb.ProductResponse, error) {
	return nil, nil
}

func (server ProductHandler) GetProduct(ctx context.Context, productID *pb.ProductID) (*pb.ProductResponse, error) {
	return nil, nil
}
