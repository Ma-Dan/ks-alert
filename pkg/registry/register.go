package registry

import (
	"fmt"
	"github.com/carmanzhang/ks-alert/pkg/utils/jsonutil"
	"github.com/golang/glog"
	etcd3 "go.etcd.io/etcd/clientv3"
	"golang.org/x/net/context"
	"log"
	"strings"
	"time"
)

// Prefix should start and end with no slash
var client *etcd3.Client
var svcKey string
var stopSignal = make(chan bool, 1)

// Register
func Register(svcName string, host string, port int, target string, interval time.Duration, ttl int) error {
	svcAddr := fmt.Sprintf("%s:%d", host, port)
	svcKey = fmt.Sprintf("/%s/%s", svcName, svcAddr)
	glog.Infoln(svcAddr, svcKey)
	// get endpoints for register dial address
	var err error
	client, err = etcd3.New(etcd3.Config{
		Endpoints: strings.Split(target, ","),
	})
	if err != nil {
		glog.Infoln(err.Error())
		return fmt.Errorf("create etcd3 client failed: %v", err)
	}
	go func() {
		// invoke self-register with ticker
		ticker := time.NewTicker(interval)
		for {
			select {
			case <-stopSignal:
				return
			case <-ticker.C:
				// minimum lease TTL is ttl-second
				resp, err := client.Grant(context.TODO(), int64(ttl))
				if err != nil {
					glog.Errorln(err.Error())
				}

				value := jsonutil.Marshal(&ServiceInfo{ServiceAddress: svcAddr, SysStatus: GetHardwareData()})
				if _, err := client.Put(context.Background(), svcKey, value, etcd3.WithLease(resp.ID)); err != nil {
					log.Printf("refresh service '%s' with ttl to etcd3 failed: %s", svcName, err.Error())
				}
			}
		}
	}()
	return nil
}

// UnRegister delete registered service from etcd
func UnRegister() error {
	stopSignal <- true
	stopSignal = make(chan bool, 1) // just a hack to avoid multi UnRegister deadlock
	var err error
	if _, err := client.Delete(context.Background(), svcKey); err != nil {
		log.Printf("unregister '%s' failed: %s", svcKey, err.Error())
	} else {
		log.Printf("unregister '%s' ok.", svcKey)
	}
	return err
}
