package client

import (
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestSendMonitoringRequest(t *testing.T) {
	Convey("test sending requst", t, func() {
		Convey("test sending requst01", func() {
			resources := []string{"i-2346", "i-uyoiu", "i-hvfbn"}
			metrics := []string{"cpu_utilization", "memory_utilization", "net_utilization"}
			res := SendMonitoringRequest("localhost:8081/cluster/local/nodes?", resources, metrics)
			fmt.Println(res)
		})
	})
}
