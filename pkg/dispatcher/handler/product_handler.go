package handler

import (
	"context"
	"kubesphere.io/ks-alert/pkg/dispatcher/pb"
)

// product
func (server Server) CreateProduct(ctx context.Context, product *pb.Product) (*pb.ProductResponse, error) {
	return nil, nil
}

func (server Server) DeleteProduct(ctx context.Context, productID *pb.ProductID) (*pb.ProductResponse, error) {
	return nil, nil
}

func (server Server) UpdateProduct(ctx context.Context, product *pb.Product) (*pb.ProductResponse, error) {
	return nil, nil
}

func (server Server) GetProduct(ctx context.Context, productID *pb.ProductID) (*pb.ProductResponse, error) {
	return nil, nil
}
