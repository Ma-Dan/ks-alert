package models

import (
	"fmt"
	"github.com/carmanzhang/ks-alert/pkg/option"
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

func UpdateAlertConfigKeepAliveTime(acID string) error {
	db, err := dbutil.DBClient()

	if err != nil {
		return Error{Text: err.Error(), Code: DBError}
	}

	now := time.Now()

	sql := fmt.Sprintf("UPDATE alert_configs SET keep_alive_at='%v', updated_at='%v' WHERE alert_config_id='%s'", now, now, acID)

	if err := db.Exec(sql).Error; err != nil {
		return Error{Text: err.Error(), Code: DBError}
	}

	return nil
}

func GetAbnormalExecutedAlertConfig(hostID string, latestReportTime time.Time, limit int) (*[]AlertConfig, error) {
	db, err := dbutil.DBClient()

	if err != nil {
		return nil, Error{Text: err.Error(), Code: DBError}
	}

	//sql := fmt.Sprintf(`SELECT * FROM alert_configs WHERE host_id='%s' AND keep_alive_at<'%v' UNION
	//	(SELECT * FROM alert_configs WHERE host_id!='%s' AND keep_alive_at<'%v' order by keep_alive_at asc limit 10)`,
	//	hostID, latestReportTime, hostID, latestReportTime)

	sql := "SELECT * FROM alert_configs WHERE host_id=? AND keep_alive_at<? " +
		"UNION " +
		"(SELECT * FROM alert_configs WHERE host_id!=? AND keep_alive_at<? order by keep_alive_at asc)"

	var alertConfigs []AlertConfig
	err = db.Debug().Raw(sql, hostID, latestReportTime, hostID, latestReportTime).Limit(limit).Scan(&alertConfigs).Error

	return &alertConfigs, err
}

func UpdateAlertConfigBindingHostAndVersion(alertConfigs *[]AlertConfig) ([]bool, error) {
	db, err := dbutil.DBClient()

	if err != nil {
		return nil, Error{Text: err.Error(), Code: DBError}
	}

	l := len(*alertConfigs)
	var b = make([]bool, l)

	for i := 0; i < l; i++ {
		ac := (*alertConfigs)[i]

		sql := fmt.Sprintf("UPDATE alert_configs SET host_id='%s', version='%d', keep_alive_at='%v', updated_at='%v' WHERE alert_config_id='%s' AND version='%d' ",
			ac.HostID, ac.Version, ac.KeepAliveAt, ac.UpdatedAt, ac.AlertConfigID, ac.Version-1)

		if db.Debug().Exec(sql).RowsAffected == 0 {
			b[i] = false
		} else {
			b[i] = true
		}
	}

	return b, nil
}

func GetAlertConfigBindingHost(acID string) (string, error) {
	db, err := dbutil.DBClient()

	if err != nil {
		return "", Error{Text: err.Error(), Code: DBError}
	}

	var alertConfig AlertConfig

	if err := db.Where("alert_config_id=?", acID).First(&alertConfig).Error; err != nil {
		return "", Error{Text: err.Error(), Code: DBError}
	}

	if alertConfig.AlertConfigID == "" {
		return "", Error{Text: "alert config id does not exist", Code: InvalidParam}
	}

	return alertConfig.HostID, nil
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

func GetAlertConfigRows() (map[string]string, error) {
	var alertConfigIDMap = make(map[string]string)

	db, err := dbutil.DBClient()
	if err != nil {
		return alertConfigIDMap, Error{Text: err.Error(), Code: DBError}
	}

	hostID := fmt.Sprintf("%s:%d", *option.ServiceHost, *option.ExecutorServicePort)

	rows, err := db.Raw("select alert_config_id from alert_configs where host_id = ?", hostID).Rows()

	defer rows.Close()

	for rows.Next() {
		var acID string
		rows.Scan(&acID)
		alertConfigIDMap[acID] = ""
	}

	return alertConfigIDMap, nil
}
