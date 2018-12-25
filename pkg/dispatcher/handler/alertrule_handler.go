package handler

import (
	"context"
	"kubesphere.io/ks-alert/pkg/dispatcher/pb"
)

// alert rule
func (server Server) CreateAlertRule(ctx context.Context, alertRule *pb.AlertRule) (*pb.AlertRuleResponse, error) {
	return nil, nil
}

func (server Server) DeleteAlertRule(ctx context.Context, alertRuleID *pb.AlertRuleID) (*pb.AlertRuleResponse, error) {
	return nil, nil
}

func (server Server) UpdateAlertRule(ctx context.Context, alertRule *pb.AlertRule) (*pb.AlertRuleResponse, error) {
	return nil, nil
}

func (server Server) GetAlertRule(ctx context.Context, alertRuleID *pb.AlertRuleID) (*pb.AlertRuleResponse, error) {
	return nil, nil
}
