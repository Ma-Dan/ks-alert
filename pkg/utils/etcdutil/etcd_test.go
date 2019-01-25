package etcdutil

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"log"
	"testing"
)

func TestConnect(t *testing.T) {
	//e := new(Etcd)
	endpoints := []string{"139.198.120.226:32379"}
	//prefix := "/" + *option.ExecutorServiceName + "/"
	//prefix := "/alert_executor_service"
	prefix := "/alert_dispatcher_service/127.0.0.1:50000"
	e, err := Connect(endpoints, "")
	resp, err := e.Get(context.Background(), prefix, clientv3.WithPrefix())

	for i := range resp.Kvs {
		kv := resp.Kvs[i]
		fmt.Println(string(kv.Key), string(kv.Value))
	}

	log.Println(err)
	if err != nil {
		t.Fatal(err)
	}
}

func TestPut(t *testing.T) {
	//e := new(Etcd)
	endpoints := []string{"139.198.120.226:32379"}
	//prefix := "/" + *option.ExecutorServiceName + "/"
	//prefix := "/alert_executor_service"
	e, err := Connect(endpoints, "")
	//resp, err := e.Get(context.Background(), prefix, etcd3.WithPrefix())
	_, err = e.Put(context.Background(), "/hello", "hahahahah")

	log.Println(err)
	if err != nil {
		t.Fatal(err)
	}
}
