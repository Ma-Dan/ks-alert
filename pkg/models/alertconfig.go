package models

import (
	"github.com/carmanzhang/ks-alert/pkg/utils/idutil"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"time"
)

// signal, used to notify goroutine with different
type Signal int32

const (
	Run    Signal = iota // value --> 0
	Create               // value --> 1
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
	Action
	AlertConfigID   string `gorm:"primary_key"`
	AlertConfigName string `gorm:"type:varchar(50);not null;"`

	AlertRuleGroupID string `gorm:"type:varchar(50);not null;"`
	ResourceGroupID  string `gorm:"type:varchar(50);not null;"`
	ReceiverGroupID  string `gorm:"type:varchar(50);not null;"`

	ReceiverGroup  *ReceiverGroup  `gorm:"-"`
	AlertRuleGroup *AlertRuleGroup `gorm:"-"`
	ResourceGroup  *ResourceGroup  `gorm:"-"`

	SeverityID string `gorm:"type:varchar(50);not null;" json:"severity_id"`
	SeverityCh string `gorm:"type:varchar(10);" json:"severity_ch"`

	EnableStart time.Time `gorm:"not null;"`
	EnableEnd   time.Time `gorm:"not null;"`

	Description string `gorm:"type:text;"`

	CreatedAt time.Time `gorm:"type:timestamp;not null;default:current_timestamp;"`
	UpdatedAt time.Time `gorm:"type:timestamp;not null on update current_timestamp;default:current_timestamp;"`

	UpdateType string `gorm:"type:varchar(10);not null;"`

	Version int `gorm:"type:int unsigned;not null;default:0;"`

	KeepAliveAt time.Time `gorm:"not null;"`
	// this alert config binding is executing on a specific `node`
	HostID string `gorm:"type:varchar(50);not null;"`
}

func (r AlertConfig) Create(tx *gorm.DB, v interface{}) (interface{}, error) {
	ac, ok := v.(*AlertConfig)

	if !ok {
		return nil, errors.Errorf("type %v assert error", ac)
	}

	ac.AlertConfigID = idutil.GetUuid36("alert-config-")

	//sql := "INSERT INTO alert_configs (alert_config_id, alert_config_name, alert_rule_group_id, " +
	//	"resource_group_id, receiver_group_id, severity_id, severity_ch, enable_start, enable_end, description, created_at, updated_at) VALUES " + item

	if err := tx.Create(ac).Error; err != nil {
		return nil, err
	}

	return nil, nil
}

func (r AlertConfig) Update(tx *gorm.DB, v interface{}) (interface{}, error) {
	ac, ok := v.(*AlertConfig)

	if !ok {
		return nil, errors.Errorf("type %v assert error", ac)
	}

	if err := tx.Where("alert_config_id=?", ac.AlertConfigID).Update(ac).Error; err != nil {
		return nil, err
	}

	return nil, nil
}

func (r AlertConfig) Get(tx *gorm.DB, v interface{}) (interface{}, error) {

	ac, ok := v.(*AlertConfig)

	if !ok {
		return nil, errors.Errorf("type %v assert error", ac)
	}

	var alertConfig AlertConfig

	if err := tx.Where("alert_config_id=?", alertConfig.AlertConfigID).First(&alertConfig).Error; err != nil {
		return nil, err
	}

	return nil, nil
}

func (r AlertConfig) Delete(tx *gorm.DB, v interface{}) (interface{}, error) {
	ac, ok := v.(*AlertConfig)

	if !ok {
		return nil, errors.Errorf("type %v assert error", ac)
	}

	if ac.AlertConfigID == "" {
		return nil, errors.New("alert config id must be specified")
	}

	if err := tx.Delete(&AlertConfig{AlertConfigID: ac.AlertConfigID}).Error; err != nil {
		return nil, err
	}

	return nil, nil
}
