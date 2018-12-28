package models

import (
	"k8s.io/klog/glog"
	"github.com/carmanzhang/ks-alert/pkg/utils/dbutil"
	"github.com/carmanzhang/ks-alert/pkg/utils/idutil"
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
	MonitorCenterPort int32          `gorm:"type:int;not null;"`
	HomePage          string         `gorm:"type:varchar(128);not null;"`
	Address           string         `gorm:"type:varchar(128);not null;"`
	Description       string         `gorm:"type:text;"`
	CreatedAt         time.Time      `gorm:"not null;"`
	UpdatedAt         time.Time      `gorm:"not null;"`
	ResourceTypes     []ResourceType `gorm:"ForeignKey:ProductID;AssociationForeignKey:ResourceTypeID"`
}

func CreateProduct(product *Product) (*Product, error) {
	db, err := dbutil.DBClient()

	if err != nil {
		glog.Errorln(err.Error())
		return nil, err
	}

	product.ProductID = idutil.GetUuid36("product-")

	err = db.Model(&Product{}).Create(product).Error
	return product, err
}

func GetProduct(product *Product) (*Product, error) {

	db, err := dbutil.DBClient()
	if err != nil {
		glog.Errorln(err.Error())
		return nil, err
	}

	var products Product
	db.Model(&Product{}).Where(product).Find(&products)
	return &products, err
}

func DeleteProduct(prod *Product) error {
	prodID := prod.ProductID
	prodName := prod.ProductName

	if _, err := GetProduct(prod); err != nil {
		return err
	}

	db, err := dbutil.DBClient()

	if err != nil {
		glog.Errorln(err.Error())
		return err
	}

	if prodID != "" {
		db.Delete(&Product{ProductID: prodID})
		glog.Errorln(db.Error)
		return db.Error
	}else if prodName != "" {
		db.Delete(&Product{ProductName: prodName})
		glog.Errorln(db.Error)
		return db.Error
	}

	// TODO need to delete related items in table `resource type` `alert rule` ...
 	return nil
}


func UpdateProduct(prod *Product) (error)  {
	db, err := dbutil.DBClient()

	if err != nil {
		glog.Errorln(err.Error())
		return err
	}

	if prod.ProductID != "" {
		err = db.Model(prod).Where("product_id = ?", prod.ProductID).Update(prod).Error

	}else if prod.ProductName != "" {
		err = db.Model(prod).Where("product_name = ?", prod.ProductName).Update(prod).Error
	}

	return err
}