package models

import (
	"goshop/restful/common"
)

func (systemGroupData *SystemGroupData) ListByGids(ids []int) (r []*SystemGroupData, err error) {
	db := common.GetDB()
	err = db.Where("gid in(?)", ids).Find(&r).Error
	return
}
