package models

import (
	"fmt"
	. "kubesphere.io/ks-alert/pkg/stderr"
	"kubesphere.io/ks-alert/pkg/utils/idutil"
	"github.com/jinzhu/gorm"
	"time"
)

var Bool2Int = map[bool]int8{true: 1, false: 0}

type AlertRule struct {
	AlertRuleID            string    `gorm:"primary_key"`
	AlertRuleName          string    `gorm:"type:varchar(50);not null;"`
	AlertRuleGroupID       string    `gorm:"type:varchar(50);not null;"`
	MetricName             string    `gorm:"type:varchar(50);"`
	ConditionType          string    `gorm:"type:varchar(10);not null;"`
	PerferSeverity         bool      `gorm:"type:boolean;"`
	Threshold              float32   `gorm:"type:float;not null;"`
	Unit                   string    `gorm:"type:varchar(10);"`
	Period                 int32     `gorm:"type:int;not null;"`
	ConsecutiveCount       int32     `gorm:"type:int;not null;"`
	InhibitRule            bool      `gorm:"type:boolean;"`
	Enable                 bool      `gorm:"type:boolean;"`
	RepeatSendType         int32     `gorm:"type:varchar(10);not null;"`
	InitRepeatSendInterval uint32    `gorm:"type:mediumint(11) unsigned;not null;"`
	MaxRepeatSendCount     uint32    `gorm:"type:mediumint(11) unsigned;not null;"`
	CreatedAt              time.Time `gorm:"not null;"`
	UpdatedAt              time.Time `gorm:"not null;"`
}

type AlertRuleGroup struct {
	Action
	AlertRuleGroupID   string       `gorm:"primary_key" json:"-"`
	AlertRuleGroupName string       `gorm:"type:varchar(50);not null;" json:"alert_rule_group_name"`
	AlertRules         []*AlertRule `gorm:"-" json:"alert_rules"`
	Description        string       `gorm:"type:text;" json:"desc"`
	SystemRule         bool         `gorm:"type:boolean;not null;"`
	ResourceTypeID     string       `gorm:"type:varchar(50);not null;"`
	CreatedAt          time.Time    `gorm:"not null;" json:"-"`
	UpdatedAt          time.Time    `gorm:"not null;" json:"-"`
}

func createAlertGroupAndRules(tx *gorm.DB, ruleGroup *AlertRuleGroup) error {

	alertRules := ruleGroup.AlertRules
	systemRule := ruleGroup.SystemRule

	// create alert rule group
	item := fmt.Sprintf("('%s','%s','%s','%v','%s','%v','%v')", ruleGroup.AlertRuleGroupID, ruleGroup.AlertRuleGroupName, ruleGroup.Description,
		Bool2Int[systemRule], ruleGroup.ResourceTypeID, ruleGroup.UpdatedAt, ruleGroup.CreatedAt)

	sql := "INSERT INTO alert_rule_groups (alert_rule_group_id, alert_rule_group_name, " +
		"description, system_rule, resource_type_id, updated_at, created_at) VALUES " + item

	if err := tx.Exec(sql).Error; err != nil {
		return Error{Text: err.Error(), Code: DBError, Where: Caller(0, true)}
	}

	err := createAlertRules(tx, alertRules, ruleGroup.AlertRuleGroupID)
	if err != nil {
		return Error{Text: err.Error(), Code: DBError, Where: Caller(0, true)}
	}

	return nil
}

func createAlertRules(tx *gorm.DB, alertRules []*AlertRule, ruleGroupID string) error {
	// create alert rules
	sql := "INSERT INTO alert_rules (alert_rule_id, " +
		"alert_rule_name, alert_rule_group_id, metric_name, condition_type, perfer_severity, " +
		"threshold, unit, period, consecutive_count, inhibit_rule, enable, repeat_send_type, " +
		"init_repeat_send_interval, max_repeat_send_count, created_at, updated_at) VALUES "
	l := len(alertRules)
	for i := 0; i < l; i++ {
		a := alertRules[i]

		a.AlertRuleGroupID = ruleGroupID
		a.AlertRuleID = idutil.GetUuid36("")

		item := fmt.Sprintf("('%s','%s','%s','%s','%s','%v','%f','%s','%d','%d','%v','%v','%d','%d','%d','%v','%v') ",
			a.AlertRuleID, a.AlertRuleName, a.AlertRuleGroupID, a.MetricName, a.ConditionType, Bool2Int[a.PerferSeverity],
			a.Threshold, a.Unit, a.Period, a.ConsecutiveCount, Bool2Int[a.InhibitRule], Bool2Int[a.Enable],
			a.RepeatSendType, a.InitRepeatSendInterval, a.MaxRepeatSendCount, a.CreatedAt, a.UpdatedAt)

		sql = sql + item
		if i != l-1 {
			sql = sql + ","
		}
	}

	return tx.Exec(sql).Error
}

