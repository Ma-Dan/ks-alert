package client

import (
	"time"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"kubesphere.io/ks-alert/pkg/registry"
)

var clientConn *grpc.ClientConn

func GetExecutorGrpcClient(svc string, etcdAddress string) (*grpc.ClientConn, error) {
	if clientConn != nil {
		return clientConn, nil
	}

	r := registry.NewResolver(svc)
	b := grpc.RoundRobin(r)
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	var err error
	clientConn, err =  grpc.DialContext(ctx, etcdAddress, grpc.WithInsecure(), grpc.WithBalancer(b))
	return clientConn, err
}
