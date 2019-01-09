package client

import (
	"fmt"
	"testing"
)

func TestSendMonitoringRequest(t *testing.T) {
	Convey("test sending requst", t, func() {
		Convey("test sending requst01", func() {
			resources := []string{"i-2346", "i-uyoiu", "i-hvfbn"}
			//metrics := []string{"cpu_utilization", "memory_utilization", "net_utilization"}
			res := SendMonitoringRequest("localhost:8081/cluster/local/nodes?", resources, "cpu_utilization", 14324535465, 12433243535, 6)
			fmt.Println(res)
		})

		Convey("test sending requst02", func() {
			resources := []string{"i-2346", "i-uyoiu", "i-hvfbn"}
			res := SendMonitoringRequest("localhost:8081/cluster/local/nodes?", resources, "cpu_utilization", 14324535465, 12433243535, 6)
			fmt.Println(res)
		})
	})
}
