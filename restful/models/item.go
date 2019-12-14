package models

import (
	"github.com/jinzhu/gorm"
)

type Item struct {
	gorm.Model
	Title       string  `gorm:"column:title"`
	Price       float64 `gorm:"column:price"`
	OriginPrice float64 `gorm:"orginPrice"`
	Num         int64   `gorm:"column:num"`       // 数量
	LimitNum    int64   `gorm:"column:limit_num"` // 限制数量
	Status      int64   `gorm:"column:status"`    // 状态
	Desc        string  `gorm:"column:desc"`      // 描述
	PicUrl      string  `gorm:"column:picUrl"`    // 图片地址
}
