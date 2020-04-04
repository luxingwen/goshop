//generate by gen
package models

import (
	"goshop/restful/common"
)

//签到记录表
type UserSign struct {
	Id      int    `gorm:"column:id"`       //
	Uid     int    `gorm:"column:uid"`      //用户uid
	Title   string `gorm:"column:title"`    //签到说明
	Number  int    `gorm:"column:number"`   //获得积分
	Balance int    `gorm:"column:balance"`  //剩余积分
	AddTime int    `gorm:"column:add_time"` //添加时间

}

//修改默认表名
func (UserSign) TableName() string {
	return "eb_user_sign"
}

func (userSign *UserSign) Insert() error {
	err := common.GetDB().Create(userSign).Error
	return err
}

func (userSign *UserSign) Patch() error {
	err := common.GetDB().Model(userSign).Updates(userSign).Error
	return err
}

func (userSign *UserSign) Update() error {
	err := common.GetDB().Save(userSign).Error
	return err
}

func (userSign *UserSign) Delete() error {
	return common.GetDB().Delete(userSign).Error
}

func (userSign *UserSign) List(rawQuery string, rawOrder string, offset int, limit int) (*[]UserSign, int, error) {
	userSigns := []UserSign{}
	total := 0
	db := common.GetDB().Model(userSign)
	db, err := buildWhere(rawQuery, db)
	if err != nil {
		return &userSigns, total, err
	}
	db, err = buildOrder(rawOrder, db)
	if err != nil {
		return &userSigns, total, err
	}
	db.Offset(offset).
		Limit(limit).
		Find(&userSigns).
		Count(&total)
	err = db.Error
	return &userSigns, total, err
}

func (userSign *UserSign) Get() (*UserSign, error) {
	err := common.GetDB().Find(&userSign).Error
	return userSign, err
}
