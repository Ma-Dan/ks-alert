package handler

import (
	"context"
	"github.com/carmanzhang/ks-alert/pkg/pb"
)

type TestGrpc struct{}

func (h TestGrpc) Ping(ctx context.Context, req *pb.Empty) (*pb.Content, error) {
	return &pb.Content{Text: "hello"}, nil
}

func (h TestGrpc) Reverse(ctx context.Context, req *pb.Content) (*pb.Content, error) {
	txt := req.Text
	resp := "hello, " + txt
	return &pb.Content{Text: resp}, nil
}
