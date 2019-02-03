package handler

import (
	"context"
	"kubesphere.io/ks-alert/pkg/pb"
)

type AlertHistoryHandler struct{}

func (h AlertHistoryHandler) GetAlertHistory(ctx context.Context, alertHistory *pb.AlertHistoryRequest) (*pb.AlertHistoryResponse, error) {
	return nil, nil
}