func (r *AlertRuleGroup) Create(tx *gorm.DB) (interface{}, error) {
	if r.ResourceTypeID == "" {
		return nil, Error{Text: "resource type id must be specified", Code: InvalidParam, Where: Caller(0, true)}
	}

	if len(r.AlertRules) == 0 || r.AlertRuleGroupName == "" {
		return nil, Error{Text: "at least one alert rule and it's name must be specified", Code: InvalidParam, Where: Caller(0, true)}
	}

	rt, _ := GetResourceType(&ResourceType{ResourceTypeID: r.ResourceTypeID})

	if rt == nil || rt.ResourceTypeID == "" {
		return nil, Error{Text: "resource type does not exist", Code: InvalidParam, Where: Caller(0, true)}
	}

	// check if the build-in alert rule group exist?
	if r.SystemRule {
		rule := AlertRuleGroup{ResourceTypeID: r.ResourceTypeID, SystemRule: true}
		ruleGroup, _ := rule.Get(tx)

		if ruleGroup != nil && ruleGroup.(*AlertRuleGroup).AlertRuleGroupID != "" {
			// build-in alert rule group exist
			return nil, Error{Text: "build-in rules exists for the current resource type", Code: InvalidParam, Where: Caller(0, true)}
		}
	}

	r.AlertRuleGroupID = idutil.GetUuid36("")

	// do create
	err := createAlertGroupAndRules(tx, r)
	if err != nil {
		return nil, err
	}

	return r, nil
}

func (r *AlertRuleGroup) Update(tx *gorm.DB) (interface{}, error) {
	// check alert_rule_group_id is exist
	if r.AlertRuleGroupID == "" || r.AlertRuleGroupName == "" {
		return nil, Error{Text: "resource type id and rule group name must be specified", Code: InvalidParam, Where: Caller(0, true)}
	}

	alertRules := r.AlertRules

	// update alert rule group
	sql := fmt.Sprintf("UPDATE alert_rule_groups SET alert_rule_group_name='%s',description='%s',"+
		"system_rule='%v', updated_at='%v' WHERE alert_rule_group_id='%s'",
		r.AlertRuleGroupName, r.Description,
		Bool2Int[r.SystemRule],
		r.UpdatedAt, r.AlertRuleGroupID)

	if err := tx.Exec(sql).Error; err != nil {
		return nil, Error{Text: err.Error(), Code: DBError, Where: Caller(0, true)}
	}

	var needCreatedAlertRule []*AlertRule
	// update alert rules
	l := len(alertRules)
	for i := 0; i < l; i++ {
		a := alertRules[i]

		if a.AlertRuleID == "" {
			needCreatedAlertRule = append(needCreatedAlertRule, a)
			continue
		}

		sql = fmt.Sprintf("UPDATE alert_rules SET "+
			"alert_rule_name='%s', metric_name='%s', condition_type='%s', perfer_severity='%v', "+
			"threshold='%f', unit='%s', period='%d', consecutive_count='%d', inhibit_rule='%v', enable='%v', repeat_send_type='%d', "+
			"init_repeat_send_interval='%d', max_repeat_send_count='%d', updated_at='%v' WHERE alert_rule_group_id='%s' AND alert_rule_id='%s'",
			a.AlertRuleName, a.MetricName, a.ConditionType, Bool2Int[a.PerferSeverity],
			a.Threshold, a.Unit, a.Period, a.ConsecutiveCount, Bool2Int[a.InhibitRule],
			Bool2Int[a.Enable], a.RepeatSendType, a.InitRepeatSendInterval,
			a.MaxRepeatSendCount, time.Now(), r.AlertRuleGroupID, a.AlertRuleID)

		if err := tx.Exec(sql).Error; err != nil {
			return nil, Error{Text: err.Error(), Code: DBError, Where: Caller(0, true)}
		}
	}

	if len(needCreatedAlertRule) > 0 {
		err := createAlertRules(tx, needCreatedAlertRule, r.AlertRuleGroupID)
		return nil, Error{Text: err.Error(), Code: DBError, Where: Caller(0, true)}
	}

	return r, nil
}

