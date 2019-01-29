package handler

import (
	"context"
	"fmt"
	"github.com/carmanzhang/ks-alert/pkg/dispatcher/client"
	"github.com/carmanzhang/ks-alert/pkg/models"
	"github.com/carmanzhang/ks-alert/pkg/option"
	"github.com/carmanzhang/ks-alert/pkg/pb"
	"github.com/carmanzhang/ks-alert/pkg/registry"
	"github.com/golang/glog"
	"strconv"
	"strings"
	"time"
)

// alert
type AlertHandler struct{}

func (h AlertHandler) CreateAlertConfig(ctx context.Context, pbac *pb.AlertConfig) (*pb.AlertConfigResponse, error) {

	ac := ConvertPB2AlertConfig(pbac)
	// the host(node) which this alert congfig whill be executed
	svcAddr, err := registry.GetIdleExecutorAddress()
	if err != nil {
		respon := getAlertConfigResponse(nil)
		respon.Error = ErrorWrapper(err)
		return respon, nil
	}

	ac.HostID = svcAddr

	v, err := DoTransactionAction(ac, AlertConfig, MethodCreate)
	respon := getAlertConfigResponse(v)
	respon.Error = ErrorWrapper(err)
	if err != nil {
		return respon, nil
	}

	//option.HostInfo
	pbErr := ExecuteAlertConfig(svcAddr, respon.AlertConfig.AlertConfigId, pb.Informer_CREATE)
	respon.Error = pbErr
	return respon, nil
}

func ExecuteAlertConfig(svcAddr string, acID string, sig pb.Informer_Signal) *pb.Error {
	var err error
	conn, err := client.GetExecutorGrpcConn(svcAddr)
	if err != nil {
		return ErrorWrapper(err)
	}

	cli := pb.NewExecutorClient(conn)
	pbErr, err := cli.Execute(context.Background(), &pb.Informer{AlertConfigId: acID, Signal: sig})
	// err adaptor
	if pbErr != nil {
		return ErrorWrapper(err)
	}

	if err != nil {
		return ErrorWrapper(err)
	}

	return nil
}

func getAlertConfigResponse(v interface{}) *pb.AlertConfigResponse {

	var respon = &pb.AlertConfigResponse{}

	if v != nil {
		ac := v.(*models.AlertConfig)
		respon.AlertConfig = ConvertAlertConfig2PB(ac)
	}
	return respon
}

func (h AlertHandler) DeleteAlertConfig(ctx context.Context, alertConfigSpec *pb.AlertConfigSpec) (*pb.AlertConfigResponse, error) {
	acID := alertConfigSpec.AlertConfigId
	hostInfo, err := models.GetAlertConfigBindingHost(acID)
	if err != nil {
		respon := getAlertConfigResponse(nil)
		respon.Error = ErrorWrapper(err)
		return respon, nil
	}

	pbErr := ExecuteAlertConfig(hostInfo, acID, pb.Informer_TERMINATE)

	if pbErr != nil {
		respon := getAlertConfigResponse(nil)
		respon.Error = ErrorWrapper(err)
		return respon, nil
	}

	v, err := DoTransactionAction(&models.AlertConfig{AlertConfigID: acID}, AlertConfig, MethodDelete)
	respon := getAlertConfigResponse(v)
	respon.Error = ErrorWrapper(err)
	return respon, nil
}

func (h AlertHandler) UpdateAlertConfig(ctx context.Context, alertConfig *pb.AlertConfig) (*pb.AlertConfigResponse, error) {
	ac := ConvertPB2AlertConfig(alertConfig)

	v, err := DoTransactionAction(ac, AlertConfig, MethodUpdate)
	respon := getAlertConfigResponse(v)
	respon.Error = ErrorWrapper(err)
	if err != nil {
		return respon, nil
	}

	acID := alertConfig.AlertConfigId

	hostID, err := models.GetAlertConfigBindingHost(acID)
	if err != nil {
		respon := getAlertConfigResponse(nil)
		respon.Error = ErrorWrapper(err)
		return respon, nil
	}

	hostInfo := strings.Split(hostID, "-")
	svcAddress := fmt.Sprintf("%s:%d", hostInfo[1], *option.ExecutorServicePort)

	pbErr := ExecuteAlertConfig(svcAddress, acID, pb.Informer_RELOAD)

	if pbErr != nil {
		respon := getAlertConfigResponse(nil)
		respon.Error = ErrorWrapper(err)
		return respon, nil
	}

	return respon, nil
}

func (h AlertHandler) GetAlertConfig(ctx context.Context, alertConfigSpec *pb.AlertConfigSpec) (*pb.AlertConfigResponse, error) {
	ac := models.AlertConfig{AlertConfigID: alertConfigSpec.AlertConfigId}
	v, err := DoTransactionAction(&ac, AlertConfig, MethodGet)
	respon := getAlertConfigResponse(v)
	respon.Error = ErrorWrapper(err)
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
