package models

import (
	"fmt"
	"github.com/carmanzhang/ks-alert/pkg/dispatcher/pb"
	"github.com/carmanzhang/ks-alert/pkg/utils/dbutil"
	"github.com/carmanzhang/ks-alert/pkg/utils/idutil"
	"github.com/pkg/errors"
	"time"
)

type Severity struct {
	SeverityID string    `gorm:"type:varchar(50);primary_key"`
	ProductID  string    `gorm:"type:varchar(50);"`
	SeverityEn string    `gorm:"type:varchar(20);"`
	SeverityCh string    `gorm:"type:varchar(20);not null;"`
	CreatedAt  time.Time `gorm:"not null;"`
	UpdatedAt  time.Time `gorm:"not null;"`
}

func CreateSeverity(severity *Severity) (*Severity, error) {

	db, err := dbutil.DBClient()

	if err != nil {
		return nil, err
	}

	severity.SeverityID = idutil.GetUuid36("severity-")

	err = db.Model(&Severity{}).Create(severity).Error

	if err != nil {
		return nil, err
	}

	return severity, nil
}

func UpdateSeverity(severity *Severity) (*Severity, error) {

	db, err := dbutil.DBClient()

	if err != nil {
		return nil, err
	}

	sql := fmt.Sprintf("UPDATE severities SET severity_en='%s',severity_ch='%s',"+
		"updated_at='%v' WHERE severity_id='%s'",
		severity.SeverityEn, severity.SeverityCh,
		time.Now(), severity.SeverityID)

	fmt.Println(sql)

	if err := db.Exec(sql).Error; err != nil {
		if db.RecordNotFound() {
			return nil, errors.New("record not found")
		}
		return nil, err
	}

	return severity, nil
}

func DeleteSeverity(sevSpec *pb.SeveritySpec) (*Severity, error) {

	db, err := dbutil.DBClient()

	if err != nil {
		return nil, err
	}

	err = db.Debug().Raw("DELETE from severities WHERE severity_id=?", sevSpec.SeverityId).Error

	if err != nil {
		return nil, err
	}

	return nil, nil
}

func GetSeverity(sevSpec *pb.SeveritySpec) (*[]Severity, error) {

	db, err := dbutil.DBClient()

	if err != nil {
		return nil, err
	}

	// get severity by severity_id
	var severities *[]Severity
	if sevSpec.SeverityId != "" {
		var sev Severity
		// db.Raw("SELECT name, age FROM users WHERE name = ?", 3).Scan(&result)
		err = db.Debug().Raw("SELECT * from severities WHERE severity_id=?", sevSpec.SeverityId).Scan(&sev).Error

		if err != nil {
			return nil, err
		}

		if sev.SeverityID != "" {
			severities = &[]Severity{sev}
		}
	}
	// get severity by product_id
	if sevSpec.ProductId != "" {
		var sevs []Severity
		//err = db.Debug().Exec("SELECT * from severities WHERE (product_id = ?)", sevSpec.ProductId).Find(&sevs).Error
		err = db.Debug().Find(&sevs, "product_id = ?", sevSpec.ProductId).Error

		if err != nil {
			return nil, err
		}

		severities = &sevs
	}

	return severities, nil
}
