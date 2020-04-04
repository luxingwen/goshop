//generate by gen
package models

import (
	"goshop/restful/common"
)

//微信缓存表
type Cache struct {
	Key     string `gorm:"column:key"`      //
	Result  string `gorm:"column:result"`   //缓存数据
	AddTime int    `gorm:"column:add_time"` //缓存时间

}

//修改默认表名
func (Cache) TableName() string {
	return "eb_cache"
}

func (cache *Cache) Insert() error {
	err := common.GetDB().Create(cache).Error
	return err
}

func (cache *Cache) Patch() error {
	err := common.GetDB().Model(cache).Updates(cache).Error
	return err
}

func (cache *Cache) Update() error {
	err := common.GetDB().Save(cache).Error
	return err
}

func (cache *Cache) Delete() error {
	return common.GetDB().Delete(cache).Error
}

func (cache *Cache) List(rawQuery string, rawOrder string, offset int, limit int) (*[]Cache, int, error) {
	caches := []Cache{}
	total := 0
	db := common.GetDB().Model(cache)
	db, err := buildWhere(rawQuery, db)
	if err != nil {
		return &caches, total, err
	}
	db, err = buildOrder(rawOrder, db)
	if err != nil {
		return &caches, total, err
	}
	db.Offset(offset).
		Limit(limit).
		Find(&caches).
		Count(&total)
	err = db.Error
	return &caches, total, err
}

func (cache *Cache) Get() (*Cache, error) {
	err := common.GetDB().Find(&cache).Error
	return cache, err
}
