package handler

import (
	"context"
	"fmt"
	"github.com/carmanzhang/ks-alert/pkg/dispatcher/pb"
	"github.com/carmanzhang/ks-alert/pkg/models"
	"github.com/golang/glog"
	"strconv"
	"time"
)

// alert
type AlertHandler struct{}

func (server AlertHandler) CreateAlertConfig(ctx context.Context, pbac *pb.AlertConfig) (*pb.AlertConfigResponse, error) {

	ac := ConvertPB2AlertConfig(pbac)

	res, err := DoTransactionAction(ac, AlertConfig, MethodCreate)

	if err != nil {
		return nil, err
	}

	v := res.([]interface{})

	alertConfig := ConvertAlertConfig2PB(v)

	return &pb.AlertConfigResponse{
		AlertConfig: alertConfig,
	}, nil
}

func ConvertPB2AlertConfig(pbac *pb.AlertConfig) *models.AlertConfig {
	if pbac == nil {
		return nil
	}
	return &models.AlertConfig{
		AlertConfigID:   pbac.AlertConfigId,
		AlertConfigName: pbac.AlertConfigName,

		ReceiverGroup:  ConvertPB2ReceiverGroup(pbac.ReceiverGroup),
		ResourceGroup:  ConvertPB2ResourceGroup(pbac.ResourceGroup),
		AlertRuleGroup: ConvertPB2AlertRuleGroup(pbac.AlertRuleGroup),

		SeverityCh: pbac.SeverityCh,
		SeverityID: pbac.SeverityId,

		EnableStart: ConvertString2Time(pbac.EnableStart),
		EnableEnd:   ConvertString2Time(pbac.EnableEnd),
		Description: pbac.Desc,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
}

func (server AlertHandler) DeleteAlertConfig(ctx context.Context, alertConfigSpec *pb.AlertConfigSpec) (*pb.AlertConfigResponse, error) {
	return nil, nil
}

func (server AlertHandler) UpdateAlertConfig(ctx context.Context, alertConfig *pb.AlertConfig) (*pb.AlertConfigResponse, error) {
	return nil, nil
}

func (server AlertHandler) GetAlertConfig(ctx context.Context, alertConfig *pb.AlertConfigSpec) (*pb.AlertConfigResponse, error) {
	return nil, nil
}

func ConvertAlertConfig2PB(v []interface{}) *pb.AlertConfig {
	// alertConfig *models.AlertConfig, ruleGroup *models.AlertRuleGroup, reveiverGroup *models.ReceiverGroup, resourceGroup *models.ResourceGroup

	var pbac = pb.AlertConfig{}

	if v[0] != nil {
		alertConfig := v[0].(*models.AlertConfig)
		pbac.AlertConfigId = alertConfig.AlertConfigID
		pbac.AlertConfigName = alertConfig.AlertConfigName
		pbac.SeverityId = alertConfig.SeverityID
		pbac.SeverityCh = alertConfig.SeverityCh
		pbac.EnableStart = ConvertTime2String(alertConfig.EnableStart)
		pbac.EnableEnd = ConvertTime2String(alertConfig.EnableEnd)
		pbac.Desc = alertConfig.Description
	}

	if v[1] != nil {
		ruleGroup := v[1].(*models.AlertRuleGroup)
		pbac.AlertRuleGroup = ConvertAlertRuleGroup2PB(ruleGroup)
	}

	if v[2] != nil {
		reveiverGroup := v[2].(*models.ReceiverGroup)
		pbac.ReceiverGroup = ConvertReceiverGroup2PB(reveiverGroup)
	}

	if v[3] != nil {
		resourceGroup := v[3].(*models.ResourceGroup)
		pbac.ResourceGroup = ConvertResourceGroup2PB(resourceGroup)
	}

	return &pbac
}

func ConvertString2Time(ts string) time.Time {
	timeInt, err := strconv.ParseInt(ts, 10, 64)
	if err != nil {
		glog.Errorf("convert second timestamp %s to minute timestamp failed", ts)
		return time.Now()
	}

	return time.Unix(timeInt, 0)
}

func ConvertTime2String(t time.Time) string {
	return fmt.Sprintf("%d", t.Unix())
}
