package models

import (
	"goshop/restful/common"
)

func (storeCategory *StoreCategory) ListByIndex(limit int) (r []*StoreCategory, err error) {
	db := common.GetDB()
	err = db.Select("id, cate_name, pid, pic").Where("pid > ? and is_show = ?", 0, 1).Order("sort DESC").Limit(limit).Find(&r).Error
	return
}
