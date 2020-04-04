//generate by gen
package models

import (
	"goshop/restful/common"
)

//购物车表
type StoreCart struct {
	Id                int    `gorm:"column:id"`                  //购物车表ID
	Uid               int    `gorm:"column:uid"`                 //用户ID
	Type              string `gorm:"column:type"`                //类型
	ProductId         int    `gorm:"column:product_id"`          //商品ID
	ProductAttrUnique string `gorm:"column:product_attr_unique"` //商品属性
	CartNum           int    `gorm:"column:cart_num"`            //商品数量
	AddTime           int    `gorm:"column:add_time"`            //添加时间
	IsPay             int    `gorm:"column:is_pay"`              //0 = 未购买 1 = 已购买
	IsDel             int    `gorm:"column:is_del"`              //是否删除
	IsNew             int    `gorm:"column:is_new"`              //是否为立即购买
	CombinationId     int    `gorm:"column:combination_id"`      //拼团id
	SeckillId         int    `gorm:"column:seckill_id"`          //秒杀产品ID
	BargainId         int    `gorm:"column:bargain_id"`          //砍价id

}

//修改默认表名
func (StoreCart) TableName() string {
	return "eb_store_cart"
}

func (storeCart *StoreCart) Insert() error {
	err := common.GetDB().Create(storeCart).Error
	return err
}

func (storeCart *StoreCart) Patch() error {
	err := common.GetDB().Model(storeCart).Updates(storeCart).Error
	return err
}

func (storeCart *StoreCart) Update() error {
	err := common.GetDB().Save(storeCart).Error
	return err
}

func (storeCart *StoreCart) Delete() error {
	return common.GetDB().Delete(storeCart).Error
}

func (storeCart *StoreCart) List(rawQuery string, rawOrder string, offset int, limit int) (*[]StoreCart, int, error) {
	storeCarts := []StoreCart{}
	total := 0
	db := common.GetDB().Model(storeCart)
	db, err := buildWhere(rawQuery, db)
	if err != nil {
		return &storeCarts, total, err
	}
	db, err = buildOrder(rawOrder, db)
	if err != nil {
		return &storeCarts, total, err
	}
	db.Offset(offset).
		Limit(limit).
		Find(&storeCarts).
		Count(&total)
	err = db.Error
	return &storeCarts, total, err
}

func (storeCart *StoreCart) Get() (*StoreCart, error) {
	err := common.GetDB().Find(&storeCart).Error
	return storeCart, err
}
