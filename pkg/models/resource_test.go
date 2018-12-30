package models

//
//import (
//	"encoding/json"
//	"fmt"
//	. "github.com/smartystreets/goconvey/convey"
//	"github.com/carmanzhang/ks-alert/pkg/utils/jsonutil"
//	"testing"
//)
//
//func TestAddSourceType(t *testing.T) {
//	Convey("add Source Type fo a product", t, func() {
//		// get product id
//		var err error
//		ent, err := GetEnterprise(&Enterprise{EnterpriseName: "北京优帆科技有限公司武汉分公司"})
//		So(err, ShouldBeNil)
//
//		product, err := GetProduct(&Product{EnterpriseID: ent.EnterpriseID})
//		So(err, ShouldBeNil)
//
//		productID := (*product).ProductID
//
//		err = CreateSourceType(productID, "clusters", "", true, &ResourceURITmpls{
//			ResourceURITmpl: []ResourceURITmpl{
//				{
//					URI: "clusters",
//				},
//			},
//		})
//		So(err, ShouldBeNil)
//
//		err = CreateSourceType(productID, "nodes", "", true, &ResourceURITmpls{
//			ResourceURITmpl: []ResourceURITmpl{
//				{
//					URI: "nodes",
//				},
//			},
//		})
//		So(err, ShouldBeNil)
//
//		err = CreateSourceType(productID, "workspaces", "", true, &ResourceURITmpls{
//			ResourceURITmpl: []ResourceURITmpl{
//				{
//					URI: "workspaces",
//				},
//				{
//					Params: map[string]string{"ws_name": ""},
//					URI:    "workspaces/{ws_name}",
//				},
//			},
//		})
//		So(err, ShouldBeNil)
//
//		err = CreateSourceType(productID, "namespaces", "", true, &ResourceURITmpls{
//			ResourceURITmpl: []ResourceURITmpl{
//				{
//					URI: "namespaces",
//				},
//				{
//					Params: map[string]string{"ns_name": ""},
//					URI:    "namespaces/{ns_name}",
//				},
//			},
//		})
//		So(err, ShouldBeNil)
//
//		err = CreateSourceType(productID, "pods", "", true, &ResourceURITmpls{
//			ResourceURITmpl: []ResourceURITmpl{
//				{
//					Params: map[string]string{"ns_name": ""},
//					URI:    "namespces/{ns_name}/pods",
//				},
//
//				{
//					Params: map[string]string{"ns_name": "", "pod_name": ""},
//					URI:    "namespces/{ns_name}/pods/{pod_name}",
//				},
//
//				{
//					Params: map[string]string{"node_id": ""},
//					URI:    "nodes/{node_id}/pods",
//				},
//
//				{
//					Params: map[string]string{"node_id": "", "pod_name": ""},
//					URI:    "nodes/{node_id}/pods/{pod_name}",
//				},
//			},
//		})
//		So(err, ShouldBeNil)
//
//		err = CreateSourceType(productID, "containers", "", true, &ResourceURITmpls{
//			ResourceURITmpl: []ResourceURITmpl{
//				{
//					Params: map[string]string{"ns_name": "", "pod_name": ""},
//					URI:    "namespces/{ns_name}/pods/{pod_name}/containers",
//				},
//
//				{
//					Params: map[string]string{"ns_name": "", "pod_name": "", "container_name": ""},
//					URI:    "namespces/{ns_name}/pods/{pod_name}/containers/{container_name}",
//				},
//
//				{
//					Params: map[string]string{"node_id": "", "pod_name": ""},
//					URI:    "nodes/{node_id}/pods/{pod_name}/containers",
//				},
//
//				{
//					Params: map[string]string{"node_id": "", "pod_name": "", "container_name": ""},
//					URI:    "nodes/{node_id}/pods/{pod_name}/containers/{container_name}",
//				},
//			},
//		})
//		So(err, ShouldBeNil)
//	})
//}
//
//func TestGetResourceTypeByProductID(t *testing.T) {
//	Convey("add Source Type fo a product", t, func() {
//
//		productID := "product-4llxr47k7q82wz"
//		resourceTypes := GetResourceType(productID)
//
//		for _, rt := range *resourceTypes {
//			endpoint := rt.ResourceURITmpls
//			var resourceURITmpls ResourceURITmpls
//			jsonutil.Unmarshal(endpoint, &resourceURITmpls)
//			if resourceTypes != nil {
//				fmt.Println(resourceURITmpls)
//			}
//		}
//	})
//}
//
////func TestCreateResourceGroup(t *testing.T) {
////	Convey("a group of nodes", t, func() {
////		CreateResourceGroup("carman_nodes_group", "description0")
////	})
////	/*
////	&ResourceURITmpls{
////				ResourceURITmpl: []ResourceURITmpl{
////					{
////						Params:    map[string]string{},
////						Resources: []string{"i-0jk3bsyk", "i-k89a62il", "i-obveybo3", "i-zteam1zt"},
////					},
////				},
////			}
////
////	,resourceTypeID := "resource_type-36586v48vk60kk"
////			&ResourceURITmpls{
////				ResourceURITmpl: []ResourceURITmpl{
////					{
////						Params:    map[string]string{"ns_name": "kubesphere-monitoring-system"},
////						Resources: []string{"prometheus-k8s-0", "kube-state-metrics-c95f7d66-wqps2", "node-exporter-4cn7k", "prometheus-operator-5f9c5d444f-k49z4"},
////					},
////				},
////			}
////	 */
////
////	Convey("a group of pods", t, func() {
////		CreateResourceGroup("caman_pods_group", "description1")
////	})
////
////}
//
//func TestAssembledResourceTmpls(t *testing.T) {
//	Convey("test resource url template", t, func() {
//		tmpls := ResourceURITmpls{
//			ResourceURITmpl: []ResourceURITmpl{
//				ResourceURITmpl{
//					Params: map[string]string{"ns_name": ""},
//					URI:    "namespces/{ns_name}/pods",
//				},
//
//				ResourceURITmpl{
//					Params: map[string]string{"node_id": "", "pod_name": ""},
//					URI:    "nodes/{node_id}/pods/{pod_name}",
//				},
//			},
//		}
//
//		bytes, err := json.Marshal(tmpls)
//
//		if err == nil {
//			fmt.Println(string(bytes))
//		}
//	})
//}
