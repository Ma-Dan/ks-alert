package models

import "time"

type Silence struct {
	ResourceRuleID            string    `gorm:"type:varchar(50);primary_key"`
	ResourceID                string    `gorm:"type:varchar(50);not null;"`
	AlertRuleID               string    `gorm:"type:varchar(50);not null;"`
	SilenceEnable             bool      `gorm:"type:bool;"`
	SilenceStartAt            time.Time `gorm:"not null;"`
	SilenceEndAt              time.Time `gorm:"not null;"`
	CurrentRepeatSendInterval uint32    `gorm:"type:mediumint unsigned;not null;"`
	CurrentRepeatSendCount    uint16    `gorm:"type:smallint unsigned;not null;"`
	CreateAt                  time.Time `gorm:"type:varchar(50);"`
	UpdateAt                  time.Time `gorm:"type:varchar(50);"`
}
