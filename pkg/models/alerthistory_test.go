package models

import (
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
	"time"
)

func TestCreateAlertHistory(t *testing.T) {
	Convey("test create alert history", t, func() {
		Convey("test01", func() {
			CreateAlertHistory(&AlertHistory{
				AlertConfigID: "xzzzzzzzzz",
				ProductID:     "ppppppppp",

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
				CurrentRepeatSendCount:    1,
				CurrentRepeatSendInterval: 12,

				SilenceStartAt: time.Now(),
				SilenceEndAt:   time.Now(),

				Cause:           "yyyyyyyyy",
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
			historyies, err := GetAlertHistory(&AlertHistory{ID: 3})
			fmt.Println(err)
			fmt.Println(historyies[0])

		})
	})
}
