package client

import (
	"context"
	. "kubesphere.io/ks-alert/pkg/stderr"
	"google.golang.org/grpc"
	"time"
)

var nfClient *grpc.ClientConn

func GetNotificationConn(svcAddress string) (*grpc.ClientConn, error) {

	if nfClient != nil {
		return nfClient, nil
	}

	ctx, _ := context.WithTimeout(context.Background(), GrpcConnectionTimeout*time.Second)
	var err error

	nfClient, err = grpc.DialContext(ctx, svcAddress, grpc.WithInsecure())

	if err != nil {
		return nil, Error{Text: err.Error(), Code: GrpcError, Where: Caller(0, true)}
	}

	return nfClient, nil
}
