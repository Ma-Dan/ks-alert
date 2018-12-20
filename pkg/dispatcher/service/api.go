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

func createEnterprise(request *restful.Request, response *restful.Response) {
	handler.CreateEnterprise(request, response)
}

func retrieveEnterprise(request *restful.Request, response *restful.Response) {
	handler.RetrieveEnterprise(request, response)
}

func updateEnterprise(request *restful.Request, response *restful.Response) {
	handler.UpdateEnterprise(request, response)
}

func deleteEnterprise(request *restful.Request, response *restful.Response) {
	handler.DeleteEnterprise(request, response)
}

func createProduct(request *restful.Request, response *restful.Response) {
	handler.CreateProduct(request, response)
}

func retrieveProduct(request *restful.Request, response *restful.Response) {
	handler.RetrieveProduct(request, response)
}

func updateProduct(request *restful.Request, response *restful.Response) {
	handler.UpdateProduct(request, response)
}

func deleteProduct(request *restful.Request, response *restful.Response) {
	handler.DeleteProduct(request, response)
}

func createResourceType(request *restful.Request, response *restful.Response) {
	handler.CreateResourceType(request, response)
}

func retrieveResourceType(request *restful.Request, response *restful.Response) {
	handler.RetrieveResourceType(request, response)
}

func updateResourceType(request *restful.Request, response *restful.Response) {
	handler.UpdateResourceType(request, response)
}

func deleteResourceType(request *restful.Request, response *restful.Response) {
	handler.DeleteResourceType(request, response)
}

func createMetric(request *restful.Request, response *restful.Response) {
	handler.CreateMetric(request, response)
}

func retrieveMetric(request *restful.Request, response *restful.Response) {
	handler.RetrieveMetric(request, response)
}

func updateMetric(request *restful.Request, response *restful.Response) {
	handler.UpdateMetric(request, response)
}

func deleteMetric(request *restful.Request, response *restful.Response) {
	handler.DeleteMetric(request, response)
}

func createAlertRule(request *restful.Request, response *restful.Response) {
	handler.CreateAlertRule(request, response)
}

func retrieveAlertRule(request *restful.Request, response *restful.Response) {
	handler.RetrieveAlertRule(request, response)
}

func updateAlertRule(request *restful.Request, response *restful.Response) {
	handler.UpdateAlertRule(request, response)
}

func deleteAlertRule(request *restful.Request, response *restful.Response) {
	handler.DeleteAlertRule(request, response)
}

