package handler

import (
	"context"
	"kubesphere.io/ks-alert/pkg/dispatcher/pb"
)

// resource type
func (server Server) CreateResourceType(ctx context.Context, resourceType *pb.ResourceType) (*pb.ResourceTypeResponse, error) {
	return nil, nil
}

func (server Server) DeleteResourceType(ctx context.Context, resourceTypeID *pb.ResourceTypeID) (*pb.ResourceTypeResponse, error) {
	return nil, nil
}

func (server Server) UpdateResourceType(ctx context.Context, resourceType *pb.ResourceType) (*pb.ResourceTypeResponse, error) {
	return nil, nil
}

func (server Server) GetResourceType(ctx context.Context, resourceTypeID *pb.ResourceTypeID) (*pb.ResourceTypeResponse, error) {
	return nil, nil
}
