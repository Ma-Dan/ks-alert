package option

import (
	"errors"
	"flag"
	"fmt"
	"github.com/carmanzhang/ks-alert/pkg/utils/idutil"
	"github.com/shirou/gopsutil/net"
	"os"
	"strings"
)

var (
	ServiceHost = flag.String("ip", "127.0.0.1", "service host")

	DispatcherServiceName = flag.String("dispatcher_service", "alert_dispatcher_service", "service name")
	DispatcherServicePort = flag.Int("dispatcher_port", 50000, "listening port")

	ExecutorServiceName = flag.String("executor_service", "alert_executor_service", "service name")
	ExecutorServicePort = flag.Int("executor_port", 50001, "listening port")

	EtcdAddress = flag.String("etcd_addr", "http://127.0.0.1:2379", "register etcd address")
)

var HostID string

func init() {
	flag.Parse()

	if *ServiceHost == "" {
		panic(errors.New("ip address must be specific"))
	}

	ifs, err := net.Interfaces()

	if err != nil {
		panic(err)
	}

	var addrs = make(map[string]string)
	for i := 0; i < len(ifs); i++ {
		if ifs[i].HardwareAddr == "" {
			continue
		}
		flags := ifs[i].Flags

		b := false
		for _, f := range flags {
			if f == "up" {
				b = true
			}
		}
		if b {
			addr := ifs[i].Addrs
			for _, a := range addr {
				if !strings.Contains(a.Addr, "::") {
					x := a.Addr[0:strings.Index(a.Addr, "/")]
					addrs[x] = ""
				}
			}
		}
	}

	// TODO need uncomment
	//if _, ok := addrs[*ServiceHost]; len(addrs) == 0 || !ok {
	//	panic(errors.New("no up network interface or giving ip does not exist"))
	//}

	HostID, err = os.Hostname()
	if err != nil {
		panic(err)
	}

	// a unique name in whole scope, mainly for distinguishing executors
	HostID = HostID + "-" + *ServiceHost + "-" + idutil.GetUuid36("")
	fmt.Println(HostID)
}
