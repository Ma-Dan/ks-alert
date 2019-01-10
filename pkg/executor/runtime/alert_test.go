package runtime

import (
	"fmt"
	"github.com/carmanzhang/ks-alert/pkg/models"
	"github.com/carmanzhang/ks-alert/pkg/utils/jsonutil"
	"testing"
)

func TestAssembeURLPrefix(t *testing.T) {
	Convey("test uri", t, func() {
		Convey("test uri", func() {
			params := map[string]string{
				"ws": "system-workspace",
				"ns": "kube-system",
				"wk": "calic",
			}

			uri, err := AssembeURLPrefix("localhost", 8080, "/workspaces/{ws}/namespaces/{ns}/workload/{wk}", params)
			fmt.Println(uri)
			if err != nil {
				fmt.Println(err)
			}
		})
	})
}

func TestIsMatch(t *testing.T) {
	Convey("test map matching", t, func() {
		Convey("test map matching01", func() {
			params01 := map[string]string{
				"ws": "system-workspace",
				"ns": "kube-system",
				"wk": "calic",
			}

			params02 := map[string]string{
				"ws": "system-workspace",
				"ns": "kube-system",
				"wk": "calic",
			}

			b := IsMatch(params01, params02)
			So(b, ShouldBeTrue)
		})

		Convey("test map matching02", func() {
			params01 := map[string]string{
				"ws": "system-workspace",
				"ns": "kube-system",
				"wk": "calic",
			}

			b := IsMatch(params01, nil)
			So(b, ShouldBeFalse)
		})

		Convey("test map matching03", func() {
			b := IsMatch(nil, nil)
			So(b, ShouldBeTrue)
		})

		Convey("test map matching04", func() {
			params01 := map[string]string{
				"ws": "system-workspace",
			}

			params02 := map[string]string{
				"ws": "system-workspace",
				"wk": "calic",
			}

			b := IsMatch(params01, params02)
			So(b, ShouldBeFalse)
		})
	})
}

func TestGetResourcesSpec(t *testing.T) {
	Convey("get spec resource", t, func() {
		Convey("test get resource", func() {
			fmt.Println(GetResourcesSpec(&models.ResourceGroup{
				ResourceGroupID: "resource_group-73xqmrxmwq59kk",
				ResourceTypeID:  "resource_type-4o3kmjz32k0zj4",
				Resources: []*models.Resource{
					{ResourceName: "i-k89a62il"},
					{ResourceName: "i-obveybo3"},
					{ResourceName: "i-zteam1zt"},
				},
				URIParams: `{}`,
			}))
		})
	})
}

func TestUnmarshalJson(t *testing.T) {
	Convey("test json", t, func() {
		Convey("test json01", func() {
			s := `{"resource_uri_tmpl":[{"uri_tmpl":"/api/v1alpha1/monitoring/nodes"}]}`
			var uriTmpls models.ResourceUriTmpls
			jsonutil.Unmarshal(s, &uriTmpls)
			fmt.Println(uriTmpls)
		})
	})
}
