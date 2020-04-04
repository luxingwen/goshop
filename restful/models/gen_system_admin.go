//generate by gen
package models

import (
	"goshop/restful/common"
)

//后台管理员表
type SystemAdmin struct {
	Id         int    `gorm:"column:id"`          //后台管理员表ID
	Account    string `gorm:"column:account"`     //后台管理员账号
	Pwd        string `gorm:"column:pwd"`         //后台管理员密码
	RealName   string `gorm:"column:real_name"`   //后台管理员姓名
	Roles      string `gorm:"column:roles"`       //后台管理员权限(menus_id)
	LastIp     string `gorm:"column:last_ip"`     //后台管理员最后一次登录ip
	LastTime   int    `gorm:"column:last_time"`   //后台管理员最后一次登录时间
	AddTime    int    `gorm:"column:add_time"`    //后台管理员添加时间
	LoginCount int    `gorm:"column:login_count"` //登录次数
	Level      int    `gorm:"column:level"`       //后台管理员级别
	Status     int    `gorm:"column:status"`      //后台管理员状态 1有效0无效
	IsDel      int    `gorm:"column:is_del"`      //

}

//修改默认表名
func (SystemAdmin) TableName() string {
	return "eb_system_admin"
}

func (systemAdmin *SystemAdmin) Insert() error {
	err := common.GetDB().Create(systemAdmin).Error
	return err
}

func (systemAdmin *SystemAdmin) Patch() error {
	err := common.GetDB().Model(systemAdmin).Updates(systemAdmin).Error
	return err
}

func (systemAdmin *SystemAdmin) Update() error {
	err := common.GetDB().Save(systemAdmin).Error
	return err
}

func (systemAdmin *SystemAdmin) Delete() error {
	return common.GetDB().Delete(systemAdmin).Error
}

func (systemAdmin *SystemAdmin) List(rawQuery string, rawOrder string, offset int, limit int) (*[]SystemAdmin, int, error) {
	systemAdmins := []SystemAdmin{}
	total := 0
	db := common.GetDB().Model(systemAdmin)
	db, err := buildWhere(rawQuery, db)
	if err != nil {
		return &systemAdmins, total, err
	}
	db, err = buildOrder(rawOrder, db)
	if err != nil {
		return &systemAdmins, total, err
	}
	db.Offset(offset).
		Limit(limit).
		Find(&systemAdmins).
		Count(&total)
	err = db.Error
	return &systemAdmins, total, err
}

func (systemAdmin *SystemAdmin) Get() (*SystemAdmin, error) {
	err := common.GetDB().Find(&systemAdmin).Error
	return systemAdmin, err
}
