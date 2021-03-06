package option

import (
	"errors"
	"flag"
	"fmt"
	"github.com/golang/glog"
	"github.com/shirou/gopsutil/net"
	"strings"
)

var (
	ServiceHost = flag.String("ip", "127.0.0.1", "service host")

	DispatcherServiceName = flag.String("dispatcher_service", "alert_dispatcher_service", "service name")
	DispatcherServicePort = flag.Int("dispatcher_port", 50000, "listening port")

	ExecutorServiceName = flag.String("executor_service", "alert_executor_service", "service name")
	ExecutorServicePort = flag.Int("executor_port", 50001, "listening port")

	EtcdAddress = flag.String("etcd_addr", "http://127.0.0.1:2379", "register etcd address")

	MysqlHost = flag.String("mysql", "127.0.0.1", "")
	MysqlPort = flag.String("mysql_port", "3306", "")

	Database = flag.String("database", "alert", "")
	User     = flag.String("user", "root", "")
	Password = flag.String("password", "password", "")

	NotificationHost = flag.String("nf_host", "alerting-notification-server.kubesphere-monitoring-system.svc", "")
	NotificationPort = flag.String("nf_port", "9201", "")
)

var HostInfo string

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

	if len(addrs) == 0 {
		panic("no up network interface")
	}

	if _, ok := addrs[*ServiceHost]; !ok {
		glog.Errorln("giving ip does not exist")

		for k := range addrs {
			*ServiceHost = k
			break
		}
	}

	// a unique name in whole scope, mainly for distinguishing executors
	HostInfo = fmt.Sprintf("%s:%d", *ServiceHost, *ExecutorServicePort)

	glog.Infoln(HostInfo)
}
