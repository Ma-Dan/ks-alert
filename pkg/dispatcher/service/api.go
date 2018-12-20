package service

import (
	"github.com/emicklei/go-restful"
	"github.com/emicklei/go-restful-openapi"
	"kubesphere.io/ks-alert/pkg/dispatcher/handler"
	"kubesphere.io/ks-alert/pkg/models"
)

type AlertAPI struct{}

func createAlert(request *restful.Request, response *restful.Response) {
	handler.CreateAlert(request, response)
}

func retrieveAlert(request *restful.Request, response *restful.Response) {
	handler.RetrieveAlert(request, response)
}

func updateAlert(request *restful.Request, response *restful.Response) {
	handler.UpdateAlert(request, response)
}

func deleteAlert(request *restful.Request, response *restful.Response) {
	handler.DeleteAlert(request, response)
}

func (u AlertAPI) WebService() *restful.WebService {
	ws := new(restful.WebService)
	ws.
		Path("/alert/v1").
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON)

	tags := []string{"alert apis"}

	ws.Route(ws.POST("/alerts").To(createAlert).
		Doc("create AlertConfig").
		Operation("create an AlertConfig operator").
		Reads(models.AlertConfig{}).
		Metadata(restfulspec.KeyOpenAPITags, tags))

	ws.Route(ws.GET("/alerts").To(retrieveAlert).
		Doc("retrieve AlertConfig").
		Operation("create an AlertConfig operator").
		//Reads(models.AlertConfig{}).
		Metadata(restfulspec.KeyOpenAPITags, tags))

	ws.Route(ws.PUT("/alerts").To(updateAlert).
		Doc("update AlertConfig").
		Operation("create an AlertConfig operator").
		//Reads(models.AlertConfig{}).
		Metadata(restfulspec.KeyOpenAPITags, tags))

	ws.Route(ws.DELETE("/alerts").To(deleteAlert).
		Doc("delete AlertConfig").
		Operation("create an AlertConfig operator").
		//Reads(models.AlertConfig{}).
		Metadata(restfulspec.KeyOpenAPITags, tags))
	//resource_group rule_group receiver_group
	//****************************************************************************************************
	//resource_group  crud
	ws.Route(ws.POST("/resource_group").To(createResourceGroup))
	ws.Route(ws.PUT("/resource_group").To(updateResourceGroup))
	ws.Route(ws.GET("/resource_group").To(retrieveResourceGroup))
	ws.Route(ws.DELETE("/resource_group").To(deleteResourceGroup))
	//****************************************************************************************************
	// receiver_group crud
	ws.Route(ws.POST("/receiver_group").To(createReceiverGroup))
	ws.Route(ws.PUT("/receiver_group").To(updateReceiverGroup))
	ws.Route(ws.GET("/receiver_group").To(retrieveReceiverGroup))
	ws.Route(ws.DELETE("/receiver_group").To(deleteReceiverGroup))
	//****************************************************************************************************
	// rule_group crud
	ws.Route(ws.POST("/rule_group").To(createRuleGroup))
	ws.Route(ws.PUT("/rule_group").To(updateRuleGroup))
	ws.Route(ws.GET("/rule_group").To(retrieveRuleGroup))
	ws.Route(ws.DELETE("/rule_group").To(deleteRuleGroup))
	//****************************************************************************************************
	// silence api
	ws.Route(ws.POST("/silence").To(createSilence))
	ws.Route(ws.PUT("/silence").To(updateSilence))
	ws.Route(ws.GET("/silence").To(retrieveSilence))
	ws.Route(ws.DELETE("/silence").To(deleteSilence))
	//****************************************************************************************************
	// enterprise crud
	ws.Route(ws.POST("/enterprise").To(createEnterprise))
	ws.Route(ws.PUT("/enterprise").To(updateEnterprise))
	ws.Route(ws.GET("/enterprise").To(retrieveEnterprise))
	ws.Route(ws.DELETE("/enterprise").To(deleteEnterprise))
	//****************************************************************************************************
	// product crud
	ws.Route(ws.POST("/product").To(createProduct))
	ws.Route(ws.PUT("/product").To(updateProduct))
	ws.Route(ws.GET("/product").To(retrieveProduct))
	ws.Route(ws.DELETE("/product").To(deleteProduct))
	//****************************************************************************************************
	// resource_type crud
	ws.Route(ws.POST("/resource_types").To(createResourceTypes))
	ws.Route(ws.PUT("/resource_types").To(updateResourceTypes))
	ws.Route(ws.GET("/resource_types").To(retrieveResourceTypes))
	ws.Route(ws.DELETE("/resource_types").To(deleteResourceTypes))
	//****************************************************************************************************
	// resource_type crud
	ws.Route(ws.POST("/metrics").To(createMetrics))
	ws.Route(ws.PUT("/metrics").To(updateMetrics))
	ws.Route(ws.GET("/metrics").To(retrieveMetrics))
	ws.Route(ws.DELETE("/metrics").To(deleteMetrics))
	//****************************************************************************************************
	// alert_rule crud
	ws.Route(ws.POST("/alert_rules").To(createAlertRules))
	ws.Route(ws.PUT("/alert_rules").To(updateAlertRules))
	ws.Route(ws.GET("/alert_rules").To(retrieveAlertRules))
	ws.Route(ws.DELETE("/alert_rules").To(deleteAlertRules))
	//// current fired alert
	//ws.Route(ws.GET("/alerts/fired")
	//
	//// alert history(resolved alert, include start-time and  end-time)
	//ws.Route(ws.GET("/alerts/history")
	return ws
}
