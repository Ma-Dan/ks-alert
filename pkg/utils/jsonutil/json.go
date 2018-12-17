package jsonutil

import (
	"k8s.io/klog/glog"
	"github.com/json-iterator/go"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

func Marshal(v interface{}) string {
	byteArray, err := json.Marshal(v)

	if err != nil {
		glog.Errorf("Marshal json %+v object failed, err: %s", v, err.Error())
		return ""
	}

	return string(byteArray)
}

func Unmarshal(res string, v interface{}) {
	err := json.Unmarshal([]byte(res), v)

	if err != nil {
		glog.Errorf("Unmarshal json %+v object failed, err: %s", v, err.Error())
	}
}
