package models

import "time"

// signal, used to notify goroutine with different
type Signal int32

const (
	Run       Signal = iota // value --> 0
	Create                  // value --> 1
	Terminate
	Reload
	Stop
)

type UpdateType string

const (
	ALertRuleUpdate UpdateType = "alert_rule"
	ResourceUpdate  UpdateType = "resource"
	ReceiverUpdate  UpdateType = "receiver"
)

type AlertConfig struct {
	AlertConfigID  string         `json:"alert_config_id, omitempty"`
	UpdateType     UpdateType     `json:"update_type, omitempty"`
	ReceiverGroup  ReceiverGroup  `json:"receiver_group, omitempty"`
	AlertRuleGroup AlertRuleGroup `json:"alert_rule_group, omitempty"`
	ResourceGroup  ResourceGroup  `json:"resource_group, omitempty"`
	URIParams      Params         `json:"resource_uri_params, omitempty"`
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

func GetAlertBindingItem(alertConfigID string) (*AlertBinding, error) {
	return nil, nil
}

func DeleteAlertBindingItem(alertConfigID string) error {
	return nil
}