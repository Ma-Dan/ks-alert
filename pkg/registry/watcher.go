package registry

import (
	"fmt"
	etcd3 "go.etcd.io/etcd/clientv3"
	"go.etcd.io/etcd/mvcc/mvccpb"
	"golang.org/x/net/context"
	"google.golang.org/grpc/naming"
)

// watcher is the implementaion of grpc.naming.Watcher
type watcher struct {
	re            *resolver // re: Etcd Resolver
	client        etcd3.Client
	isInitialized bool
}

// Close do nothing
func (w *watcher) Close() {
}

// Next to return the updates
func (w *watcher) Next() ([]*naming.Update, error) {
	// prefix is the etcd prefix/value to watch
	prefix := fmt.Sprintf("/%s/", w.re.serviceName)
	// check if is initialized
	if !w.isInitialized {
		// query addresses from etcd
		resp, err := w.client.Get(context.Background(), prefix, etcd3.WithPrefix())
		w.isInitialized = true
		if err == nil {
			addrs := GetAllExecutorServiceInfo(resp).ExtractServiceAddrs()
			//if not empty, return the updates or watcher new dir
			if l := len(addrs); l != 0 {
				updates := make([]*naming.Update, l)
				i := 0
				for add := range addrs {
					updates[i] = &naming.Update{Op: naming.Add, Addr: add}
					i = i + 1
				}
				return updates, nil
			}
		}
	}
	// generate etcd Watcher
	rch := w.client.Watch(context.Background(), prefix, etcd3.WithPrefix())
	for wresp := range rch {
		for _, ev := range wresp.Events {
			switch ev.Type {
			case mvccpb.PUT:
				return []*naming.Update{{Op: naming.Add, Addr: string(ev.Kv.Value)}}, nil
			case mvccpb.DELETE:
				return []*naming.Update{{Op: naming.Delete, Addr: string(ev.Kv.Value)}}, nil
			}
		}
	}
	return nil, nil
}
