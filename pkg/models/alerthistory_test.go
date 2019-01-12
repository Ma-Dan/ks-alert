package models

import (
	"fmt"
	"github.com/carmanzhang/ks-alert/pkg/utils/jsonutil"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
	"time"
)

func TestCreateAlertHistory(t *testing.T) {
	Convey("test create alert history", t, func() {
		Convey("test01", func() {
			CreateAlertHistory(&AlertHistory{
				AlertConfigID: "xzzzzzzzzz",

				ReceiverGroupID:   "xxxxxxxx",
				ReceiverGroup:     "ReceiverGroup",
				ReceiverGroupName: "ReceiverGroupName",

				ResourceGroupID:   "yyyyyyy",
				ResourceGroupName: "dcccccccc",
				AlertedResource:   "AlertedResource",

				AlertRuleGroupID:   "bbbbbbbbb",
				TriggerMetricName:  "xsacfdv",
				AlertRuleGroupName: "cdsvfdvfdv",

				SeverityID: "fffffff",
				SeverityCh: "严重",

				RepeatSendType:            1,
				InitRepeatSendInterval:    2,
				MaxRepeatSendCount:        1,
				CumulateRepeatSendCount:   1,
				CurrentRepeatSendInterval: 12,

				SilenceStartAt: time.Now(),
				SilenceEndAt:   time.Now(),

				MetricData:      "yyyyyyyyy",
				AlertRecoveryAt: time.Now(),
				AlertFiredAt:    time.Now(),

				RequestNotificationStatus: "xdsfrv",
				NotificationSendAt:        time.Now(),
				CreatedAt:                 time.Now(),
				UpdatedAt:                 time.Now(),
			})
		})
	})
}

func TestGetAlertHistory(t *testing.T) {
	Convey("test create alert history", t, func() {
		Convey("test01", func() {
			historyies, err := GetAlertHistory(&AlertHistory{ID: 6})
			fmt.Println(err)
			fmt.Println(jsonutil.Marshal(historyies[0]))
		})

		Convey("test02", func() {
			historyies, err := GetAlertHistory(&AlertHistory{AlertConfigID: "alert-config-435kj7zrn4jrwz"})
			fmt.Println(err)
			fmt.Println(jsonutil.Marshal(historyies[0]))
		})
	})
}

func TestUpdateAlertSendStatus(t *testing.T) {
	Convey("test update notification status", t, func() {
		Convey("test01", func() {
			err := UpdateAlertSendStatus(&AlertHistory{ID: 12}, "hahahahahahahahaha")
			fmt.Println(err)
		})
	})
}
