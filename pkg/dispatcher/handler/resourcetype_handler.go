package handler

import (
	"context"
	"github.com/carmanzhang/ks-alert/pkg/dispatcher/pb"
)

type ResourceTypeHandler struct{}

// resource type
func (server ResourceTypeHandler) CreateResourceType(ctx context.Context, resourceType *pb.ResourceType) (*pb.ResourceTypeResponse, error) {
	return nil, nil
}

func (server ResourceTypeHandler) DeleteResourceType(ctx context.Context, resourceTypeID *pb.ResourceTypeID) (*pb.ResourceTypeResponse, error) {
	return nil, nil
}

func (server ResourceTypeHandler) UpdateResourceType(ctx context.Context, resourceType *pb.ResourceType) (*pb.ResourceTypeResponse, error) {
	return nil, nil
}

func (server ResourceTypeHandler) GetResourceType(ctx context.Context, resourceTypeID *pb.ResourceTypeID) (*pb.ResourceTypeResponse, error) {
	return nil, nil
}
