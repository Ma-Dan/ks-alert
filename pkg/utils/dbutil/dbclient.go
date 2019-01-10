package dbutil

import (
	"errors"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// TODO need to support db connection pool
var db *gorm.DB

func init() {
	var err error
	//db, err = gorm.Open("mysql", "root:password@tcp(139.198.190.141:33306)/alert?charset=utf8&parseTime=True&loc=Local")
	db, err = gorm.Open("mysql", "root:password@tcp(127.0.0.1:3306)/alert?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
}

func DBClient() (*gorm.DB, error) {
	if db == nil {
		return db, errors.New("db conection init failed")
	}
	return db, nil
}
