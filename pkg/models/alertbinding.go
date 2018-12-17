package models

import "time"

type AlertConfig struct {
	ReceiverGroup  ReceiverGroup  `json:"receiver_group"`
	AlertRuleGroup AlertRuleGroup `json:"alert_rule_group"`
	ResourceGroup  ResourceGroup  `json:"resource_group"`
	URIParams  Params      `json:"resource_uri_params, omitempty"`
}

type AlertBinding struct {
	AlertBindingID string `gorm:"primary_key"`
	AlertName      string `gorm:"type:varchar(50);not null;"`

	AlertRuleGroupID string `gorm:"type:varchar(50);not null;"`
	ResourceGroupID  string `gorm:"type:varchar(50);not null;"`
	ReceiverGroupID  string `gorm:"type:varchar(50);not null;"`

	ProductID string `gorm:"type:varchar(50);not null;"`
	// repeat send
	RepeatSendType         string `gorm:"type:varchar(10);not null;"`
	RepeatSendInterval     uint   `gorm:"type:mediumint(11) unsigned;not null;"`
	InitRepeatSendInterval uint   `gorm:"type:mediumint(11) unsigned;not null;"`
	MaxRepeatSendInterval  uint   `gorm:"type:mediumint(11) unsigned;not null;"`

	Enable    bool      `gorm:"type:boolean;not null;default:true;"`
	EnableAt  time.Time `gorm:"not null;"`
	DisableAt time.Time `gorm:"not null;"`

	Description string `gorm:"type:text;"`

	CreatedAt time.Time `gorm:"type:timestamp;not null;default:current_timestamp;"`
	UpdatedAt time.Time `gorm:"type:timestamp;not null on update current_timestamp;default:current_timestamp;"`

	UpdateType string `gorm:"type:varchar(10);not null;"`

	Version int `gorm:"type:int unsigned;not null;default:0;"`

	KeepAliveAt time.Time `gorm:"not null;"`
	// this alert config binding is executing on a specific `node`
	HostID string `gorm:"type:varchar(50);not null;"`
}
