package client

import (
	"github.com/golang/glog"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

var client = &http.Client{}

const DefaultScheme = "http"

func SendMonitoringRequest(uri string, resources []string, metrics []string) string {

	var parmas = url.Values{
		"metrics_filter":   []string{strings.Join(metrics, "|")},
		"resources_filter": []string{strings.Join(resources, "|")},
	}

	urlStr := DefaultScheme + "://" + uri + parmas.Encode()

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
