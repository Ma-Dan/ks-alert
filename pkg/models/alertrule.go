package models

import (
	"errors"
	"fmt"
	"github.com/carmanzhang/ks-alert/pkg/dispatcher/pb"
	"github.com/carmanzhang/ks-alert/pkg/utils/dbutil"
	"github.com/carmanzhang/ks-alert/pkg/utils/idutil"
	"github.com/golang/glog"
	"time"
)

var Bool2Int = map[bool]int8{true: 1, false: 0}

type AlertRule struct {
	AlertRuleID      string `gorm:"primary_key" json:"-"`
	AlertRuleName    string `gorm:"type:varchar(50);not null;" json:"-"`
	AlertRuleGroupID string `gorm:"type:varchar(50);not null;" json:"-"`

	MetricName string `gorm:"type:varchar(50);" json:"metric_name"`

	ConditionType string `gorm:"type:varchar(10);not null;" json:"condition_type"`
	// a flag which use to indicate that relationship between Severity and Threshold
	PerferSeverity bool `gorm:"type:boolean;" json:"perfer_severity"`

	Threshold float32 `gorm:"type:float;not null;" json:"threshold"`

	//ThresholdSeverityString string              `gorm:"type:text;not null;" json:"-"`
	//ThresholdSeverity       []ThresholdSeverity `gorm:"-" json:"threshold_severity"`

	Unit string `gorm:"type:varchar(10);" json:"unit"`

	Period           int32 `gorm:"type:int;not null;" json:"period"`
	ConsecutiveCount int32 `gorm:"type:int;not null;" json:"consecutive_count"`

	InhibitRule bool `gorm:"type:boolean;" json:"inhibit_rule_enable"`
	Enable      bool `gorm:"type:boolean;" json:"enable"`
	SystemRule  bool `gorm:"type:boolean;" json:"enable"`
	// repeat send
	RepeatSendType         int32  `gorm:"type:varchar(10);not null;"`
	InitRepeatSendInterval uint32 `gorm:"type:mediumint(11) unsigned;not null;"`
	MaxRepeatSendCount     uint32 `gorm:"type:mediumint(11) unsigned;not null;"`

	CreatedAt time.Time `gorm:"not null;" json:"-"`
	UpdatedAt time.Time `gorm:"not null;" json:"-"`
}

type AlertRuleGroup struct {
	AlertRuleGroupID   string      `gorm:"primary_key" json:"-"`
	AlertRuleGroupName string      `gorm:"type:varchar(50);not null;" json:"alert_rule_group_name"`
	AlertRules         []AlertRule `gorm:"-" json:"alert_rules"`
	Description        string      `gorm:"type:text;" json:"desc"`
	SystemRule         bool        `gorm:"type:boolean;not null;"`
	ResourceTypeID     string      `gorm:"type:varchar(50);not null;"`
	CreatedAt          time.Time   `gorm:"not null;" json:"-"`
	UpdatedAt          time.Time   `gorm:"not null;" json:"-"`
}

func CreateAlertRuleGroup(alertRuleGroup *AlertRuleGroup) (*AlertRuleGroup, error) {
	// check if the build-in alert rule group exist?
	if alertRuleGroup.SystemRule {
		r := pb.AlertRuleGroupSpec{ResourceTypeId: alertRuleGroup.ResourceTypeID, SystemRule: true}
		ruleGroup, _ := GetAlertRuleGroup(&r)

		if ruleGroup != nil && ruleGroup.AlertRuleGroupID != "" {
			// build-in alert rule group exist
			return nil, errors.New("build-in rules exists for the current resource type")
		}
	}

	alertRuleGroup.AlertRuleGroupID = idutil.GetUuid36("rule_group-")

	// do create
	err := createAlertGroupAndRules(alertRuleGroup)
	if err != nil {
		return nil, err
	}

	return alertRuleGroup, nil
}

