package models

import (
	"fmt"
	"kubesphere.io/ks-alert/pkg/utils/jsonutil"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
	"time"
)

func TestGetAlertConfigRows(t *testing.T) {
	Convey("test get alert config", t, func() {
		Convey("test01", func() {
			maps, err := GetAlertConfigRows()
			fmt.Println(maps, err)
		})
	})
}

func TestGetAbnormalExecutedAlertConfig(t *testing.T) {
	Convey("test get alert config", t, func() {
		Convey("test01", func() {
			latestReportTime := time.Now().Add(-time.Minute * 3)
			hostID := "10.244.4.210:50001"
			limit := 3
			als, err := GetAbnormalExecutedAlertConfig(hostID, latestReportTime, limit)
			fmt.Println(err)
			for _, ac := range *als {
				fmt.Println(jsonutil.Marshal(ac))
			}
		})
	})
}

func TestUpdateAlertConfigBindingHostAndVersion(t *testing.T) {
	Convey("test update alert config info", t, func() {
		Convey("test01", func() {
			bools, err := UpdateAlertConfigBindingHostAndVersion(&[]AlertConfig{AlertConfig{
				AlertConfigID: "alert-config-658etrjgewteyh",
				HostID:        "yyyy:8080",
				Version:       3,
				KeepAliveAt:   time.Now(),
				UpdatedAt:     time.Now(),
			}})
			fmt.Println(bools)
			fmt.Println(err)
		})
	})
}
