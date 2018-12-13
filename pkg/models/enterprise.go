package models

import (
	"kubesphere.io/ks-alert/pkg/utils/dbutil"
	"kubesphere.io/ks-alert/pkg/utils/idutil"
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

func CreateEnterprise(enterprise *Enterprise) error {
	db, err := dbutil.DBClient()
	if err != nil {
		panic(err)
	}

	enterprise.EnterpriseID = idutil.GetUuid36("enterprise-")

	err = db.Model(&Enterprise{}).Create(enterprise).Error
	return err
}

func GetEnterprise(enterprise *Enterprise) (*Enterprise, error) {
	db, err := dbutil.DBClient()
	if err != nil {
		panic(err)
	}
	var ent Enterprise
	db.Model(&Enterprise{}).Where(enterprise).First(&ent)
	return &ent, err
}
