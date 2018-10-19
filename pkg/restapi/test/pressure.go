package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"sync"
	"time"
)

var successCount int32
var failureCount int32
var mutex sync.Mutex

var client = &http.Client{}

//func testUnit(wg sync.WaitGroup, poststr string, params string) {
//	defer wg.Done()
//	resp, err := client.Get("http://139.198.190.141:8087/api/v1alpha1/monitoring/" + poststr + params)
//	if resp != nil {
//		defer resp.Body.Close()
//	}
//	if err != nil {
//		fmt.Println("error:", err)
//		return
//	}
//
//	body, err := ioutil.ReadAll(resp.Body)
//	if err != nil {
//		fmt.Println("error:", err)
//		return
//	}
//	mutex.Lock()
//	fmt.Println(string(body))
//	if strings.Count(string(body), "metric_name") == 20 {
//		successCount += 1
//	}else {
//		failureCount += 1
//		fmt.Println("---- 失败 -> ", failureCount)
//	}
//	defer mutex.Unlock()
//
//}

func main() {
	successCount = 0
	failureCount = 0
	timestart := time.Now()
	var wg sync.WaitGroup
	for i := 0; i < 4000; i++ { // time查询 7000  range_query 6600
		wg.Add(1)
		go func() {
			defer wg.Done()
			resp, err := client.Get("http://139.198.190.141:8087/api/v1alpha1/monitoring/nodes?time=1539331481&metrics_filter=node_pod_count")
			//resp, err := client.Get("http://139.198.190.141:8087/api/v1alpha1/monitoring/nodes?start=1539331400&end=1539331481&step=20s&metrics_filter=node_pod_count")
			//resp, err := client.Get("http://139.198.190.141:8087/api/v1alpha1/monitoring/nodes")
			if resp != nil {
				defer resp.Body.Close()
			} else {
				mutex.Lock()
				failureCount += 1
				fmt.Println("---- 失败le -> ", failureCount)
				defer mutex.Unlock()
				return
			}
			if err != nil {
				fmt.Println("error:", err)
				return
			}
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				fmt.Println("error:", err)
				return
			}
			mutex.Lock()
			if strings.Count(string(body), "metric_name") == 1 {
				successCount += 1
				if strings.Contains(string(body), "err") || strings.Contains(string(body), "failed") {
					fmt.Println(string(body))
				}
			} else {
				failureCount += 1
				fmt.Println("---- 失败 -> ", failureCount)
			}
			defer mutex.Unlock()
		}()
	}
	wg.Wait()
	fmt.Println("successCount", successCount)
	fmt.Println("failureCount", failureCount)
	//fmt.Println(time.Since(timestart).Nanoseconds())
	fmt.Println(time.Since(timestart).Seconds())
}

//func main1() {
//	successCount = 0
//	failureCount = 0
//	timestart := time.Now()
//	var wg sync.WaitGroup
//	for i := 0; i < 50; i++ {
//		wg.Add(1)
//		go testUnit(wg, "clusters", "?time=1539327156")
//	}
//	wg.Wait()
//	fmt.Println("successCount", successCount)
//	fmt.Println("failureCount", failureCount)
//	//fmt.Println(time.Since(timestart).Nanoseconds())
//	fmt.Println(time.Since(timestart).Seconds())
//}
