//generate by gen
package models

import (
	"goshop/restful/common"
)

//商品属性详情表
type StoreCombinationAttrResult struct {
	ProductId  int    `gorm:"column:product_id"`  //商品ID
	Result     string `gorm:"column:result"`      //商品属性参数
	ChangeTime int    `gorm:"column:change_time"` //上次修改时间

}

//修改默认表名
func (StoreCombinationAttrResult) TableName() string {
	return "eb_store_combination_attr_result"
}

func (storeCombinationAttrResult *StoreCombinationAttrResult) Insert() error {
	err := common.GetDB().Create(storeCombinationAttrResult).Error
	return err
}

func (storeCombinationAttrResult *StoreCombinationAttrResult) Patch() error {
	err := common.GetDB().Model(storeCombinationAttrResult).Updates(storeCombinationAttrResult).Error
	return err
}

func (storeCombinationAttrResult *StoreCombinationAttrResult) Update() error {
	err := common.GetDB().Save(storeCombinationAttrResult).Error
	return err
}

func (storeCombinationAttrResult *StoreCombinationAttrResult) Delete() error {
	return common.GetDB().Delete(storeCombinationAttrResult).Error
}

func (storeCombinationAttrResult *StoreCombinationAttrResult) List(rawQuery string, rawOrder string, offset int, limit int) (*[]StoreCombinationAttrResult, int, error) {
	storeCombinationAttrResults := []StoreCombinationAttrResult{}
	total := 0
	db := common.GetDB().Model(storeCombinationAttrResult)
	db, err := buildWhere(rawQuery, db)
	if err != nil {
		return &storeCombinationAttrResults, total, err
	}
	db, err = buildOrder(rawOrder, db)
	if err != nil {
		return &storeCombinationAttrResults, total, err
	}
	db.Offset(offset).
		Limit(limit).
		Find(&storeCombinationAttrResults).
		Count(&total)
	err = db.Error
	return &storeCombinationAttrResults, total, err
}

func (storeCombinationAttrResult *StoreCombinationAttrResult) Get() (*StoreCombinationAttrResult, error) {
	err := common.GetDB().Find(&storeCombinationAttrResult).Error
	return storeCombinationAttrResult, err
}
