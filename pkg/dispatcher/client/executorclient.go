package client

import (
	. "github.com/carmanzhang/ks-alert/pkg/stderr"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"time"
)

const GrpcConnectionTimeout = 10

//var clientLBConn *grpc.ClientConn

// cached grpc client map, key is hostname, value is corresponding grpc connection
var cachedGrpcConn = make(map[string]*grpc.ClientConn)

//func GetExecutorGrpcLoadBalancerConn(address ...string) (*grpc.ClientConn, error) {
//	var svc string
//	var etcd string
//	if address != nil && len(address) > 0 {
//		svc = address[0]
//		if len(address) > 1 {
//			etcd = address[1]
//		}
//	} else {
//		svc = *option.ExecutorServiceName
//		etcd = *option.EtcdAddress
//	}
//
//	if clientLBConn != nil {
//		return clientLBConn, nil
//	}
//
//	r := registry.NewResolver(svc)
//	b := grpc.RoundRobin(r)
//	ctx, _ := context.WithTimeout(context.Background(), GrpcConnectionTimeout*time.Second)
//	var err error
//	clientLBConn, err = grpc.DialContext(ctx, etcd, grpc.WithInsecure(), grpc.WithBalancer(b))
//	return clientLBConn, err
//}

func GetExecutorGrpcConn(svcAddress string) (*grpc.ClientConn, error) {

	if conn, ok := cachedGrpcConn[svcAddress]; ok {
		return conn, nil
	}

	ctx, _ := context.WithTimeout(context.Background(), GrpcConnectionTimeout*time.Second)
	conn, err := grpc.DialContext(ctx, svcAddress, grpc.WithInsecure())

	if err != nil {
		return nil, Error{Text: err.Error(), Code: GrpcError, Where: Caller(1, true)}
	}

	cachedGrpcConn[svcAddress] = conn

	return conn, nil
}
