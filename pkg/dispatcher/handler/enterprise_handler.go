package handler

import (
	"context"
	"github.com/carmanzhang/ks-alert/pkg/dispatcher/pb"
)

type EnterpriseHandler struct{}
// enterprise
func (server EnterpriseHandler) CreateEnterprise(ctx context.Context, enterprise *pb.Enterprise) (*pb.EnterpriseResponse, error) {
	return nil, nil
}

func (server EnterpriseHandler) DeleteEnterprise(ctx context.Context, entID *pb.EnterpriseID) (*pb.EnterpriseResponse, error) {
	return nil, nil
}

func (server EnterpriseHandler) UpdateEnterprise(ctx context.Context, ent *pb.Enterprise) (*pb.EnterpriseResponse, error) {
	return nil, nil
}
func (server EnterpriseHandler) GetEnterprise(ctx context.Context, entID *pb.EnterpriseID) (*pb.EnterpriseResponse, error) {
	return nil, nil
}
