package handler

import (
	"github.com/emicklei/go-restful"
	"kubesphere.io/ks-alert/pkg/models"
	"k8s.io/klog/glog"
)

func CreateAlert(request *restful.Request, response *restful.Response) {

	var alertConfig models.AlertConfig
	err := request.ReadEntity(&alertConfig)
	if err != nil {
		glog.Errorln(err)
	}

	alertRuleGroup, err := models.CreateAlertRuleGroup(&alertConfig.AlertRuleGroup)

	// alertRules
	_, err = models.CreateAlertRules(&alertConfig.AlertRuleGroup.AlertRules, alertRuleGroup.AlertRuleGroupID)
	if err != nil {
		glog.Errorln(err)
	}

	resourceGroup, err := models.CreateResourceGroup(alertConfig.ResourceGroup.ResourceGroupName, alertConfig.ResourceGroup.Description)
	if err != nil {
		glog.Errorln(err)
	}

	err = models.CreateResources(&alertConfig.ResourceGroup.Resources, resourceGroup, &alertConfig.URIParams)
	if err != nil {
		glog.Errorln(err)
	}

	receiverGroup, err := models.CreateReceiverGroup(&alertConfig.ReceiverGroup)
	if err != nil {
		glog.Errorln(err)
	}

	receivers, err := models.CreateReceivers(&alertConfig.ReceiverGroup.Receivers)
	if err != nil {
		glog.Errorln(err)
	}

	err = models.CreateReceiverBindingGroupItem(receivers, receiverGroup)
	if err != nil {
		glog.Errorln(err)
	}

}

func RetrieveAlert(request *restful.Request, response *restful.Response) {

}

func UpdateAlert(request *restful.Request, response *restful.Response) {

}

func DeleteAlert(request *restful.Request, response *restful.Response) {

}
