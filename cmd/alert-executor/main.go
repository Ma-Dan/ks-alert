package main

import (
	"github.com/carmanzhang/ks-alert/pkg/executor/service"
	"github.com/golang/glog"
)

func main() {
	defer glog.Flush()
	service.Run()
}
