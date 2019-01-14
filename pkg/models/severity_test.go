package models

import (
	"fmt"
	"github.com/carmanzhang/ks-alert/pkg/pb"
	"testing"
	"time"
)

func TestCreateSeverities(t *testing.T) {
	Convey("create severities", t, func() {
		_, err := CreateSeverity(&Severity{
			ProductID:  "product-m57o28v7101rwz",
			SeverityCh: "严重",
			SeverityEn: "Critical",
			UpdatedAt:  time.Now(),
			CreatedAt:  time.Now(),
		})

		So(err, ShouldBeNil)

		_, err = CreateSeverity(&Severity{
			ProductID:  "product-m57o28v7101rwz",
			SeverityCh: "较危险",
			SeverityEn: "Major",
			UpdatedAt:  time.Now(),
			CreatedAt:  time.Now(),
		})

		So(err, ShouldBeNil)

		_, err = CreateSeverity(&Severity{
			ProductID:  "product-m57o28v7101rwz",
			SeverityCh: "危险",
			SeverityEn: "Minor",
			UpdatedAt:  time.Now(),
			CreatedAt:  time.Now(),
		})
		So(err, ShouldBeNil)

		_, err = CreateSeverity(&Severity{
			ProductID:  "product-m57o28v7101rwz",
			SeverityCh: "警告",
			SeverityEn: "Warn",
			UpdatedAt:  time.Now(),
			CreatedAt:  time.Now(),
		})
		So(err, ShouldBeNil)
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
