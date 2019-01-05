package handler

import (
	"context"
	"fmt"
	"github.com/carmanzhang/ks-alert/pkg/dispatcher/pb"
	"github.com/carmanzhang/ks-alert/pkg/executor/handler"
	p "github.com/carmanzhang/ks-alert/pkg/executor/pb"
	"github.com/carmanzhang/ks-alert/pkg/models"
	"github.com/golang/glog"
	"strconv"
	"time"
)

// alert
type AlertHandler struct{}

func (server AlertHandler) CreateAlertConfig(ctx context.Context, pbac *pb.AlertConfig) (*pb.AlertConfigResponse, error) {

	ac := ConvertPB2AlertConfig(pbac)

	v, err := DoTransactionAction(ac, AlertConfig, MethodCreate)
	respon := getAlertConfigResponse(v, err)
	executor := handler.Executor{}
	// TODO need add error adaptor
	executor.Execute(ctx, &p.Message{AlertConfigId: respon.AlertConfig.AlertConfigId, Signal: p.Message_CREATE})

	return respon, nil
}

func getAlertConfigResponse(v interface{}, err error) *pb.AlertConfigResponse {

	var respon = &pb.AlertConfigResponse{}

	if v != nil {
		vv := v.([]interface{})
		respon.AlertConfig = ConvertAlertConfig2PB(vv)
	}

	if err != nil {
		glog.Errorln(err.Error())
		respon.Error = ErrorConverter(err)
	} else {
		respon.Error = ErrorConverter(*models.NewError(0, models.Success))
	}

	return respon
}

func (server AlertHandler) DeleteAlertConfig(ctx context.Context, alertConfigSpec *pb.AlertConfigSpec) (*pb.AlertConfigResponse, error) {
	ac := models.AlertConfig{AlertConfigID: alertConfigSpec.AlertConfigId}

	acIDs, _ := DoTransactionAction(&ac, AlertConfig, MethodDelete, false)
	alertConfig := acIDs.([]interface{})[0].(*models.AlertConfig)

	ac.AlertRuleGroupID = alertConfig.AlertRuleGroupID
	ac.ResourceGroupID = alertConfig.ResourceGroupID
	ac.ReceiverGroupID = alertConfig.ReceiverGroupID

	v, err := DoTransactionAction(&ac, AlertConfig, MethodDelete)
	respon := getAlertConfigResponse(v, err)
	return respon, nil
}

func (server AlertHandler) UpdateAlertConfig(ctx context.Context, alertConfig *pb.AlertConfig) (*pb.AlertConfigResponse, error) {
	ac := ConvertPB2AlertConfig(alertConfig)

	v, err := DoTransactionAction(ac, AlertConfig, MethodUpdate)
	respon := getAlertConfigResponse(v, err)
	return respon, nil
}

func (server AlertHandler) GetAlertConfig(ctx context.Context, alertConfigSpec *pb.AlertConfigSpec) (*pb.AlertConfigResponse, error) {
	ac := models.AlertConfig{AlertConfigID: alertConfigSpec.AlertConfigId}

	acIDs, _ := DoTransactionAction(&ac, AlertConfig, MethodGet, false)
	alertConfig := acIDs.([]interface{})[0].(*models.AlertConfig)

	ac.AlertRuleGroup = &models.AlertRuleGroup{AlertRuleGroupID: alertConfig.AlertRuleGroupID}
	ac.ResourceGroup = &models.ResourceGroup{ResourceGroupID: alertConfig.ResourceGroupID}
	ac.ReceiverGroup = &models.ReceiverGroup{ReceiverGroupID: alertConfig.ReceiverGroupID}

	v, err := DoTransactionAction(&ac, AlertConfig, MethodGet)
	respon := getAlertConfigResponse(v, err)
	return respon, nil
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
