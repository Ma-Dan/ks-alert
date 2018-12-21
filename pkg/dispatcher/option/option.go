package option

import "flag"


var (
	ExecutorServiceName = flag.String("service", "alert_executor_service", "service name")
	EtcdAddr            = flag.String("etcd", "http://127.0.0.1:2379", "register etcd address")
)

type ServerRunOptions struct {
	executorServiceName string
	etcdEndpoints       string
}

// NewServerRunOptions creates a new ServerRunOptions object with default parameters
func NewServerRunOptions() *ServerRunOptions {
	s := ServerRunOptions{
		etcdEndpoints:       *EtcdAddr,
		executorServiceName: *ExecutorServiceName,
	}
	return &s
}

func (s *ServerRunOptions) GetExecutorServiceName() string {
	return s.executorServiceName
}

func (s *ServerRunOptions) GetETCDAddresses() string {
	return s.etcdEndpoints
}
