package main

import (
	"kubesphere.io/ks-alert/pkg/dispatcher/service"
	"github.com/golang/glog"
)

func main() {
	defer glog.Flush()
	service.Run()
}
