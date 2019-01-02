package models

import "time"

type AlertHistory struct {
	// this ID is used for Paging
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

	SeverityID string `gorm:"type:varchar(50);not null;" json:"severity_id"`
	SeverityCh string `gorm:"type:varchar(10);" json:"severity_ch"`

	Cause string `gorm:"type:text;"`

	//SilenceEnable  bool      `gorm:"type:boolean;not null;default:false;"`
	SilenceStartAt time.Time `gorm:"not null;"`
	SilenceEndAt   time.Time `gorm:"not null;"`

	RepeatSendType            string `gorm:"type:varchar(10);not null;"`
	CurrentRepeatSendInterval uint32 `gorm:"type:int unsigned;not null;"`
	CurrentRepeatSendCount    uint32 `gorm:"type:int unsigned;not null;"`
	InitRepeatSendInterval    uint32 `gorm:"type:int unsigned;not null;"`
	MaxRepeatSendCount        uint32 `gorm:"type:int unsigned;not null;"`

	RequestNotificationStatus string    `gorm:"type:text;"`
	NotificationSendAt        time.Time `gorm:"not null;"`

	CreatedAt time.Time `gorm:"not null;"`
	UpdatedAt time.Time `gorm:"not null;"`
}
