//generate by gen
package models

import (
	"goshop/restful/common"
)

//用户分组表
type UserGroup struct {
	Id        int    `gorm:"column:id"`         //
	GroupName string `gorm:"column:group_name"` //用户分组名称

}

//修改默认表名
func (UserGroup) TableName() string {
	return "eb_user_group"
}

func (userGroup *UserGroup) Insert() error {
	err := common.GetDB().Create(userGroup).Error
	return err
}

func (userGroup *UserGroup) Patch() error {
	err := common.GetDB().Model(userGroup).Updates(userGroup).Error
	return err
}

func (userGroup *UserGroup) Update() error {
	err := common.GetDB().Save(userGroup).Error
	return err
}

func (userGroup *UserGroup) Delete() error {
	return common.GetDB().Delete(userGroup).Error
}

func (userGroup *UserGroup) List(rawQuery string, rawOrder string, offset int, limit int) (*[]UserGroup, int, error) {
	userGroups := []UserGroup{}
	total := 0
	db := common.GetDB().Model(userGroup)
	db, err := buildWhere(rawQuery, db)
	if err != nil {
		return &userGroups, total, err
	}
	db, err = buildOrder(rawOrder, db)
	if err != nil {
		return &userGroups, total, err
	}
	db.Offset(offset).
		Limit(limit).
		Find(&userGroups).
		Count(&total)
	err = db.Error
	return &userGroups, total, err
}

func (userGroup *UserGroup) Get() (*UserGroup, error) {
	err := common.GetDB().Find(&userGroup).Error
	return userGroup, err
}
