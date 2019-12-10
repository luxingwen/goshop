package models

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Username string
	Password string
	Phone    string
	Email    string
	Sex      string
	Address  string
	Status   int
	Desc     string
	Img      string
}
