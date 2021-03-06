package models

import (
	"fmt"
	"kubesphere.io/ks-alert/pkg/pb"
	. "kubesphere.io/ks-alert/pkg/stderr"
	"kubesphere.io/ks-alert/pkg/utils/dbutil"
	"kubesphere.io/ks-alert/pkg/utils/idutil"
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
		return nil, Error{Text: err.Error(), Code: DBError, Where: Caller(0, true)}
	}

	// check product exists
	product, err := GetProduct(&Product{ProductID: severity.ProductID})

	if err != nil {
		return nil, err
	}

	if product.ProductID == "" {
		return nil, Error{Text: "product not found", Code: InvalidParam, Where: Caller(0, true)}
	}

	severity.SeverityID = idutil.GetUuid36("")

	err = db.Model(&Severity{}).Create(severity).Error

	if err != nil {
		return nil, Error{Text: err.Error(), Code: DBError, Where: Caller(0, true)}
	}

	return severity, nil
}

func UpdateSeverity(severity *Severity) (*Severity, error) {

	db, err := dbutil.DBClient()

	if err != nil {
		return nil, Error{Text: err.Error(), Code: DBError, Where: Caller(0, true)}
	}

	sql := fmt.Sprintf("UPDATE severities SET severity_en='%s',severity_ch='%s',"+
		"updated_at='%v' WHERE severity_id='%s'",
		severity.SeverityEn, severity.SeverityCh,
		time.Now(), severity.SeverityID)

	if err := db.Exec(sql).Error; err != nil {
		if db.RecordNotFound() {
			return nil, Error{Text: "record not found", Code: DBError, Where: Caller(0, true)}
		}
		return nil, Error{Text: err.Error(), Code: DBError, Where: Caller(0, true)}
	}

	return severity, nil
}

func DeleteSeverity(sevSpec *pb.SeveritySpec) (*Severity, error) {

	db, err := dbutil.DBClient()

	if err != nil {
		return nil, Error{Text: err.Error(), Code: DBError, Where: Caller(0, true)}
	}

	sID := sevSpec.SeverityId

	if sID == "" {
		return nil, Error{Text: "severity id must be specified", Code: InvalidParam, Where: Caller(0, true)}
	}

	err = db.Exec("DELETE from severities WHERE severity_id=?", sID).Error

	if err != nil {
		return nil, Error{Text: err.Error(), Code: DBError, Where: Caller(0, true)}
	}

	return nil, nil
}

func GetSeverity(sevSpec *pb.SeveritySpec) (*[]Severity, error) {

	db, err := dbutil.DBClient()

	if err != nil {
		return nil, Error{Text: err.Error(), Code: DBError, Where: Caller(0, true)}
	}

	// get severity by severity_id
	var severities *[]Severity
	if sevSpec.SeverityId != "" {
		var sev Severity
		// db.Raw("SELECT name, age FROM users WHERE name = ?", 3).Scan(&result)
		err = db.Raw("SELECT * from severities WHERE severity_id=?", sevSpec.SeverityId).Scan(&sev).Error

		if err != nil {
			return nil, Error{Text: err.Error(), Code: DBError, Where: Caller(0, true)}
		}

		if sev.SeverityID != "" {
			severities = &[]Severity{sev}
		}

		return severities, nil
	}

	// get severity by product_id
	if sevSpec.ProductId != "" {
		var sevs []Severity
		//err = db.Exec("SELECT * from severities WHERE (product_id = ?)", sevSpec.ProductId).Find(&sevs).Error
		err = db.Find(&sevs, "product_id = ?", sevSpec.ProductId).Error

		if err != nil {
			return nil, Error{Text: err.Error(), Code: DBError, Where: Caller(0, true)}
		}

		severities = &sevs
	}

	return severities, nil
}
