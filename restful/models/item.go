package models

import (
	"github.com/jinzhu/gorm"
)

type Item struct {
	gorm.Model
	Title    string  `gorm:"column:title"`
	Price    float64 `gorm:"column:price"`
	Num      int64   `gorm:"column:num"`
	LimitNum int64   `gorm:"column:limit_num"`
	Status   int64   `gorm:"column:status"`
}
