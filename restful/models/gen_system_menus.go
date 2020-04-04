//generate by gen
package models

import (
	"goshop/restful/common"
)

//菜单表
type SystemMenus struct {
	Id         int    `gorm:"column:id"`         //菜单ID
	Pid        int    `gorm:"column:pid"`        //父级id
	Icon       string `gorm:"column:icon"`       //图标
	MenuName   string `gorm:"column:menu_name"`  //按钮名
	Module     string `gorm:"column:module"`     //模块名
	Controller string `gorm:"column:controller"` //控制器
	Action     string `gorm:"column:action"`     //方法名
	Params     string `gorm:"column:params"`     //参数
	Sort       int    `gorm:"column:sort"`       //排序
	IsShow     int    `gorm:"column:is_show"`    //是否显示
	Access     int    `gorm:"column:access"`     //子管理员是否可用

}

//修改默认表名
func (SystemMenus) TableName() string {
	return "eb_system_menus"
}

func (systemMenus *SystemMenus) Insert() error {
	err := common.GetDB().Create(systemMenus).Error
	return err
}

func (systemMenus *SystemMenus) Patch() error {
	err := common.GetDB().Model(systemMenus).Updates(systemMenus).Error
	return err
}

func (systemMenus *SystemMenus) Update() error {
	err := common.GetDB().Save(systemMenus).Error
	return err
}

func (systemMenus *SystemMenus) Delete() error {
	return common.GetDB().Delete(systemMenus).Error
}

func (systemMenus *SystemMenus) List(rawQuery string, rawOrder string, offset int, limit int) (*[]SystemMenus, int, error) {
	systemMenuss := []SystemMenus{}
	total := 0
	db := common.GetDB().Model(systemMenus)
	db, err := buildWhere(rawQuery, db)
	if err != nil {
		return &systemMenuss, total, err
	}
	db, err = buildOrder(rawOrder, db)
	if err != nil {
		return &systemMenuss, total, err
	}
	db.Offset(offset).
		Limit(limit).
		Find(&systemMenuss).
		Count(&total)
	err = db.Error
	return &systemMenuss, total, err
}

func (systemMenus *SystemMenus) Get() (*SystemMenus, error) {
	err := common.GetDB().Find(&systemMenus).Error
	return systemMenus, err
}
