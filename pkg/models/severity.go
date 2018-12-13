package models

import "time"

type Severity struct {
	SeverityID string    `gorm:"type:varchar(50);primary_key"`
	SeverityEn string    `gorm:"type:varchar(20);"`
	SeverityCh string    `gorm:"type:varchar(20);"`
	CreatedBy  string    `gorm:"type:varchar(20);"`
	CreatedAt  time.Time `gorm:"not null;"`
	UpdatedAt  time.Time `gorm:"not null;"`
}
