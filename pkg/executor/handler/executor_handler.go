package handler

import (
	"context"
	"github.com/carmanzhang/ks-alert/pkg/executor/pb"
	"github.com/carmanzhang/ks-alert/pkg/executor/runtime"
)

type Executor struct{}

// Executor is used to implement ExecuteAlertConfig.
// ExecuteAlertConfig(context.Context, *AlertConfig) (*Error, error)
func (s *Executor) Execute(ctx context.Context, alertConfig *pb.Message) (*pb.Error, error) {
	//id := alertConfig.AlertConfigId
	return nil, runtime.Action(ctx, alertConfig)
}
