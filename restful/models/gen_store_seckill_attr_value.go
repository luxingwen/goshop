//generate by gen
package models

import (
	"goshop/restful/common"
)

//秒杀商品属性值表
type StoreSeckillAttrValue struct {
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
func (StoreSeckillAttrValue) TableName() string {
	return "eb_store_seckill_attr_value"
}

func (storeSeckillAttrValue *StoreSeckillAttrValue) Insert() error {
	err := common.GetDB().Create(storeSeckillAttrValue).Error
	return err
}

func (storeSeckillAttrValue *StoreSeckillAttrValue) Patch() error {
	err := common.GetDB().Model(storeSeckillAttrValue).Updates(storeSeckillAttrValue).Error
	return err
}

func (storeSeckillAttrValue *StoreSeckillAttrValue) Update() error {
	err := common.GetDB().Save(storeSeckillAttrValue).Error
	return err
}

func (storeSeckillAttrValue *StoreSeckillAttrValue) Delete() error {
	return common.GetDB().Delete(storeSeckillAttrValue).Error
}

func (storeSeckillAttrValue *StoreSeckillAttrValue) List(rawQuery string, rawOrder string, offset int, limit int) (*[]StoreSeckillAttrValue, int, error) {
	storeSeckillAttrValues := []StoreSeckillAttrValue{}
	total := 0
	db := common.GetDB().Model(storeSeckillAttrValue)
	db, err := buildWhere(rawQuery, db)
	if err != nil {
		return &storeSeckillAttrValues, total, err
	}
	db, err = buildOrder(rawOrder, db)
	if err != nil {
		return &storeSeckillAttrValues, total, err
	}
	db.Offset(offset).
		Limit(limit).
		Find(&storeSeckillAttrValues).
		Count(&total)
	err = db.Error
	return &storeSeckillAttrValues, total, err
}

func (storeSeckillAttrValue *StoreSeckillAttrValue) Get() (*StoreSeckillAttrValue, error) {
	err := common.GetDB().Find(&storeSeckillAttrValue).Error
	return storeSeckillAttrValue, err
}
