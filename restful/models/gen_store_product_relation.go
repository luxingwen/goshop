//generate by gen
package models

import (
	"goshop/restful/common"
)

//商品点赞和收藏表
type StoreProductRelation struct {
	Uid       int    `gorm:"column:uid"`        //用户ID
	ProductId int    `gorm:"column:product_id"` //商品ID
	Type      string `gorm:"column:type"`       //类型(收藏(collect）、点赞(like))
	Category  string `gorm:"column:category"`   //某种类型的商品(普通商品、秒杀商品)
	AddTime   int    `gorm:"column:add_time"`   //添加时间

}

//修改默认表名
func (StoreProductRelation) TableName() string {
	return "eb_store_product_relation"
}

func (storeProductRelation *StoreProductRelation) Insert() error {
	err := common.GetDB().Create(storeProductRelation).Error
	return err
}

func (storeProductRelation *StoreProductRelation) Patch() error {
	err := common.GetDB().Model(storeProductRelation).Updates(storeProductRelation).Error
	return err
}

func (storeProductRelation *StoreProductRelation) Update() error {
	err := common.GetDB().Save(storeProductRelation).Error
	return err
}

func (storeProductRelation *StoreProductRelation) Delete() error {
	return common.GetDB().Delete(storeProductRelation).Error
}

func (storeProductRelation *StoreProductRelation) List(rawQuery string, rawOrder string, offset int, limit int) (*[]StoreProductRelation, int, error) {
	storeProductRelations := []StoreProductRelation{}
	total := 0
	db := common.GetDB().Model(storeProductRelation)
	db, err := buildWhere(rawQuery, db)
	if err != nil {
		return &storeProductRelations, total, err
	}
	db, err = buildOrder(rawOrder, db)
	if err != nil {
		return &storeProductRelations, total, err
	}
	db.Offset(offset).
		Limit(limit).
		Find(&storeProductRelations).
		Count(&total)
	err = db.Error
	return &storeProductRelations, total, err
}

func (storeProductRelation *StoreProductRelation) Get() (*StoreProductRelation, error) {
	err := common.GetDB().Find(&storeProductRelation).Error
	return storeProductRelation, err
}

// 获取用户收藏的所有产品的个数
func (storeProductRelation *StoreProductRelation) GetUserIdCollect(uid int) (count int, err error) {
	db := common.GetDB()
	err = db.Table(storeProductRelation.TableName()).Where("uid = ? AND type = ?", uid, "collect").Count(&count).Error
	return
}

// 获取用户点赞所有产品的个数
func (storeProductRelation *StoreProductRelation) GetUserIdLike(uid int) (count int, err error) {
	db := common.GetDB()
	err = db.Table(storeProductRelation.TableName()).Where("uid = ? AND type = ?", uid, "like").Count(&count).Error
	return
}
