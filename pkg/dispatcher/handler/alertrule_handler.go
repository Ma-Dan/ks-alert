package handler

import (
	"context"
	"github.com/carmanzhang/ks-alert/pkg/dispatcher/pb"
)

type AlertRuleHandler struct {}
// alert rule
func (server AlertRuleHandler) CreateAlertRule(ctx context.Context, alertRule *pb.AlertRule) (*pb.AlertRuleResponse, error) {
	return nil, nil
}

func (server AlertRuleHandler) DeleteAlertRule(ctx context.Context, alertRuleID *pb.AlertRuleID) (*pb.AlertRuleResponse, error) {
	return nil, nil
}

func (server AlertRuleHandler) UpdateAlertRule(ctx context.Context, alertRule *pb.AlertRule) (*pb.AlertRuleResponse, error) {
	return nil, nil
}

func (server AlertRuleHandler) GetAlertRule(ctx context.Context, alertRuleID *pb.AlertRuleID) (*pb.AlertRuleResponse, error) {
	return nil, nil
}
