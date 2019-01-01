package models

import (
	"fmt"
	"github.com/carmanzhang/ks-alert/pkg/dispatcher/pb"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestCreateSeverities(t *testing.T) {
	Convey("create severities", t, func() {

	})
}

func TestDeleteSeverity(t *testing.T) {
	Convey("delete severity", t, func() {
		x, err := DeleteSeverity(&pb.SeveritySpec{
			SeverityId: "severity-004wnqz602qnnn",
		})
		fmt.Println(x)
		fmt.Println(err)
	})
}

func TestGetSeverity(t *testing.T) {
	Convey("delete severity", t, func() {
		severity, err := GetSeverity(&pb.SeveritySpec{
			SeverityId: "severity-10npz0p3no5wwy",
			//ProductId:"string",
		})

		fmt.Println(err)
		fmt.Println(severity)
		//fmt.Println(len(*severity))
	})
}
