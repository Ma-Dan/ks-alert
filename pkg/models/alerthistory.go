package models

import (
	"github.com/carmanzhang/ks-alert/pkg/utils/dbutil"
	"github.com/carmanzhang/ks-alert/pkg/utils/idutil"
	"github.com/golang/glog"
	"time"
)

type AlertHistory struct {
	//AlertHistoryID string `gorm:"primary_key"`
	// this ID is used for Paging
	ID                 uint64 `gorm:"primary_key;type:int(11) unsigned auto_increment;"`
	AlertHistoryID     string `gorm:"type:varchar(50);not null;"`
	AlertConfigID      string `gorm:"type:varchar(50);not null;"`
	AlertRuleGroupName string `gorm:"type:varchar(50);not null;"`

	//ProductID         string `gorm:"type:varchar(50);not null;"`
	ResourceGroupID   string `gorm:"type:varchar(50);not null;"`
	ResourceGroupName string `gorm:"type:varchar(50);not null;"`
	AlertedResource   string `gorm:"type:text;not null;"`

	ReceiverGroupID   string `gorm:"type:varchar(50);not null;"`
	ReceiverGroupName string `gorm:"type:varchar(50);not null;"`
	ReceiverGroup     string `gorm:"type:text;not null;"`

	AlertRuleGroupID  string `gorm:"type:varchar(50);not null;"`
	TriggerMetricName string `gorm:"type:text;not null;"`

	SeverityID string `gorm:"type:varchar(50);not null;" json:"severity_id"`
	SeverityCh string `gorm:"type:varchar(10);" json:"severity_ch"`

	MetricData string `gorm:"type:text;"`

	//SilenceEnable  bool      `gorm:"type:boolean;not null;default:false;"`
	SilenceStartAt time.Time `gorm:""`
	SilenceEndAt   time.Time `gorm:""`

	AlertFiredAt    time.Time `gorm:""`
	AlertRecoveryAt time.Time `gorm:""`

	RepeatSendType            uint32 `gorm:"type:varchar(10);not null;"`
	CurrentRepeatSendInterval uint32 `gorm:"type:int unsigned;not null;"`
	CumulateRepeatSendCount   uint32 `gorm:"type:int unsigned;not null;"`
	InitRepeatSendInterval    uint32 `gorm:"type:int unsigned;not null;"`
	MaxRepeatSendCount        uint32 `gorm:"type:int unsigned;not null;"`

	RequestNotificationStatus string    `gorm:"type:text;"`
	NotificationSendAt        time.Time `gorm:""`

	CreatedAt time.Time `gorm:"not null;"`
	UpdatedAt time.Time `gorm:"not null;"`
}

func CreateAlertHistory(ah *AlertHistory) (*AlertHistory, error) {
	db, err := dbutil.DBClient()

	if err != nil {
		glog.Errorln(err.Error())
		return nil, err
	}
	ah.AlertHistoryID = idutil.GetUuid36("alert_history_")
	err = db.Model(&AlertHistory{}).Create(ah).Error
	return ah, err
}

//func CreateAlertHistory1(rtAlert *runtime.RuntimeAlertConfig) (*AlertHistory, error) {
//	db, err := dbutil.DBClient()
//
//	if err != nil {
//		glog.Errorln(err.Error())
//		return nil, err
//	}
//
//	err = db.Model(&AlertHistory{}).Create(ah).Error
//	return nil, err
//}

// TODO need to implement
func GetAlertHistory(ah *AlertHistory) ([]*AlertHistory, error) {
	db, err := dbutil.DBClient()
	if err != nil {
		return nil, err
	}

	var alertHistories []AlertHistory
	db.Model(&AlertHistory{}).Where(ah).Find(&alertHistories)

	var als []*AlertHistory
	l := len(alertHistories)
	for i := 0; i < l; i++ {
		als = append(als, &alertHistories[i])
	}
	return als, nil
}

func UpdateAlertSendStatus(ah *AlertHistory, sendStatus string) error {
	db, err := dbutil.DBClient()
	if err != nil {
		return err
	}

	db.Model(ah).Where("alert_history_id = ?", ah.AlertHistoryID).Update("request_notification_status", sendStatus)

	if db.Error != nil {
		return err
	}

	return nil
}
