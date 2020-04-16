package models

import (
	"goshop/restful/common"
)

func (systemConfig *SystemConfig) AllMenuNameKey() (mMap map[string]*SystemConfig, err error) {
	db := common.GetDB()
	rList := make([]*SystemConfig, 0)
	err = db.Find(&rList).Error
	if err != nil {
		return
	}

	mMap = make(map[string]*SystemConfig, 0)

	for _, item := range rList {
		mMap[item.MenuName] = item
	}
	return
}
