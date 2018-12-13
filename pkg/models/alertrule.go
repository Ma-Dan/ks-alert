package models

import "time"

type AlertRule struct {
	AlertRuleID      string `gorm:"primary_key"`
	AlertRuleGroupID string `gorm:"type:varchar(50);not null;"`

	SeverityID string `gorm:"type:varchar(50);not null;"`
	SeverityCh string `gorm:"type:varchar(10);"`

	// ResourceTypeID string
	MetricID   string `gorm:"type:varchar(50);not null;"`
	MetricName string `gorm:"type:varchar(50);"`

	ConditionType string  `gorm:"type:varchar(10);not null;"`
	Threshold     float32 `gorm:"type:float;not null;"`
	Unit          string  `gorm:"type:varchar(10);"`

	Period           int `gorm:"type:int;not null;"`
	ConsecutiveCount int `gorm:"type:int;not null;"`

	InhibitRule bool `gorm:"type:boolean;"`
	Enable      bool `gorm:"type:boolean;"`

	CreatedAt time.Time `gorm:"not null;"`
	UpdatedAt time.Time `gorm:"not null;"`
	Version   int       `gorm:"type:int;not null;"`

	RefAlertRuleID string `gorm:"type:varchar(50);"`
}

type AlertRuleGroup struct {
	AlertRuleGroupID   string    `gorm:"primary_key"`
	AlertRuleGroupName string    `gorm:"type:varchar(50);not null;"`
	Enable             bool      `gorm:"type:boolean;not null;"`
	CreatedAt          time.Time `gorm:"not null;"`
	UpdatedAt          time.Time `gorm:"not null;"`
}
