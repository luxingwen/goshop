//generate by gen
package models

import (
	"goshop/restful/common"
)

//客服用户对话记录表
type StoreServiceLog struct {
	Id      int    `gorm:"column:id"`       //客服用户对话记录表ID
	MerId   int    `gorm:"column:mer_id"`   //商户id
	Msn     string `gorm:"column:msn"`      //消息内容
	Uid     int    `gorm:"column:uid"`      //发送人uid
	ToUid   int    `gorm:"column:to_uid"`   //接收人uid
	AddTime int    `gorm:"column:add_time"` //发送时间
	Type    int    `gorm:"column:type"`     //是否已读（0：否；1：是；）
	Remind  int    `gorm:"column:remind"`   //是否提醒过

}

//修改默认表名
func (StoreServiceLog) TableName() string {
	return "eb_store_service_log"
}

func (storeServiceLog *StoreServiceLog) Insert() error {
	err := common.GetDB().Create(storeServiceLog).Error
	return err
}

func (storeServiceLog *StoreServiceLog) Patch() error {
	err := common.GetDB().Model(storeServiceLog).Updates(storeServiceLog).Error
	return err
}

func (storeServiceLog *StoreServiceLog) Update() error {
	err := common.GetDB().Save(storeServiceLog).Error
	return err
}

func (storeServiceLog *StoreServiceLog) Delete() error {
	return common.GetDB().Delete(storeServiceLog).Error
}

func (storeServiceLog *StoreServiceLog) List(rawQuery string, rawOrder string, offset int, limit int) (*[]StoreServiceLog, int, error) {
	storeServiceLogs := []StoreServiceLog{}
	total := 0
	db := common.GetDB().Model(storeServiceLog)
	db, err := buildWhere(rawQuery, db)
	if err != nil {
		return &storeServiceLogs, total, err
	}
	db, err = buildOrder(rawOrder, db)
	if err != nil {
		return &storeServiceLogs, total, err
	}
	db.Offset(offset).
		Limit(limit).
		Find(&storeServiceLogs).
		Count(&total)
	err = db.Error
	return &storeServiceLogs, total, err
}

func (storeServiceLog *StoreServiceLog) Get() (*StoreServiceLog, error) {
	err := common.GetDB().Find(&storeServiceLog).Error
	return storeServiceLog, err
}
