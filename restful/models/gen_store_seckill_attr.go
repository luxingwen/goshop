//generate by gen
package models

import (
	"goshop/restful/common"
)

//秒杀商品属性表
type StoreSeckillAttr struct {
	ProductId  int    `gorm:"column:product_id"`  //商品ID
	AttrName   string `gorm:"column:attr_name"`   //属性名
	AttrValues string `gorm:"column:attr_values"` //属性值

}

//修改默认表名
func (StoreSeckillAttr) TableName() string {
	return "eb_store_seckill_attr"
}

func (storeSeckillAttr *StoreSeckillAttr) Insert() error {
	err := common.GetDB().Create(storeSeckillAttr).Error
	return err
}

func (storeSeckillAttr *StoreSeckillAttr) Patch() error {
	err := common.GetDB().Model(storeSeckillAttr).Updates(storeSeckillAttr).Error
	return err
}

func (storeSeckillAttr *StoreSeckillAttr) Update() error {
	err := common.GetDB().Save(storeSeckillAttr).Error
	return err
}

func (storeSeckillAttr *StoreSeckillAttr) Delete() error {
	return common.GetDB().Delete(storeSeckillAttr).Error
}

func (storeSeckillAttr *StoreSeckillAttr) List(rawQuery string, rawOrder string, offset int, limit int) (*[]StoreSeckillAttr, int, error) {
	storeSeckillAttrs := []StoreSeckillAttr{}
	total := 0
	db := common.GetDB().Model(storeSeckillAttr)
	db, err := buildWhere(rawQuery, db)
	if err != nil {
		return &storeSeckillAttrs, total, err
	}
	db, err = buildOrder(rawOrder, db)
	if err != nil {
		return &storeSeckillAttrs, total, err
	}
	db.Offset(offset).
		Limit(limit).
		Find(&storeSeckillAttrs).
		Count(&total)
	err = db.Error
	return &storeSeckillAttrs, total, err
}

func (storeSeckillAttr *StoreSeckillAttr) Get() (*StoreSeckillAttr, error) {
	err := common.GetDB().Find(&storeSeckillAttr).Error
	return storeSeckillAttr, err
}
