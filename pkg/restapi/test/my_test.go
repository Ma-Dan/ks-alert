package main

import (
	"testing"
	"time"
	"fmt"
	"strconv"
	"github.com/pkg/errors"
	"github.com/golang/glog"
	"strings"
	"net/http"
	"github.com/emicklei/go-restful"
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

func TestReadEntityJsonCharset(t *testing.T) {
	bodyReader := strings.NewReader(`{"Value" : "42"}`)
	httpRequest, _ := http.NewRequest("GET", "/test", bodyReader)
	httpRequest.Header.Set("Content-Type", "application/json; charset=UTF-8")
	request := restful.NewRequest(httpRequest)
	sam := new(Sample)
	request.ReadEntity(sam)
	if sam.Value != "42" {
		t.Fatal("read failed")
	}
	fmt.Println(sam)
}

type Sample struct {
	Value string
}

func TestTimer(t *testing.T) {
	//初始化通道
	ch11 := make(chan int, 1000)
	sign := make(chan byte, 1)

	//给ch11通道写入数据
	for i := 0; i < 1000; i++ {
		ch11 <- i
	}

	//单独起一个Goroutine执行select
	go func(){
		var e int
		ok := true
		//首先声明一个*time.Timer类型的值，然后在相关case之后声明的匿名函数中尽可能的复用它
		var timer *time.Timer

		for{
			select {
			case e = <- ch11:
				fmt.Printf("ch11 -> %d\n",e)
			case <- func() <-chan time.Time {
				if timer == nil{
					//初始化到期时间据此间隔1ms的定时器
					timer = time.NewTimer(time.Millisecond)
				}else {
					//复用，通过Reset方法重置定时器
					timer.Reset(time.Millisecond)
				}
				//得知定时器到期事件来临时，返回结果
				return timer.C
			}():
				fmt.Println("Timeout.")
				ok = false
				break
			}
			//终止for循环
			if !ok {
				sign <- 0
				break
			}
		}

	}()

	//惯用手法，读取sign通道数据，为了等待select的Goroutine执行。
	<- sign
}