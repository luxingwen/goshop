package models

import (
	"github.com/jinzhu/gorm"
)

type ItemDesc struct {
	gorm.Model
	ItemId   int64
	ItemDesc string
}