func (u AlertAPI) WebService() *restful.WebService {
	ws := new(restful.WebService)
	ws.
		Path("/alert/v1").
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON)

	tags := []string{"alert apis"}

	ws.Route(ws.POST("/alerts").To(handler.CreateAlert).
		Doc("create AlertConfig").
		Operation("create an AlertConfig operator").
		Reads(models.AlertConfig{}).
		Writes(models.Error{}).
		Metadata(restfulspec.KeyOpenAPITags, tags))

	ws.Route(ws.GET("/alerts").To(retrieveAlert).
		Doc("retrieve AlertConfig").
		Param(restful.QueryParameter("alert_id", "get alert config by id")).
		Writes(models.AlertConfig{}).
		Operation("retrieve an AlertConfig operator").
		Metadata(restfulspec.KeyOpenAPITags, tags))

	ws.Route(ws.PUT("/alerts").To(updateAlert).
		Doc("update AlertConfig").
		Operation("create an AlertConfig operator").
		Reads(models.AlertConfig{}).
		Writes(models.Error{}).
		Metadata(restfulspec.KeyOpenAPITags, tags))

	ws.Route(ws.DELETE("/alerts").To(deleteAlert).
		Doc("delete AlertConfig").
		Operation("delete an AlertConfig operator").
		Param(restful.QueryParameter("alert_id", "delete alert config by id")).
		Writes(models.Error{}).
		Metadata(restfulspec.KeyOpenAPITags, tags))
	//resource_group rule_group receiver_group
	//****************************************************************************************************
	//resource_group  crud
	//ws.Route(ws.POST("/resource_group").To(createResourceGroup))
	//ws.Route(ws.PUT("/resource_group").To(updateResourceGroup))
	//ws.Route(ws.GET("/resource_group").To(retrieveResourceGroup))
	//ws.Route(ws.DELETE("/resource_group").To(deleteResourceGroup))
	////****************************************************************************************************
	//// receiver_group crud
	//ws.Route(ws.POST("/receiver_group").To(createReceiverGroup))
	//ws.Route(ws.PUT("/receiver_group").To(updateReceiverGroup))
	//ws.Route(ws.GET("/receiver_group").To(retrieveReceiverGroup))
	//ws.Route(ws.DELETE("/receiver_group").To(deleteReceiverGroup))
	////****************************************************************************************************
	//// rule_group crud
	//ws.Route(ws.POST("/rule_group").To(createRuleGroup))
	//ws.Route(ws.PUT("/rule_group").To(updateRuleGroup))
	//ws.Route(ws.GET("/rule_group").To(retrieveRuleGroup))
	//ws.Route(ws.DELETE("/rule_group").To(deleteRuleGroup))
	////****************************************************************************************************
	//// silence api
	//ws.Route(ws.POST("/silence").To(createSilence))
	//ws.Route(ws.PUT("/silence").To(updateSilence))
	//ws.Route(ws.GET("/silence").To(retrieveSilence))
	//ws.Route(ws.DELETE("/silence").To(deleteSilence))

	//****************************************************************************************************
	// enterprise crud
	tags = []string{"enterprise apis"}

	ws.Route(ws.POST("/enterprise").To(createEnterprise).
		Doc("create enterprise").
		Reads(models.Enterprise{}).
		Writes(models.Error{}).
		Metadata(restfulspec.KeyOpenAPITags, tags))

	ws.Route(ws.GET("/enterprise").To(retrieveEnterprise).
		Doc("retrieve enterprise").
		Param(restful.QueryParameter("enterprise_id", "get enterprise by id")).
		Writes(models.Enterprise{}).
		Metadata(restfulspec.KeyOpenAPITags, tags))

	ws.Route(ws.PUT("/enterprise").To(updateEnterprise).
		Doc("update enterprise").
		Reads(models.Enterprise{}).
		Writes(models.Error{}).
		Metadata(restfulspec.KeyOpenAPITags, tags))

	ws.Route(ws.DELETE("/enterprise").To(deleteEnterprise).
		Doc("delete enterprise").
		Param(restful.QueryParameter("enterprise_id", "delete enterprise by id")).
		Writes(models.Error{}).
		Metadata(restfulspec.KeyOpenAPITags, tags))

	//****************************************************************************************************
	// product crud
	tags = []string{"product apis"}

	ws.Route(ws.POST("/product").To(createProduct).
		Doc("create product").
		Reads(models.Product{}).
		Writes(models.Error{}).
		Metadata(restfulspec.KeyOpenAPITags, tags))

	ws.Route(ws.GET("/product").To(retrieveProduct).
		Doc("retrieve product").
		Param(restful.QueryParameter("product_id", "get product by id")).
		Writes(models.Product{}).
		Metadata(restfulspec.KeyOpenAPITags, tags))

	ws.Route(ws.PUT("/product").To(updateProduct).
		Doc("update product").
		Reads(models.Product{}).
		Writes(models.Error{}).
		Metadata(restfulspec.KeyOpenAPITags, tags))

	ws.Route(ws.DELETE("/product").To(deleteProduct).
		Doc("delete product").
		Param(restful.QueryParameter("product_id", "delete product by id")).
		Writes(models.Error{}).
		Metadata(restfulspec.KeyOpenAPITags, tags))

	//****************************************************************************************************
	// resource_type crud
	tags = []string{"resource type apis"}

	ws.Route(ws.POST("/resource_type").To(createResourceType).
		Doc("create resource_type").
		Reads(models.ResourceType{}).
		Writes(models.Error{}).
		Metadata(restfulspec.KeyOpenAPITags, tags))

	ws.Route(ws.GET("/resource_type").To(retrieveResourceType).
		Doc("retrieve resource_type").
		Param(restful.QueryParameter("resource_type_id", "get resource_type by id")).
		Writes(models.ResourceType{}).
		Metadata(restfulspec.KeyOpenAPITags, tags))

	ws.Route(ws.PUT("/resource_type").To(updateResourceType).
		Doc("update resource_type").
		Reads(models.ResourceType{}).
		Writes(models.Error{}).
		Metadata(restfulspec.KeyOpenAPITags, tags))

	ws.Route(ws.DELETE("/resource_type").To(deleteResourceType).
		Doc("delete resource_type").
		Param(restful.QueryParameter("resource_type_id", "delete resource_type by id")).
		Writes(models.Error{}).
		Metadata(restfulspec.KeyOpenAPITags, tags))

	//****************************************************************************************************
	// resource_type crud
	tags = []string{"metric apis"}

	ws.Route(ws.POST("/metric").To(createMetric).
		Doc("create metric").
		Reads(models.Metric{}).
		Writes(models.Error{}).
		Metadata(restfulspec.KeyOpenAPITags, tags))

	ws.Route(ws.GET("/metric").To(retrieveMetric).
		Doc("retrieve metric").
		Param(restful.QueryParameter("metric_id", "get metric by id")).
		Writes(models.Metric{}).
		Metadata(restfulspec.KeyOpenAPITags, tags))

	ws.Route(ws.PUT("/metric").To(updateMetric).
		Doc("update metric").
		Reads(models.Metric{}).
		Writes(models.Error{}).
		Metadata(restfulspec.KeyOpenAPITags, tags))

	ws.Route(ws.DELETE("/metric").To(deleteMetric).
		Doc("delete metric").
		Param(restful.QueryParameter("metric_id", "delete metric by id")).
		Writes(models.Error{}).
		Metadata(restfulspec.KeyOpenAPITags, tags))

	//****************************************************************************************************
	// alert_rule crud
	tags = []string{"alert rule apis"}

	ws.Route(ws.POST("/alert_rule").To(createAlertRule).
		Doc("create alert_rule").
		Reads(models.AlertRule{}).
		Writes(models.Error{}).
		Metadata(restfulspec.KeyOpenAPITags, tags))

	ws.Route(ws.GET("/alert_rule").To(retrieveAlertRule).
		Doc("retrieve alert_rule").
		Param(restful.QueryParameter("alert_rule_id", "get alert_rule by id")).
		Writes(models.AlertRule{}).
		Metadata(restfulspec.KeyOpenAPITags, tags))

	ws.Route(ws.PUT("/alert_rule").To(updateAlertRule).
		Doc("update alert_rule").
		Reads(models.AlertRule{}).
		Writes(models.Error{}).
		Metadata(restfulspec.KeyOpenAPITags, tags))

	ws.Route(ws.DELETE("/alert_rule").To(deleteAlertRule).
		Doc("delete alert_rule").
		Param(restful.QueryParameter("alert_rule_id", "delete alert_rule by id")).
		Writes(models.Error{}).
		Metadata(restfulspec.KeyOpenAPITags, tags))

	//ws.Route(ws.POST("/alert_rules").To(createAlertRules))
	//ws.Route(ws.PUT("/alert_rules").To(updateAlertRules))
	//ws.Route(ws.//GET("/alert_rules").To(retrieveAlertRules))
	//ws.Route(ws.DELETE("/alert_rules").To(deleteAlertRules))

	//// current fired alert
	//ws.Route(ws.GET("/alerts/fired")
	//
	//// alert history(resolved alert, include start-time and  end-time)
	//ws.Route(ws.GET("/alerts/history")
	return ws
}
