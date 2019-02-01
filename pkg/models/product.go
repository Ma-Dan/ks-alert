package models

import (
	. "github.com/carmanzhang/ks-alert/pkg/stderr"
	"github.com/carmanzhang/ks-alert/pkg/utils/dbutil"
	"github.com/carmanzhang/ks-alert/pkg/utils/idutil"
	"time"
)

type Product struct {
	ProductID         string    `gorm:"primary_key"`
	ProductName       string    `gorm:"type:varchar(50);not null;unique"`
	EnterpriseID      string    `gorm:"not null;"`
	Phone             string    `gorm:"type:varchar(50);not null;"`
	Contacts          string    `gorm:"type:varchar(50);not null;"`
	Email             string    `gorm:"type:varchar(50);not null;"`
	MonitorCenterHost string    `gorm:"type:varchar(128);not null;"`
	MonitorCenterPort int32     `gorm:"type:int;not null;"`
	HomePage          string    `gorm:"type:varchar(128);not null;"`
	Address           string    `gorm:"type:varchar(128);not null;"`
	Description       string    `gorm:"type:text;"`
	CreatedAt         time.Time `gorm:"not null;"`
	UpdatedAt         time.Time `gorm:"not null;"`

	// TODO each product may has it's own webhook address, this webhook mainly used to dispaly fired alert on UI
	Webhook       string `gorm:"type:varchar(50);"`
	WebhookEnable bool   `gorm:"type:bool;"`
}

func CreateProduct(product *Product) (*Product, error) {
	db, err := dbutil.DBClient()

	if err != nil {
		return nil, Error{Text: err.Error(), Code: DBError, Where: Caller(0, true)}
	}

	product.ProductID = idutil.GetUuid36("")

	err = db.Model(&Product{}).Create(product).Error

	if err != nil {
		return nil, Error{Text: err.Error(), Code: DBError, Where: Caller(0, true)}
	}

	return product, err
}

func GetProduct(product *Product) (*Product, error) {

	db, err := dbutil.DBClient()

	if err != nil {
		return nil, Error{Text: err.Error(), Code: DBError, Where: Caller(0, true)}
	}

	var products Product
	err = db.Model(&Product{}).Where(product).First(&products).Error

	if err != nil {
		return nil, Error{Text: err.Error(), Code: DBError, Where: Caller(0, true)}
	}

	return &products, err
}

func DeleteProduct(prod *Product) error {
	prodID := prod.ProductID
	prodName := prod.ProductName

	if _, err := GetProduct(prod); err != nil {
		return Error{Text: err.Error(), Code: DBError, Where: Caller(0, true)}
	}

	db, err := dbutil.DBClient()

	if err != nil {
		return Error{Text: err.Error(), Code: DBError, Where: Caller(0, true)}
	}

	if prodID != "" {
		err = db.Delete(&Product{ProductID: prodID}).Error
	} else if prodName != "" {
		err = db.Delete(&Product{ProductName: prodName}).Error
	}

	if err != nil {
		return Error{Text: err.Error(), Code: DBError, Where: Caller(0, true)}
	}

	// TODO need to delete related items in table `resource type` `alert rule` ...
	return nil
}

func UpdateProduct(prod *Product) error {
	db, err := dbutil.DBClient()

	if err != nil {
		return Error{Text: err.Error(), Code: DBError, Where: Caller(0, true)}
	}

	if prod.ProductID != "" {
		err = db.Model(prod).Where("product_id = ?", prod.ProductID).Update(prod).Error
	} else if prod.ProductName != "" {
		err = db.Model(prod).Where("product_name = ?", prod.ProductName).Update(prod).Error
	}

	if err != nil {
		return Error{Text: err.Error(), Code: DBError, Where: Caller(0, true)}
	}

	return err
}
