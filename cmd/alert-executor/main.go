package main

import (
	"kubesphere.io/ks-alert/pkg/executor/service"
	"github.com/golang/glog"
)

func main() {
	defer glog.Flush()
	service.Run()
}
