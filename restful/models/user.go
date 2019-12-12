/*
 * @Author: kslamp
 * @Date: 2019-12-11 20:11:22
 * @LastEditTime: 2019-12-12 13:37:33
 * @FilePath: /goshop/restful/models/user.go
 * @Description:
 */
package models

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"column:username" json:"username" faker:"username"`
	Password string `gorm:"column:password" json:"password" faker:"password"`
	Name     string `gorm:"column:name" json:"name" faker:"name"`
	Phone    string `gorm:"column:phone" json:"phone" faker:"phone_number"`
	Email    string `gorm:"column:email" json:"email" faker:"email"`
	Sex      string `gorm:"column:sex" json:"sex"`
	Address  string `gorm:"column:address" json:"address"`
	Status   int    `gorm:"column:status" json:"status"`
	Desc     string `gorm:"column:desc" json:"desc" faker:"word"`
	Img      string `gorm:"column:img" json:"img" faker:"url"`
}
