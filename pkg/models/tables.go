package models

import (
	"fmt"
	"github.com/carmanzhang/ks-alert/pkg/utils/dbutil"
	"github.com/jinzhu/gorm"
)

func init() {
	db, err := dbutil.DBClient()
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

	if !db.HasTable(&Resource{}) {
		if err := db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").CreateTable(&Resource{}).Error; err != nil {
			panic(err)
		}
	}

	if !db.HasTable(&ResourceGroup{}) {
		if err := db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").CreateTable(&ResourceGroup{}).Error; err != nil {
			panic(err)
		}
	}

	if !db.HasTable(&Receiver{}) {
		if err := db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").CreateTable(&Receiver{}).Error; err != nil {
			panic(err)
		}
	}

	if !db.HasTable(&ReceiverBindingGroup{}) {
		if err := db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").CreateTable(&ReceiverBindingGroup{}).Error; err != nil {
			panic(err)
		}
	}

	if !db.HasTable(&ReceiverGroup{}) {
		if err := db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").CreateTable(&ReceiverGroup{}).Error; err != nil {
			panic(err)
		}
	}

	if !db.HasTable(&Severity{}) {
		if err := db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").CreateTable(&Severity{}).Error; err != nil {
			panic(err)
		}
	}

	if !db.HasTable(&AlertRule{}) {
		if err := db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").CreateTable(&AlertRule{}).Error; err != nil {
			panic(err)
		}
	}

	if !db.HasTable(&AlertRuleGroup{}) {
		if err := db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").CreateTable(&AlertRuleGroup{}).Error; err != nil {
			panic(err)
		}
	}

	if !db.HasTable(&AlertConfig{}) {
		if err := db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").CreateTable(&AlertConfig{}).Error; err != nil {
			fmt.Print(err)
			panic(err)
		}
	}

	if !db.HasTable(&AlertHistory{}) {
		if err := db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").CreateTable(&AlertHistory{}).Error; err != nil {
			panic(err)
		}
	}

	if !db.HasTable(&Suggestion{}) {
		if err := db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").CreateTable(&Suggestion{}).Error; err != nil {
			panic(err)
		}
	}

	if !db.HasTable(&Silence{}) {
		if err := db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").CreateTable(&Silence{}).Error; err != nil {
			panic(err)
		}
	}

	if !db.HasTable(&SendPolicy{}) {
		if err := db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").CreateTable(&SendPolicy{}).Error; err != nil {
			panic(err)
		}
	}

}
