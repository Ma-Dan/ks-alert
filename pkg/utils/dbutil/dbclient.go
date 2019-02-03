package dbutil

import (
	"errors"
	"kubesphere.io/ks-alert/pkg/option"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// TODO need to support db connection pool
var db *gorm.DB

func DBClient() (*gorm.DB, error) {
	if db != nil {
		return db, nil
	}

	var err error
	user := *option.User
	pwd := *option.Password
	host := *option.MysqlHost
	port := *option.MysqlPort
	database := *option.Database

	prefix := user + ":" + pwd + "@tcp(" + host + ":" + port + ")/" + database
	db, err = gorm.Open("mysql", prefix+"?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		return nil, err
	}

	if db == nil {
		return db, errors.New("db conection init failed")
	}

	return db, nil
}
