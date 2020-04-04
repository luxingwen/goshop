//generate by gen
package models

import (
	"goshop/restful/common"
)

//通知模板表
type SystemNotice struct {
	Id         int    `gorm:"column:id"`          //通知模板id
	Title      string `gorm:"column:title"`       //通知标题
	Type       string `gorm:"column:type"`        //通知类型
	Icon       string `gorm:"column:icon"`        //图标
	Url        string `gorm:"column:url"`         //链接
	TableTitle string `gorm:"column:table_title"` //通知数据
	Template   string `gorm:"column:template"`    //通知模板
	PushAdmin  string `gorm:"column:push_admin"`  //通知管理员id
	Status     int    `gorm:"column:status"`      //状态

}

//修改默认表名
func (SystemNotice) TableName() string {
	return "eb_system_notice"
}

func (systemNotice *SystemNotice) Insert() error {
	err := common.GetDB().Create(systemNotice).Error
	return err
}

func (systemNotice *SystemNotice) Patch() error {
	err := common.GetDB().Model(systemNotice).Updates(systemNotice).Error
	return err
}

func (systemNotice *SystemNotice) Update() error {
	err := common.GetDB().Save(systemNotice).Error
	return err
}

func (systemNotice *SystemNotice) Delete() error {
	return common.GetDB().Delete(systemNotice).Error
}

func (systemNotice *SystemNotice) List(rawQuery string, rawOrder string, offset int, limit int) (*[]SystemNotice, int, error) {
	systemNotices := []SystemNotice{}
	total := 0
	db := common.GetDB().Model(systemNotice)
	db, err := buildWhere(rawQuery, db)
	if err != nil {
		return &systemNotices, total, err
	}
	db, err = buildOrder(rawOrder, db)
	if err != nil {
		return &systemNotices, total, err
	}
	db.Offset(offset).
		Limit(limit).
		Find(&systemNotices).
		Count(&total)
	err = db.Error
	return &systemNotices, total, err
}

func (systemNotice *SystemNotice) Get() (*SystemNotice, error) {
	err := common.GetDB().Find(&systemNotice).Error
	return systemNotice, err
}
