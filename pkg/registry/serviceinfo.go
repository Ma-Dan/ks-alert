package registry

import "sort"

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

func (s ServiceInfoArray) Sort(reverse bool) {
	if reverse {
		sort.Sort(StatusWrapper{s, func(p, q *Status) bool {
			return p.CpuUtilization*(1.0+float32(p.NumberGoroutine)) > q.CpuUtilization*(1.0+float32(q.NumberGoroutine))
		}})
	} else {
		sort.Sort(StatusWrapper{s, func(p, q *Status) bool {
			return p.CpuUtilization*(1.0+float32(p.NumberGoroutine)) < q.CpuUtilization*(1.0+float32(q.NumberGoroutine))
		}})
	}
}

func (s ServiceInfoArray) TopK(k int) ServiceInfoArray {
	l := len(s)
	if k < 0 || k > l {
		return s
	}

	return s[:k]
}
