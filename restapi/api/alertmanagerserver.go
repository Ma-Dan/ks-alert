package main

import (
	"github.com/emicklei/go-restful"
	"github.com/emicklei/go-restful/log"
	"github.com/emicklei/go-restful-openapi"
	"net/http"
	"github.com/go-openapi/spec"
)

func (u MonitorResource) sayHello(request *restful.Request, response *restful.Response) {
	name := request.PathParameter("name")
	response.WriteAsJson("hi, " + name)
}

func (u MonitorResource) sayBye(request *restful.Request, response *restful.Response) {
	name := request.PathParameter("name")
	response.WriteAsJson("bye, " + name)

}

type MonitorResource struct {
}

func main() {
	u := MonitorResource{}
	restful.DefaultContainer.Add(u.WebService())
	restful.TraceLogger(log.Logger)
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
		WebServices:                   restful.RegisteredWebServices(), // you control what services are visible
		APIPath:                       "/apidocs.json",
		PostBuildSwaggerObjectHandler: enrichSwaggerObject}
	restful.DefaultContainer.Add(restfulspec.NewOpenAPIService(config))
	// Open http://localhost:8080/apidocs/?url=http://localhost:8080/apidocs.json
	http.Handle("/apidocs/", http.StripPrefix("/apidocs/", http.FileServer(http.Dir("C:/Users/Carman/Desktop/eventhub-dev/prometheus-k8s-restapi/swagger-ui/dist"))))
}

func enrichSwaggerObject(swo *spec.Swagger) {
	swo.Info = &spec.Info{
		InfoProps: spec.InfoProps{
			Title:       "kubesphere alertmanager restful apis",
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
		Path("/v1/monitoring").
		Consumes(restful.MIME_XML, restful.MIME_JSON).
		Produces(restful.MIME_JSON, restful.MIME_XML)

	tags := []string{"monitoring apis"}

	myws := MyWebService{ws}
	myws.Route(myws.GET("/hi/{name}").To(u.sayHello).
		Doc("test01").
		Param(ws.PathParameter("name", "your name").DataType("string").Required(true).DefaultValue("carman")).
		Metadata(restfulspec.KeyOpenAPITags, tags)).
		Produces(restful.MIME_JSON)

	myws.Route(ws.GET("/bye/{name}").To(u.sayBye).
		Doc("test01").
		Param(ws.PathParameter("name", "your name").DataType("string").Required(true).DefaultValue("carman")).
		Metadata(restfulspec.KeyOpenAPITags, tags)).
		Produces(restful.MIME_JSON)

	return ws
}

type MyWebService struct {
	*restful.WebService
}


