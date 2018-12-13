package models

import (
	"github.com/jinzhu/gorm"
	"kubesphere.io/ks-alert/pkg/client"
)

func init() {
	db, err := client.DBClient()
	if err != nil {
		panic(err)
	}
	createTables(db)
}

func createTables(db *gorm.DB) {

	if !db.HasTable(&Enterprise{}) {
		if err := db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").CreateTable(&Enterprise{}).Error; err != nil {
			panic(err)
		}
	}
	if !db.HasTable(&Product{}) {
		if err := db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").CreateTable(&Product{}).Error; err != nil {
			panic(err)
		}
	}

	if !db.HasTable(&ResourceType{}) {
		if err := db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").CreateTable(&ResourceType{}).Error; err != nil {
			panic(err)
		}
	}

	if !db.HasTable(&Metric{}) {
		if err := db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").CreateTable(&Metric{}).Error; err != nil {
			panic(err)
		}
	}

	if !db.HasTable(&ResourceGroup{}) {
		if err := db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").CreateTable(&ResourceGroup{}).Error; err != nil {
			panic(err)
		}
	}

}
