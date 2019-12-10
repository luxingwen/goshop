package models

import (
	"github.com/jinzhu/gorm"
)

type Item struct {
	gorm.Model
	Title    string
	Price    float64
	Num      int64
	LimitNum int64
	Status   int64
}
