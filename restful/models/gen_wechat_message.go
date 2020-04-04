//generate by gen
package models

import (
	"goshop/restful/common"
)

//用户行为记录表
type WechatMessage struct {
	Id      int    `gorm:"column:id"`       //用户行为记录id
	Openid  string `gorm:"column:openid"`   //用户openid
	Type    string `gorm:"column:type"`     //操作类型
	Result  string `gorm:"column:result"`   //操作详细记录
	AddTime int    `gorm:"column:add_time"` //操作时间

}

//修改默认表名
func (WechatMessage) TableName() string {
	return "eb_wechat_message"
}

func (wechatMessage *WechatMessage) Insert() error {
	err := common.GetDB().Create(wechatMessage).Error
	return err
}

func (wechatMessage *WechatMessage) Patch() error {
	err := common.GetDB().Model(wechatMessage).Updates(wechatMessage).Error
	return err
}

func (wechatMessage *WechatMessage) Update() error {
	err := common.GetDB().Save(wechatMessage).Error
	return err
}

func (wechatMessage *WechatMessage) Delete() error {
	return common.GetDB().Delete(wechatMessage).Error
}

func (wechatMessage *WechatMessage) List(rawQuery string, rawOrder string, offset int, limit int) (*[]WechatMessage, int, error) {
	wechatMessages := []WechatMessage{}
	total := 0
	db := common.GetDB().Model(wechatMessage)
	db, err := buildWhere(rawQuery, db)
	if err != nil {
		return &wechatMessages, total, err
	}
	db, err = buildOrder(rawOrder, db)
	if err != nil {
		return &wechatMessages, total, err
	}
	db.Offset(offset).
		Limit(limit).
		Find(&wechatMessages).
		Count(&total)
	err = db.Error
	return &wechatMessages, total, err
}

func (wechatMessage *WechatMessage) Get() (*WechatMessage, error) {
	err := common.GetDB().Find(&wechatMessage).Error
	return wechatMessage, err
}
