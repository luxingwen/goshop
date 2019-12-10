package models

import (
	"github.com/jinzhu/gorm"
)

type Order struct {
	gorm.Model
	OrderId string
	PayMent float64
	Status  int
	PayTime int64
	UserId  int64
}
