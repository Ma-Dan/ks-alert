package main

import (
	"flag"
	"fmt"
	"github.com/carmanzhang/ks-alert/pkg/executor/handler"
	"github.com/carmanzhang/ks-alert/pkg/executor/pb"
	"github.com/carmanzhang/ks-alert/pkg/registry"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const (
	DefaultServiceHost          = "127.0.0.1"
	DefaultServiceName          = "alert_executor_service"
	ServiceRegistrationInterval = 10
)

var (
	serviceName = flag.String("service", "", "service name")
	serviceHost = flag.String("service_host", "", "service host")
	servicePort = flag.Int("service_port", 50001, "listening port")
	etcdAddress = flag.String("etcd_addr", "http://127.0.0.1:2379", "register etcd address")
)

func main() {
	flag.Parse()

	if *serviceHost == "" {
		//panic(errors.New("executor service host ip is not specified"))
		*serviceHost = DefaultServiceHost
	}

	if *serviceName == "" {
		*serviceName = DefaultServiceName
	}

	lis, err := net.Listen("tcp", fmt.Sprintf(*serviceHost+":%d", *servicePort))
	if err != nil {
		panic(err)
	}

	// register executor service information in etcd
	err = registry.Register(*serviceName, *serviceHost, *servicePort, *etcdAddress, time.Second*ServiceRegistrationInterval, 15)
	if err != nil {
		panic(err)
	}

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL, syscall.SIGHUP, syscall.SIGQUIT)

	go func() {
		s := <-ch
		log.Printf("receive signal '%v'", s)
		registry.UnRegister()
		os.Exit(1)
	}()
	log.Printf("starting executor service at %s:%d", *serviceHost, *servicePort)
	s := grpc.NewServer()
	pb.RegisterExecutorServer(s, &handler.Executor{})
	s.Serve(lis)
}
