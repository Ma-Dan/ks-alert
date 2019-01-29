package models

import (
	"fmt"
	"github.com/carmanzhang/ks-alert/pkg/pb"
	"github.com/carmanzhang/ks-alert/pkg/utils/dbutil"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
	"time"
)

func TestCreateAlertRuleGroup(t *testing.T) {
	Convey("test create alert rule group", t, func() {
		Convey("test create cluster rule group", func() {
			db, _ := dbutil.DBClient()
			group := &AlertRuleGroup{
				AlertRuleGroupName: "ks_cluster_rules",
				SystemRule:         true,
				UpdatedAt:          time.Now(),
				CreatedAt:          time.Now(),
				Description:        "",
				ResourceTypeID:     "lvzkj25vol1zk5",
				AlertRules: []*AlertRule{
					&AlertRule{
						AlertRuleName:          "cluster_cpu_utilisation",
						MetricName:             "cluster_cpu_utilisation",
						ConditionType:          ">",
						PerferSeverity:         true,
						Threshold:              80,
						Period:                 3,
						Unit:                   "%",
						ConsecutiveCount:       3,
						InhibitRule:            false,
						Enable:                 true,
						RepeatSendType:         int32(pb.RepeatSendType_Exponential),
						InitRepeatSendInterval: 60,
						MaxRepeatSendCount:     10,
						CreatedAt:              time.Now(),
						UpdatedAt:              time.Now(),
					},

					&AlertRule{
						AlertRuleName:          "cluster_memory_utilisation",
						MetricName:             "cluster_memory_utilisation",
						ConditionType:          ">",
						PerferSeverity:         true,
						Threshold:              80,
						Period:                 3,
						Unit:                   "%",
						ConsecutiveCount:       3,
						InhibitRule:            false,
						Enable:                 true,
						RepeatSendType:         int32(pb.RepeatSendType_Exponential),
						InitRepeatSendInterval: 60,
						MaxRepeatSendCount:     10,
						CreatedAt:              time.Now(),
						UpdatedAt:              time.Now(),
					},

					&AlertRule{
						AlertRuleName:          "cluster_net_utilisation",
						MetricName:             "cluster_net_utilisation",
						ConditionType:          ">",
						PerferSeverity:         true,
						Threshold:              200,
						Period:                 3,
						Unit:                   "bps",
						ConsecutiveCount:       3,
						InhibitRule:            false,
						Enable:                 true,
						RepeatSendType:         int32(pb.RepeatSendType_Exponential),
						InitRepeatSendInterval: 60,
						MaxRepeatSendCount:     10,
						CreatedAt:              time.Now(),
						UpdatedAt:              time.Now(),
					},

					&AlertRule{
						AlertRuleName:          "cluster_node_offline",
						MetricName:             "cluster_node_offline",
						ConditionType:          ">",
						PerferSeverity:         true,
						Threshold:              1,
						Period:                 3,
						Unit:                   "",
						ConsecutiveCount:       3,
						InhibitRule:            false,
						Enable:                 true,
						RepeatSendType:         int32(pb.RepeatSendType_Exponential),
						InitRepeatSendInterval: 60,
						MaxRepeatSendCount:     10,
						CreatedAt:              time.Now(),
						UpdatedAt:              time.Now(),
					},

					&AlertRule{
						AlertRuleName:          "cluster_namespace_count",
						MetricName:             "cluster_namespace_count",
						ConditionType:          ">",
						PerferSeverity:         true,
						Threshold:              25,
						Period:                 3,
						Unit:                   "",
						ConsecutiveCount:       3,
						InhibitRule:            false,
						Enable:                 true,
						RepeatSendType:         int32(pb.RepeatSendType_Exponential),
						InitRepeatSendInterval: 60,
						MaxRepeatSendCount:     10,
						CreatedAt:              time.Now(),
						UpdatedAt:              time.Now(),
					},
				},
			}
			v, err := group.Create(db)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(v)
		})

		// node
		Convey("test create node rule group", func() {
			db, _ := dbutil.DBClient()
			group := &AlertRuleGroup{
				AlertRuleGroupName: "ks_node_rules",
				SystemRule:         true,
				UpdatedAt:          time.Now(),
				CreatedAt:          time.Now(),
				Description:        "",
				ResourceTypeID:     "l07941wpjzvyll",
				AlertRules: []*AlertRule{
					&AlertRule{
						AlertRuleName:          "node_cpu_utilisation",
						MetricName:             "node_cpu_utilisation",
						ConditionType:          ">",
						PerferSeverity:         true,
						Threshold:              60,
						Period:                 4,
						Unit:                   "%",
						ConsecutiveCount:       3,
						InhibitRule:            false,
						Enable:                 true,
						RepeatSendType:         int32(pb.RepeatSendType_Fixed),
						InitRepeatSendInterval: 120,
						MaxRepeatSendCount:     10,
						CreatedAt:              time.Now(),
						UpdatedAt:              time.Now(),
					},

					&AlertRule{
						AlertRuleName:          "node_memory_utilisation",
						MetricName:             "node_memory_utilisation",
						ConditionType:          ">",
						PerferSeverity:         true,
						Threshold:              85,
						Period:                 3,
						Unit:                   "%",
						ConsecutiveCount:       3,
						InhibitRule:            false,
						Enable:                 true,
						RepeatSendType:         int32(pb.RepeatSendType_Exponential),
						InitRepeatSendInterval: 100,
						MaxRepeatSendCount:     10,
						CreatedAt:              time.Now(),
						UpdatedAt:              time.Now(),
					},

					&AlertRule{
						AlertRuleName:          "node_disk_inode_utilisation",
						MetricName:             "node_disk_inode_utilisation",
						ConditionType:          ">",
						PerferSeverity:         true,
						Threshold:              90,
						Period:                 3,
						Unit:                   "%",
						ConsecutiveCount:       3,
						InhibitRule:            false,
						Enable:                 true,
						RepeatSendType:         int32(pb.RepeatSendType_Exponential),
						InitRepeatSendInterval: 60,
						MaxRepeatSendCount:     10,
						CreatedAt:              time.Now(),
						UpdatedAt:              time.Now(),
					},
				},
			}
			v, err := group.Create(db)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(v)
		})

		// namespace
		Convey("test create namespace rule group", func() {
			db, _ := dbutil.DBClient()
			group := &AlertRuleGroup{
				AlertRuleGroupName: "ks_namespace_rules",
				SystemRule:         true,
				UpdatedAt:          time.Now(),
				CreatedAt:          time.Now(),
				Description:        "",
				ResourceTypeID:     "q0lw3n1pn04yjp",
				AlertRules: []*AlertRule{
					&AlertRule{
						AlertRuleName:          "namespace_cpu_usage",
						MetricName:             "namespace_cpu_usage",
						ConditionType:          ">",
						PerferSeverity:         true,
						Threshold:              120,
						Period:                 4,
						Unit:                   "m",
						ConsecutiveCount:       3,
						InhibitRule:            false,
						Enable:                 true,
						RepeatSendType:         int32(pb.RepeatSendType_Fixed),
						InitRepeatSendInterval: 360,
						MaxRepeatSendCount:     10,
						CreatedAt:              time.Now(),
						UpdatedAt:              time.Now(),
					},

					&AlertRule{
						AlertRuleName:          "namespace_memory_usage_wo_cache",
						MetricName:             "namespace_memory_usage_wo_cache",
						ConditionType:          ">",
						PerferSeverity:         true,
						Threshold:              500,
						Period:                 3,
						Unit:                   "m",
						ConsecutiveCount:       3,
						InhibitRule:            false,
						Enable:                 true,
						RepeatSendType:         int32(pb.RepeatSendType_Exponential),
						InitRepeatSendInterval: 60,
						MaxRepeatSendCount:     10,
						CreatedAt:              time.Now(),
						UpdatedAt:              time.Now(),
					},

					&AlertRule{
						AlertRuleName:          "namespace_pod_count",
						MetricName:             "namespace_pod_count",
						ConditionType:          ">",
						PerferSeverity:         true,
						Threshold:              70,
						Period:                 3,
						Unit:                   "",
						ConsecutiveCount:       3,
						InhibitRule:            false,
						Enable:                 true,
						RepeatSendType:         int32(pb.RepeatSendType_Exponential),
						InitRepeatSendInterval: 100,
						MaxRepeatSendCount:     10,
						CreatedAt:              time.Now(),
						UpdatedAt:              time.Now(),
					},
				},
			}
			v, err := group.Create(db)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(v)
		})

		// workload
		Convey("test create workload rule group", func() {
			db, _ := dbutil.DBClient()
			group := &AlertRuleGroup{
				AlertRuleGroupName: "ks_workload_rules",
				SystemRule:         true,
				UpdatedAt:          time.Now(),
				CreatedAt:          time.Now(),
				Description:        "",
				ResourceTypeID:     "k744yo5vol1zk5",
				AlertRules: []*AlertRule{
					&AlertRule{
						AlertRuleName:          "workload_pod_cpu_usage",
						MetricName:             "workload_pod_cpu_usage",
						ConditionType:          ">",
						PerferSeverity:         true,
						Threshold:              600,
						Period:                 4,
						Unit:                   "m",
						ConsecutiveCount:       3,
						InhibitRule:            false,
						Enable:                 true,
						RepeatSendType:         int32(pb.RepeatSendType_Fixed),
						InitRepeatSendInterval: 120,
						MaxRepeatSendCount:     10,
						CreatedAt:              time.Now(),
						UpdatedAt:              time.Now(),
					},

					&AlertRule{
						AlertRuleName:          "workload_pod_memory_usage_wo_cache",
						MetricName:             "workload_pod_memory_usage_wo_cache",
						ConditionType:          ">",
						PerferSeverity:         true,
						Threshold:              600,
						Period:                 3,
						Unit:                   "m",
						ConsecutiveCount:       3,
						InhibitRule:            false,
						Enable:                 true,
						RepeatSendType:         int32(pb.RepeatSendType_Exponential),
						InitRepeatSendInterval: 100,
						MaxRepeatSendCount:     10,
						CreatedAt:              time.Now(),
						UpdatedAt:              time.Now(),
					},
				},
			}
			v, err := group.Create(db)
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
			group := &AlertRuleGroup{
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
						RepeatSendType:         int32(pb.RepeatSendType_Exponential),
						InitRepeatSendInterval: 15,
						MaxRepeatSendCount:     10,
						//AlertRuleID:            "rule_id-qn4m4xpry0633x",
					},
				},
			}
			v, err := group.Update(db)
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
			arg := AlertRuleGroup{AlertRuleGroupID: "rule_group-n3no33k98nw330"}
			v, err := arg.Get(db)

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

			group := &AlertRuleGroup{
				AlertRuleGroupName: "cluster_type_rule_group",
				ResourceTypeID:     "resource_type-lvzkj25vol1zk5",
				Description:        "",
				SystemRule:         true,
				UpdatedAt:          time.Now(),
				CreatedAt:          time.Now(),
				AlertRules:         []*AlertRule{&AlertRule{AlertRuleName: "workspace", MetricName: "workspace_memory", ConditionType: ">", PerferSeverity: true, Threshold: 0.6, Period: 20, Unit: "", ConsecutiveCount: 3, InhibitRule: false, Enable: true, RepeatSendType: int32(pb.RepeatSendType_Exponential), InitRepeatSendInterval: 15, MaxRepeatSendCount: 10, CreatedAt: time.Now(), UpdatedAt: time.Now()}},
			}
			_, err := group.Create(db)

			So(err, ShouldBeNil)

			ruleGroup := &AlertRuleGroup{
				AlertRuleGroupName: "node_type_rule_group",
				ResourceTypeID:     "resource_type-l07941wpjzvyll",
				Description:        "",
				SystemRule:         true,
				UpdatedAt:          time.Now(),
				CreatedAt:          time.Now(),
				AlertRules:         []*AlertRule{&AlertRule{AlertRuleName: "workspace", MetricName: "workspace_memory", ConditionType: ">", PerferSeverity: true, Threshold: 0.6, Period: 20, Unit: "", ConsecutiveCount: 3, InhibitRule: false, Enable: true, RepeatSendType: int32(pb.RepeatSendType_Exponential), InitRepeatSendInterval: 15, MaxRepeatSendCount: 10, CreatedAt: time.Now(), UpdatedAt: time.Now()}},
			}
			_, err = ruleGroup.Create(db)

			So(err, ShouldBeNil)
			alertRuleGroup := &AlertRuleGroup{
				AlertRuleGroupName: "workspace_type_rule_group",
				ResourceTypeID:     "resource_type-lz6o6vjqxlroww",
				Description:        "",
				SystemRule:         true,
				UpdatedAt:          time.Now(),
				CreatedAt:          time.Now(),
				AlertRules:         []*AlertRule{&AlertRule{AlertRuleName: "workspace", MetricName: "workspace_memory", ConditionType: ">", PerferSeverity: true, Threshold: 0.6, Period: 20, Unit: "", ConsecutiveCount: 3, InhibitRule: false, Enable: true, RepeatSendType: int32(pb.RepeatSendType_Exponential), InitRepeatSendInterval: 15, MaxRepeatSendCount: 10, CreatedAt: time.Now(), UpdatedAt: time.Now()}},
			}
			_, err = alertRuleGroup.Create(db)

			So(err, ShouldBeNil)
			alertRuleGroup = &AlertRuleGroup{
				AlertRuleGroupName: "namespace_type_rule_group",
				ResourceTypeID:     "resource_type-q0lw3n1pn04yjp",
				Description:        "",
				SystemRule:         true,
				UpdatedAt:          time.Now(),
				CreatedAt:          time.Now(),
				AlertRules:         []*AlertRule{&AlertRule{AlertRuleName: "workspace", MetricName: "workspace_memory", ConditionType: ">", PerferSeverity: true, Threshold: 0.6, Period: 20, Unit: "", ConsecutiveCount: 3, InhibitRule: false, Enable: true, RepeatSendType: int32(pb.RepeatSendType_Exponential), InitRepeatSendInterval: 15, MaxRepeatSendCount: 10, CreatedAt: time.Now(), UpdatedAt: time.Now()}},
			}
			_, err = alertRuleGroup.Create(db)

			So(err, ShouldBeNil)

			alertRuleGroup = &AlertRuleGroup{
				AlertRuleGroupName: "workload_type_rule_group",
				ResourceTypeID:     "resource_type-k744yo5vol1zk5",
				Description:        "",
				SystemRule:         true,
				UpdatedAt:          time.Now(),
				CreatedAt:          time.Now(),
				AlertRules:         []*AlertRule{&AlertRule{AlertRuleName: "workspace", MetricName: "workspace_memory", ConditionType: ">", PerferSeverity: true, Threshold: 0.6, Period: 20, Unit: "", ConsecutiveCount: 3, InhibitRule: false, Enable: true, RepeatSendType: int32(pb.RepeatSendType_Exponential), InitRepeatSendInterval: 15, MaxRepeatSendCount: 10, CreatedAt: time.Now(), UpdatedAt: time.Now()}},
			}
			_, err = alertRuleGroup.Create(db)

			So(err, ShouldBeNil)

			alertRuleGroup = &AlertRuleGroup{
				AlertRuleGroupName: "pod_type_rule_group",
				ResourceTypeID:     "resource_type-z3485jwpjzvyll",
				Description:        "",
				SystemRule:         true,
				UpdatedAt:          time.Now(),
				CreatedAt:          time.Now(),
				AlertRules:         []*AlertRule{&AlertRule{AlertRuleName: "workspace", MetricName: "workspace_memory", ConditionType: ">", PerferSeverity: true, Threshold: 0.6, Period: 20, Unit: "", ConsecutiveCount: 3, InhibitRule: false, Enable: true, RepeatSendType: int32(pb.RepeatSendType_Exponential), InitRepeatSendInterval: 15, MaxRepeatSendCount: 10, CreatedAt: time.Now(), UpdatedAt: time.Now()}},
			}
			_, err = alertRuleGroup.Create(db)

			So(err, ShouldBeNil)

			alertRuleGroup = &AlertRuleGroup{
				AlertRuleGroupName: "container_type_rule_group",
				ResourceTypeID:     "resource_type-o1jxqpk3vp9zo9",
				Description:        "",
				SystemRule:         true,
				UpdatedAt:          time.Now(),
				CreatedAt:          time.Now(),
				AlertRules:         []*AlertRule{&AlertRule{AlertRuleName: "workspace", MetricName: "workspace_memory", ConditionType: ">", PerferSeverity: true, Threshold: 0.6, Period: 20, Unit: "", ConsecutiveCount: 3, InhibitRule: false, Enable: true, RepeatSendType: int32(pb.RepeatSendType_Exponential), InitRepeatSendInterval: 15, MaxRepeatSendCount: 10, CreatedAt: time.Now(), UpdatedAt: time.Now()}},
			}
			_, err = alertRuleGroup.Create(db)

			So(err, ShouldBeNil)

		})
	})
}
