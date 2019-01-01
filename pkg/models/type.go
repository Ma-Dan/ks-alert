package models

import "github.com/jinzhu/gorm"

type Action interface {
	Create(*gorm.DB, interface{}) (interface{}, error)
	Update(*gorm.DB, interface{}) (interface{}, error)
	Get(*gorm.DB, interface{}) (interface{}, error)
	Delete(*gorm.DB, interface{}) (interface{}, error)
}
