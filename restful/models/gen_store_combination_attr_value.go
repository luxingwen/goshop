//generate by gen
package models

import (
	"goshop/restful/common"
)

//商品属性值表
type StoreCombinationAttrValue struct {
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
func (StoreCombinationAttrValue) TableName() string {
	return "eb_store_combination_attr_value"
}

func (storeCombinationAttrValue *StoreCombinationAttrValue) Insert() error {
	err := common.GetDB().Create(storeCombinationAttrValue).Error
	return err
}

func (storeCombinationAttrValue *StoreCombinationAttrValue) Patch() error {
	err := common.GetDB().Model(storeCombinationAttrValue).Updates(storeCombinationAttrValue).Error
	return err
}

func (storeCombinationAttrValue *StoreCombinationAttrValue) Update() error {
	err := common.GetDB().Save(storeCombinationAttrValue).Error
	return err
}

func (storeCombinationAttrValue *StoreCombinationAttrValue) Delete() error {
	return common.GetDB().Delete(storeCombinationAttrValue).Error
}

func (storeCombinationAttrValue *StoreCombinationAttrValue) List(rawQuery string, rawOrder string, offset int, limit int) (*[]StoreCombinationAttrValue, int, error) {
	storeCombinationAttrValues := []StoreCombinationAttrValue{}
	total := 0
	db := common.GetDB().Model(storeCombinationAttrValue)
	db, err := buildWhere(rawQuery, db)
	if err != nil {
		return &storeCombinationAttrValues, total, err
	}
	db, err = buildOrder(rawOrder, db)
	if err != nil {
		return &storeCombinationAttrValues, total, err
	}
	db.Offset(offset).
		Limit(limit).
		Find(&storeCombinationAttrValues).
		Count(&total)
	err = db.Error
	return &storeCombinationAttrValues, total, err
}

func (storeCombinationAttrValue *StoreCombinationAttrValue) Get() (*StoreCombinationAttrValue, error) {
	err := common.GetDB().Find(&storeCombinationAttrValue).Error
	return storeCombinationAttrValue, err
}
