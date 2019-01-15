package handler

import (
	"context"
	"github.com/carmanzhang/ks-alert/pkg/models"
	"github.com/carmanzhang/ks-alert/pkg/pb"
	"time"
)

type SilenceHandler struct{}

// silence
func (h SilenceHandler) CreateSilence(ctx context.Context, silence *pb.Silence) (*pb.SilenceResponse, error) {
	err := models.CreateSendPolicy(ConvertSilence2SendPolicy(silence))
	return getSilenceResponse(nil, err), nil
}

func getSilenceResponse(severity *models.SendPolicy, err error) *pb.SilenceResponse {
	arg := ConvertSendPolicy2Silence(severity)
	var respon = pb.SilenceResponse{Silence: arg}
	respon.Error = ErrorWrapper(err)

	return &respon
}

func (h SilenceHandler) DeleteSilence(ctx context.Context, silence *pb.Silence) (*pb.SilenceResponse, error) {
	sp := ConvertSilence2SendPolicy(silence)
	sp.SilenceStartAt = time.Now()
	sp.SilenceEndAt = time.Now()
	err := models.CreateOrUpdateSendPolicy(sp)
	return getSilenceResponse(nil, err), nil
}

func (h SilenceHandler) UpdateSilence(ctx context.Context, silence *pb.Silence) (*pb.SilenceResponse, error) {
	sp := ConvertSilence2SendPolicy(silence)
	err := models.UpdateSendPolicySilenceRule(sp)
	return getSilenceResponse(nil, err), nil
}

func (h SilenceHandler) GetSilence(ctx context.Context, silence *pb.Silence) (*pb.SilenceResponse, error) {
	sp := ConvertSilence2SendPolicy(silence)
	sendPolicy, err := models.GetSendPolicy(sp)
	return getSilenceResponse(sendPolicy, err), nil
}

func ConvertSilence2SendPolicy(sl *pb.Silence) *models.SendPolicy {
	if sl == nil {
		return nil
	}

	silenceStartAt := time.Unix(sl.StartTimestamp, 0)
	d := time.Second * time.Duration(int64(sl.Dutation))
	silenceEndAt := silenceStartAt.Add(d)
	return &models.SendPolicy{
		AlertRuleID:    sl.AlertRuleId,
		ResourceID:     sl.ResourceId,
		SilenceStartAt: silenceStartAt,
		SilenceEndAt:   silenceEndAt,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}
}

func ConvertSendPolicy2Silence(sp *models.SendPolicy) *pb.Silence {
	if sp == nil {
		return nil
	}
	return &pb.Silence{
		AlertRuleId:    sp.AlertRuleID,
		ResourceId:     sp.ResourceID,
		StartTimestamp: sp.SilenceStartAt.Unix(),
		Dutation:       int32(sp.SilenceEndAt.Unix() - sp.SilenceStartAt.Unix()),
	}
}
