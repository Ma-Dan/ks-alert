package models

import (
	"fmt"
	"kubesphere.io/ks-alert/pkg/option"
	. "kubesphere.io/ks-alert/pkg/stderr"
	"kubesphere.io/ks-alert/pkg/utils/dbutil"
	"kubesphere.io/ks-alert/pkg/utils/idutil"
	"github.com/jinzhu/gorm"
	"time"
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

func (r *AlertConfig) Create(tx *gorm.DB) (interface{}, error) {

	r.AlertConfigID = idutil.GetUuid36("")

	//sql := "INSERT INTO alert_configs (alert_config_id, alert_config_name, alert_rule_group_id, " +
	//	"resource_group_id, receiver_group_id, severity_id, severity_ch, enable_start, enable_end, description, created_at, updated_at) VALUES " + item

	ruleGroup, err := r.AlertRuleGroup.Create(tx)

	if err != nil {
		return nil, err
	}

	resGroup, err := r.ResourceGroup.Create(tx)

	if err != nil {
		return nil, err
	}

	recvGroup, err := r.ReceiverGroup.Create(tx)

	if err != nil {
		return nil, err
	}

	r.AlertRuleGroupID = ruleGroup.(*AlertRuleGroup).AlertRuleGroupID
	r.ResourceGroupID = resGroup.(*ResourceGroup).ResourceGroupID
	r.ReceiverGroupID = recvGroup.(*ReceiverGroup).ReceiverGroupID

	if err := tx.Create(r).Error; err != nil {
		return nil, Error{Text: err.Error(), Code: DBError, Where: Caller(0, true)}
	}

	return r, nil
}

func UpdateAlertConfigKeepAliveTime(acID string) error {
	db, err := dbutil.DBClient()

	if err != nil {
		return Error{Text: err.Error(), Code: DBError, Where: Caller(0, true)}
	}

	now := time.Now()

	sql := fmt.Sprintf("UPDATE alert_configs SET keep_alive_at='%v', updated_at='%v' WHERE alert_config_id='%s'", now, now, acID)

	if err := db.Exec(sql).Error; err != nil {
		return Error{Text: err.Error(), Code: DBError, Where: Caller(0, true)}
	}

	return nil
}

func GetAbnormalExecutedAlertConfig(hostID string, latestReportTime time.Time, limit int) (*[]AlertConfig, error) {
	db, err := dbutil.DBClient()

	if err != nil {
		return nil, Error{Text: err.Error(), Code: DBError, Where: Caller(0, true)}
	}

	//sql := fmt.Sprintf(`SELECT * FROM alert_configs WHERE host_id='%s' AND keep_alive_at<'%v' UNION
	//	(SELECT * FROM alert_configs WHERE host_id!='%s' AND keep_alive_at<'%v' order by keep_alive_at asc limit 10)`,
	//	hostID, latestReportTime, hostID, latestReportTime)

	sql := "SELECT * FROM alert_configs WHERE host_id=? AND keep_alive_at<? " +
		"UNION " +
		"(SELECT * FROM alert_configs WHERE host_id!=? AND keep_alive_at<? order by keep_alive_at asc)"

	var alertConfigs []AlertConfig
	err = db.Raw(sql, hostID, latestReportTime, hostID, latestReportTime).Limit(limit).Scan(&alertConfigs).Error

	if err != nil {
		return nil, Error{Text: err.Error(), Code: DBError, Where: Caller(0, true)}
	}

	return &alertConfigs, nil
}

func UpdateAlertConfigBindingHostAndVersion(alertConfigs *[]AlertConfig) ([]bool, error) {
	db, err := dbutil.DBClient()

	if err != nil {
		return nil, Error{Text: err.Error(), Code: DBError, Where: Caller(0, true)}
	}

	l := len(*alertConfigs)
	var b = make([]bool, l)

	for i := 0; i < l; i++ {
		ac := (*alertConfigs)[i]

		sql := fmt.Sprintf("UPDATE alert_configs SET host_id='%s', version='%d', keep_alive_at='%v', updated_at='%v' WHERE alert_config_id='%s' AND version='%d' ",
			ac.HostID, ac.Version, ac.KeepAliveAt, ac.UpdatedAt, ac.AlertConfigID, ac.Version-1)

		if db.Exec(sql).RowsAffected == 0 {
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
		return "", Error{Text: err.Error(), Code: DBError, Where: Caller(0, true)}
	}

	var alertConfig AlertConfig

	if err := db.Where("alert_config_id=?", acID).First(&alertConfig).Error; err != nil {
		return "", Error{Text: err.Error(), Code: DBError, Where: Caller(0, true)}
	}

	if alertConfig.AlertConfigID == "" {
		return "", Error{Text: "alert config id does not exist", Code: InvalidParam, Where: Caller(0, true)}
	}

	return alertConfig.HostID, nil
}

func (r *AlertConfig) Update(tx *gorm.DB) (interface{}, error) {

	var alertConfig AlertConfig

	if err := tx.Where("alert_config_id=?", r.AlertConfigID).First(&alertConfig).Error; err != nil {
		return nil, Error{Text: err.Error(), Code: DBError, Where: Caller(0, true)}
	}

	if alertConfig.AlertConfigID == "" {
		return nil, Error{Text: "alert config id does not exist", Code: InvalidParam, Where: Caller(0, true)}
	}

	r.AlertRuleGroup.AlertRuleGroupID = alertConfig.AlertRuleGroupID
	ruleGroup, err := r.AlertRuleGroup.Update(tx)

	if err != nil {
		return nil, err
	}

	r.ResourceGroup.ResourceGroupID = alertConfig.ResourceGroupID
	resGroup, err := r.ResourceGroup.Update(tx)

	if err != nil {
		return nil, err
	}

	r.ReceiverGroup.ReceiverGroupID = alertConfig.ReceiverGroupID
	recvGroup, err := r.ReceiverGroup.Update(tx)

	if err != nil {
		return nil, err
	}

	r.AlertRuleGroupID = ruleGroup.(*AlertRuleGroup).AlertRuleGroupID
	r.ResourceGroupID = resGroup.(*ResourceGroup).ResourceGroupID
	r.ReceiverGroupID = recvGroup.(*ReceiverGroup).ReceiverGroupID

	if err := tx.Model(AlertConfig{}).Where("alert_config_id=?", r.AlertConfigID).Update(r).Error; err != nil {
		return nil, Error{Text: err.Error(), Code: DBError, Where: Caller(0, true)}
	}

	return r, nil
}

func GetAlertConfig(ac *AlertConfig) (*AlertConfig, error) {
	db, err := dbutil.DBClient()
	if err != nil {
		return nil, Error{Text: err.Error(), Code: DBError, Where: Caller(0, true)}
	}

	v, err := ac.Get(db)

	if err != nil {
		return nil, err
	}

	alertConfig := v.(*AlertConfig)
	return alertConfig, nil
}

func (r *AlertConfig) Get(tx *gorm.DB) (interface{}, error) {

	var alertConfig AlertConfig

	if err := tx.Where("alert_config_id=?", r.AlertConfigID).First(&alertConfig).Error; err != nil {
		return nil, Error{Text: err.Error(), Code: DBError, Where: Caller(0, true)}
	}

	if alertConfig.AlertConfigID == "" {
		return nil, Error{Text: "alert config id does not exist", Code: InvalidParam, Where: Caller(0, true)}
	}

	ruleGroup, err := (&AlertRuleGroup{AlertRuleGroupID: alertConfig.AlertRuleGroupID}).Get(tx)

	if err != nil {
		return nil, err
	}

	resGroup, err := (&ResourceGroup{ResourceGroupID: alertConfig.ResourceGroupID}).Get(tx)

	if err != nil {
		return nil, err
	}

	recvGroup, err := (&ReceiverGroup{ReceiverGroupID: alertConfig.ReceiverGroupID}).Get(tx)

	if err != nil {
		return nil, err
	}

	alertConfig.AlertRuleGroup = ruleGroup.(*AlertRuleGroup)
	alertConfig.ResourceGroup = resGroup.(*ResourceGroup)
	alertConfig.ReceiverGroup = recvGroup.(*ReceiverGroup)

	return &alertConfig, nil
}

func (r *AlertConfig) Delete(tx *gorm.DB) (interface{}, error) {

	if r.AlertConfigID == "" {
		return nil, Error{Text: "alert config id must be specified", Code: InvalidParam, Where: Caller(0, true)}
	}

	// firstly, get alert config
	var alertConfig AlertConfig

	if err := tx.Where("alert_config_id=?", r.AlertConfigID).First(&alertConfig).Error; err != nil {
		return nil, Error{Text: err.Error(), Code: DBError, Where: Caller(0, true)}
	}

	if alertConfig.AlertConfigID == "" {
		return nil, Error{Text: "alert config id does not exist", Code: InvalidParam, Where: Caller(0, true)}
	}

	// secondly, delete three groups
	_, err := (&AlertRuleGroup{AlertRuleGroupID: alertConfig.AlertRuleGroupID}).Delete(tx)

	if err != nil {
		return nil, err
	}

	_, err = (&ResourceGroup{ResourceGroupID: alertConfig.ResourceGroupID}).Delete(tx)

	if err != nil {
		return nil, err
	}

	_, err = (&ReceiverGroup{ReceiverGroupID: alertConfig.ReceiverGroupID}).Delete(tx)

	if err != nil {
		return nil, err
	}

	if err := tx.Delete(&AlertConfig{AlertConfigID: alertConfig.AlertConfigID}).Error; err != nil {
		return nil, Error{Text: err.Error(), Code: DBError, Where: Caller(0, true)}
	}

	return nil, nil
}

func GetAlertConfigRows() (map[string]string, error) {
	var alertConfigIDMap = make(map[string]string)

	db, err := dbutil.DBClient()
	if err != nil {
		return alertConfigIDMap, Error{Text: err.Error(), Code: DBError, Where: Caller(0, true)}
	}

	rows, err := db.Raw("select alert_config_id from alert_configs where host_id = ?", option.HostInfo).Rows()

	defer rows.Close()

	for rows.Next() {
		var acID string
		rows.Scan(&acID)
		alertConfigIDMap[acID] = ""
	}

	return alertConfigIDMap, nil
}
