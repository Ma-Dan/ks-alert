package client

import (
	. "github.com/carmanzhang/ks-alert/pkg/stderr"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"time"
)

const GrpcConnectionTimeout = 10

// cached grpc client map, key is hostname, value is corresponding grpc connection
var cachedGrpcConn = make(map[string]*grpc.ClientConn)

func GetExecutorGrpcConn(svcAddress string) (*grpc.ClientConn, error) {

	if conn, ok := cachedGrpcConn[svcAddress]; ok {
		return conn, nil
	}

	ctx, _ := context.WithTimeout(context.Background(), GrpcConnectionTimeout*time.Second)
	conn, err := grpc.DialContext(ctx, svcAddress, grpc.WithInsecure())

	if err != nil {
		return nil, Error{Text: err.Error(), Code: GrpcError, Where: Caller(0, true)}
	}

	cachedGrpcConn[svcAddress] = conn

	return conn, nil
}
