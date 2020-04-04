//generate by gen
package models

import (
	"goshop/restful/common"
)

//秒杀商品属性详情表
type StoreSeckillAttrResult struct {
	ProductId  int    `gorm:"column:product_id"`  //商品ID
	Result     string `gorm:"column:result"`      //商品属性参数
	ChangeTime int    `gorm:"column:change_time"` //上次修改时间

}

//修改默认表名
func (StoreSeckillAttrResult) TableName() string {
	return "eb_store_seckill_attr_result"
}

func (storeSeckillAttrResult *StoreSeckillAttrResult) Insert() error {
	err := common.GetDB().Create(storeSeckillAttrResult).Error
	return err
}

func (storeSeckillAttrResult *StoreSeckillAttrResult) Patch() error {
	err := common.GetDB().Model(storeSeckillAttrResult).Updates(storeSeckillAttrResult).Error
	return err
}

func (storeSeckillAttrResult *StoreSeckillAttrResult) Update() error {
	err := common.GetDB().Save(storeSeckillAttrResult).Error
	return err
}

func (storeSeckillAttrResult *StoreSeckillAttrResult) Delete() error {
	return common.GetDB().Delete(storeSeckillAttrResult).Error
}

func (storeSeckillAttrResult *StoreSeckillAttrResult) List(rawQuery string, rawOrder string, offset int, limit int) (*[]StoreSeckillAttrResult, int, error) {
	storeSeckillAttrResults := []StoreSeckillAttrResult{}
	total := 0
	db := common.GetDB().Model(storeSeckillAttrResult)
	db, err := buildWhere(rawQuery, db)
	if err != nil {
		return &storeSeckillAttrResults, total, err
	}
	db, err = buildOrder(rawOrder, db)
	if err != nil {
		return &storeSeckillAttrResults, total, err
	}
	db.Offset(offset).
		Limit(limit).
		Find(&storeSeckillAttrResults).
		Count(&total)
	err = db.Error
	return &storeSeckillAttrResults, total, err
}

func (storeSeckillAttrResult *StoreSeckillAttrResult) Get() (*StoreSeckillAttrResult, error) {
	err := common.GetDB().Find(&storeSeckillAttrResult).Error
	return storeSeckillAttrResult, err
}
