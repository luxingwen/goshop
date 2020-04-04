//generate by gen
package models

import (
	"goshop/restful/common"
)

//商品属性表
type StoreCombinationAttr struct {
	ProductId  int    `gorm:"column:product_id"`  //商品ID
	AttrName   string `gorm:"column:attr_name"`   //属性名
	AttrValues string `gorm:"column:attr_values"` //属性值

}

//修改默认表名
func (StoreCombinationAttr) TableName() string {
	return "eb_store_combination_attr"
}

func (storeCombinationAttr *StoreCombinationAttr) Insert() error {
	err := common.GetDB().Create(storeCombinationAttr).Error
	return err
}

func (storeCombinationAttr *StoreCombinationAttr) Patch() error {
	err := common.GetDB().Model(storeCombinationAttr).Updates(storeCombinationAttr).Error
	return err
}

func (storeCombinationAttr *StoreCombinationAttr) Update() error {
	err := common.GetDB().Save(storeCombinationAttr).Error
	return err
}

func (storeCombinationAttr *StoreCombinationAttr) Delete() error {
	return common.GetDB().Delete(storeCombinationAttr).Error
}

func (storeCombinationAttr *StoreCombinationAttr) List(rawQuery string, rawOrder string, offset int, limit int) (*[]StoreCombinationAttr, int, error) {
	storeCombinationAttrs := []StoreCombinationAttr{}
	total := 0
	db := common.GetDB().Model(storeCombinationAttr)
	db, err := buildWhere(rawQuery, db)
	if err != nil {
		return &storeCombinationAttrs, total, err
	}
	db, err = buildOrder(rawOrder, db)
	if err != nil {
		return &storeCombinationAttrs, total, err
	}
	db.Offset(offset).
		Limit(limit).
		Find(&storeCombinationAttrs).
		Count(&total)
	err = db.Error
	return &storeCombinationAttrs, total, err
}

func (storeCombinationAttr *StoreCombinationAttr) Get() (*StoreCombinationAttr, error) {
	err := common.GetDB().Find(&storeCombinationAttr).Error
	return storeCombinationAttr, err
}
