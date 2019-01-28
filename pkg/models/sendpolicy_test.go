package models

import (
	"fmt"
	"github.com/carmanzhang/ks-alert/pkg/utils/jsonutil"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
	"time"
)

func TestGetSendPolicy(t *testing.T) {
	Convey("test get fired alert", t, func() {
		Convey("test01", func() {
			policy, err := GetSendPolicy(&SendPolicy{
				AlertRuleID: "xxxxx",
				ResourceID:  "errrrrrrr",
			})
			fmt.Println(jsonutil.Marshal(policy))
			fmt.Println(err)
			//fmt.Println(policy.SilenceStartAt.Before(time.Now()))
		})
	})
}

func TestUpdateSendPolicy(t *testing.T) {
	Convey("test get fired alert", t, func() {
		Convey("test01", func() {
			err := UpdateSendPolicySilenceRule(&SendPolicy{
				AlertRuleID:             "rule_id-yyyyy",
				ResourceID:              "rule_id-yyyyy",
				CumulateRepeatSendCount: 100,
				NextRepeatSendInterval:  88,
				SilenceStartAt:          time.Now(),
				SilenceEndAt:            time.Now().Add(time.Minute * 5),
				UpdatedAt:               time.Now(),
			})
			fmt.Println(err)
		})
	})
}

func TestCreateSendPolicy(t *testing.T) {
	Convey("test get fired alert", t, func() {
		Convey("test01", func() {

			str := "0001-01-01T00:00:00Z"
			ti, err := time.Parse(str, str)
			fmt.Println(ti, err)

			e := CreateSendPolicy(&SendPolicy{
				AlertRuleID: "rule_id-yyyyy",
				ResourceID:  "rule_id-yyyyy",
				//SendPolicyID:              "tttttttt",
				NextRepeatSendInterval:  1,
				CumulateRepeatSendCount: 0,
			})
			if e != nil {
				fmt.Println(e.Error())
			}
		})
	})
}
