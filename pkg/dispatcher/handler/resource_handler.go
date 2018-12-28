package handler

import (
	"github.com/carmanzhang/ks-alert/pkg/dispatcher/pb"
	"context"
)

type ResourceHandler struct{}
// alert rule
func (server ResourceHandler) CreateResource(ctx context.Context, Resource *pb.ResourceGroup) (*pb.ResourceGroupResponse, error) {
	return nil, nil
}

func (server ResourceHandler) DeleteResource(ctx context.Context, ResourceID *pb.ResourceGroupSpec) (*pb.ResourceGroupResponse, error) {
	return nil, nil
}

func (server ResourceHandler) UpdateResource(ctx context.Context, Resource *pb.ResourceGroup) (*pb.ResourceGroupResponse, error) {
	return nil, nil
}

func (server ResourceHandler) GetResource(ctx context.Context, ResourceID *pb.ResourceGroupSpec) (*pb.ResourceGroupResponse, error) {
	return nil, nil
}


