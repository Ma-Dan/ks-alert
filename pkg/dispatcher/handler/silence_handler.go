package handler

import (
	"context"
	"github.com/carmanzhang/ks-alert/pkg/pb"
)

type SilenceHandler struct{}

// silence
func (server SilenceHandler) CreateSilence(ctx context.Context, silence *pb.Silence) (*pb.SilenceResponse, error) {
	return nil, nil
}

func (server SilenceHandler) DeleteSilence(ctx context.Context, silence *pb.Silence) (*pb.SilenceResponse, error) {
	return nil, nil
}

func (server SilenceHandler) UpdateSilence(ctx context.Context, silence *pb.Silence) (*pb.SilenceResponse, error) {
	return nil, nil
}

func (server SilenceHandler) GetSilence(ctx context.Context, silence *pb.Silence) (*pb.SilenceResponse, error) {
	return nil, nil
}
