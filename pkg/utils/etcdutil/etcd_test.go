package etcdutil

import (
	"context"
	"fmt"
	etcd3 "go.etcd.io/etcd/clientv3"
	"log"
	"testing"

	"github.com/carmanzhang/ks-alert/pkg/option"
)

func TestConnect(t *testing.T) {
	//e := new(Etcd)
	endpoints := []string{"127.0.0.1:2379"}
	prefix := "/" + *option.ExecutorServiceName + "/"
	//prefix := "/alert_executor_service"
	e, err := Connect(endpoints, "")
	resp, err := e.Get(context.Background(), prefix, etcd3.WithPrefix())

	for i := range resp.Kvs {
		kv := resp.Kvs[i]
		fmt.Println(string(kv.Key))
		fmt.Println(string(kv.Value))
	}

	log.Println(err)
	if err != nil {
		t.Fatal(err)
	}

}

func TestNewQueue(t *testing.T) {
	//endpoints:=[]string{"192.168.0.7:2379,192.168.0.8:2379,192.168.0.6:2379"}
	endpoints := []string{"192.168.0.7:2379"}
	prefix := "test"
	e, err := Connect(endpoints, prefix)
	log.Println(e)
	if err != nil {
		t.Fatal(err)
	}

	q := e.NewQueue("notification")
	q.Enqueue("ssss")
}

func TestEnqueue(t *testing.T) {
	endpoints := []string{"192.168.0.7:2379"}
	prefix := "test"
	e, err := Connect(endpoints, prefix)
	if err != nil {
		t.Fatal(err)
	}
	queue := e.NewQueue("notification")
	go func() {
		for i := 0; i < 100; i++ {
			err := queue.Enqueue(fmt.Sprintf("%d", i))
			if err != nil {
				t.Fatal(err)
			}
			t.Logf("Push message to queue, worker number [%d]", i)
		}

	}()
	for i := 0; i < 100; i++ {
		n, err := queue.Dequeue()
		if err != nil {
			t.Fatal(err)
		}
		t.Logf("Got message [%s] from queue, worker number [%d]", n, i)
	}
}

func TestEnqueue2(t *testing.T) {
	endpoints := []string{"192.168.0.7:2379"}
	prefix := "nf_"
	e, err := Connect(endpoints, prefix)
	if err != nil {
		t.Fatal(err)
	}
	queue := e.NewQueue("nf_")
	//go func() {
	//	for i := 0; i < 100; i++ {
	//		err := queue.Enqueue(fmt.Sprintf("%d", i))
	//		if err != nil {
	//			t.Fatal(err)
	//		}
	//		t.Logf("Push message to queue, worker number [%d]", i)
	//	}
	//
	//}()
	for i := 0; i < 100; i++ {
		n, err := queue.Dequeue()
		if err != nil {
			t.Fatal(err)
		}
		t.Logf("Got message [%s] from queue, worker number [%d]", n, i)
	}
}
