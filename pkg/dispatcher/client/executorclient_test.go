package client

import (
	"context"
	"fmt"
	"github.com/carmanzhang/ks-alert/pkg/pb"
	"testing"
	"time"
)

func TestGetExecutorGrpcClient(t *testing.T) {
	Convey("test get executor grpc client", t, func() {
		conn, err := GetExecutorGrpcConn("127.0.0.1:50001")
		if err != nil {
			panic(err)
		}
		// sleep a few millsecond for grpc dial etcd
		time.Sleep(time.Millisecond * 500)
		clientX := pb.NewExecutorClient(conn)

		resp, err := clientX.Execute(context.Background(), &pb.Informer{Signal: pb.Informer_RELOAD, AlertConfigId: "alert-config-3jvxm4343roxkk"})
		//resp, err := clientX.Execute(context.Background(), &pb.Informer{Signal: pb.Informer_TERMINATE, AlertConfigId: "alert-config-0wjmywop7pwn8v"})
		fmt.Println(resp, err)
	})
}
