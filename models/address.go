package models

import (
	"github.com/jinzhu/gorm"
)

type Address struct {
	gorm.Model
	UserId     int64  // 用户id
	Username   string // 收件人姓名
	StreetName string // 街道地址
	IsDefault  bool   // 是否是默认地址
}
