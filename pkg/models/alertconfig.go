package models

import (
	"fmt"
	"github.com/carmanzhang/ks-alert/pkg/utils/dbutil"
	"github.com/carmanzhang/ks-alert/pkg/utils/idutil"
	"github.com/jinzhu/gorm"
	"k8s.io/klog/glog"
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
		return nil, Error{Text: fmt.Sprintf("type %v assert error", ac), Code: AssertError}
	}

	ac.AlertConfigID = idutil.GetUuid36("alert-config-")

	//sql := "INSERT INTO alert_configs (alert_config_id, alert_config_name, alert_rule_group_id, " +
	//	"resource_group_id, receiver_group_id, severity_id, severity_ch, enable_start, enable_end, description, created_at, updated_at) VALUES " + item

	ruleGroup, err := AlertRuleGroup{}.Create(tx, ac.AlertRuleGroup)

	if err != nil {
		return nil, err
	}

	resGroup, err := ResourceGroup{}.Create(tx, ac.ResourceGroup)

	if err != nil {
		return nil, err
	}

	recvGroup, err := ReceiverGroup{}.Create(tx, ac.ReceiverGroup)

	if err != nil {
		return nil, err
	}

	ac.AlertRuleGroupID = ruleGroup.(*AlertRuleGroup).AlertRuleGroupID
	ac.ResourceGroupID = resGroup.(*ResourceGroup).ResourceGroupID
	ac.ReceiverGroupID = recvGroup.(*ReceiverGroup).ReceiverGroupID

	if err := tx.Create(ac).Error; err != nil {
		return nil, Error{Text: err.Error(), Code: DBError}
	}

	return ac, nil
}

func (r AlertConfig) Update(tx *gorm.DB, v interface{}) (interface{}, error) {
	ac, ok := v.(*AlertConfig)

	if !ok {
		return nil, Error{Text: fmt.Sprintf("type %v assert error", ac), Code: AssertError}
	}

	var alertConfig AlertConfig

	if err := tx.Where("alert_config_id=?", ac.AlertConfigID).First(&alertConfig).Error; err != nil {
		return nil, Error{Text: err.Error(), Code: DBError}
	}

	if alertConfig.AlertConfigID == "" {
		return nil, Error{Text: "alert config id does not exist", Code: InvalidParam}
	}

	ac.AlertRuleGroup.AlertRuleGroupID = alertConfig.AlertRuleGroupID
	ruleGroup, err := AlertRuleGroup{}.Update(tx, ac.AlertRuleGroup)

	if err != nil {
		return nil, err
	}

	ac.ResourceGroup.ResourceGroupID = alertConfig.ResourceGroupID
	resGroup, err := ResourceGroup{}.Update(tx, ac.ResourceGroup)

	if err != nil {
		return nil, err
	}

	ac.ReceiverGroup.ReceiverGroupID = alertConfig.ReceiverGroupID
	recvGroup, err := ReceiverGroup{}.Update(tx, ac.ReceiverGroup)

	if err != nil {
		return nil, err
	}

	ac.AlertRuleGroupID = ruleGroup.(*AlertRuleGroup).AlertRuleGroupID
	ac.ResourceGroupID = resGroup.(*ResourceGroup).ResourceGroupID
	ac.ReceiverGroupID = recvGroup.(*ReceiverGroup).ReceiverGroupID

	if err := tx.Debug().Model(AlertConfig{}).Where("alert_config_id=?", ac.AlertConfigID).Update(ac).Error; err != nil {
		return nil, Error{Text: err.Error(), Code: DBError}
	}

	return ac, nil
}

func GetAlertConfig(ac *AlertConfig) (*AlertConfig, error) {
	db, err := dbutil.DBClient()
	if err != nil {
		return nil, Error{Text: err.Error(), Code: DBError}
	}

	v, err := ac.Get(db, ac)

	if err != nil {
		return nil, err
	}

	alertConfig := v.(*AlertConfig)
	return alertConfig, nil
}

func (r AlertConfig) Get(tx *gorm.DB, v interface{}) (interface{}, error) {

	ac, ok := v.(*AlertConfig)

	if !ok {
		return nil, Error{Text: fmt.Sprintf("type %v assert error", ac), Code: AssertError}
	}

	var alertConfig AlertConfig

	if err := tx.Where("alert_config_id=?", ac.AlertConfigID).First(&alertConfig).Error; err != nil {
		return nil, Error{Text: err.Error(), Code: DBError}
	}

	if alertConfig.AlertConfigID == "" {
		return nil, Error{Text: "alert config id does not exist", Code: InvalidParam}
	}

	ruleGroup, err := AlertRuleGroup{}.Get(tx, &AlertRuleGroup{AlertRuleGroupID: alertConfig.AlertRuleGroupID})

	// TODO maybe need return this err?
	if err != nil {
		glog.Errorln(err.Error())
	}

	resGroup, err := ResourceGroup{}.Get(tx, &ResourceGroup{ResourceGroupID: alertConfig.ResourceGroupID})

	if err != nil {
		glog.Errorln(err.Error())
	}

	recvGroup, err := ReceiverGroup{}.Get(tx, &ReceiverGroup{ReceiverGroupID: alertConfig.ReceiverGroupID})

	if err != nil {
		glog.Errorln(err.Error())
	}

	alertConfig.AlertRuleGroup = ruleGroup.(*AlertRuleGroup)
	alertConfig.ResourceGroup = resGroup.(*ResourceGroup)
	alertConfig.ReceiverGroup = recvGroup.(*ReceiverGroup)

	return &alertConfig, nil
}

func (r AlertConfig) Delete(tx *gorm.DB, v interface{}) (interface{}, error) {
	ac, ok := v.(*AlertConfig)

	if !ok {
		return nil, Error{Text: fmt.Sprintf("type %v assert error", ac), Code: AssertError}
	}

	if ac.AlertConfigID == "" {
		return nil, Error{Text: "alert config id must be specified", Code: InvalidParam}
	}

	// firstly, get alert config
	var alertConfig AlertConfig

	if err := tx.Where("alert_config_id=?", ac.AlertConfigID).First(&alertConfig).Error; err != nil {
		return nil, Error{Text: err.Error(), Code: DBError}
	}

	if alertConfig.AlertConfigID == "" {
		return nil, Error{Text: "alert config id does not exist", Code: InvalidParam}
	}

	// secondly, delete three groups
	_, err := AlertRuleGroup{}.Delete(tx, &AlertRuleGroup{AlertRuleGroupID: alertConfig.AlertRuleGroupID})

	if err != nil {
		return nil, err
	}

	_, err = ResourceGroup{}.Delete(tx, &ResourceGroup{ResourceGroupID: alertConfig.ResourceGroupID})

	if err != nil {
		return nil, err
	}

	_, err = ReceiverGroup{}.Delete(tx, &ReceiverGroup{ReceiverGroupID: alertConfig.ReceiverGroupID})

	if err != nil {
		return nil, err
	}

	if err := tx.Delete(&AlertConfig{AlertConfigID: alertConfig.AlertConfigID}).Error; err != nil {
		return nil, Error{Text: err.Error(), Code: DBError}
	}

	return nil, nil
}
