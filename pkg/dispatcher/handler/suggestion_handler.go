package handler

import (
	"context"
	"kubesphere.io/ks-alert/pkg/models"
	"kubesphere.io/ks-alert/pkg/pb"
	"kubesphere.io/ks-alert/pkg/stderr"
	"kubesphere.io/ks-alert/pkg/utils/jsonutil"
	"time"
)

type SuggestionHandler struct{}

// can only execute updating and getting operation
// for creating and deleting,

func (server SuggestionHandler) UpdateSuggestion(ctx context.Context, suggestion *pb.Suggestion) (*pb.SuggestionResponse, error) {
	if suggestion.AlertConfigId == "" || suggestion.AlertRuleId == "" || suggestion.ResourceId == "" {
		return getSuggestionResponse(nil, stderr.Error{
			Code: stderr.InvalidParam,
			Text: "alert config id and alert rule id and resource id must be specified",
		}), nil
	}

	sug, err := models.UpdateSuggestion(ConvertPB2Suggestion(suggestion))

	return getSuggestionResponse(sug, err), nil
}

func getSuggestionResponse(suggestion *models.Suggestion, err error) *pb.SuggestionResponse {
	arg := ConvertSuggestion2PB(suggestion)
	var respon = pb.SuggestionResponse{Suggestion: arg}
	respon.Error = stderr.ErrorWrapper(err)

	return &respon
}

func (server SuggestionHandler) GetSuggestion(ctx context.Context, suggestion *pb.Suggestion) (*pb.SuggestionResponse, error) {
	if suggestion.AlertConfigId == "" || suggestion.AlertRuleId == "" || suggestion.ResourceId == "" {
		return getSuggestionResponse(nil, stderr.Error{
			Code: stderr.InvalidParam,
			Text: "alert config id and alert rule id and resource id must be specified",
		}), nil
	}

	sug, err := models.GetSuggestion(ConvertPB2Suggestion(suggestion))
	return getSuggestionResponse(sug, err), nil
}

func ConvertSuggestion2PB(s *models.Suggestion) *pb.Suggestion {
	if s == nil {
		return nil
	}

	var message []*pb.Message

	jsonutil.Unmarshal(s.Message, &message)

	return &pb.Suggestion{
		AlertConfigId: s.AlertRuleID,
		AlertRuleId:   s.AlertRuleID,
		ResourceId:    s.ResourceID,
		Messages:      message,
	}
}

func ConvertPB2Suggestion(s *pb.Suggestion) *models.Suggestion {

	if s == nil {
		return nil
	}

	return &models.Suggestion{
		AlertConfigID: s.AlertConfigId,
		AlertRuleID:   s.AlertRuleId,
		ResourceID:    s.ResourceId,
		Message:       jsonutil.Marshal(s.Messages),
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}
}
