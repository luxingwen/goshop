package models

import (
	"goshop/restful/common"
)

func (systemGroup *SystemGroup) All() (r []*SystemGroup, err error) {

	db := common.GetDB()
	r = make([]*SystemGroup, 0)
	err = db.Find(&r).Error
	return
}
