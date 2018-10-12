package main

import (
	"testing"
	"time"
	"fmt"
	"strconv"
	"github.com/pkg/errors"
	"github.com/golang/glog"
)

func TestMap(testing *testing.T) {

	//timestamp := time.Now().Unix()
	//fmt.Println(timestamp)
	//secondTime := time.Now().Minute()
	//fmt.Println(secondTime)
	ts := "1539320671.12134"
	_, err := strconv.ParseInt(ts, 10, 0)
	if err != nil {
		glog.Error(errors.Errorf("Parse request time %s failed", ts))
	}

	timeFlot, _ := strconv.ParseFloat(ts, 64)
	timeInt := int64(timeFlot)
	fmt.Println(timeInt)
	secondTime1 := time.Unix(timeInt, 0).Truncate(time.Minute).Unix()
	fmt.Println(secondTime1)

	secondTime2 := time.Unix(timeInt, 0).Round(time.Minute).Unix()
	fmt.Println(secondTime2)
	//strconv.FormatInt(int64(time.Now().Unix()), 10)
	//secondTime1 := time.Now().Truncate(time.Minute).Unix()
	//fmt.Println(secondTime1)
	//var devopsProject []string
	//devopsProject = nil
	//fmt.Println(devopsProject)
	//for _, dp := range devopsProject {
	//	fmt.Println("dp: " + dp)
	//}
	fmt.Println(convertTimeGranularity("1539320671.12134"))
}

func convertTimeGranularity(ts string) string {
	timeFloat, err := strconv.ParseFloat(ts, 64)
	if err != nil {
		glog.Errorf("convert second timestamp %s to minute timestamp failed", ts)
		return strconv.FormatInt(int64(time.Now().Unix()), 10)
	}
	timeInt := int64(timeFloat)
	// convert second timestamp to minute timestamp
	secondTime := time.Unix(timeInt, 0).Truncate(time.Minute).Unix()
	return strconv.FormatInt(secondTime, 10)
}