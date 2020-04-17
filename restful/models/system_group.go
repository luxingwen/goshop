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

func (systemGroup *SystemGroup) GetByConfigName(name string) (err error) {
	db := common.GetDB()
	err = db.Where("config_name = ?", name).Find(&systemGroup).Error
	return
}

func (systemGroup *SystemGroup) GetByConfigNameKey() (r map[string]*SystemGroup, err error) {
	db := common.GetDB()
	list := make([]*SystemGroup, 0)
	err = db.Find(&list).Error
	if err != nil {
		return
	}
	r = make(map[string]*SystemGroup, 0)

	for _, item := range list {
		r[item.ConfigName] = item
	}
	return
}
