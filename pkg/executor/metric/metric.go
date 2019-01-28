package metric

import (
	"encoding/json"
	"fmt"
	"k8s.io/klog/glog"
)

const (
	MetricTypeVector = "vector"
	MetricTypeMatrix = "matrix"
)

type Metrics struct {
	MetricsLevel string   `json:"metrics_level"`
	Results      []Metric `json:"results"`
}

type Metric struct {
	MetricName string     `json:"metric_name, omitempty"`
	Status     string     `json:"status"`
	Data       MetricData `json:"data, omitempty"`
}

type MetricData struct {
	Result     json.RawMessage
	ResultType string `json:"resultType"`
}

type InstantMetric struct {
	MetricInfo  map[string]interface{} `json:"metric"`
	MetricValue []interface{}          `json:"value"`
}

type TimeSeriesMetric struct {
	MetricInfo   map[string]interface{} `json:"metric"`
	MetricValues [][]interface{}        `json:"values"`
}

type TV struct {
	T float64
	V string
}

type ResourceMetrics struct {
	ResourceMetric map[string][]TV
	RuleName       string
	RuleIndx       int
}

const (
	ConsecutiveCount    = 3
	MaxConsecutiveCount = 15
	Step                = 5
	MaxStep             = 60
)

// a single metric for many resources
func GetResourceTimeSeriesMetric(metricStr string, metricName string, startTime int64, endTime int64) *ResourceMetrics {
	//fmt.Println(metricStr)
	fmt.Println(startTime, endTime)
	var metrics Metrics
	err := json.Unmarshal([]byte(metricStr), &metrics)
	if err != nil {
		glog.Errorln(err.Error())
	}
	l := len(metrics.Results)

	var m *ResourceMetrics

	for i := 0; i < l; i++ {
		if metrics.Results[i].MetricName != metricName {
			continue
		}

		tp := metrics.Results[i].Data.ResultType
		r := metrics.Results[i].Data.Result
		if tp == MetricTypeVector {
			var instantMetrics []InstantMetric
			err := json.Unmarshal(r, &instantMetrics)
			if err != nil {
				fmt.Println(err.Error())
			}

			item := make(map[string][]TV)

			for i := 0; i < len(instantMetrics); i++ {
				metricInfo := instantMetrics[i].MetricInfo
				resourceName := metricInfo["resource_name"].(string)

				metricValue := instantMetrics[i].MetricValue
				t := metricValue[0].(float64)
				if t < float64(startTime) || t > float64(endTime) {
					continue
				}
				v := metricValue[1].(string)
				tv := TV{T: t, V: v}
				item[resourceName] = []TV{tv}
			}

			m = &ResourceMetrics{
				ResourceMetric: item,
				RuleName:       metricName,
			}

		} else if tp == MetricTypeMatrix {
			var timeRangeMetrics []TimeSeriesMetric
			err := json.Unmarshal(r, &timeRangeMetrics)
			if err != nil {
				glog.Errorln(err.Error())
			}

			item := make(map[string][]TV)

			for i := 0; i < len(timeRangeMetrics); i++ {
				metricItem := timeRangeMetrics[i]
				metricInfo := metricItem.MetricInfo
				resourceName := metricInfo["resource_name"].(string)
				metricValues := metricItem.MetricValues
				var tvs []TV
				for i := 0; i < len(metricValues); i++ {
					t := metricValues[i][0].(float64)
					if t < float64(startTime) || t > float64(endTime) {
						continue
					}
					v := metricValues[i][1].(string)
					tv := TV{T: t, V: v}
					tvs = append(tvs, tv)
				}
				if len(tvs) != 0 {
					item[resourceName] = tvs
				}
			}
			m = &ResourceMetrics{
				ResourceMetric: item,
				RuleName:       metricName,
			}
		}
	}
	return m
}
