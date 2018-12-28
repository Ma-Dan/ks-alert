package handler

import (
	"context"
	"github.com/carmanzhang/ks-alert/pkg/dispatcher/pb"
)

type AlertRuleHandler struct {}
// alert rule
func (server AlertRuleHandler) CreateAlertRule(ctx context.Context, alertRule *pb.AlertRuleGroup) (*pb.AlertRuleGroupResponse, error) {
	return nil, nil
}

func (server AlertRuleHandler) DeleteAlertRule(ctx context.Context, alertRuleID *pb.AlertRuleGroupSpec) (*pb.AlertRuleGroupResponse, error) {
	return nil, nil
}

func (server AlertRuleHandler) UpdateAlertRule(ctx context.Context, alertRule *pb.AlertRuleGroup) (*pb.AlertRuleGroupResponse, error) {
	return nil, nil
}

func (server AlertRuleHandler) GetAlertRule(ctx context.Context, alertRuleID *pb.AlertRuleGroupSpec) (*pb.AlertRuleGroupResponse, error) {
	return nil, nil
}
