package handler

import (
	"context"
	"github.com/carmanzhang/ks-alert/pkg/dispatcher/pb"
)

type ResourceHandler struct{}

// suggestion
func (server ResourceHandler) GetResource(ctx context.Context, resource *pb.Resource) (*pb.Resource, error) {
	return nil, nil
}

