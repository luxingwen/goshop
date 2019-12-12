/*
 * @Author: kslamp
 * @Date: 2019-12-11 20:11:22
 * @LastEditTime: 2019-12-12 13:37:33
 * @FilePath: /goshop/restful/models/user.go
 * @Description:
 */
package models

import (
	"errors"
	"goshop/restful/common"

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

type ReqUser struct {
	Query
	Name string `form:"name"`
}

func (this *User) IsUserExist() (err error) {
	err = common.GetDB().Find(&this, User{Username: this.Username, Password: this.Password}).Error
	return
}

func (this *User) RegisterUser() (err error) {
	db := common.GetDB()
	if db.NewRecord(this) {
		err = db.Create(&this).Error
		return
	}
	return errors.New("can not create record")
}

func (req *ReqUser) UserList() (r []*User, count int, err error) {
	limit := 10
	page := 0

	if req.Page != 0 {
		page = 0
	}
	if req.PageNum != 0 {
		limit = 0
	}

	offset := limit * page

	db := common.GetDB()
	if req.Name != "" {
		db = db.Where("Username = ?", req.Name)
	}

	r = make([]*User, offset)
	err = db.Offset(offset).Limit(limit).Find(&r).Error
	if err != nil {
		return
	}
	err = db.Model(&User{}).Count(&count).Error
	return
}
