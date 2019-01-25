package service

import (
	"context"
	"fmt"
	"github.com/carmanzhang/ks-alert/pkg/executor/handler"
	"github.com/carmanzhang/ks-alert/pkg/executor/runtime"
	"github.com/carmanzhang/ks-alert/pkg/models"
	"github.com/carmanzhang/ks-alert/pkg/option"
	"github.com/carmanzhang/ks-alert/pkg/pb"
	"github.com/carmanzhang/ks-alert/pkg/registry"
	"github.com/carmanzhang/ks-alert/pkg/utils/etcdutil"
	"github.com/golang/glog"
	"go.etcd.io/etcd/clientv3"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const (
	ServiceRegistrationInterval = 30
)

func Run() {
	lis, err := net.Listen("tcp", fmt.Sprintf(*option.ServiceHost+":%d", *option.ExecutorServicePort))
	if err != nil {
		panic(err)
	}

	// register executor service information in etcd
	err = registry.Register(*option.ExecutorServiceName, *option.ServiceHost, *option.ExecutorServicePort, *option.EtcdAddress, time.Second*ServiceRegistrationInterval, 60)
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

	// check
	go func() {
		prefix := "/" + *option.ExecutorServiceName + "/"
		//prefix := "/alert_executor_service"
		var e *etcdutil.Etcd
		for {
			time.Sleep(time.Second * 5)
			var err error
			e, err = etcdutil.Connect([]string{*option.EtcdAddress}, "")
			if err != nil {
				glog.Errorln(err.Error())
			}
			if e == nil {
				continue
			} else {
				break
			}
		}

		timer := time.NewTicker(time.Minute * 1)
		var counter = 0
		for {
			select {
			case <-timer.C:
				resp, _ := e.Get(context.Background(), prefix, clientv3.WithPrefix())
				ss := registry.GetAllExecutorServiceInfo(resp)
				addrMap := ss.ExtractServiceAddrs()
				svcAddr := fmt.Sprintf("%s:%d", *option.ServiceHost, *option.ExecutorServicePort)
				if _, ok := addrMap[svcAddr]; !ok {
					counter = counter + 1
				} else {
					counter = 0
				}

				if counter > 3 {
					// means this host does not register successfully, main thread will be terminated
					panic("this host does not register successfully in 3 consecutive checks, terminated main thread")
				}
			}
		}
	}()

	// ensure `hostid` in alert_config running at the right node
	go func() {
		//timer := time.NewTicker(time.Minute * 1)
		timer := time.NewTicker(time.Second * 15)
		hostID := fmt.Sprintf("%s:%d", *option.ServiceHost, *option.ExecutorServicePort)
		recordLimit := 3
		for {
			select {
			case <-timer.C:
				latestReportTime := time.Now().Add(-time.Minute * runtime.MaxKeepAliveReportInterval)
				alertConfigs, err := models.GetAbnormalExecutedAlertConfig(hostID, latestReportTime, recordLimit)

				if err != nil {
					glog.Errorln(err.Error())
				}

				for i := range *alertConfigs {
					ac := &(*alertConfigs)[i]
					if ac.HostID != hostID {
						ac.HostID = hostID
					}
					ac.KeepAliveAt = time.Now()
					ac.Version += 1
				}

				bools, err := models.UpdateAlertConfigBindingHostAndVersion(alertConfigs)

				if err != nil {
					glog.Errorln(err.Error())
				}

				for i, b := range bools {
					if b {
						acID := (*alertConfigs)[i].AlertConfigID
						fmt.Println("executing", acID)
						err := runtime.Action(context.Background(), &pb.Informer{Signal: pb.Informer_CREATE, AlertConfigId: acID})
						if err != nil {
							glog.Errorln(err.Error())
						}
					}
				}
			}
		}
	}()

	go ListenHttpApi()

	// TODO need launch a daemon, for indeed checking or ensuring alert_configs are executing in this node
	log.Printf("starting executor service at %s:%d", *option.ServiceHost, *option.ExecutorServicePort)
	s := grpc.NewServer()
	pb.RegisterExecutorServer(s, &handler.Executor{})
	s.Serve(lis)
}

func ListenHttpApi() {
	// http api for probe
	mux := http.NewServeMux()

	// livenessProbe
	mux.HandleFunc("/-/healthy", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("healthy"))
	})

	// readinessProbe
	mux.HandleFunc("/-/ready", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("ready"))
	})

	port := 80
	log.Printf("executor probe service at localhost:%d", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), mux)
}
