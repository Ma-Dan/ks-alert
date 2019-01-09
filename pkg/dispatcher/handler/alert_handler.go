package handler

import (
	"context"
	"fmt"
	"github.com/carmanzhang/ks-alert/pkg/dispatcher/client"
	"github.com/carmanzhang/ks-alert/pkg/dispatcher/pb"
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

	conn, err := client.GetExecutorGrpcLoadBalancerClient()
	cli := p.NewExecutorClient(conn)
	// TODO need add error adaptor
	pbErr, _ := cli.Execute(ctx, &p.Message{AlertConfigId: respon.AlertConfig.AlertConfigId, Signal: p.Message_CREATE})
	if pbErr != nil {
		fmt.Println(pbErr.Code, pbErr.Text)
	}
	return respon, nil
}

func getAlertConfigResponse(v interface{}, err error) *pb.AlertConfigResponse {

	var respon = &pb.AlertConfigResponse{}

	if v != nil {
		ac := v.(*models.AlertConfig)
		respon.AlertConfig = ConvertAlertConfig2PB(ac)
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
	v, err := DoTransactionAction(&ac, AlertConfig, MethodGet)
	respon := getAlertConfigResponse(v, err)
	return respon, nil
}

func ConvertPB2AlertConfig(pbac *pb.AlertConfig) *models.AlertConfig {
	if pbac == nil {
		return nil
	}

	ac := &models.AlertConfig{
		AlertConfigID:   pbac.AlertConfigId,
		AlertConfigName: pbac.AlertConfigName,

		ReceiverGroup:  ConvertPB2ReceiverGroup(pbac.ReceiverGroup),
		ResourceGroup:  ConvertPB2ResourceGroup(pbac.ResourceGroup),
		AlertRuleGroup: ConvertPB2AlertRuleGroup(pbac.AlertRuleGroup),

		SeverityCh: pbac.SeverityCh,
		SeverityID: pbac.SeverityId,

		EnableStart: ConvertString2Time(pbac.EnableStart),
		Description: pbac.Desc,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	if pbac.EnableEnd != "" {
		ac.EnableEnd = ConvertString2Time(pbac.EnableEnd)
	} else {
		// future 2100 year's timestamp
		ac.EnableEnd = time.Unix(4102416000, 0)
	}
	return ac
}

func ConvertAlertConfig2PB(ac *models.AlertConfig) *pb.AlertConfig {
	var pbac = pb.AlertConfig{}

	if ac != nil {
		pbac.AlertConfigId = ac.AlertConfigID
		pbac.AlertConfigName = ac.AlertConfigName
		pbac.SeverityId = ac.SeverityID
		pbac.SeverityCh = ac.SeverityCh
		pbac.EnableStart = ConvertTime2String(ac.EnableStart)
		pbac.EnableEnd = ConvertTime2String(ac.EnableEnd)
		pbac.Desc = ac.Description
		pbac.AlertRuleGroup = ConvertAlertRuleGroup2PB(ac.AlertRuleGroup)
		pbac.ReceiverGroup = ConvertReceiverGroup2PB(ac.ReceiverGroup)
		pbac.ResourceGroup = ConvertResourceGroup2PB(ac.ResourceGroup)
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
