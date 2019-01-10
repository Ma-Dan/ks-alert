package registry

import (
	"context"
	"encoding/json"
	"github.com/carmanzhang/ks-alert/pkg/option"
	"github.com/carmanzhang/ks-alert/pkg/utils/etcdutil"
	"github.com/pkg/errors"
	etcd3 "go.etcd.io/etcd/clientv3"
	"k8s.io/klog/glog"
	"sort"
)

type StatusWrapper struct {
	serviceInfoArray ServiceInfoArray
	by               func(p, q *Status) bool
}

func (w StatusWrapper) Len() int {
	return len(w.serviceInfoArray)
}

func (w StatusWrapper) Less(i, j int) bool {
	s1 := w.serviceInfoArray[i].SysStatus
	s2 := w.serviceInfoArray[j].SysStatus
	return w.by(s1, s2)
}

func (w StatusWrapper) Swap(i, j int) {
	w.serviceInfoArray[i], w.serviceInfoArray[j] = w.serviceInfoArray[j], w.serviceInfoArray[i]
}

type ServiceInfoArray []*ServiceInfo

type ServiceInfo struct {
	// host:port
	ServiceAddress string
	SysStatus      *Status
}

func (s ServiceInfoArray) Sort(reverse bool) ServiceInfoArray {
	if reverse {
		sort.Sort(StatusWrapper{s, func(p, q *Status) bool {
			return p.CpuUtilization*(1.0+float32(p.NumberGoroutine)) > q.CpuUtilization*(1.0+float32(q.NumberGoroutine))
		}})
	} else {
		sort.Sort(StatusWrapper{s, func(p, q *Status) bool {
			return p.CpuUtilization*(1.0+float32(p.NumberGoroutine)) < q.CpuUtilization*(1.0+float32(q.NumberGoroutine))
		}})
	}
	return s
}

func (s ServiceInfoArray) TopK(k int) ServiceInfoArray {
	l := len(s)
	if k < 0 || k > l {
		return s
	}

	return s[:k]
}

// get all executor service info from etcd
func GetAllExecutorServiceInfo(resp *etcd3.GetResponse) ServiceInfoArray {
	var sia ServiceInfoArray

	for i := range resp.Kvs {
		kv := resp.Kvs[i]
		var si ServiceInfo
		err := json.Unmarshal(kv.Value, &si)
		if err != nil {
			glog.Errorln(err.Error())
		}

		sia = append(sia, &si)
	}

	return sia
}

// TODO need to change to watch mode, a global variable Map used for saving service addresses
func (s ServiceInfoArray) ExtractServiceAddrs() map[string]string {
	var svcAddrs = make(map[string]string)
	for i, _ := range s {
		addr := s[i].ServiceAddress
		if addr != "" {
			svcAddrs[addr] = ""
		}
	}
	return svcAddrs
}

func GetIdleExecutorAddress(endpoints ...string) (string, error) {
	if len(endpoints) == 0 {
		endpoints = []string{*option.EtcdAddress}
	}

	e, err := etcdutil.Connect(endpoints, "")
	if err != nil {
		return "", err
	}

	prefix := "/" + *option.ExecutorServiceName + "/"
	resp, err := e.Get(context.Background(), prefix, etcd3.WithPrefix())
	if err != nil {
		return "", err
	}

	ss := GetAllExecutorServiceInfo(resp)
	addr := ss.Sort(false).TopK(1).ExtractServiceAddrs()
	if len(addr) == 0 {
		return "", errors.New("no executor service address")
	}

	var value string
	for k := range addr {
		value = addr[k]
	}
	return value, nil
}
