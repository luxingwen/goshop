//generate by gen
package models

import (
	"goshop/restful/common"
)

//商品属性表
type StoreProductAttr struct {
	ProductId  int    `gorm:"column:product_id"`  //商品ID
	AttrName   string `gorm:"column:attr_name"`   //属性名
	AttrValues string `gorm:"column:attr_values"` //属性值

}

//修改默认表名
func (StoreProductAttr) TableName() string {
	return "eb_store_product_attr"
}

func (storeProductAttr *StoreProductAttr) Insert() error {
	err := common.GetDB().Create(storeProductAttr).Error
	return err
}

func (storeProductAttr *StoreProductAttr) Patch() error {
	err := common.GetDB().Model(storeProductAttr).Updates(storeProductAttr).Error
	return err
}

func (storeProductAttr *StoreProductAttr) Update() error {
	err := common.GetDB().Save(storeProductAttr).Error
	return err
}

func (storeProductAttr *StoreProductAttr) Delete() error {
	return common.GetDB().Delete(storeProductAttr).Error
}

func (storeProductAttr *StoreProductAttr) List(rawQuery string, rawOrder string, offset int, limit int) (*[]StoreProductAttr, int, error) {
	storeProductAttrs := []StoreProductAttr{}
	total := 0
	db := common.GetDB().Model(storeProductAttr)
	db, err := buildWhere(rawQuery, db)
	if err != nil {
		return &storeProductAttrs, total, err
	}
	db, err = buildOrder(rawOrder, db)
	if err != nil {
		return &storeProductAttrs, total, err
	}
	db.Offset(offset).
		Limit(limit).
		Find(&storeProductAttrs).
		Count(&total)
	err = db.Error
	return &storeProductAttrs, total, err
}

func (storeProductAttr *StoreProductAttr) Get() (*StoreProductAttr, error) {
	err := common.GetDB().Find(&storeProductAttr).Error
	return storeProductAttr, err
}

func (storeProductAttr *StoreProductAttr) GetByProductId(id int) (err error) {
	db := common.GetDB()
	err = db.Where("product_id = ?", id).First(&storeProductAttr).Error
	return
}
