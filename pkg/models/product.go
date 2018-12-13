package models

import (
	"k8s.io/klog/glog"
	"kubesphere.io/ks-alert/pkg/utils/dbutil"
	"kubesphere.io/ks-alert/pkg/utils/idutil"
	"time"
)

type Product struct {
	ProductID         string         `gorm:"primary_key"`
	ProductName       string         `gorm:"type:varchar(50);not null;unique"`
	EnterpriseID      string         `gorm:"not null;"`
	Phone             string         `gorm:"type:varchar(50);not null;"`
	Contacts          string         `gorm:"type:varchar(50);not null;"`
	Email             string         `gorm:"type:varchar(50);not null;"`
	MonitorCenterHost string         `gorm:"type:varchar(128);not null;"`
	MonitorCenterPort int            `gorm:"type:int;not null;"`
	HomePage          string         `gorm:"type:varchar(128);not null;"`
	Description       string         `gorm:"type:text;"`
	CreatedAt         time.Time      `gorm:"not null;"`
	UpdatedAt         time.Time      `gorm:"not null;"`
	ResourceTypes     []ResourceType `gorm:"ForeignKey:ProductID;AssociationForeignKey:ResourceTypeID"`
}

func CreateProduct(product *Product) error {
	db, err := dbutil.DBClient()
	if err != nil {
		panic(err)
	}

	product.ProductID = idutil.GetUuid36("product-")

	err = db.Model(&Product{}).Create(product).Error
	return err
}

func GetProducts(product *Product) (*[]Product, error) {

	db, err := dbutil.DBClient()
	if err != nil {
		panic(err)
	}

	var products []Product
	db.Model(&Product{}).Where(product).Find(&products)
	return &products, err
}

func GetProductsByEnterprise(enterprise *Enterprise) (*[]Product, error) {
	ent, err := GetEnterprise(enterprise)

	if err != nil {
		glog.Errorln(err.Error())
	}

	db, err := dbutil.DBClient()
	if err != nil {
		panic(err)
	}
	var products []Product
	db.Model(&Product{}).Where(&Product{EnterpriseID: ent.EnterpriseID}).Find(&products)
	return &products, err
}
