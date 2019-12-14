package models

import (
	"github.com/jinzhu/gorm"
)

type Address struct {
	gorm.Model
	UserId     int64  `gorm:"column:user_id" json:"userId"`         // 用户id
	Name       string `gorm:"column:name" json:"name"`              // 收件人姓名
	Mobile     string `gorm:"column:address" json:"mobile"`         // 手机
	Address    string `gorm:"column:address" json:"address"`        // 地区
	StreetName string `gorm:"column:street_name" json:"streetName"` // 街道地址
	IsDefault  bool   `gorm:"column:is_default" json:"isDefault"`   // 是否是默认地址
}