func (r *AlertRuleGroup) Get(tx *gorm.DB) (interface{}, error) {
	groupID := r.AlertRuleGroupID
	// means to get supported alert rule for the resource type
	typeID := r.ResourceTypeID
	//systemRule := alertRuleSpec.SystemRule

	if groupID == "" && typeID == "" {
		return nil, Error{Text: "rsource type id or rule group id must be specified", Code: InvalidParam, Where: Caller(0, true)}
	}

	var rg AlertRuleGroup

	// get alert rule group
	if r.AlertRuleGroupID != "" {
		tx.Model(&AlertRuleGroup{}).Where(&AlertRuleGroup{AlertRuleGroupID: r.AlertRuleGroupID}).First(&rg)
	} else if r.ResourceTypeID != "" {
		//x := &AlertRuleGroup{ResourceTypeID: r.ResourceTypeId, SystemRule: true}
		tx.Model(&AlertRuleGroup{}).Where("resource_type_id=? AND system_rule=?", r.ResourceTypeID, true).First(&rg)
	}

	if tx.RecordNotFound() {
		return nil, Error{Text: "record not found", Code: DBError, Where: Caller(0, true)}
	}

	// get alert rules
	if rg.AlertRuleGroupID != "" {
		var alertRules []AlertRule
		err := tx.Find(&alertRules, "alert_rule_group_id=?", rg.AlertRuleGroupID).Error
		//err := db.Where(&AlertRule{AlertRuleGroupID: r.AlertRuleGroupId}).Find(&alertRules).Error
		//db.Exec("SELECT * FROM alert_rules WHERE alert_rule_group_id=?", r.AlertRuleGroupId).First(&alertRules)

		if err != nil {
			return &rg, Error{Text: err.Error(), Code: DBError, Where: Caller(0, true)}
		}

		var rules []*AlertRule
		for i := 0; i < len(alertRules); i++ {
			rules = append(rules, &alertRules[i])
		}
		rg.AlertRules = rules
	}

	return &rg, nil
}

func (r *AlertRuleGroup) Delete(tx *gorm.DB) (interface{}, error) {
	groupID := r.AlertRuleGroupID
	// means to get supported alert rule for the resource type
	typeID := r.ResourceTypeID
	//systemRule := alertRuleSpec.SystemRule

	// only can remove one alert rule once
	var ruleID string
	l := len(r.AlertRules)
	for i := 0; i < l; i++ {
		id := r.AlertRules[i].AlertRuleID
		if id != "" {
			ruleID = id
			break
		}
	}

	if groupID == "" && typeID == "" && ruleID != "" {
		return nil, Error{Text: "resource type id or rule group id or rule id must be specified", Code: InvalidParam, Where: Caller(0, true)}
	}

	if ruleID != "" {
		sql := "DELETE ar FROM alert_rules as ar WHERE ar.alert_rule_id=?"

		if err := tx.Exec(sql, ruleID).Error; err != nil {
			return nil, Error{Text: err.Error(), Code: DBError, Where: Caller(0, true)}
		}
	} else if groupID != "" {
		sql := "DELETE arg, ar FROM alert_rule_groups as arg LEFT JOIN alert_rules as ar ON arg.alert_rule_group_id=ar.alert_rule_group_id WHERE arg.alert_rule_group_id=?"

		if err := tx.Exec(sql, groupID).Error; err != nil {
			return nil, Error{Text: err.Error(), Code: DBError, Where: Caller(0, true)}
		}
	} else if typeID != "" {
		sql := "DELETE arg, ar FROM alert_rule_groups as arg LEFT JOIN alert_rules as ar ON arg.resource_type_id=ar.resource_type_id WHERE arg.resource_type_id=? AND arg.system_rule"

		if err := tx.Exec(sql, typeID, true).Error; err != nil {
			return nil, Error{Text: err.Error(), Code: DBError, Where: Caller(0, true)}
		}
	}

	return nil, nil
}
