package grpcutil

import (
	"google.golang.org/grpc"
	"kubesphere.io/ks-alert/pkg/executor"
	"k8s.io/klog/glog"
)

var grpcConnection *grpc.ClientConn

func GetExecutorGRPCClient(address... string) (*pb.ExecutorClient, error){
	var err error
	var addr string
	if grpcConnection == nil {
		if address == nil {
			//addr =
		}else {
			addr = address[0]
		}

		grpcConnection, err = grpc.Dial(addr, grpc.WithInsecure())
		if err != nil {
			glog.Errorln(err.Error())
		}
		return nil, err
	}
	//defer conn.Close()
	client := pb.NewExecutorClient(grpcConnection)
	return &client, nil
}

func CloseExecutorGRPCConnection(grpcConnection *grpc.ClientConn) {
	if grpcConnection != nil {
		err := grpcConnection.Close()
		glog.Errorln(err.Error())
	}
}