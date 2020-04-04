//generate by gen
package models

import (
	"goshop/restful/common"
)

//用户通知表
type UserNotice struct {
	Id       int    `gorm:"column:id"`        //
	Uid      string `gorm:"column:uid"`       //接收消息的用户id（类型：json数据）
	Type     int    `gorm:"column:type"`      //消息通知类型（1：系统消息；2：用户通知）
	User     string `gorm:"column:user"`      //发送人
	Title    string `gorm:"column:title"`     //通知消息的标题信息
	Content  string `gorm:"column:content"`   //通知消息的内容
	AddTime  int    `gorm:"column:add_time"`  //通知消息发送的时间
	IsSend   int    `gorm:"column:is_send"`   //是否发送（0：未发送；1：已发送）
	SendTime int    `gorm:"column:send_time"` //发送时间

}

//修改默认表名
func (UserNotice) TableName() string {
	return "eb_user_notice"
}

func (userNotice *UserNotice) Insert() error {
	err := common.GetDB().Create(userNotice).Error
	return err
}

func (userNotice *UserNotice) Patch() error {
	err := common.GetDB().Model(userNotice).Updates(userNotice).Error
	return err
}

func (userNotice *UserNotice) Update() error {
	err := common.GetDB().Save(userNotice).Error
	return err
}

func (userNotice *UserNotice) Delete() error {
	return common.GetDB().Delete(userNotice).Error
}

func (userNotice *UserNotice) List(rawQuery string, rawOrder string, offset int, limit int) (*[]UserNotice, int, error) {
	userNotices := []UserNotice{}
	total := 0
	db := common.GetDB().Model(userNotice)
	db, err := buildWhere(rawQuery, db)
	if err != nil {
		return &userNotices, total, err
	}
	db, err = buildOrder(rawOrder, db)
	if err != nil {
		return &userNotices, total, err
	}
	db.Offset(offset).
		Limit(limit).
		Find(&userNotices).
		Count(&total)
	err = db.Error
	return &userNotices, total, err
}

func (userNotice *UserNotice) Get() (*UserNotice, error) {
	err := common.GetDB().Find(&userNotice).Error
	return userNotice, err
}
