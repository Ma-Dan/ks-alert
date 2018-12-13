package api

import (
	"github.com/emicklei/go-restful"
	"github.com/emicklei/go-restful-openapi"
	"github.com/emicklei/go-restful/log"
	"github.com/go-openapi/spec"
	"net/http"
)

func (u MonitorResource) senderAlertConfig(request *restful.Request, response *restful.Response) {
	//dbutil.SenderAlertConfig(request, response)
}

func (u MonitorResource) sayBye(request *restful.Request, response *restful.Response) {
	name := request.PathParameter("name")
	response.WriteAsJson("bye, " + name)

}

type MonitorResource struct {
}

func Run() {
	u := MonitorResource{}
	restful.DefaultContainer.Add(u.WebService())
	handleSwagger()
	enableCORS()

	log.Printf("Get the API using http://localhost:8080/apidocs.json")
	log.Printf("Open Swagger UI using http://localhost:8080/apidocs/") // ?url=http://localhost:8080/apidocs.json
	log.Print(http.ListenAndServe(":8080", nil))
}

func enableCORS() {
	// Optionally, you may need to enable CORS for the UI to work.
	cors := restful.CrossOriginResourceSharing{
		AllowedHeaders: []string{"Content-Type", "Accept"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
		CookiesAllowed: false,
		Container:      restful.DefaultContainer}
	restful.DefaultContainer.Filter(cors.Filter)
}

func handleSwagger() {
	config := restfulspec.Config{
		WebServices: restful.RegisteredWebServices(), // you control what services are visible
		APIPath:     "/apidocs.json",
		PostBuildSwaggerObjectHandler: enrichSwaggerObject}
	restful.DefaultContainer.Add(restfulspec.NewOpenAPIService(config))
	// Open http://localhost:8080/apidocs/?url=http://localhost:8080/apidocs.json
	// C:\Users\Carman\go\src\kubesphere.io\alert-kubesphere-plugin\swagger-ui
	http.Handle("/apidocs/", http.StripPrefix("/apidocs/", http.FileServer(http.Dir("C:/Users/Carman/go/src/kubesphere.io/alert-kubesphere-plugin/swagger-ui/dist"))))
}

func enrichSwaggerObject(swo *spec.Swagger) {
	swo.Info = &spec.Info{
		InfoProps: spec.InfoProps{
			Title: "kubesphere alertmanager restful apis",
			Contact: &spec.ContactInfo{
				Name:  "carman",
				Email: "carmanzhang@yunify.com",
				URL:   "",
			},
			License: &spec.License{
				Name: "MIT License",
				URL:  "http://mit.org",
			},
			Version: "1.0.0",
		},
	}
}

func (u MonitorResource) WebService() *restful.WebService {
	ws := new(restful.WebService)
	ws.
		Path("/v1/alert").
		Consumes(restful.MIME_XML, restful.MIME_JSON).
		Produces(restful.MIME_JSON, restful.MIME_XML)

	//tags := []string{"monitoring apis"}
	// resource_group rule_group receiver_group
	// ****************************************************************************************************
	// resource_group  crud
	//ws.Route(ws.PUT("/resource_group")
	//// receiver_group crud
	//ws.Route(ws.PUT("/receiver_group")
	//// rule_group crud
	//ws.Route(ws.PUT("/rule_group")
	//// ****************************************************************************************************
	//
	//// 这里的 alert 包含 resource_group  receiver_group  rule_group，如果已经创建，则使用已创建的，如果未创建，则通过 create alert api 也可以创建
	//// ****************************************************************************************************
	//// modify alert
	//ws.Route(ws.POST("/alerts/{alert_id}")
	//// create alert
	//ws.Route(ws.PUT("/alerts")
	//// delete alert
	//ws.Route(ws.DELETE("/alerts/{alert_id}")
	//// get alert config
	//ws.Route(ws.GET("/alerts/{alert_id}")
	//ws.Route(ws.GET("/alerts")
	//// get alert status
	//ws.Route(ws.GET("/alerts/{alert_id}/status")
	//// current fired alert
	//ws.Route(ws.GET("/alerts/fired")
	//// alert history(resolved alert, include start-time and  end-time)
	//ws.Route(ws.GET("/alerts/history")
	//// ****************************************************************************************************
	//
	//// silence api and repeat send
	//// ****************************************************************************************************
	//// create silence
	//ws.Route(ws.PUT("/silence")
	//// modify silence
	//ws.Route(ws.POST("/silence")
	//// get silence
	//ws.Route(ws.GET("/silence")
	//// delete silence
	//ws.Route(ws.DELETE("/silence")
	//
	//ws.Route(ws.POST("/repeat_send")
	//
	//// ****************************************************************************************************
	//
	//
	//// enterprise api
	//// ****************************************************************************************************
	//// enterprise register/modify/delete/get
	//ws.Route(ws.PUT("/enterprise")
	//// product register/modify/delete/get
	//ws.Route(ws.POST("/enterprise/{enterprise_id}/product")
	//// resource register/modify/delete/get
	//ws.Route(ws.POST("/enterprise/{enterprise_id}/product/{product_id}/resource")
	//// rule register/modify/delete/get
	//ws.Route(ws.POST("/enterprise/{enterprise_id}/product/{product_id}/resource/{resource_id}/rule")
	//// ****************************************************************************************************

	return ws
}
