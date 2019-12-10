package models

import (
	"github.com/jinzhu/gorm"
)

type Address struct {
	gorm.Model
	UserId     int64  `gorm:"column:user_id" json:"userId"`         // 用户id
	Username   string `gorm:"column:username" json:"username"`      // 收件人姓名
	StreetName string `gorm:"column:street_name" json:"streetName"` // 街道地址
	IsDefault  bool   `gorm:"column:is_default" json:"isDefault"`   // 是否是默认地址
}
