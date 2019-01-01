package jsonutil

import (
	"fmt"
	"github.com/carmanzhang/ks-alert/pkg/dispatcher/pb"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestMarshal(t *testing.T) {
	Convey("test marshal", t, func() {
		x := pb.Suggestion{
			AlertConfigId: "qq",
			AlertRuleId:   "dsadfv",
			ResourceId:    "xxx",
			Messages: []*pb.Message{
				{
					Text:      "hahaah",
					UserId:    "cdvcv",
					Timestamp: 234356654,
					UserName:  "carman",
				},
				{
					Text:      "ertrhy",
					UserId:    "1346",
					Timestamp: 6754,
					UserName:  "zhang",
				},
			},
		}

		s := Marshal(x)
		fmt.Println(s)
	})
}

func TestUnmarshal(t *testing.T) {

	Convey("test unmarshal", t, func() {
		s := `{"alert_config_id":"qq","resource_id":"xxx","alert_rule_id":"dsadfv","messages":[{"user_name":"carman","user_id":"cdvcv","timestamp":234356654,"text":"hahaah"},{"user_name":"zhang","user_id":"1346","timestamp":6754,"text":"ertrhy"}]}`

		var p pb.Suggestion
		Unmarshal(s, &p)

		fmt.Println(p)
	})
}
