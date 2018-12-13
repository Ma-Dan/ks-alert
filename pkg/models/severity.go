package models

import (
	"kubesphere.io/ks-alert/pkg/utils/dbutil"
	"kubesphere.io/ks-alert/pkg/utils/idutil"
	"time"
)

type Severity struct {
	SeverityID     string    `gorm:"type:varchar(50);primary_key"`
	ProductID      string    `gorm:"type:varchar(50);"`
	SeverityEn     string    `gorm:"type:varchar(20);"`
	SeverityCh     string    `gorm:"type:varchar(20);not null;"`
	SeverityDegree int       `gorm:"type:int;"`
	CreatedBy      string    `gorm:"type:varchar(20);"`
	CreatedAt      time.Time `gorm:"not null;"`
	UpdatedAt      time.Time `gorm:"not null;"`
}

// Create a batch of Severitys
func CreateSeverities(severitys *[]Severity) error {
	db, err := dbutil.DBClient()
	if err != nil {
		panic(err)
	}

	tx := db.Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if tx.Error != nil {
		return err
	}

	for _, severity := range *severitys {
		severity.SeverityID = idutil.GetUuid36("severity-")
		err = db.Model(&Severity{}).Create(severity).Error
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit().Error
}

// Create Single Severity
func CreateSeverity(severity *Severity) (*Severity, error) {

	db, err := dbutil.DBClient()
	if err != nil {
		panic(err)
	}

	severity.SeverityID = idutil.GetUuid36("severity-")

	err = db.Model(&Severity{}).Create(severity).Error

	return severity, err

}

func GetSeveritiesByProductID(productID string) *[]Severity {

	db, err := dbutil.DBClient()
	if err != nil {
		panic(err)
	}

	var severities []Severity
	db.Model(&Severity{}).Where(&Severity{ProductID: productID}).Find(&severities)

	return &severities
}
