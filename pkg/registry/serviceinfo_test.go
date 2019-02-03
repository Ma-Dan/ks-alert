package registry

import (
	"context"
	"fmt"
	"kubesphere.io/ks-alert/pkg/option"
	"kubesphere.io/ks-alert/pkg/utils/etcdutil"
	"kubesphere.io/ks-alert/pkg/utils/jsonutil"
	etcd3 "go.etcd.io/etcd/clientv3"
	"testing"
)

func TestGetHardwareData(t *testing.T) {
	Convey("test get hardware data", t, func() {
		status := GetHardwareData()
		fmt.Println(jsonutil.Marshal(status))
	})
}

func TestGetAllExecutorServiceInfo(t *testing.T) {
	Convey("test get hardware data", t, func() {
		endpoints := []string{"127.0.0.1:2379"}
		prefix := "/" + *option.ExecutorServiceName + "/"
		//prefix := "/alert_executor_service"
		e, _ := etcdutil.Connect(endpoints, "")
		resp, _ := e.Get(context.Background(), prefix, etcd3.WithPrefix())
		ss := GetAllExecutorServiceInfo(resp)
		addr := ss.Sort(false).TopK(2).ExtractServiceAddrs()
		fmt.Println(addr)
	})
}

func TestGetIdleExecutorAddress(t *testing.T) {
	Convey("TestGetIdleExecutorAddress", t, func() {
		s, e := GetIdleExecutorAddress()
		fmt.Println(s, e)
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
