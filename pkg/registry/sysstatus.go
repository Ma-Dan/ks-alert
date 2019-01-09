package registry

import (
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
	"k8s.io/klog/glog"
	"runtime"
	"time"
)

type Status struct {
	TotalMemory       int
	MemoryUtilization float32
	TotalCores        int
	CpuUtilization    float32
	HostName          string
	Uptime            int
	NumberProcess     int
	OS                string
	NumberGoroutine   int
}

func GetHardwareData() *Status {
	vmStat, err := mem.VirtualMemory()
	cpuStat, err := cpu.Info()
	percentage, err := cpu.Percent(time.Duration(time.Second*1), true)
	hostStat, err := host.Info()
	numofGoroutine := runtime.NumGoroutine()

	if err != nil {
		glog.Errorln(err.Error())
	}

	avgPercent := 0.0
	for _, cpupercent := range percentage {
		avgPercent += cpupercent
	}
	avgPercent = avgPercent * 0.01 / float64(cpuStat[0].Cores) // [0].Cores
	return &Status{
		CpuUtilization:    float32(avgPercent),
		HostName:          hostStat.OS,
		MemoryUtilization: float32(vmStat.UsedPercent * 0.01),
		NumberGoroutine:   numofGoroutine,
		NumberProcess:     int(hostStat.Procs),
		OS:                hostStat.OS,
		TotalCores:        int(cpuStat[0].Cores),
		TotalMemory:       int(vmStat.Total / 1024 / 1024),
		Uptime:            int(hostStat.Uptime),
	}
}
