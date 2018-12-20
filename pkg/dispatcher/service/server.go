package service

import (
	"github.com/emicklei/go-restful"
	"net/http"
	"github.com/emicklei/go-restful-openapi"
	"github.com/go-openapi/spec"
	"log"
	"flag"
)

var (
	ExecutorServiceName = flag.String("service", "alert_executor_service", "service name")
	EtcdAddr            = flag.String("etcd", "http://127.0.0.1:2379", "register etcd address")
)

func Run() {
	flag.Parse()

	api := AlertAPI{}
	restful.DefaultContainer.Add(api.WebService())
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
	http.Handle("/apidocs/", http.StripPrefix("/apidocs/", http.FileServer(http.Dir("C:/Users/Carman/go/src/kubesphere.io/ks-alert/pkg/dispatcher/swagger-ui/dist"))))
}

func enrichSwaggerObject(swo *spec.Swagger) {
	swo.Info = &spec.Info{
		InfoProps: spec.InfoProps{
			Title: "kubesphere alert restful apis",
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
