package handler

import (
	"context"
	"kubesphere.io/ks-alert/pkg/dispatcher/pb"
)

// suggestion
func (server Server) CreateSuggestion(ctx context.Context, suggestion *pb.Suggestion) (*pb.SuggestionResponse, error) {
	return nil, nil
}

func (server Server) DeleteSuggestion(ctx context.Context, suggestion *pb.Suggestion) (*pb.SuggestionResponse, error) {
	return nil, nil
}

func (server Server) UpdateSuggestion(ctx context.Context, suggestion *pb.Suggestion) (*pb.SuggestionResponse, error) {
	return nil, nil
}

func (server Server) GetSuggestion(ctx context.Context, suggestion *pb.Suggestion) (*pb.SuggestionResponse, error) {
	return nil, nil
}
