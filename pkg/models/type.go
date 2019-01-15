package models

import "github.com/jinzhu/gorm"

type Action interface {
	Create(*gorm.DB) (interface{}, error)
	Update(*gorm.DB) (interface{}, error)
	Get(*gorm.DB) (interface{}, error)
	Delete(*gorm.DB) (interface{}, error)
}
