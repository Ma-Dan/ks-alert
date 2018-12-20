package models

import (
	"errors"
	"github.com/jinzhu/gorm"
	"kubesphere.io/ks-alert/pkg/clientient"
	. "kubesphere.io/ks-alert/pkg/models"
)

// table `user_alert_bindings` record alerts created by user
// associated table of table `user` and table `alert`
type UserAlertBinding struct {
	gorm.Model
	UserID        string `gorm:"type:varchar(50);not null;"`
	AlertConfigID string `gorm:"type:varchar(50);not null;"`
	ProductID     string `gorm:"type:varchar(50);not null;"`
}

// registed enterprise
type RegisteredEnterprise struct {
	gorm.Model
	EnterpriseID string `gorm:"type:varchar(50);not null;"`
}

// registed product
type RegisteredProduct struct {
	gorm.Model
	ProductID    string `gorm:"type:varchar(50);not null;"`
	EnterpriseID string `gorm:"type:varchar(50);not null;"`
}

func CreateRegisteredEnterpriseItem(ent *RegisteredEnterprise) error {

	if ent.EnterpriseID == "" {
		return errors.New("registered enterprise has no enterprise_id," +
			" perhaps this enterprise does not register")
	}

	db, err := client.DBClient()
	if err != nil {
		panic(err)
	}

	err = db.Model(&RegisteredEnterprise{}).Create(ent).Error
	return err
}

func CreateRegisteredProductItem(product *RegisteredProduct) error {

	if product.ProductID == "" {
		return errors.New("registered product has no product_id," +
			" perhaps this product does not register")
	}

	if product.EnterpriseID == "" {
		return errors.New("registered enterprise has no enterprise_id," +
			" perhaps this enterprise does not register")
	}

	db, err := client.DBClient()
	if err != nil {
		panic(err)
	}

	err = db.Model(&RegisteredProduct{}).Create(product).Error
	return err
}

// receiver group
// alert_rule group
// resource group
func CreateUserAlertBindingItem(receivers *[]Receiver, resources *ResourceGroup) error {

	db, err := client.DBClient()
	if err != nil {
		panic(err)
	}

	tx := db.Begin()

	if tx.Error != nil {
		return err
	}

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// TODO

	return tx.Commit().Error
}
