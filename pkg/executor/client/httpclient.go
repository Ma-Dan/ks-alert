package client

import (
	"fmt"
	"github.com/golang/glog"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

var client = &http.Client{}

const (
	DefaultScheme = "http"
)

func SendMonitoringRequest(uri string, resources []string, metricName string, startTime int64, endTime int64, stepInMinute int32) string {
	startTs := strconv.FormatInt(startTime, 10)
	endTs := strconv.FormatInt(endTime, 10)
	step := fmt.Sprintf("%dm", stepInMinute)

	var parmas = url.Values{
		//"metrics_filter":   []string{strings.Join(metrics, "|")},
		"metrics_filter":   []string{metricName},
		"resources_filter": []string{strings.Join(resources, "|")},
		"start":            []string{startTs},
		"end":              []string{endTs},
		"step":             []string{step},
	}

	var urlStr string
	if strings.HasPrefix(uri, "http://") || strings.HasPrefix(uri, "https://") {
		urlStr = uri + "?" + parmas.Encode()

	} else {
		urlStr = DefaultScheme + "://" + uri + "?" + parmas.Encode()
	}

	glog.Info(urlStr)

	response, err := client.Get(urlStr)
	if err != nil {
		glog.Error(err)
	} else {
		defer response.Body.Close()

		contents, err := ioutil.ReadAll(response.Body)

		if err != nil {
			glog.Errorln(err.Error())
		}

		return string(contents)
	}
	return ""
}
