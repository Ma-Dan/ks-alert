package models

import (
	"fmt"
	"github.com/carmanzhang/ks-alert/pkg/pb"
	"github.com/carmanzhang/ks-alert/pkg/utils/jsonutil"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
	"time"
)

func TestAddSourceType(t *testing.T) {
	Convey("add Source Type fo a product", t, func() {
		// get product id
		var err error
		ent, err := GetEnterprise(&Enterprise{EnterpriseID: "enterprise-6y19xy9pwm24oo"})
		So(err, ShouldBeNil)

		product, err := GetProduct(&Product{EnterpriseID: ent.EnterpriseID})
		So(err, ShouldBeNil)

		productID := (*product).ProductID
		fmt.Println(productID)
		// cluster
		_, err = CreateResourceType(&ResourceType{
			CreatedAt:         time.Now(),
			UpdatedAt:         time.Now(),
			ProductID:         "product-m57o28v7101rwz",
			Description:       "",
			Enable:            true,
			ResourceTypeName:  "cluster",
			MonitorCenterPort: 8087,
			MonitorCenterHost: "http://139.198.190.141",
			ResourceURITmpls: jsonutil.Marshal(&pb.ResourceUriTmpls{
				ResourceUriTmpl: []*pb.ResourceUriTmpl{
					&pb.ResourceUriTmpl{
						QueryParams:  "",
						PathParams:   map[string]string{},
						UriTmpl:      "/clusters",
						ResourceName: nil,
					},
				},
			}),
		})

		So(err, ShouldBeNil)

		// node
		_, err = CreateResourceType(&ResourceType{
			CreatedAt:         time.Now(),
			UpdatedAt:         time.Now(),
			ProductID:         "product-m57o28v7101rwz",
			Description:       "",
			Enable:            true,
			ResourceTypeName:  "node",
			MonitorCenterPort: 8087,
			MonitorCenterHost: "http://139.198.190.141",
			ResourceURITmpls: jsonutil.Marshal(&pb.ResourceUriTmpls{
				ResourceUriTmpl: []*pb.ResourceUriTmpl{
					&pb.ResourceUriTmpl{
						QueryParams:  "",
						PathParams:   map[string]string{},
						UriTmpl:      "/nodes",
						ResourceName: nil,
					},
				},
			}),
		})

		So(err, ShouldBeNil)

		// workspace
		_, err = CreateResourceType(&ResourceType{
			CreatedAt:         time.Now(),
			UpdatedAt:         time.Now(),
			ProductID:         "product-m57o28v7101rwz",
			Description:       "",
			Enable:            true,
			ResourceTypeName:  "workspace",
			MonitorCenterPort: 8087,
			MonitorCenterHost: "http://139.198.190.141",
			ResourceURITmpls: jsonutil.Marshal(&pb.ResourceUriTmpls{
				ResourceUriTmpl: []*pb.ResourceUriTmpl{
					&pb.ResourceUriTmpl{
						QueryParams:  "",
						PathParams:   map[string]string{},
						UriTmpl:      "/workspaces",
						ResourceName: nil,
					},
				},
			}),
		})

		So(err, ShouldBeNil)

		// namespace
		_, err = CreateResourceType(&ResourceType{
			CreatedAt:         time.Now(),
			UpdatedAt:         time.Now(),
			ProductID:         "product-m57o28v7101rwz",
			Description:       "",
			Enable:            true,
			ResourceTypeName:  "namespace",
			MonitorCenterPort: 8087,
			MonitorCenterHost: "http://139.198.190.141",
			ResourceURITmpls: jsonutil.Marshal(&pb.ResourceUriTmpls{
				ResourceUriTmpl: []*pb.ResourceUriTmpl{
					&pb.ResourceUriTmpl{
						QueryParams:  "",
						PathParams:   map[string]string{"ws_name": ""},
						UriTmpl:      "/workspaces/{ws_name}/namespaces",
						ResourceName: nil,
					},

					&pb.ResourceUriTmpl{
						QueryParams:  "",
						PathParams:   map[string]string{},
						UriTmpl:      "/namespaces",
						ResourceName: nil,
					},
				},
			}),
		})

		So(err, ShouldBeNil)

		// workload
		_, err = CreateResourceType(&ResourceType{
			CreatedAt:         time.Now(),
			UpdatedAt:         time.Now(),
			ProductID:         "product-m57o28v7101rwz",
			Description:       "",
			Enable:            true,
			ResourceTypeName:  "workload",
			MonitorCenterPort: 8087,
			MonitorCenterHost: "http://139.198.190.141",
			ResourceURITmpls: jsonutil.Marshal(&pb.ResourceUriTmpls{
				ResourceUriTmpl: []*pb.ResourceUriTmpl{
					&pb.ResourceUriTmpl{
						QueryParams:  "",
						PathParams:   map[string]string{"ns": "", "wl_kind": ""},
						UriTmpl:      "/namespaces/{ns}/{wl_kind}",
						ResourceName: nil,
					},
				},
			}),
		})

		So(err, ShouldBeNil)

		// pod
		_, err = CreateResourceType(&ResourceType{
			CreatedAt:         time.Now(),
			UpdatedAt:         time.Now(),
			ProductID:         "product-m57o28v7101rwz",
			Description:       "",
			Enable:            true,
			ResourceTypeName:  "pod",
			MonitorCenterPort: 8087,
			MonitorCenterHost: "http://139.198.190.141",
			ResourceURITmpls: jsonutil.Marshal(&pb.ResourceUriTmpls{
				ResourceUriTmpl: []*pb.ResourceUriTmpl{
					&pb.ResourceUriTmpl{
						QueryParams:  "",
						PathParams:   map[string]string{"ns_name": ""},
						UriTmpl:      "/namespaces/{ns_name}/pods",
						ResourceName: nil,
					},
					&pb.ResourceUriTmpl{
						QueryParams:  "",
						PathParams:   map[string]string{"node_id": ""},
						UriTmpl:      "/nodes/{node_id}/pods",
						ResourceName: nil,
					},
				},
			}),
		})

		So(err, ShouldBeNil)

		// container
		_, err = CreateResourceType(&ResourceType{
			CreatedAt:         time.Now(),
			UpdatedAt:         time.Now(),
			ProductID:         "product-m57o28v7101rwz",
			Description:       "",
			Enable:            true,
			ResourceTypeName:  "container",
			MonitorCenterPort: 8087,
			MonitorCenterHost: "http://139.198.190.141",
			ResourceURITmpls: jsonutil.Marshal(&pb.ResourceUriTmpls{
				ResourceUriTmpl: []*pb.ResourceUriTmpl{
					&pb.ResourceUriTmpl{
						QueryParams:  "",
						PathParams:   map[string]string{"ns_name": "", "pod_name": ""},
						UriTmpl:      "/namespaces/{ns_name}/pods/{pod_name}/containers",
						ResourceName: nil,
					},
					&pb.ResourceUriTmpl{
						QueryParams:  "",
						PathParams:   map[string]string{"node_id": "", "pod_name": ""},
						UriTmpl:      "/nodes/{node_id}/pods/{pod_name}/containers",
						ResourceName: nil,
					},
				},
			}),
		})

		So(err, ShouldBeNil)
	})
}
