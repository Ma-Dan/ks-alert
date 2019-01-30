package models

import (
	"fmt"
	. "github.com/carmanzhang/ks-alert/pkg/stderr"
	"github.com/carmanzhang/ks-alert/pkg/utils/dbutil"
	"github.com/carmanzhang/ks-alert/pkg/utils/idutil"
	"time"
)

type Enterprise struct {
	EnterpriseID   string    `gorm:"primary_key"`
	EnterpriseName string    `gorm:"type:varchar(50);not null;unique"`
	Contacts       string    `gorm:"type:varchar(50);not null;"`
	Email          string    `gorm:"type:varchar(50);not null;"`
	Phone          string    `gorm:"type:varchar(50);not null;"`
	Address        string    `gorm:"type:varchar(128);not null;"`
	HomePage       string    `gorm:"type:varchar(128);not null;"`
	Description    string    `gorm:"type:text;"`
	CreatedAt      time.Time `gorm:"not null;"`
	UpdatedAt      time.Time `gorm:"not null;"`
	Products       []Product `gorm:"ForeignKey:EnterpriseID;AssociationForeignKey:ProductID"`
}

func CreateEnterprise(enterprise *Enterprise) (*Enterprise, error) {
	db, err := dbutil.DBClient()

	if err != nil {
		return nil, Error{Text: err.Error(), Code: DBError, Where: Caller(1, true)}
	}

	enterprise.EnterpriseID = idutil.GetUuid36("")

	err = db.Model(&Enterprise{}).Create(enterprise).Error

	if err != nil {
		return nil, Error{Text: err.Error(), Code: DBError, Where: Caller(1, true)}
	}

	return enterprise, nil
}

func GetEnterprise(enterprise *Enterprise) (*Enterprise, error) {
	entID := enterprise.EnterpriseID
	entName := enterprise.EnterpriseName

	db, err := dbutil.DBClient()

	if err != nil {
		return nil, Error{Text: err.Error(), Code: DBError, Where: Caller(1, true)}
	}

	var ent Enterprise

	if entID != "" {
		db.Model(&Enterprise{}).Where(&Enterprise{EnterpriseID: entID}).First(&ent)
	}

	if ent.EnterpriseID == "" && entName != "" {
		db.Model(&Enterprise{}).Where(&Enterprise{EnterpriseName: entName}).First(&ent)
	}

	if ent.EnterpriseID == "" {
		errStr := fmt.Sprintf("can not find the enterprise with enterprise_id: %s or enterprise_name: %s", entID, entName)
		return nil, Error{Text: errStr, Code: DBError, Where: Caller(1, true)}
	}
	return &ent, nil
}

func DeleteEnterprise(enterprise *Enterprise) error {
	entID := enterprise.EnterpriseID
	entName := enterprise.EnterpriseName

	if _, err := GetEnterprise(enterprise); err != nil {
		return err
	}

	db, err := dbutil.DBClient()

	if err != nil {
		return Error{Text: err.Error(), Code: DBError, Where: Caller(1, true)}
	}

	if entID != "" {
		err = db.Delete(&Enterprise{EnterpriseID: entID}).Error
	} else if entName != "" {
		err = db.Delete(&Enterprise{EnterpriseName: entName}).Error
	}

	if err != nil {
		return Error{Text: err.Error(), Code: DBError, Where: Caller(1, true)}
	}

	return nil
}

func UpdateEnterprise(ent *Enterprise) error {
	db, err := dbutil.DBClient()

	if err != nil {
		return Error{Text: err.Error(), Code: DBError, Where: Caller(1, true)}
	}

	if ent.EnterpriseID != "" {
		err = db.Model(ent).Where("enterprise_id = ?", ent.EnterpriseID).Update(ent).Error

	} else if ent.EnterpriseName != "" {
		err = db.Model(ent).Where("enterprise_name = ?", ent.EnterpriseName).Update(ent).Error
	}

	if err != nil {
		return Error{Text: err.Error(), Code: DBError, Where: Caller(1, true)}
	}

	return nil
}
