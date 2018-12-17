package models

import (
	"time"
	"kubesphere.io/ks-alert/pkg/utils/dbutil"
	"errors"
	"kubesphere.io/ks-alert/pkg/utils/idutil"
)

type AlertRule struct {
	AlertRuleID      string `gorm:"primary_key" json:"-"`
	AlertRuleGroupID string `gorm:"type:varchar(50);not null;" json:"-"`

	SeverityID string `gorm:"type:varchar(50);not null;" json:"severity_id"`
	SeverityCh string `gorm:"type:varchar(10);" json:"severity_ch"`

	// ResourceTypeID string
	MetricID   string `gorm:"type:varchar(50);not null;" json:"metric_id"`
	MetricName string `gorm:"type:varchar(50);" json:"metric_name"`

	ConditionType string  `gorm:"type:varchar(10);not null;" json:"condition_type"`
	Threshold     float32 `gorm:"type:float;not null;" json:"threshold"`
	Unit          string  `gorm:"type:varchar(10);" json:"unit"`

	Period           int `gorm:"type:int;not null;" json:"period"`
	ConsecutiveCount int `gorm:"type:int;not null;" json:"consecutive_count"`

	InhibitRule bool `gorm:"type:boolean;" json:"inhibit_rule_enable"`
	Enable      bool `gorm:"type:boolean;" json:"enable"`

	CreatedAt time.Time `gorm:"not null;" json:"-"`
	UpdatedAt time.Time `gorm:"not null;" json:"-"`
	Version   int       `gorm:"type:int;not null;" json:"-"`

	RefAlertRuleID string `gorm:"type:varchar(50);" json:"ref_alert_rule_id"`
}

type AlertRuleGroup struct {
	AlertRuleGroupID   string      `gorm:"primary_key" json:"-"`
	AlertRuleGroupName string      `gorm:"type:varchar(50);not null;" json:"alert_rule_group_name"`
	AlertRules         []AlertRule `gorm:"-" json:"alert_rules"`
	Description        string      `gorm:"type:text;" json:"desc"`
	CreatedAt          time.Time   `gorm:"not null;" json:"-"`
	UpdatedAt          time.Time   `gorm:"not null;" json:"-"`
}

func CreateAlertRuleGroup(alertRuleGroup *AlertRuleGroup) (*AlertRuleGroup, error) {

	if alertRuleGroup.AlertRuleGroupName == "" {
		return nil, errors.New("resource Group Name is not given")
	}

	db, err := dbutil.DBClient()
	if err != nil {
		panic(err)
	}

	alertRuleGroup.AlertRuleGroupID = idutil.GetUuid36("alert_rule_group-")

	err = db.Model(&AlertRuleGroup{}).Create(alertRuleGroup).Error
	return alertRuleGroup, err
}

func CreateAlertRules(alertRules *[]AlertRule, alertRuleGroupID string) (*[]AlertRule, error) {
	var createdAlertRule []AlertRule

	for i := 0; i < len(*alertRules); i++ {
		alertRule, err := CreateAlertRule(&(*alertRules)[i], alertRuleGroupID)
		if err != nil {
			return &createdAlertRule, err
		}
		createdAlertRule = append(createdAlertRule, *alertRule)
	}

	return &createdAlertRule, nil
}

func CreateAlertRule(alertRule *AlertRule, alertRuleGroupID string) (*AlertRule, error) {

	db, err := dbutil.DBClient()
	if err != nil {
		panic(err)
	}

	refAlertRuleID := alertRule.AlertRuleID
	alertRule.AlertRuleID = idutil.GetUuid36("alert_rule-")
	alertRule.AlertRuleGroupID = alertRuleGroupID
	alertRule.RefAlertRuleID = refAlertRuleID

	err = db.Model(&AlertRule{}).Create(alertRule).Error

	return alertRule, err
}
