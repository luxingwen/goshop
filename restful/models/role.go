package models

import (
	"github.com/jinzhu/gorm"
)

type Role struct {
	gorm.Model
	Name string `gorm:"column:name" json:"name"`
	Desc string `gorm:"column:desc" json:"desc"`
}
