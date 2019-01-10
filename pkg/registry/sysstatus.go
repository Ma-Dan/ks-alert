package registry

import (
	"context"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
	"k8s.io/klog/glog"
	"runtime"
	"sync"
	"time"
)

type Status struct {
	TotalMemory       int
	MemoryUtilization float32
	TotalCores        int
	CpuUtilization    float32
	//NumberProcess     int
	NumberGoroutine int
	HostInfo        *host.InfoStat
}

const DefaultTimeout = 3

func GetHardwareData() *Status {

	ctx := context.Background()

	var vmStat *mem.VirtualMemoryStat
	var wg = sync.WaitGroup{}
	wg.Add(1)
	go func() {
		ctx, cancel := context.WithTimeout(ctx, time.Second*DefaultTimeout)
		defer cancel()
		var err error
		vmStat, err = mem.VirtualMemoryWithContext(ctx)
		if err != nil {
			glog.Errorln(err.Error())
		}
		wg.Done()
	}()

	var cpuStates []cpu.InfoStat
	wg.Add(1)
	go func() {
		ctx, cancel := context.WithTimeout(ctx, time.Second*DefaultTimeout)
		defer cancel()
		var err error
		cpuStates, err = cpu.InfoWithContext(ctx)
		if err != nil {
			glog.Errorln(err.Error())
		}
		wg.Done()
	}()

	var hostStat *host.InfoStat
	wg.Add(1)
	go func() {
		ctx, cancel := context.WithTimeout(ctx, time.Second*DefaultTimeout)
		defer cancel()
		var err error
		hostStat, err = host.InfoWithContext(ctx)
		if err != nil {
			glog.Errorln(err.Error())
		}
		wg.Done()
	}()

	var avgPercent = 5.0
	wg.Add(1)
	go func() {
		ctx, cancel := context.WithTimeout(ctx, time.Second*DefaultTimeout)
		defer cancel()
		avgPercents, err := cpu.PercentWithContext(ctx, time.Duration(time.Second*1), false)
		if avgPercents != nil && len(avgPercents) != 0 {
			avgPercent = avgPercents[0]
		}
		if err != nil {
			glog.Errorln(err.Error())
		}
		wg.Done()
	}()

	wg.Wait()

	memUse := 0.1
	var memTotal uint64 = 4000
	if vmStat != nil {
		memUse = vmStat.UsedPercent * 0.01
		memTotal = vmStat.Total
	}

	core := 1
	if cpuStates != nil && len(cpuStates) != 0 {
		core = int(cpuStates[0].Cores)
	}

	numofGoroutine := runtime.NumGoroutine()

	return &Status{
		CpuUtilization:    float32(avgPercent * 0.01),
		MemoryUtilization: float32(memUse),
		NumberGoroutine:   numofGoroutine,
		TotalCores:        core,
		TotalMemory:       int(memTotal / 1024 / 1024),
		HostInfo:          hostStat,
	}
}
