package models

import (
	"github.com/jinzhu/gorm"
)

type ItemDesc struct {
	gorm.Model
	ItemId   int64  `gorm:"column:item_id" json:"itemId"`
	ItemDesc string `gorm:"column:item_desc" json:"itemDesc"`
}
