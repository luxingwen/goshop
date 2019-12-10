package models

import (
	"github.com/jinzhu/gorm"
)

type Order struct {
	gorm.Model
	OrderId string  `gorm:"column:order_id" json:"orderId"`
	PayMent float64 `gorm:"column:pay_ment" json:"payMent"`
	Status  int     `gorm:"column:status" json:"status"`
	PayTime int64   `gorm:"column:pay_time" json:"payTime"`
	UserId  int64   `gorm:"column:user_id" json:"userId"`
}
