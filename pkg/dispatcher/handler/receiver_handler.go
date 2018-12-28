package handler

import (
	"github.com/carmanzhang/ks-alert/pkg/dispatcher/pb"
	"context"
)

type ReceiverHandler struct {}
// alert rule
func (server ReceiverHandler) CreateReceiver(ctx context.Context, Receiver *pb.ReceiverGroup) (*pb.ReceiverGroupResponse, error) {
	return nil, nil
}

func (server ReceiverHandler) DeleteReceiver(ctx context.Context, ReceiverID *pb.ReceiverGroupSpec) (*pb.ReceiverGroupResponse, error) {
	return nil, nil
}

func (server ReceiverHandler) UpdateReceiver(ctx context.Context, Receiver *pb.ReceiverGroup) (*pb.ReceiverGroupResponse, error) {
	return nil, nil
}

func (server ReceiverHandler) GetReceiver(ctx context.Context, ReceiverID *pb.ReceiverGroupSpec) (*pb.ReceiverGroupResponse, error) {
	return nil, nil
}

