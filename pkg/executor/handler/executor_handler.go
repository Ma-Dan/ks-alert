package handler

import (
	"context"
	"kubesphere.io/ks-alert/pkg/executor/runtime"
	"kubesphere.io/ks-alert/pkg/pb"
	"kubesphere.io/ks-alert/pkg/stderr"
)

type Executor struct{}

func (s *Executor) Execute(ctx context.Context, informer *pb.Informer) (*pb.Error, error) {
	err := runtime.ExecuteAlertConfig(ctx, informer)
	return stderr.ErrorWrapper(err), nil
}
