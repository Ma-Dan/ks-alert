package models

import (
	"fmt"
	"github.com/carmanzhang/ks-alert/pkg/pb"
	"github.com/carmanzhang/ks-alert/pkg/utils/dbutil"
	"testing"
	"time"
)

func TestCreateAlertRuleGroup(t *testing.T) {
	Convey("test create alert rule group", t, func() {
		Convey("test create", func() {
			db, _ := dbutil.DBClient()
			arg := AlertRuleGroup{}
			v, err := arg.Create(db, &AlertRuleGroup{
				AlertRuleGroupName: "xxx",
				SystemRule:         true,
				UpdatedAt:          time.Now(),
				CreatedAt:          time.Now(),
				Description:        "desc",
				ResourceTypeID:     "xxxxxxxxxxxxxxx",
				AlertRules: []*AlertRule{
					&AlertRule{
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
						RepeatSendType:         int32(pb.RepeatSendType_Fixed),
						InitRepeatSendInterval: 60,
						MaxRepeatSendCount:     4,
						CreatedAt:              time.Now(),
						UpdatedAt:              time.Now(),
					},
					&AlertRule{
						AlertRuleName:          "namespace",
						MetricName:             "namespace_memory",
						ConditionType:          ">",
						PerferSeverity:         true,
						Threshold:              0.6,
						Period:                 20,
						Unit:                   "",
						ConsecutiveCount:       3,
						InhibitRule:            false,
						Enable:                 true,
						SystemRule:             true,
						RepeatSendType:         int32(pb.RepeatSendType_Exponential),
						InitRepeatSendInterval: 15,
						MaxRepeatSendCount:     10,
						CreatedAt:              time.Now(),
						UpdatedAt:              time.Now(),
					},
				},
			})
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(v)
		})
	})
}

func TestUpdateAlertRuleGroup(t *testing.T) {
	Convey("test update alert rule group", t, func() {
		Convey("test update", func() {
			db, _ := dbutil.DBClient()
			arg := AlertRuleGroup{}
			v, err := arg.Update(db, &AlertRuleGroup{
				AlertRuleGroupName: "kkkkkkk",
				SystemRule:         false,
				UpdatedAt:          time.Now(),
				CreatedAt:          time.Now(),
				Description:        "descdescdesc",
				ResourceTypeID:     "yyyyyyyyy",
				AlertRuleGroupID:   "rule_group-w7p5wm31j92330",
				AlertRules: []*AlertRule{
					&AlertRule{
						AlertRuleName:          "workspace",
						MetricName:             "workspace_cpu",
						ConditionType:          ">",
						PerferSeverity:         true,
						Threshold:              80,
						Period:                 3,
						Unit:                   "%",
						ConsecutiveCount:       3,
						InhibitRule:            false,
						Enable:                 true,
						SystemRule:             false,
						RepeatSendType:         int32(pb.RepeatSendType_Fixed),
						InitRepeatSendInterval: 60,
						MaxRepeatSendCount:     4,
						AlertRuleID:            "rule_id-n6mklr01j92330",
					},
					&AlertRule{
						AlertRuleName:          "workspace",
						MetricName:             "workspace_memory",
						ConditionType:          ">",
						PerferSeverity:         true,
						Threshold:              0.6,
						Period:                 20,
						Unit:                   "",
						ConsecutiveCount:       3,
						InhibitRule:            false,
						Enable:                 true,
						SystemRule:             true,
						RepeatSendType:         int32(pb.RepeatSendType_Exponential),
						InitRepeatSendInterval: 15,
						MaxRepeatSendCount:     10,
						//AlertRuleID:            "rule_id-qn4m4xpry0633x",
					},
				},
			})
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(v)
		})
	})
}

func TestGetAlertRuleGroup(t *testing.T) {
	Convey("test get alert rule group", t, func() {
		Convey("test get", func() {
			db, _ := dbutil.DBClient()
			arg := AlertRuleGroup{}
			v, err := arg.Get(db, &pb.AlertRuleGroupSpec{
				AlertRuleGroupId: "rule_group-n3no33k98nw330",
			})

			if err != nil {
				fmt.Println(err.Error())
			}
			fmt.Println(v)
		})
	})
}

