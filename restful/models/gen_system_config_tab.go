//generate by gen
package models

import (
	"goshop/restful/common"
)

//配置分类表
type SystemConfigTab struct {
	Id       int    `gorm:"column:id"`        //配置分类id
	Title    string `gorm:"column:title"`     //配置分类名称
	EngTitle string `gorm:"column:eng_title"` //配置分类英文名称
	Status   int    `gorm:"column:status"`    //配置分类状态
	Info     int    `gorm:"column:info"`      //配置分类是否显示
	Icon     string `gorm:"column:icon"`      //图标
	Type     int    `gorm:"column:type"`      //配置类型

}

//修改默认表名
func (SystemConfigTab) TableName() string {
	return "eb_system_config_tab"
}

func (systemConfigTab *SystemConfigTab) Insert() error {
	err := common.GetDB().Create(systemConfigTab).Error
	return err
}

func (systemConfigTab *SystemConfigTab) Patch() error {
	err := common.GetDB().Model(systemConfigTab).Updates(systemConfigTab).Error
	return err
}

func (systemConfigTab *SystemConfigTab) Update() error {
	err := common.GetDB().Save(systemConfigTab).Error
	return err
}

func (systemConfigTab *SystemConfigTab) Delete() error {
	return common.GetDB().Delete(systemConfigTab).Error
}

func (systemConfigTab *SystemConfigTab) List(rawQuery string, rawOrder string, offset int, limit int) (*[]SystemConfigTab, int, error) {
	systemConfigTabs := []SystemConfigTab{}
	total := 0
	db := common.GetDB().Model(systemConfigTab)
	db, err := buildWhere(rawQuery, db)
	if err != nil {
		return &systemConfigTabs, total, err
	}
	db, err = buildOrder(rawOrder, db)
	if err != nil {
		return &systemConfigTabs, total, err
	}
	db.Offset(offset).
		Limit(limit).
		Find(&systemConfigTabs).
		Count(&total)
	err = db.Error
	return &systemConfigTabs, total, err
}

func (systemConfigTab *SystemConfigTab) Get() (*SystemConfigTab, error) {
	err := common.GetDB().Find(&systemConfigTab).Error
	return systemConfigTab, err
}
