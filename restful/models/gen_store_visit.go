//generate by gen
package models

import (
	"goshop/restful/common"
)

//产品浏览分析表
type StoreVisit struct {
	Id          int    `gorm:"column:id"`           //
	ProductId   int    `gorm:"column:product_id"`   //产品ID
	ProductType string `gorm:"column:product_type"` //产品类型
	CateId      int    `gorm:"column:cate_id"`      //产品分类ID
	Type        string `gorm:"column:type"`         //产品类型
	Uid         int    `gorm:"column:uid"`          //用户ID
	Count       int    `gorm:"column:count"`        //访问次数
	Content     string `gorm:"column:content"`      //备注描述
	AddTime     int    `gorm:"column:add_time"`     //添加时间

}

//修改默认表名
func (StoreVisit) TableName() string {
	return "eb_store_visit"
}

func (storeVisit *StoreVisit) Insert() error {
	err := common.GetDB().Create(storeVisit).Error
	return err
}

func (storeVisit *StoreVisit) Patch() error {
	err := common.GetDB().Model(storeVisit).Updates(storeVisit).Error
	return err
}

func (storeVisit *StoreVisit) Update() error {
	err := common.GetDB().Save(storeVisit).Error
	return err
}

func (storeVisit *StoreVisit) Delete() error {
	return common.GetDB().Delete(storeVisit).Error
}

func (storeVisit *StoreVisit) List(rawQuery string, rawOrder string, offset int, limit int) (*[]StoreVisit, int, error) {
	storeVisits := []StoreVisit{}
	total := 0
	db := common.GetDB().Model(storeVisit)
	db, err := buildWhere(rawQuery, db)
	if err != nil {
		return &storeVisits, total, err
	}
	db, err = buildOrder(rawOrder, db)
	if err != nil {
		return &storeVisits, total, err
	}
	db.Offset(offset).
		Limit(limit).
		Find(&storeVisits).
		Count(&total)
	err = db.Error
	return &storeVisits, total, err
}

func (storeVisit *StoreVisit) Get() (*StoreVisit, error) {
	err := common.GetDB().Find(&storeVisit).Error
	return storeVisit, err
}
