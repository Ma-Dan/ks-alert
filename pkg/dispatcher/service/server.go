package service

import (
	"fmt"
	"github.com/carmanzhang/ks-alert/pkg/dispatcher/handler"
	"github.com/carmanzhang/ks-alert/pkg/option"
	"github.com/carmanzhang/ks-alert/pkg/pb"
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
	DispatcherServiceRegistrationInterval = 30
)

func Run() {
	lis, err := net.Listen("tcp", fmt.Sprintf(*option.ServiceHost+":%d", *option.DispatcherServicePort))
	if err != nil {
		panic(err)
	}
	// register executor service information in etcd
	err = registry.Register(*option.DispatcherServiceName, *option.ServiceHost, *option.DispatcherServicePort, *option.EtcdAddress, time.Second*DispatcherServiceRegistrationInterval, 60)
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
	log.Printf("starting dispatcher service at %s:%d", *option.ServiceHost, *option.DispatcherServicePort)
	s := grpc.NewServer()

	pb.RegisterAlertConfigHandlerServer(s, &handler.AlertHandler{})
	pb.RegisterAlertHistoryHandlerServer(s, &handler.AlertHistoryHandler{})

	pb.RegisterEnterpriseHandlerServer(s, &handler.EnterpriseHandler{})
	pb.RegisterProductHandlerServer(s, &handler.ProductHandler{})
	pb.RegisterResourceTypeHandlerServer(s, &handler.ResourceTypeHandler{})

	pb.RegisterAlertRuleHandlerServer(s, &handler.AlertRuleHandler{})
	pb.RegisterResourceHandlerServer(s, &handler.ResourceHandler{})
	pb.RegisterReceiverHandlerServer(s, &handler.ReceiverHandler{})

	pb.RegisterSilenceHandlerServer(s, &handler.SilenceHandler{})
	pb.RegisterSuggestionHandlerServer(s, &handler.SuggestionHandler{})
	pb.RegisterSeverityHandlerServer(s, &handler.SeverityHandler{})

	s.Serve(lis)
}
