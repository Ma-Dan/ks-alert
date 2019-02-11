package notification

import (
	"fmt"
	"kubesphere.io/ks-alert/pkg/executor/metric"
	"kubesphere.io/ks-alert/pkg/models"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
	"time"
)

func TestSender_Send(t *testing.T) {
	Convey("test send info", t, func() {
		Convey("test 01", func() {
			notice := Notice{
				TriggerTime:         time.Now(),
				CumulateReSendCount: 2,
				ResourceName:        "node_cpu",
				Metrics: &[]metric.TV{
					metric.TV{
						T: 1243465,
						V: "435.4354",
					},
					metric.TV{
						T: 4356567,
						V: "546756.4354",
					},
					metric.TV{
						T: 659809,
						V: "4354.4354",
					},
				},
				Rule: &models.AlertRule{
					AlertRuleName:          "namespace",
					MetricName:             "namespace_cpu",
					ConditionType:          ">",
					PerferSeverity:         true,
					Threshold:              80,
					Period:                 3,
					Unit:                   "%",
					ConsecutiveCount:       3,
					InhibitRule:            false,
					Enable:                 true,
					RepeatSendType:         1,
					InitRepeatSendInterval: 60,
					MaxRepeatSendCount:     4,
					CreatedAt:              time.Now(),
					UpdatedAt:              time.Now(),
				},
			}

			noticeStr := notice.MakeNotice(false)
			fmt.Println(noticeStr)
			fmt.Println(`{"email": ["1262758612@qq.com", "513590612@qq.com"]}`)
			sender := Sender{}
			sendStatus := sender.Send(&[]models.Receiver{
				{
					ReceiverName: "bbbbbbbbbbb",
					Email:        "zlahu@foxmail.com",
					Phone:        "33333333333333",
					Wechat:       "zhangliaish",
				},
				{
					ReceiverName: "aqaaaaaaaaaaaa",
					Email:        "13156518189@163.com",
					Phone:        "33333333333333",
					Wechat:       "haahaah",
				},
				{
					ReceiverName: "aqaaaaaaaaaaaa",
					Email:        "carmanzhang@yunify.com",
					Phone:        "33333333333333",
					Wechat:       "haahaah",
				},
			}, "{\"threshold\":80,\"time_series_metrics\":[{\"T\":1243465,\"V\":\"435.4354\"},{\"T\":1243465,\"V\":\"435.4354\"}]}")

			fmt.Println(sendStatus)
		})
	})
}
