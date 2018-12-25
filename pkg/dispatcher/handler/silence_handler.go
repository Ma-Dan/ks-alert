package handler

import (
	"context"
	"kubesphere.io/ks-alert/pkg/dispatcher/pb"
)

// silence
func (server Server) CreateSilence(ctx context.Context, silence *pb.Silence) (*pb.SilenceResponse, error) {
	return nil, nil
}

func (server Server) DeleteSilence(ctx context.Context, silence *pb.Silence) (*pb.SilenceResponse, error) {
	return nil, nil
}

func (server Server) UpdateSilence(ctx context.Context, silence *pb.Silence) (*pb.SilenceResponse, error) {
	return nil, nil
}

func (server Server) GetSilence(ctx context.Context, silence *pb.Silence) (*pb.SilenceResponse, error) {
	return nil, nil
}
