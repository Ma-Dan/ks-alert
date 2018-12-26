package handler

import (
	"context"
	"kubesphere.io/ks-alert/pkg/dispatcher/pb"
)

// enterprise
func (server Server) CreateEnterprise(ctx context.Context, enterprise *pb.Enterprise) (*pb.EnterpriseResponse, error) {
	return nil, nil
}

func (server Server) DeleteEnterprise(ctx context.Context, entID *pb.EnterpriseID) (*pb.EnterpriseResponse, error) {
	return nil, nil
}

func (server Server) UpdateEnterprise(ctx context.Context, ent *pb.Enterprise) (*pb.EnterpriseResponse, error) {
	return nil, nil
}
func (server Server) GetEnterprise(ctx context.Context, entID *pb.EnterpriseID) (*pb.GetEnterpriseResponse, error) {
	return nil, nil
}
