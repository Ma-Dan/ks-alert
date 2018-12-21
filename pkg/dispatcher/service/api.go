package service

import (
	"github.com/emicklei/go-restful"
	"github.com/emicklei/go-restful-openapi"
	"kubesphere.io/ks-alert/pkg/dispatcher/handler"
	"kubesphere.io/ks-alert/pkg/models"
)

type AlertAPI struct{}

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

	ws.Route(ws.GET("/alerts").To(handler.RetrieveAlert).
		Doc("retrieve AlertConfig").
		Param(restful.QueryParameter("alert_id", "get alert config by id")).
		Writes(models.AlertConfig{}).
		Operation("retrieve an AlertConfig operator").
		Metadata(restfulspec.KeyOpenAPITags, tags))

	ws.Route(ws.PUT("/alerts").To(handler.UpdateAlert).
		Doc("update AlertConfig").
		Operation("create an AlertConfig operator").
		Reads(models.AlertConfig{}).
		Writes(models.Error{}).
		Metadata(restfulspec.KeyOpenAPITags, tags))

	ws.Route(ws.DELETE("/alerts").To(handler.DeleteAlert).
		Doc("delete AlertConfig").
		Operation("delete an AlertConfig operator").
		Param(restful.QueryParameter("alert_id", "delete alert config by id")).
		Writes(models.Error{}).
		Metadata(restfulspec.KeyOpenAPITags, tags))
	//resource_group rule_group receiver_group
	//****************************************************************************************************
	//resource_group  crud
	//ws.Route(ws.POST("/resource_group").To(handler.CreateResourceGroup))
	//ws.Route(ws.PUT("/resource_group").To(handler.UpdateResourceGroup))
	//ws.Route(ws.GET("/resource_group").To(handler.RetrieveResourceGroup))
	//ws.Route(ws.DELETE("/resource_group").To(handler.DeleteResourceGroup))
	////****************************************************************************************************
	//// receiver_group crud
	//ws.Route(ws.POST("/receiver_group").To(handler.CreateReceiverGroup))
	//ws.Route(ws.PUT("/receiver_group").To(handler.UpdateReceiverGroup))
	//ws.Route(ws.GET("/receiver_group").To(handler.RetrieveReceiverGroup))
	//ws.Route(ws.DELETE("/receiver_group").To(handler.DeleteReceiverGroup))
	////****************************************************************************************************
	//// rule_group crud
	//ws.Route(ws.POST("/rule_group").To(handler.CreateRuleGroup))
	//ws.Route(ws.PUT("/rule_group").To(handler.UpdateRuleGroup))
	//ws.Route(ws.GET("/rule_group").To(handler.RetrieveRuleGroup))
	//ws.Route(ws.DELETE("/rule_group").To(handler.DeleteRuleGroup))
	////****************************************************************************************************
	//// silence api
	//ws.Route(ws.POST("/silence").To(handler.CreateSilence))
	//ws.Route(ws.PUT("/silence").To(handler.UpdateSilence))
	//ws.Route(ws.GET("/silence").To(handler.RetrieveSilence))
	//ws.Route(ws.DELETE("/silence").To(handler.DeleteSilence))

	//****************************************************************************************************
	// enterprise crud
	tags = []string{"enterprise apis"}

	ws.Route(ws.POST("/enterprise").To(handler.CreateEnterprise).
		Doc("create enterprise").
		Reads(models.Enterprise{}).
		Writes(models.Error{}).
		Metadata(restfulspec.KeyOpenAPITags, tags))

	ws.Route(ws.GET("/enterprise").To(handler.RetrieveEnterprise).
		Doc("retrieve enterprise").
		Param(restful.QueryParameter("enterprise_id", "get enterprise by id")).
		Writes(models.Enterprise{}).
		Metadata(restfulspec.KeyOpenAPITags, tags))

	ws.Route(ws.PUT("/enterprise").To(handler.UpdateEnterprise).
		Doc("update enterprise").
		Reads(models.Enterprise{}).
		Writes(models.Error{}).
		Metadata(restfulspec.KeyOpenAPITags, tags))

	ws.Route(ws.DELETE("/enterprise").To(handler.DeleteEnterprise).
		Doc("delete enterprise").
		Param(restful.QueryParameter("enterprise_id", "delete enterprise by id")).
		Writes(models.Error{}).
		Metadata(restfulspec.KeyOpenAPITags, tags))

	//****************************************************************************************************
	// product crud
	tags = []string{"product apis"}

	ws.Route(ws.POST("/product").To(handler.CreateProduct).
		Doc("create product").
		Reads(models.Product{}).
		Writes(models.Error{}).
		Metadata(restfulspec.KeyOpenAPITags, tags))

	ws.Route(ws.GET("/product").To(handler.RetrieveProduct).
		Doc("retrieve product").
		Param(restful.QueryParameter("product_id", "get product by id")).
		Writes(models.Product{}).
		Metadata(restfulspec.KeyOpenAPITags, tags))

	ws.Route(ws.PUT("/product").To(handler.UpdateProduct).
		Doc("update product").
		Reads(models.Product{}).
		Writes(models.Error{}).
		Metadata(restfulspec.KeyOpenAPITags, tags))

	ws.Route(ws.DELETE("/product").To(handler.DeleteProduct).
		Doc("delete product").
		Param(restful.QueryParameter("product_id", "delete product by id")).
		Writes(models.Error{}).
		Metadata(restfulspec.KeyOpenAPITags, tags))

	//****************************************************************************************************
	// resource_type crud
	tags = []string{"resource type apis"}

	ws.Route(ws.POST("/resource_type").To(handler.CreateResourceType).
		Doc("create resource_type").
		Reads(models.ResourceType{}).
		Writes(models.Error{}).
		Metadata(restfulspec.KeyOpenAPITags, tags))

	ws.Route(ws.GET("/resource_type").To(handler.RetrieveResourceType).
		Doc("retrieve resource_type").
		Param(restful.QueryParameter("resource_type_id", "get resource_type by id")).
		Writes(models.ResourceType{}).
		Metadata(restfulspec.KeyOpenAPITags, tags))

	ws.Route(ws.PUT("/resource_type").To(handler.UpdateResourceType).
		Doc("update resource_type").
		Reads(models.ResourceType{}).
		Writes(models.Error{}).
		Metadata(restfulspec.KeyOpenAPITags, tags))

	ws.Route(ws.DELETE("/resource_type").To(handler.DeleteResourceType).
		Doc("delete resource_type").
		Param(restful.QueryParameter("resource_type_id", "delete resource_type by id")).
		Writes(models.Error{}).
		Metadata(restfulspec.KeyOpenAPITags, tags))

	//****************************************************************************************************
	// resource_type crud
	tags = []string{"metric apis"}

	ws.Route(ws.POST("/metric").To(handler.CreateMetric).
		Doc("create metric").
		Reads(models.Metric{}).
		Writes(models.Error{}).
		Metadata(restfulspec.KeyOpenAPITags, tags))

	ws.Route(ws.GET("/metric").To(handler.RetrieveMetric).
		Doc("retrieve metric").
		Param(restful.QueryParameter("metric_id", "get metric by id")).
		Writes(models.Metric{}).
		Metadata(restfulspec.KeyOpenAPITags, tags))

	ws.Route(ws.PUT("/metric").To(handler.UpdateMetric).
		Doc("update metric").
		Reads(models.Metric{}).
		Writes(models.Error{}).
		Metadata(restfulspec.KeyOpenAPITags, tags))

	ws.Route(ws.DELETE("/metric").To(handler.DeleteMetric).
		Doc("delete metric").
		Param(restful.QueryParameter("metric_id", "delete metric by id")).
		Writes(models.Error{}).
		Metadata(restfulspec.KeyOpenAPITags, tags))

	//****************************************************************************************************
	// alert_rule crud
	tags = []string{"alert rule apis"}

	ws.Route(ws.POST("/alert_rule").To(handler.CreateAlertRule).
		Doc("create alert_rule").
		Reads(models.AlertRule{}).
		Writes(models.Error{}).
		Metadata(restfulspec.KeyOpenAPITags, tags))

	ws.Route(ws.GET("/alert_rule").To(handler.RetrieveAlertRule).
		Doc("retrieve alert_rule").
		Param(restful.QueryParameter("alert_rule_id", "get alert_rule by id")).
		Writes(models.AlertRule{}).
		Metadata(restfulspec.KeyOpenAPITags, tags))

	ws.Route(ws.PUT("/alert_rule").To(handler.UpdateAlertRule).
		Doc("update alert_rule").
		Reads(models.AlertRule{}).
		Writes(models.Error{}).
		Metadata(restfulspec.KeyOpenAPITags, tags))

	ws.Route(ws.DELETE("/alert_rule").To(handler.DeleteAlertRule).
		Doc("delete alert_rule").
		Param(restful.QueryParameter("alert_rule_id", "delete alert_rule by id")).
		Writes(models.Error{}).
		Metadata(restfulspec.KeyOpenAPITags, tags))

	//ws.Route(ws.POST("/alert_rules").To(handler.CreateAlertRules))
	//ws.Route(ws.PUT("/alert_rules").To(handler.UpdateAlertRules))
	//ws.Route(ws.//GET("/alert_rules").To(handler.RetrieveAlertRules))
	//ws.Route(ws.DELETE("/alert_rules").To(handler.DeleteAlertRules))

	//// current fired alert
	//ws.Route(ws.GET("/alerts/fired")
	//
	//// alert history(resolved alert, include start-time and  end-time)
	//ws.Route(ws.GET("/alerts/history")
	return ws
}
