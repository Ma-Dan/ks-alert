package client

import (
	"context"
	"fmt"
	"github.com/carmanzhang/ks-alert/pkg/executor/pb"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
	"time"
)

func TestGetExecutorGrpcClient(t *testing.T) {
	Convey("test get executor grpc client", t, func() {
		conn, err := GetExecutorGrpcLoadBalancerClient("alert_executor_service", "http://127.0.0.1:2379")
		if err != nil {
			panic(err)
		}
		// sleep a few millsecond for grpc dial etcd
		time.Sleep(time.Millisecond * 500)
		clientX := pb.NewExecutorClient(conn)

		resp, err := clientX.Execute(context.Background(), &pb.Message{Signal: pb.Message_RELOAD, AlertConfigId: "alert-config-v2y8w603r0j9kk"})
		fmt.Println(resp, err)
	})
}
