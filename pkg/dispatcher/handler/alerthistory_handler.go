package handler

import (
	"context"
	"github.com/carmanzhang/ks-alert/pkg/pb"
)

type AlertHistoryHandler struct{}

func (h AlertHistoryHandler) GetAlertHistory(ctx context.Context, alertHistory *pb.AlertHistoryRequest) (*pb.AlertHistoryResponse, error) {
	return nil, nil
}
