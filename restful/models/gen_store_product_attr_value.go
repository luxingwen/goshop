//generate by gen
package models

import (
	"goshop/restful/common"
)

//商品属性值表
type StoreProductAttrValue struct {
	ProductId int     `gorm:"column:product_id" json:"product_id"` //商品ID
	Suk       string  `gorm:"column:suk" json:"suk"`               //商品属性索引值 (attr_value|attr_value[|....])
	Stock     int     `gorm:"column:stock" json:"stock"`           //属性对应的库存
	Sales     int     `gorm:"column:sales" json:"sales"`           //销量
	Price     float64 `gorm:"column:price" json:"price"`           //属性金额
	Image     string  `gorm:"column:image" json:"image"`           //图片
	Unique    string  `gorm:"column:unique" json:"unique"`         //唯一值
	Cost      float64 `gorm:"column:cost" json:"cost"`             //成本价

}

//修改默认表名
func (StoreProductAttrValue) TableName() string {
	return "eb_store_product_attr_value"
}

func (storeProductAttrValue *StoreProductAttrValue) Insert() error {
	err := common.GetDB().Create(storeProductAttrValue).Error
	return err
}

func (storeProductAttrValue *StoreProductAttrValue) Patch() error {
	err := common.GetDB().Model(storeProductAttrValue).Updates(storeProductAttrValue).Error
	return err
}

func (storeProductAttrValue *StoreProductAttrValue) Update() error {
	err := common.GetDB().Save(storeProductAttrValue).Error
	return err
}

func (storeProductAttrValue *StoreProductAttrValue) Delete() error {
	return common.GetDB().Delete(storeProductAttrValue).Error
}

func (storeProductAttrValue *StoreProductAttrValue) List(rawQuery string, rawOrder string, offset int, limit int) (*[]StoreProductAttrValue, int, error) {
	storeProductAttrValues := []StoreProductAttrValue{}
	total := 0
	db := common.GetDB().Model(storeProductAttrValue)
	db, err := buildWhere(rawQuery, db)
	if err != nil {
		return &storeProductAttrValues, total, err
	}
	db, err = buildOrder(rawOrder, db)
	if err != nil {
		return &storeProductAttrValues, total, err
	}
	db.Offset(offset).
		Limit(limit).
		Find(&storeProductAttrValues).
		Count(&total)
	err = db.Error
	return &storeProductAttrValues, total, err
}

func (storeProductAttrValue *StoreProductAttrValue) Get() (*StoreProductAttrValue, error) {
	err := common.GetDB().Find(&storeProductAttrValue).Error
	return storeProductAttrValue, err
}

func (storeProductAttrValue *StoreProductAttrValue) ListByProductId(id int) (r []*StoreProductAttrValue, err error) {
	db := common.GetDB()
	err = db.Where("product_id = ?", id).Find(&r).Error
	return
}

func (storeProductAttrValue *StoreProductAttrValue) UniqueByStock(unique string) (count int, err error) {
	db := common.GetDB().Table(storeProductAttrValue.TableName())
	err = db.Where("unique = ?", unique).Scan(&storeProductAttrValue).Error
	if err != nil {
		return
	}
	count = storeProductAttrValue.Stock
	return
}
