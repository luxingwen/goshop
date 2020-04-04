//generate by gen
package models

import (
	"goshop/restful/common"
)

//身份管理表
type SystemRole struct {
	Id       int    `gorm:"column:id"`        //身份管理id
	RoleName string `gorm:"column:role_name"` //身份管理名称
	Rules    string `gorm:"column:rules"`     //身份管理权限(menus_id)
	Level    int    `gorm:"column:level"`     //
	Status   int    `gorm:"column:status"`    //状态

}

//修改默认表名
func (SystemRole) TableName() string {
	return "eb_system_role"
}

func (systemRole *SystemRole) Insert() error {
	err := common.GetDB().Create(systemRole).Error
	return err
}

func (systemRole *SystemRole) Patch() error {
	err := common.GetDB().Model(systemRole).Updates(systemRole).Error
	return err
}

func (systemRole *SystemRole) Update() error {
	err := common.GetDB().Save(systemRole).Error
	return err
}

func (systemRole *SystemRole) Delete() error {
	return common.GetDB().Delete(systemRole).Error
}

func (systemRole *SystemRole) List(rawQuery string, rawOrder string, offset int, limit int) (*[]SystemRole, int, error) {
	systemRoles := []SystemRole{}
	total := 0
	db := common.GetDB().Model(systemRole)
	db, err := buildWhere(rawQuery, db)
	if err != nil {
		return &systemRoles, total, err
	}
	db, err = buildOrder(rawOrder, db)
	if err != nil {
		return &systemRoles, total, err
	}
	db.Offset(offset).
		Limit(limit).
		Find(&systemRoles).
		Count(&total)
	err = db.Error
	return &systemRoles, total, err
}

func (systemRole *SystemRole) Get() (*SystemRole, error) {
	err := common.GetDB().Find(&systemRole).Error
	return systemRole, err
}
