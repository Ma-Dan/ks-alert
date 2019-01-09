package jsonutil

import (
	"github.com/json-iterator/go"
	"k8s.io/klog/glog"
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

func UnmarshalBinary(bys []byte, v interface{}) {
	err := json.Unmarshal(bys, v)

	if err != nil {
		glog.Errorf("Unmarshal json %+v object failed, err: %s", v, err.Error())
	}
}
