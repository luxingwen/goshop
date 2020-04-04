//generate by gen
package models

import (
	"goshop/restful/common"
)

//订单操作记录表
type StoreOrderStatus struct {
	Oid           int    `gorm:"column:oid"`            //订单id
	ChangeType    string `gorm:"column:change_type"`    //操作类型
	ChangeMessage string `gorm:"column:change_message"` //操作备注
	ChangeTime    int    `gorm:"column:change_time"`    //操作时间

}

//修改默认表名
func (StoreOrderStatus) TableName() string {
	return "eb_store_order_status"
}

func (storeOrderStatus *StoreOrderStatus) Insert() error {
	err := common.GetDB().Create(storeOrderStatus).Error
	return err
}

func (storeOrderStatus *StoreOrderStatus) Patch() error {
	err := common.GetDB().Model(storeOrderStatus).Updates(storeOrderStatus).Error
	return err
}

func (storeOrderStatus *StoreOrderStatus) Update() error {
	err := common.GetDB().Save(storeOrderStatus).Error
	return err
}

func (storeOrderStatus *StoreOrderStatus) Delete() error {
	return common.GetDB().Delete(storeOrderStatus).Error
}

func (storeOrderStatus *StoreOrderStatus) List(rawQuery string, rawOrder string, offset int, limit int) (*[]StoreOrderStatus, int, error) {
	storeOrderStatuss := []StoreOrderStatus{}
	total := 0
	db := common.GetDB().Model(storeOrderStatus)
	db, err := buildWhere(rawQuery, db)
	if err != nil {
		return &storeOrderStatuss, total, err
	}
	db, err = buildOrder(rawOrder, db)
	if err != nil {
		return &storeOrderStatuss, total, err
	}
	db.Offset(offset).
		Limit(limit).
		Find(&storeOrderStatuss).
		Count(&total)
	err = db.Error
	return &storeOrderStatuss, total, err
}

func (storeOrderStatus *StoreOrderStatus) Get() (*StoreOrderStatus, error) {
	err := common.GetDB().Find(&storeOrderStatus).Error
	return storeOrderStatus, err
}
