package client

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
	"time"
	"fmt"
	"context"
	"github.com/carmanzhang/ks-alert/pkg/executor/pb"
	"github.com/carmanzhang/ks-alert/pkg/models"
)

func TestGetExecutorGrpcClient(t *testing.T) {
	Convey("test get executor grpc client", t, func() {
		conn, err := GetExecutorGrpcLoadBalancerClient("alert_executor_service", "http://127.0.0.1:2379")
		if err != nil {
			panic(err)
		}
		ticker := time.NewTicker(1 * time.Second)
		for t := range ticker.C {
			clientX := pb.NewExecutorClient(conn)
			resp, err := clientX.ExecuteAlertConfig(context.Background(), &pb.AlertConfig{Signal: pb.AlertConfig_Signal(models.Create), AlertConfigId: "world"})

			if err == nil {
				fmt.Printf("%v: Reply is %s\n", t, resp.Text)
			}
		}
	})
}