func TestAlertRuleGroup_Create(t *testing.T) {
	Convey("test get alert rule group", t, func() {
		Convey("test get", func() {
			db, _ := dbutil.DBClient()
			arg := AlertRuleGroup{}

			_, err := arg.Create(db, &AlertRuleGroup{
				AlertRuleGroupName: "cluster_type_rule_group",
				ResourceTypeID:     "resource_type-lvzkj25vol1zk5",
				Description:        "",
				SystemRule:         true,
				UpdatedAt:          time.Now(),
				CreatedAt:          time.Now(),
				AlertRules:         []*AlertRule{&AlertRule{AlertRuleName: "workspace", MetricName: "workspace_memory", ConditionType: ">", PerferSeverity: true, Threshold: 0.6, Period: 20, Unit: "", ConsecutiveCount: 3, InhibitRule: false, Enable: true, SystemRule: true, RepeatSendType: int32(pb.RepeatSendType_Exponential), InitRepeatSendInterval: 15, MaxRepeatSendCount: 10, CreatedAt: time.Now(), UpdatedAt: time.Now()}},
			})

			So(err, ShouldBeNil)

			_, err = arg.Create(db, &AlertRuleGroup{
				AlertRuleGroupName: "node_type_rule_group",
				ResourceTypeID:     "resource_type-l07941wpjzvyll",
				Description:        "",
				SystemRule:         true,
				UpdatedAt:          time.Now(),
				CreatedAt:          time.Now(),
				AlertRules:         []*AlertRule{&AlertRule{AlertRuleName: "workspace", MetricName: "workspace_memory", ConditionType: ">", PerferSeverity: true, Threshold: 0.6, Period: 20, Unit: "", ConsecutiveCount: 3, InhibitRule: false, Enable: true, SystemRule: true, RepeatSendType: int32(pb.RepeatSendType_Exponential), InitRepeatSendInterval: 15, MaxRepeatSendCount: 10, CreatedAt: time.Now(), UpdatedAt: time.Now()}},
			})

			So(err, ShouldBeNil)
			_, err = arg.Create(db, &AlertRuleGroup{
				AlertRuleGroupName: "workspace_type_rule_group",
				ResourceTypeID:     "resource_type-lz6o6vjqxlroww",
				Description:        "",
				SystemRule:         true,
				UpdatedAt:          time.Now(),
				CreatedAt:          time.Now(),
				AlertRules:         []*AlertRule{&AlertRule{AlertRuleName: "workspace", MetricName: "workspace_memory", ConditionType: ">", PerferSeverity: true, Threshold: 0.6, Period: 20, Unit: "", ConsecutiveCount: 3, InhibitRule: false, Enable: true, SystemRule: true, RepeatSendType: int32(pb.RepeatSendType_Exponential), InitRepeatSendInterval: 15, MaxRepeatSendCount: 10, CreatedAt: time.Now(), UpdatedAt: time.Now()}},
			})

			So(err, ShouldBeNil)
			_, err = arg.Create(db, &AlertRuleGroup{
				AlertRuleGroupName: "namespace_type_rule_group",
				ResourceTypeID:     "resource_type-q0lw3n1pn04yjp",
				Description:        "",
				SystemRule:         true,
				UpdatedAt:          time.Now(),
				CreatedAt:          time.Now(),
				AlertRules:         []*AlertRule{&AlertRule{AlertRuleName: "workspace", MetricName: "workspace_memory", ConditionType: ">", PerferSeverity: true, Threshold: 0.6, Period: 20, Unit: "", ConsecutiveCount: 3, InhibitRule: false, Enable: true, SystemRule: true, RepeatSendType: int32(pb.RepeatSendType_Exponential), InitRepeatSendInterval: 15, MaxRepeatSendCount: 10, CreatedAt: time.Now(), UpdatedAt: time.Now()}},
			})

			So(err, ShouldBeNil)
			_, err = arg.Create(db, &AlertRuleGroup{
				AlertRuleGroupName: "workload_type_rule_group",
				ResourceTypeID:     "resource_type-k744yo5vol1zk5",
				Description:        "",
				SystemRule:         true,
				UpdatedAt:          time.Now(),
				CreatedAt:          time.Now(),
				AlertRules:         []*AlertRule{&AlertRule{AlertRuleName: "workspace", MetricName: "workspace_memory", ConditionType: ">", PerferSeverity: true, Threshold: 0.6, Period: 20, Unit: "", ConsecutiveCount: 3, InhibitRule: false, Enable: true, SystemRule: true, RepeatSendType: int32(pb.RepeatSendType_Exponential), InitRepeatSendInterval: 15, MaxRepeatSendCount: 10, CreatedAt: time.Now(), UpdatedAt: time.Now()}},
			})

			So(err, ShouldBeNil)

			_, err = arg.Create(db, &AlertRuleGroup{
				AlertRuleGroupName: "pod_type_rule_group",
				ResourceTypeID:     "resource_type-z3485jwpjzvyll",
				Description:        "",
				SystemRule:         true,
				UpdatedAt:          time.Now(),
				CreatedAt:          time.Now(),
				AlertRules:         []*AlertRule{&AlertRule{AlertRuleName: "workspace", MetricName: "workspace_memory", ConditionType: ">", PerferSeverity: true, Threshold: 0.6, Period: 20, Unit: "", ConsecutiveCount: 3, InhibitRule: false, Enable: true, SystemRule: true, RepeatSendType: int32(pb.RepeatSendType_Exponential), InitRepeatSendInterval: 15, MaxRepeatSendCount: 10, CreatedAt: time.Now(), UpdatedAt: time.Now()}},
			})

			So(err, ShouldBeNil)

			_, err = arg.Create(db, &AlertRuleGroup{
				AlertRuleGroupName: "container_type_rule_group",
				ResourceTypeID:     "resource_type-o1jxqpk3vp9zo9",
				Description:        "",
				SystemRule:         true,
				UpdatedAt:          time.Now(),
				CreatedAt:          time.Now(),
				AlertRules:         []*AlertRule{&AlertRule{AlertRuleName: "workspace", MetricName: "workspace_memory", ConditionType: ">", PerferSeverity: true, Threshold: 0.6, Period: 20, Unit: "", ConsecutiveCount: 3, InhibitRule: false, Enable: true, SystemRule: true, RepeatSendType: int32(pb.RepeatSendType_Exponential), InitRepeatSendInterval: 15, MaxRepeatSendCount: 10, CreatedAt: time.Now(), UpdatedAt: time.Now()}},
			})

			So(err, ShouldBeNil)

		})
	})
}
