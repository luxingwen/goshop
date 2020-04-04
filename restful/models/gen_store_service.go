//generate by gen
package models

import (
	"goshop/restful/common"
)

//客服表
type StoreService struct {
	Id       int    `gorm:"column:id"`       //客服id
	MerId    int    `gorm:"column:mer_id"`   //商户id
	Uid      int    `gorm:"column:uid"`      //客服uid
	Avatar   string `gorm:"column:avatar"`   //客服头像
	Nickname string `gorm:"column:nickname"` //代理名称
	AddTime  int    `gorm:"column:add_time"` //添加时间
	Status   int    `gorm:"column:status"`   //0隐藏1显示
	Notify   int    `gorm:"column:notify"`   //订单通知1开启0关闭

}

//修改默认表名
func (StoreService) TableName() string {
	return "eb_store_service"
}

func (storeService *StoreService) Insert() error {
	err := common.GetDB().Create(storeService).Error
	return err
}

func (storeService *StoreService) Patch() error {
	err := common.GetDB().Model(storeService).Updates(storeService).Error
	return err
}

func (storeService *StoreService) Update() error {
	err := common.GetDB().Save(storeService).Error
	return err
}

func (storeService *StoreService) Delete() error {
	return common.GetDB().Delete(storeService).Error
}

func (storeService *StoreService) List(rawQuery string, rawOrder string, offset int, limit int) (*[]StoreService, int, error) {
	storeServices := []StoreService{}
	total := 0
	db := common.GetDB().Model(storeService)
	db, err := buildWhere(rawQuery, db)
	if err != nil {
		return &storeServices, total, err
	}
	db, err = buildOrder(rawOrder, db)
	if err != nil {
		return &storeServices, total, err
	}
	db.Offset(offset).
		Limit(limit).
		Find(&storeServices).
		Count(&total)
	err = db.Error
	return &storeServices, total, err
}

func (storeService *StoreService) Get() (*StoreService, error) {
	err := common.GetDB().Find(&storeService).Error
	return storeService, err
}