func UpdateAlertRuleGroup(ruleGroup *AlertRuleGroup) error {
	alertRules := ruleGroup.AlertRules

	db, err := dbutil.DBClient()

	if err != nil {
		glog.Errorln(err.Error())
		return err
	}

	tx := db.Begin()

	// update alert rule group
	sql := fmt.Sprintf("UPDATE alert_rule_groups SET alert_rule_group_name='%s',description='%s',"+
		"system_rule='%v', updated_at='%v' WHERE alert_rule_group_id='%s'",
		ruleGroup.AlertRuleGroupName, ruleGroup.Description,
		Bool2Int[ruleGroup.SystemRule],
		ruleGroup.UpdatedAt, ruleGroup.AlertRuleGroupID)

	fmt.Println(sql)

	if err := tx.Exec(sql).Error; err != nil {
		tx.Rollback()
		return err
	}

	// update alert rules
	l := len(alertRules)
	for i := 0; i < l; i++ {
		a := alertRules[i]

		if a.AlertRuleID == "" {
			continue
		}

		sql = fmt.Sprintf("UPDATE alert_rules SET "+
			"alert_rule_name='%s', metric_name='%s', condition_type='%s', perfer_severity='%v', "+
			"threshold='%f', unit='%s', period='%d', consecutive_count='%d', inhibit_rule='%v', enable='%v', system_rule='%v', repeat_send_type='%d', "+
			"init_repeat_send_interval='%d', max_repeat_send_count='%d', updated_at='%v' WHERE alert_rule_group_id='%s' AND alert_rule_id='%s'",
			a.AlertRuleName, a.MetricName, a.ConditionType, Bool2Int[a.PerferSeverity],
			a.Threshold, a.Unit, a.Period, a.ConsecutiveCount, Bool2Int[a.InhibitRule],
			Bool2Int[a.Enable], Bool2Int[ruleGroup.SystemRule],
			a.RepeatSendType, a.InitRepeatSendInterval, a.MaxRepeatSendCount, time.Now(),
			ruleGroup.AlertRuleGroupID, a.AlertRuleID)

		fmt.Println(sql)
		if err := tx.Exec(sql).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	tx.Commit()

	return nil
}

func GetAlertRuleGroup(ruleGroupSpec *pb.AlertRuleGroupSpec) (*AlertRuleGroup, error) {
	db, err := dbutil.DBClient()

	if err != nil {
		glog.Errorln(err.Error())
		return nil, err
	}
	var r AlertRuleGroup

	// get alert rule group
	if ruleGroupSpec.AlertRuleGroupId != "" {
		db = db.Model(&AlertRuleGroup{}).Where(&AlertRuleGroup{AlertRuleGroupID: ruleGroupSpec.AlertRuleGroupId}).First(&r)
	} else if ruleGroupSpec.ResourceTypeId != "" {
		//x := &AlertRuleGroup{ResourceTypeID: ruleGroupSpec.ResourceTypeId, SystemRule: true}
		err = db.Model(&AlertRuleGroup{}).Where("resource_type_id=? AND system_rule=?", ruleGroupSpec.ResourceTypeId, true).First(&r).Error
	}

	exist := db.RecordNotFound()

	if exist {
		return nil, errors.New("record not found")
	}

	// get alert rules
	if r.AlertRuleGroupID != "" {
		var alertRules []AlertRule
		err := db.Debug().Find(&alertRules, "alert_rule_group_id=?", r.AlertRuleGroupID).Error
		//err := db.Where(&AlertRule{AlertRuleGroupID: ruleGroupSpec.AlertRuleGroupId}).Find(&alertRules).Error
		//db.Exec("SELECT * FROM alert_rules WHERE alert_rule_group_id=?", ruleGroupSpec.AlertRuleGroupId).First(&alertRules)

		if err != nil {
			return &r, err
		}

		//var rulePtr []*AlertRule
		//for i := 0; i < len(alertRules); i++ {
		//	rulePtr = append(rulePtr, &alertRules[i])
		//}

		r.AlertRules = alertRules

		return &r, db.Error
	}

	return &r, nil
}

func DeleteAlertRuleGroup(ruleGroupSpec *pb.AlertRuleGroupSpec) error {
	db, err := dbutil.DBClient()

	if err != nil {
		return err
	}

	tx := db.Begin()

	err = tx.Delete(&AlertRuleGroup{AlertRuleGroupID: ruleGroupSpec.AlertRuleGroupId}).Error
	//err = db.Delete(&AlertRuleGroup{ResourceTypeID: ruleGroupSpec.ResourceTypeId, SystemRule: true}).Error
	//err = tx.Exec("DELETE from alert_rule_group WHERE alert_rule_group_id=?", ruleGroupSpec.AlertRuleGroupId).Error

	if err != nil {
		tx.Rollback()
		return err
	}

	// delete all alert rule
	err = tx.Delete(&AlertRule{AlertRuleGroupID: ruleGroupSpec.AlertRuleGroupId}).Error
	//err = tx.Exec("DELETE from alert_rule WHERE alert_rule_group_id=?", ruleGroupSpec.AlertRuleGroupId).Error

	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()

	return nil
}

func createAlertGroupAndRules(ruleGroup *AlertRuleGroup) error {

	alertRules := ruleGroup.AlertRules

	db, err := dbutil.DBClient()

	if err != nil {
		glog.Errorln(err.Error())
		return err
	}

	tx := db.Begin()

	systemRule := ruleGroup.SystemRule

	// create alert rule group
	item := fmt.Sprintf("('%s','%s','%s','%v','%s','%v','%v')", ruleGroup.AlertRuleGroupID, ruleGroup.AlertRuleGroupName, ruleGroup.Description,
		Bool2Int[systemRule], ruleGroup.ResourceTypeID, ruleGroup.CreatedAt, ruleGroup.UpdatedAt)

	sql := "INSERT INTO alert_rule_groups (alert_rule_group_id, alert_rule_group_name, " +
		"description, system_rule, resource_type_id, created_at, updated_at) VALUES " + item

	if err := tx.Exec(sql).Error; err != nil {
		tx.Rollback()
		return err
	}

	// create alert rules
	sql = "INSERT INTO alert_rules (alert_rule_id, " +
		"alert_rule_name, alert_rule_group_id, metric_name, condition_type, perfer_severity, " +
		"threshold, unit, period, consecutive_count, inhibit_rule, enable, system_rule, repeat_send_type, " +
		"init_repeat_send_interval, max_repeat_send_count, created_at, updated_at) VALUES "

	l := len(alertRules)
	for i := 0; i < l; i++ {
		a := alertRules[i]

		a.AlertRuleGroupID = ruleGroup.AlertRuleGroupID
		a.AlertRuleID = idutil.GetUuid36("rule_id-")

		item := fmt.Sprintf("('%s','%s','%s','%s','%s','%v','%f','%s','%d','%d','%v','%v','%v','%d','%d','%d','%v','%v') ",
			a.AlertRuleID, a.AlertRuleName, a.AlertRuleGroupID, a.MetricName, a.ConditionType, Bool2Int[a.PerferSeverity],
			a.Threshold, a.Unit, a.Period, a.ConsecutiveCount, Bool2Int[a.InhibitRule], Bool2Int[a.Enable], Bool2Int[systemRule],
			a.RepeatSendType, a.InitRepeatSendInterval, a.MaxRepeatSendCount, a.CreatedAt, a.UpdatedAt)

		sql = sql + item
		if i != l-1 {
			sql = sql + ","
		}
	}

	fmt.Println(sql)
	if err := tx.Exec(sql).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()

	return nil

}
