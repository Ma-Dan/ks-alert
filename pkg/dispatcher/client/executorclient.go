package client

import (
	"time"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"github.com/carmanzhang/ks-alert/pkg/registry"
)

var clientLBConn *grpc.ClientConn


//type ClientGrpcClient struct {
//	clientLBConn *grpc.ClientConn
//	host         string
//	ipAddr       string
//}
//
//var cachedGrpcClient = make(map[string]*ClientGrpcClient)

// cached grpc client map, key is hostname, value is corresponding grpc connection
var cachedGrpcClient = make(map[string]*grpc.ClientConn)

//
func GetExecutorGrpcLoadBalancerClient(svc string, etcdAddress string) (*grpc.ClientConn, error) {
	if clientLBConn != nil {
		return clientLBConn, nil
	}

	r := registry.NewResolver(svc)
	b := grpc.RoundRobin(r)
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	var err error
	clientLBConn, err = grpc.DialContext(ctx, etcdAddress, grpc.WithInsecure(), grpc.WithBalancer(b))
	return clientLBConn, err
}

func GetExecutorGrpcClient(ipAddr string) (*grpc.ClientConn, error) {

	if grpcClient, ok := cachedGrpcClient[ipAddr]; ok {
		if grpcClient != nil {
			return grpcClient, nil
		}
	}

	return nil, nil
	//r := registry.NewResolver(svc)
	//b := grpc.RoundRobin(r)
	//ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	//var err error
	//clientConn, err = grpc.DialContext(ctx, etcdAddress, grpc.WithInsecure(), grpc.WithBalancer(b))
	//return clientConn, err
}
