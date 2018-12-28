package models

import (
	"github.com/carmanzhang/ks-alert/pkg/utils/dbutil"
	"github.com/carmanzhang/ks-alert/pkg/utils/idutil"
	"time"
	"k8s.io/klog/glog"
	"github.com/pkg/errors"
	"fmt"
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
		glog.Errorln(err.Error())
		return nil, err
	}

	enterprise.EnterpriseID = idutil.GetUuid36("enterprise-")

	err = db.Model(&Enterprise{}).Create(enterprise).Error
	return enterprise, err
}

func GetEnterprise(enterprise *Enterprise) (*Enterprise, error) {
	entID := enterprise.EnterpriseID
	entName := enterprise.EnterpriseName

	db, err := dbutil.DBClient()
	if err != nil {
		glog.Errorln(err.Error())
		return &Enterprise{}, err
	}

	var ent Enterprise

	if entID != "" {
		db.Model(&Enterprise{}).Where(&Enterprise{EnterpriseID:entID}).First(&ent)
	}

	if ent.EnterpriseID == "" && entName != ""{
		db.Model(&Enterprise{}).Where(&Enterprise{EnterpriseName:entName}).First(&ent)
	}

	if ent.EnterpriseID == "" {
		errStr := fmt.Sprintf("can not find the enterprise with enterprise_id: %s or enterprise_name: %s", entID, entName)
		glog.Infoln(errStr)
		return &Enterprise{}, errors.New(errStr)
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
		glog.Errorln(err.Error())
		return err
	}

	if entID != "" {
		db.Delete(&Enterprise{EnterpriseID: entID})
		glog.Errorln(db.Error)
		return db.Error
	}else if entName != "" {
		db.Delete(&Enterprise{EnterpriseName:entName})
		glog.Errorln(db.Error)
		return db.Error
	}

	return nil
}


func UpdateEnterprise(ent *Enterprise) (error)  {
	db, err := dbutil.DBClient()

	if err != nil {
		glog.Errorln(err.Error())
		return err
	}


	if ent.EnterpriseID != "" {
		err = db.Model(ent).Where("enterprise_id = ?", ent.EnterpriseID).Update(ent).Error

	}else if ent.EnterpriseName != "" {
		err = db.Model(ent).Where("enterprise_name = ?", ent.EnterpriseName).Update(ent).Error
	}

	return err
}