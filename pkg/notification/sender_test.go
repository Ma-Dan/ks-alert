package notification

import (
	"fmt"
	"github.com/carmanzhang/ks-alert/pkg/executor/metric"
	"github.com/carmanzhang/ks-alert/pkg/models"
	"github.com/carmanzhang/ks-alert/pkg/utils/jsonutil"
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
					SystemRule:             false,
					RepeatSendType:         1,
					InitRepeatSendInterval: 60,
					MaxRepeatSendCount:     4,
					CreatedAt:              time.Now(),
					UpdatedAt:              time.Now(),
				},
			}

			noticeStr := notice.MakeNotice(false)
			fmt.Println(noticeStr)
			sender := Sender{}
			sendStatusMap := sender.Send(&[]models.Receiver{
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
			}, noticeStr)

			fmt.Println(jsonutil.Marshal(sendStatusMap))
		})
	})
}
