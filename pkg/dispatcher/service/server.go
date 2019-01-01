package service

import (
	"flag"
	"fmt"
	"github.com/carmanzhang/ks-alert/pkg/dispatcher/handler"
	"github.com/carmanzhang/ks-alert/pkg/dispatcher/pb"
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
	DefaultDispatcherServiceHost          = "127.0.0.1"
	DefaultDispatcherServiceName          = "alert_dispatcher_service"
	DispatcherServiceRegistrationInterval = 10
)

var (
	dispatcherServiceName = flag.String("dispatcher_service", "", "service name")
	dispatcherServiceHost = flag.String("dispatcher_service_host", "", "service host")
	dispatcherServicePort = flag.Int("dispatcher_service_port", 50000, "listening port")
	dispatcherEtcdAddress = flag.String("dispatcher_etcd_addr", "http://127.0.0.1:2379", "register etcd address")
)

func Run() {
	flag.Parse()
	if *dispatcherServiceHost == "" {
		//panic(errors.New("executor service host ip is not specified"))
		*dispatcherServiceHost = DefaultDispatcherServiceHost
	}

	if *dispatcherServiceName == "" {
		*dispatcherServiceName = DefaultDispatcherServiceName
	}

	lis, err := net.Listen("tcp", fmt.Sprintf(*dispatcherServiceHost+":%d", *dispatcherServicePort))
	if err != nil {
		panic(err)
	}

	// register executor service information in etcd
	err = registry.Register(*dispatcherServiceName, *dispatcherServiceHost, *dispatcherServicePort, *dispatcherEtcdAddress, time.Second*DispatcherServiceRegistrationInterval, 15)
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
	log.Printf("starting executor service at %s:%d", *dispatcherServiceHost, *dispatcherServicePort)
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
