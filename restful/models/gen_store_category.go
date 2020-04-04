//generate by gen
package models

import (
	"goshop/restful/common"
)

//商品分类表
type StoreCategory struct {
	Id       int    `gorm:"column:id"`        //商品分类表ID
	Pid      int    `gorm:"column:pid"`       //父id
	CateName string `gorm:"column:cate_name"` //分类名称
	Sort     int    `gorm:"column:sort"`      //排序
	Pic      string `gorm:"column:pic"`       //图标
	IsShow   int    `gorm:"column:is_show"`   //是否推荐
	AddTime  int    `gorm:"column:add_time"`  //添加时间

}

//修改默认表名
func (StoreCategory) TableName() string {
	return "eb_store_category"
}

func (storeCategory *StoreCategory) Insert() error {
	err := common.GetDB().Create(storeCategory).Error
	return err
}

func (storeCategory *StoreCategory) Patch() error {
	err := common.GetDB().Model(storeCategory).Updates(storeCategory).Error
	return err
}

func (storeCategory *StoreCategory) Update() error {
	err := common.GetDB().Save(storeCategory).Error
	return err
}

func (storeCategory *StoreCategory) Delete() error {
	return common.GetDB().Delete(storeCategory).Error
}

func (storeCategory *StoreCategory) List(rawQuery string, rawOrder string, offset int, limit int) (*[]StoreCategory, int, error) {
	storeCategorys := []StoreCategory{}
	total := 0
	db := common.GetDB().Model(storeCategory)
	db, err := buildWhere(rawQuery, db)
	if err != nil {
		return &storeCategorys, total, err
	}
	db, err = buildOrder(rawOrder, db)
	if err != nil {
		return &storeCategorys, total, err
	}
	db.Offset(offset).
		Limit(limit).
		Find(&storeCategorys).
		Count(&total)
	err = db.Error
	return &storeCategorys, total, err
}

func (storeCategory *StoreCategory) Get() (*StoreCategory, error) {
	err := common.GetDB().Find(&storeCategory).Error
	return storeCategory, err
}
