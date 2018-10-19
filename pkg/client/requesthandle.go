package client

import (
	"github.com/emicklei/go-restful"
	"alert-kubesphere-plugin/pkg/models"
	"github.com/golang/glog"
	"fmt"
	"alert-kubesphere-plugin/pkg/prometheus"
	"encoding/json"
	"time"
	"strconv"
)

func SenderAlertConfig(request *restful.Request, response *restful.Response) {

	//var resEntity models.UserRequest
	resEntity := new(models.UserRequest)
	err := request.ReadEntity(resEntity)
	if err != nil {
		glog.Errorln(err)
	}
	fmt.Println("readed user request: ", resEntity)
	// convert to rule
	metricName := resEntity.MetricName
	//userId := resEntity.UserId
	// find rule template by metric name
	ruleOrPromQL := enrichRulePromQL(metricName, resEntity.ResourceName, resEntity.ResourceType)
	// call prometheus api directly instead of calling enclosured apis
	paramsMap, _, err := prometheus.ParseRequestHeader(request)
	if err != nil {
		glog.Error(err)
	}

	paramsMap.Set("query", ruleOrPromQL)
	params := paramsMap.Encode()
	// instance query
	postfix := "query?"
	//response.WriteAsJson(metric)
	// call alert strategy, push alert to alert manager when this alert fired
	go monitorAndAlert(postfix, params)





}

func monitorAndAlert(query string, params string) {
	var tk *time.Ticker = time.NewTicker(10 * time.Second)
	var lastValue int64 = 0
	var avgValue int64 = 0
	for t := range tk.C {
		res := prometheus.SendRequest(query, params)
		// alarm judegment, there are many stratges
		var metric prometheus.CommonMetricsResult
		json.Unmarshal([]byte(res), &metric)
		fmt.Println(metric)
		if metric.Status == "success" {
			t.Unix()
			//resultType := metric.Data.ResultType
			for _, resultItem := range metric.Data.Result {
				curValue, _ := strconv.ParseInt(resultItem.Value[1].(string), 10, 0)
				avgValue = (curValue + lastValue) / 2
				lastValue = curValue
				if avgValue > 40 {
					// push alert

				}
			}
		}
	}


}




func enrichRulePromQL(metricName string, resourceName models.ResourceName, resourceType string) string {
	var ruleOrPromQL string
	switch resourceType {
	case "workspace":
		ruleOrPromQL = prometheus.MakeWorkspacePromQL(metricName, resourceName)   // resourceName regex support
	case "namespace":
		ruleOrPromQL = prometheus.MakeNamespacePromQL(metricName, resourceName)
	case "workload" :
		ruleOrPromQL = prometheus.MakeWorkloadRule(metricName, resourceName)
	case "pod" :
		ruleOrPromQL = prometheus.MakePodPromQL(metricName, resourceName)
	case "container" :
		ruleOrPromQL = prometheus.MakeContainerPromQL(metricName, resourceName)
	}
	return ruleOrPromQL
}