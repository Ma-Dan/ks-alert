package option

import (
	"errors"
	"flag"
	"github.com/carmanzhang/ks-alert/pkg/utils/idutil"
	"os"
)

var (
	DispatcherServiceName = flag.String("dispatcher_service", "alert_dispatcher_service", "service name")
	DispatcherServiceHost = flag.String("dispatcher_service_host", "localhost", "service host")
	DispatcherServicePort = flag.Int("dispatcher_service_port", 50000, "listening port")

	ExecutorServiceName = flag.String("executor_service", "alert_executor_service", "service name")
	ExecutorServiceHost = flag.String("executor_service_host", "localhost", "service host")
	ExecutorServicePort = flag.Int("executor_service_port", 50001, "listening port")

	EtcdAddress = flag.String("dispatcher_etcd_addr", "http://127.0.0.1:2379", "register etcd address")
	Ip          = flag.String("ip", "127.0.0.1", "register etcd address")
)

var HostName string

func init() {
	flag.Parse()

	var err error
	HostName, err = os.Hostname()

	if err != nil {
		panic(err)
	}

	// a unique name in whole scope, mainly for distinguishing executors
	HostName = HostName + "-" + idutil.GetUuid36("")

	if *Ip == "" {
		panic(errors.New("ip address must be specific"))
	}
}
