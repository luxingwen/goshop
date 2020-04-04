//generate by gen
package models

import (
	"goshop/restful/common"
)

//产品分类辅助表
type StoreProductCate struct {
	Id        int `gorm:"column:id"`         //
	ProductId int `gorm:"column:product_id"` //产品id
	CateId    int `gorm:"column:cate_id"`    //分类id
	AddTime   int `gorm:"column:add_time"`   //添加时间

}

//修改默认表名
func (StoreProductCate) TableName() string {
	return "eb_store_product_cate"
}

func (storeProductCate *StoreProductCate) Insert() error {
	err := common.GetDB().Create(storeProductCate).Error
	return err
}

func (storeProductCate *StoreProductCate) Patch() error {
	err := common.GetDB().Model(storeProductCate).Updates(storeProductCate).Error
	return err
}

func (storeProductCate *StoreProductCate) Update() error {
	err := common.GetDB().Save(storeProductCate).Error
	return err
}

func (storeProductCate *StoreProductCate) Delete() error {
	return common.GetDB().Delete(storeProductCate).Error
}

func (storeProductCate *StoreProductCate) List(rawQuery string, rawOrder string, offset int, limit int) (*[]StoreProductCate, int, error) {
	storeProductCates := []StoreProductCate{}
	total := 0
	db := common.GetDB().Model(storeProductCate)
	db, err := buildWhere(rawQuery, db)
	if err != nil {
		return &storeProductCates, total, err
	}
	db, err = buildOrder(rawOrder, db)
	if err != nil {
		return &storeProductCates, total, err
	}
	db.Offset(offset).
		Limit(limit).
		Find(&storeProductCates).
		Count(&total)
	err = db.Error
	return &storeProductCates, total, err
}

func (storeProductCate *StoreProductCate) Get() (*StoreProductCate, error) {
	err := common.GetDB().Find(&storeProductCate).Error
	return storeProductCate, err
}
