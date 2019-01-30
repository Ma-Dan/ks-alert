package handler

import (
	"context"
	"github.com/carmanzhang/ks-alert/pkg/executor/runtime"
	"github.com/carmanzhang/ks-alert/pkg/pb"
	"github.com/carmanzhang/ks-alert/pkg/stderr"
)

type Executor struct{}

func (s *Executor) Execute(ctx context.Context, informer *pb.Informer) (*pb.Error, error) {
	err := runtime.ExecuteAlertConfig(ctx, informer)
	return stderr.ErrorWrapper(err), nil
}
