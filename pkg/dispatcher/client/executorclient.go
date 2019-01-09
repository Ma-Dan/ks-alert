package client

import (
	"github.com/carmanzhang/ks-alert/pkg/option"
	"github.com/carmanzhang/ks-alert/pkg/registry"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"time"
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

func GetExecutorGrpcLoadBalancerClient(address ...string) (*grpc.ClientConn, error) {
	var svc string
	var etcd string
	if address != nil && len(address) > 0 {
		svc = address[0]
		if len(address) > 1 {
			etcd = address[1]
		}
	} else {
		svc = *option.ExecutorServiceName
		etcd = *option.EtcdAddress
	}

	if clientLBConn != nil {
		return clientLBConn, nil
	}

	r := registry.NewResolver(svc)
	b := grpc.RoundRobin(r)
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	var err error
	clientLBConn, err = grpc.DialContext(ctx, etcd, grpc.WithInsecure(), grpc.WithBalancer(b))
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
