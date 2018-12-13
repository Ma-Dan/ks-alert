package models

import "time"

type AlertBinding struct {
	AlertBindingID string `gorm:"primary_key"`
	AlertName      string `gorm:"type:varchar(50);not null;"`

	AlertRuleGroupID string `gorm:"type:varchar(50);not null;"`
	ResourceGroupID  string `gorm:"type:varchar(50);not null;"`
	ReceiverGroupID  string `gorm:"type:varchar(50);not null;"`

	ProductID string `gorm:"type:varchar(50);not null;"`

	RepeatSendType         string `gorm:"type:varchar(10);not null;"`
	RepeatSendInterval     uint   `gorm:"type:mediumint(11) unsigned;not null;"`
	InitRepeatSendInterval uint   `gorm:"type:mediumint(11) unsigned;not null;"`
	MaxRepeatSendInterval  uint   `gorm:"type:mediumint(11) unsigned;not null;"`

	EnableAt    time.Time `gorm:"not null;"`
	DisableAt   time.Time `gorm:"not null;"`
	Enable      bool      `gorm:"type:boolean;not null;default:true;"`
	Description string    `gorm:"type:text;"`

	CreatedAt time.Time `gorm:"type:timestamp;not null;default:current_timestamp;"`
	UpdatedAt time.Time `gorm:"type:timestamp;not null on update current_timestamp;default:current_timestamp;"`

	Version int `gorm:"type:int unsigned;not null;default:0;"`

	UpdateType  string    `gorm:"type:varchar(10);not null;"`
	KeepAliveAt time.Time `gorm:"not null;"`
	// this alert config binding is executing on a specific `node`
	NodeID string `gorm:"type:varchar(50);not null;"`
}
