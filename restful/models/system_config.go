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

func (systemConfig *SystemConfig) GetByMenuName(name string) (r *SystemConfig, err error) {
	db := common.GetDB()
	r = new(SystemConfig)
	err = db.Table(systemConfig.TableName()).Where("menu_name = ?", name).First(&r).Error
	return
}
