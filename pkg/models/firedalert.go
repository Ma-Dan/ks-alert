package models

import (
	"github.com/carmanzhang/ks-alert/pkg/utils/dbutil"
	"time"
)

// use to control time to send fired alert to user or webhooks(use webhook and system wehhook)
type FiredAlert struct {
	FiredAlertID              string    `gorm:"primary_key"`
	ResourceID                string    `gorm:"type:varchar(50);not null;"`
	AlertRuleID               string    `gorm:"type:varchar(50);not null;"`
	CurrentRepeatSendInterval uint32    `gorm:"type:int unsigned;not null;"`
	CurrentRepeatSendCount    uint32    `gorm:"type:int unsigned;not null;"`
	SilenceStartAt            time.Time `gorm:"not null;"`
	SilenceEndAt              time.Time `gorm:"not null;"`
	CreatedAt                 time.Time `gorm:"not null;"`
	UpdatedAt                 time.Time `gorm:"not null;"`
}

func GetFiredAlert(alert *FiredAlert) (*FiredAlert, error) {
	db, err := dbutil.DBClient()

	if err != nil {
		return nil, Error{Text: err.Error(), Code: DBError}
	}

	// get fired alert by
	resID := alert.ResourceID
	ruleID := alert.AlertRuleID

	if resID == "" || ruleID == "" {
		return nil, Error{Text: "resource id and rule id must be specified", Code: InvalidParam}
	}

	var firedAlert FiredAlert
	if err := db.Where("resource_id=? AND alert_rule_id=?", resID, ruleID).First(&firedAlert).Error; err != nil {
		return nil, Error{Text: err.Error(), Code: DBError}
	}

	return &firedAlert, nil
}
