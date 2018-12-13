package models

import (
	"kubesphere.io/ks-alert/client"
)

func init() {

	db, err := client.DBClient()

	if err != nil {
		panic(err)
	}

	if !db.HasTable(&RegisteredEnterprise{}) {
		if err := db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").CreateTable(&RegisteredEnterprise{}).Error; err != nil {
			panic(err)
		}
	}

	if !db.HasTable(&RegisteredProduct{}) {
		if err := db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").CreateTable(&RegisteredProduct{}).Error; err != nil {
			panic(err)
		}
	}

	if !db.HasTable(&UserAlertBinding{}) {
		if err := db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").CreateTable(&UserAlertBinding{}).Error; err != nil {
			panic(err)
		}
	}
}
