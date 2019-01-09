package service

import (
	"fmt"
	"github.com/carmanzhang/ks-alert/pkg/executor/handler"
	"github.com/carmanzhang/ks-alert/pkg/executor/pb"
	"github.com/carmanzhang/ks-alert/pkg/option"
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
	ServiceRegistrationInterval = 30
)

func Run() {
	lis, err := net.Listen("tcp", fmt.Sprintf(*option.ExecutorServiceHost+":%d", *option.ExecutorServicePort))
	if err != nil {
		panic(err)
	}

	// register executor service information in etcd
	err = registry.Register(*option.ExecutorServiceName, *option.ExecutorServiceHost, *option.ExecutorServicePort, *option.EtcdAddress, time.Second*ServiceRegistrationInterval, 60)
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
	log.Printf("starting executor service at %s:%d", *option.ExecutorServiceHost, *option.ExecutorServicePort)
	s := grpc.NewServer()
	pb.RegisterExecutorServer(s, &handler.Executor{})
	s.Serve(lis)
}
