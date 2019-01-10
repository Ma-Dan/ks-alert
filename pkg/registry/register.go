package registry

import (
	"fmt"
	"github.com/carmanzhang/ks-alert/pkg/utils/jsonutil"
	etcd3 "go.etcd.io/etcd/clientv3"
	"go.etcd.io/etcd/etcdserver/api/v3rpc/rpctypes"
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

	// get endpoints for register dial address
	var err error
	client, err = etcd3.New(etcd3.Config{
		Endpoints: strings.Split(target, ","),
	})
	if err != nil {
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
				resp, _ := client.Grant(context.TODO(), int64(ttl))
				// should get first, if not exist, set it
				_, err := client.Get(context.Background(), svcKey)
				if err != nil {
					if err == rpctypes.ErrKeyNotFound {
						if _, err := client.Put(context.TODO(), svcKey, svcAddr, etcd3.WithLease(resp.ID)); err != nil {
							log.Printf("set service '%s' with ttl to etcd3 failed: %s", svcName, err.Error())
						}
					} else {
						log.Printf("service '%s' connect to etcd3 failed: %s", svcName, err.Error())
					}
				} else {
					// refresh set to true for not notifying the watcher
					// TODO this calling is time consuming
					value := jsonutil.Marshal(&ServiceInfo{ServiceAddress: svcAddr, SysStatus: GetHardwareData()})
					if _, err := client.Put(context.Background(), svcKey, value, etcd3.WithLease(resp.ID)); err != nil {
						log.Printf("refresh service '%s' with ttl to etcd3 failed: %s", svcName, err.Error())
					}
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
