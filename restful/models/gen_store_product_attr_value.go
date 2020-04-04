//generate by gen
package models

import (
	"goshop/restful/common"
)

//商品属性值表
type StoreProductAttrValue struct {
	ProductId int     `gorm:"column:product_id"` //商品ID
	Suk       string  `gorm:"column:suk"`        //商品属性索引值 (attr_value|attr_value[|....])
	Stock     int     `gorm:"column:stock"`      //属性对应的库存
	Sales     int     `gorm:"column:sales"`      //销量
	Price     float64 `gorm:"column:price"`      //属性金额
	Image     string  `gorm:"column:image"`      //图片
	Unique    string  `gorm:"column:unique"`     //唯一值
	Cost      float64 `gorm:"column:cost"`       //成本价

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
