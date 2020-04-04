//generate by gen
package models

import (
	"goshop/restful/common"
)

//商品属性详情表
type StoreProductAttrResult struct {
	ProductId  int    `gorm:"column:product_id"`  //商品ID
	Result     string `gorm:"column:result"`      //商品属性参数
	ChangeTime int    `gorm:"column:change_time"` //上次修改时间

}

//修改默认表名
func (StoreProductAttrResult) TableName() string {
	return "eb_store_product_attr_result"
}

func (storeProductAttrResult *StoreProductAttrResult) Insert() error {
	err := common.GetDB().Create(storeProductAttrResult).Error
	return err
}

func (storeProductAttrResult *StoreProductAttrResult) Patch() error {
	err := common.GetDB().Model(storeProductAttrResult).Updates(storeProductAttrResult).Error
	return err
}

func (storeProductAttrResult *StoreProductAttrResult) Update() error {
	err := common.GetDB().Save(storeProductAttrResult).Error
	return err
}

func (storeProductAttrResult *StoreProductAttrResult) Delete() error {
	return common.GetDB().Delete(storeProductAttrResult).Error
}

func (storeProductAttrResult *StoreProductAttrResult) List(rawQuery string, rawOrder string, offset int, limit int) (*[]StoreProductAttrResult, int, error) {
	storeProductAttrResults := []StoreProductAttrResult{}
	total := 0
	db := common.GetDB().Model(storeProductAttrResult)
	db, err := buildWhere(rawQuery, db)
	if err != nil {
		return &storeProductAttrResults, total, err
	}
	db, err = buildOrder(rawOrder, db)
	if err != nil {
		return &storeProductAttrResults, total, err
	}
	db.Offset(offset).
		Limit(limit).
		Find(&storeProductAttrResults).
		Count(&total)
	err = db.Error
	return &storeProductAttrResults, total, err
}

func (storeProductAttrResult *StoreProductAttrResult) Get() (*StoreProductAttrResult, error) {
	err := common.GetDB().Find(&storeProductAttrResult).Error
	return storeProductAttrResult, err
}
