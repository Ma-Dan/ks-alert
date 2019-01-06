package runtime

import (
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
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
