package handler

import (
	"context"
	"github.com/carmanzhang/ks-alert/pkg/models"
	"github.com/carmanzhang/ks-alert/pkg/pb"
	"github.com/pkg/errors"
	"time"
)

type SilenceHandler struct{}

// silence
func (server SilenceHandler) CreateSilence(ctx context.Context, silence *pb.Silence) (*pb.SilenceResponse, error) {
	err := models.CreateSendPolicy(ConvertSilence2SendPolicy(silence))
	var respon = pb.SilenceResponse{}
	if err != nil {
		respon.Error = ErrorWrapper(err)
	} else {
		respon.Error = ErrorWrapper(errors.New("success"))
	}
	return &respon, nil
}

func (server SilenceHandler) DeleteSilence(ctx context.Context, silence *pb.Silence) (*pb.SilenceResponse, error) {
	sp := ConvertSilence2SendPolicy(silence)
	sp.SilenceStartAt = time.Now()
	sp.SilenceEndAt = time.Now()
	err := models.CreateOrUpdateSendPolicy(sp)
	var respon = pb.SilenceResponse{}
	if err != nil {
		respon.Error = ErrorWrapper(err)
	} else {
		respon.Error = ErrorWrapper(errors.New("success"))
	}
	return &respon, nil
}

func (server SilenceHandler) UpdateSilence(ctx context.Context, silence *pb.Silence) (*pb.SilenceResponse, error) {
	sp := ConvertSilence2SendPolicy(silence)
	err := models.UpdateSendPolicySilenceRule(sp)
	var respon = pb.SilenceResponse{}
	if err != nil {
		respon.Error = ErrorWrapper(err)
	} else {
		respon.Error = ErrorWrapper(errors.New("success"))
	}
	return &respon, nil
}

func (server SilenceHandler) GetSilence(ctx context.Context, silence *pb.Silence) (*pb.SilenceResponse, error) {
	sp := ConvertSilence2SendPolicy(silence)
	sendPolicy, err := models.GetSendPolicy(sp)
	silence = ConvertSendPolicy2Silence(sendPolicy)
	var respon = pb.SilenceResponse{}
	if err != nil {
		respon.Error = ErrorWrapper(err)
	} else {
		respon.Silence = silence
		respon.Error = ErrorWrapper(errors.New("success"))
	}
	return &respon, nil
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
