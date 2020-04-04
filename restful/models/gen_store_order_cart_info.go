//generate by gen
package models

import (
	"goshop/restful/common"
)

//订单购物详情表
type StoreOrderCartInfo struct {
	Oid       int    `gorm:"column:oid"`        //订单id
	CartId    int    `gorm:"column:cart_id"`    //购物车id
	ProductId int    `gorm:"column:product_id"` //商品ID
	CartInfo  string `gorm:"column:cart_info"`  //购买东西的详细信息
	Unique    string `gorm:"column:unique"`     //唯一id

}

//修改默认表名
func (StoreOrderCartInfo) TableName() string {
	return "eb_store_order_cart_info"
}

func (storeOrderCartInfo *StoreOrderCartInfo) Insert() error {
	err := common.GetDB().Create(storeOrderCartInfo).Error
	return err
}

func (storeOrderCartInfo *StoreOrderCartInfo) Patch() error {
	err := common.GetDB().Model(storeOrderCartInfo).Updates(storeOrderCartInfo).Error
	return err
}

func (storeOrderCartInfo *StoreOrderCartInfo) Update() error {
	err := common.GetDB().Save(storeOrderCartInfo).Error
	return err
}

func (storeOrderCartInfo *StoreOrderCartInfo) Delete() error {
	return common.GetDB().Delete(storeOrderCartInfo).Error
}

func (storeOrderCartInfo *StoreOrderCartInfo) List(rawQuery string, rawOrder string, offset int, limit int) (*[]StoreOrderCartInfo, int, error) {
	storeOrderCartInfos := []StoreOrderCartInfo{}
	total := 0
	db := common.GetDB().Model(storeOrderCartInfo)
	db, err := buildWhere(rawQuery, db)
	if err != nil {
		return &storeOrderCartInfos, total, err
	}
	db, err = buildOrder(rawOrder, db)
	if err != nil {
		return &storeOrderCartInfos, total, err
	}
	db.Offset(offset).
		Limit(limit).
		Find(&storeOrderCartInfos).
		Count(&total)
	err = db.Error
	return &storeOrderCartInfos, total, err
}

func (storeOrderCartInfo *StoreOrderCartInfo) Get() (*StoreOrderCartInfo, error) {
	err := common.GetDB().Find(&storeOrderCartInfo).Error
	return storeOrderCartInfo, err
}
