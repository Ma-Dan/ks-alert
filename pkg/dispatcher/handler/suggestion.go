package handler

import (
	"context"
	"github.com/carmanzhang/ks-alert/pkg/dispatcher/pb"
)

type SuggestionHandler struct{}

// suggestion
func (server SuggestionHandler) CreateSuggestion(ctx context.Context, suggestion *pb.Suggestion) (*pb.SuggestionResponse, error) {
	return nil, nil
}

func (server SuggestionHandler) DeleteSuggestion(ctx context.Context, suggestion *pb.Suggestion) (*pb.SuggestionResponse, error) {
	return nil, nil
}

func (server SuggestionHandler) UpdateSuggestion(ctx context.Context, suggestion *pb.Suggestion) (*pb.SuggestionResponse, error) {
	return nil, nil
}

func (server SuggestionHandler) GetSuggestion(ctx context.Context, suggestion *pb.Suggestion) (*pb.SuggestionResponse, error) {
	return nil, nil
}
