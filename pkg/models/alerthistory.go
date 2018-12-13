package models

import "time"

type AlertHistory struct {
	ID             uint32 `gorm:"primary_key;type:int(11) unsigned auto_increment;"`
	AlertBindingID string `gorm:"type:varchar(50);not null;"`
	AlertName      string `gorm:"type:varchar(50);not null;"`

	ProductID         string `gorm:"type:varchar(50);not null;"`
	ResourceGroupID   string `gorm:"type:varchar(50);not null;"`
	ResourceGroupName string `gorm:"type:varchar(50);not null;"`
	AlertedResource   string `gorm:"type:text;not null;"`

	ReceiverGroupID   string `gorm:"type:varchar(50);not null;"`
	ReceiverGroupName string `gorm:"type:varchar(50);not null;"`
	ReceiverGroup     string `gorm:"type:text;not null;"`

	AlertRuleGroupID string `gorm:"type:varchar(50);not null;"`
	TriggerAlertRule string `gorm:"type:text;not null;"`
	Cause            string `gorm:"type:text;"`

	SilenceEnable  bool      `gorm:"type:boolean;not null;default:false;"`
	SilenceStartAt time.Time `gorm:"not null;"`
	SilenceEndAt   time.Time `gorm:"not null;"`

	RepeatSendType         string `gorm:"type:varchar(10);not null;"`
	RepeatSendInterval     uint32 `gorm:"type:int unsigned;not null;"`
	InitRepeatSendInterval uint32 `gorm:"type:int unsigned;not null;"`
	MaxRepeatSendInterval  uint32 `gorm:"type:int unsigned;not null;"`

	EnableAt    time.Time `gorm:"not null;"`
	DisableAt   time.Time `gorm:"not null;"`
	Enable      bool      `gorm:"type:boolean;not null;default:true;"`
	Description string    `gorm:"type:text;"`

	CreatedAt time.Time `gorm:"not null;"`
	UpdatedAt time.Time `gorm:"not null;"`

	SendStatus  string `gorm:"type:text;"`
	AlertStatus string `gorm:"type:text;not null;"`
}
