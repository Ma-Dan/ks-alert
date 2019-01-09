package option

import "flag"

var (
	DispatcherServiceName = flag.String("dispatcher_service", "alert_dispatcher_service", "service name")
	DispatcherServiceHost = flag.String("dispatcher_service_host", "localhost", "service host")
	DispatcherServicePort = flag.Int("dispatcher_service_port", 50000, "listening port")

	ExecutorServiceName = flag.String("executor_service", "alert_executor_service", "service name")
	ExecutorServiceHost = flag.String("executor_service_host", "localhost", "service host")
	ExecutorServicePort = flag.Int("executor_service_port", 50001, "listening port")

	EtcdAddress = flag.String("dispatcher_etcd_addr", "http://127.0.0.1:2379", "register etcd address")
)

type ServerRunOptions struct {
	executorServiceName string
	etcdEndpoints       string
}

func init() {
	flag.Parse()
}
