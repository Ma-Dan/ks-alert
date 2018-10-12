package client

import (
	"github.com/emicklei/go-restful"
	"alert-kubesphere-plugin/pkg/models"
	"github.com/golang/glog"
	"fmt"
)

func SenderAlertConfig(request *restful.Request, response *restful.Response) {

	var resEntity models.UserRequest
	err := request.ReadEntity(&resEntity)
	if err != nil {
		glog.Errorln(err)
	}

	fmt.Println(resEntity)
	response.WriteAsJson(resEntity)
}
