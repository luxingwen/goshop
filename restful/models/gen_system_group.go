//generate by gen
package models

import (
	"goshop/restful/common"
)

//组合数据表
type SystemGroup struct {
	Id         int    `gorm:"column:id"`          //组合数据ID
	Name       string `gorm:"column:name"`        //数据组名称
	Info       string `gorm:"column:info"`        //数据提示
	ConfigName string `gorm:"column:config_name"` //数据字段
	Fields     string `gorm:"column:fields"`      //数据组字段以及类型（json数据）

}

//修改默认表名
func (SystemGroup) TableName() string {
	return "eb_system_group"
}

func (systemGroup *SystemGroup) Insert() error {
	err := common.GetDB().Create(systemGroup).Error
	return err
}

func (systemGroup *SystemGroup) Patch() error {
	err := common.GetDB().Model(systemGroup).Updates(systemGroup).Error
	return err
}

func (systemGroup *SystemGroup) Update() error {
	err := common.GetDB().Save(systemGroup).Error
	return err
}

func (systemGroup *SystemGroup) Delete() error {
	return common.GetDB().Delete(systemGroup).Error
}

func (systemGroup *SystemGroup) List(rawQuery string, rawOrder string, offset int, limit int) (*[]SystemGroup, int, error) {
	systemGroups := []SystemGroup{}
	total := 0
	db := common.GetDB().Model(systemGroup)
	db, err := buildWhere(rawQuery, db)
	if err != nil {
		return &systemGroups, total, err
	}
	db, err = buildOrder(rawOrder, db)
	if err != nil {
		return &systemGroups, total, err
	}
	db.Offset(offset).
		Limit(limit).
		Find(&systemGroups).
		Count(&total)
	err = db.Error
	return &systemGroups, total, err
}

func (systemGroup *SystemGroup) Get() (*SystemGroup, error) {
	err := common.GetDB().Find(&systemGroup).Error
	return systemGroup, err
}
