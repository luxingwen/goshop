//generate by gen
package models

import (
	"goshop/restful/common"
)

//用户通知发送记录表
type UserNoticeSee struct {
	Id      int `gorm:"column:id"`       //
	Nid     int `gorm:"column:nid"`      //查看的通知id
	Uid     int `gorm:"column:uid"`      //查看通知的用户id
	AddTime int `gorm:"column:add_time"` //查看通知的时间

}

//修改默认表名
func (UserNoticeSee) TableName() string {
	return "eb_user_notice_see"
}

func (userNoticeSee *UserNoticeSee) Insert() error {
	err := common.GetDB().Create(userNoticeSee).Error
	return err
}

func (userNoticeSee *UserNoticeSee) Patch() error {
	err := common.GetDB().Model(userNoticeSee).Updates(userNoticeSee).Error
	return err
}

func (userNoticeSee *UserNoticeSee) Update() error {
	err := common.GetDB().Save(userNoticeSee).Error
	return err
}

func (userNoticeSee *UserNoticeSee) Delete() error {
	return common.GetDB().Delete(userNoticeSee).Error
}

func (userNoticeSee *UserNoticeSee) List(rawQuery string, rawOrder string, offset int, limit int) (*[]UserNoticeSee, int, error) {
	userNoticeSees := []UserNoticeSee{}
	total := 0
	db := common.GetDB().Model(userNoticeSee)
	db, err := buildWhere(rawQuery, db)
	if err != nil {
		return &userNoticeSees, total, err
	}
	db, err = buildOrder(rawOrder, db)
	if err != nil {
		return &userNoticeSees, total, err
	}
	db.Offset(offset).
		Limit(limit).
		Find(&userNoticeSees).
		Count(&total)
	err = db.Error
	return &userNoticeSees, total, err
}

func (userNoticeSee *UserNoticeSee) Get() (*UserNoticeSee, error) {
	err := common.GetDB().Find(&userNoticeSee).Error
	return userNoticeSee, err
}
