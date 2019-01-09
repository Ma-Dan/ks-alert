package registry

import (
	"fmt"
	"github.com/carmanzhang/ks-alert/pkg/utils/jsonutil"
	"testing"
)

func TestGetHardwareData(t *testing.T) {
	Convey("test get hardware data", t, func() {
		status := GetHardwareData()
		fmt.Println(status)
	})
}
func TestStatusArray_Sort(t *testing.T) {
	Convey("test status array sort", t, func() {
		statusArray := ServiceInfoArray{
			&ServiceInfo{
				SysStatus: &Status{
					NumberGoroutine: 9,
					CpuUtilization:  0.2,
				},
				ServiceAddress: "localhost:8080",
			}, &ServiceInfo{
				SysStatus: &Status{
					NumberGoroutine: 89,
					CpuUtilization:  0.03,
				},
				ServiceAddress: "localhost:8080",
			}, &ServiceInfo{
				SysStatus: &Status{
					NumberGoroutine: 29,
					CpuUtilization:  0.7,
				},
				ServiceAddress: "localhost:8080",
			},
		}

		statusArray.Sort(true)
		fmt.Println(jsonutil.Marshal(statusArray))

		statusArray.Sort(false)
		fmt.Println(jsonutil.Marshal(statusArray))

		arrays := statusArray.TopK(20)
		fmt.Println(jsonutil.Marshal(arrays))
	})
}
