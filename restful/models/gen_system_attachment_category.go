//generate by gen
package models

import (
	"goshop/restful/common"
)

//附件分类表
type SystemAttachmentCategory struct {
	Id     int    `gorm:"column:id"`     //
	Pid    int    `gorm:"column:pid"`    //父级ID
	Name   string `gorm:"column:name"`   //分类名称
	Enname string `gorm:"column:enname"` //分类目录

}

//修改默认表名
func (SystemAttachmentCategory) TableName() string {
	return "eb_system_attachment_category"
}

func (systemAttachmentCategory *SystemAttachmentCategory) Insert() error {
	err := common.GetDB().Create(systemAttachmentCategory).Error
	return err
}

func (systemAttachmentCategory *SystemAttachmentCategory) Patch() error {
	err := common.GetDB().Model(systemAttachmentCategory).Updates(systemAttachmentCategory).Error
	return err
}

func (systemAttachmentCategory *SystemAttachmentCategory) Update() error {
	err := common.GetDB().Save(systemAttachmentCategory).Error
	return err
}

func (systemAttachmentCategory *SystemAttachmentCategory) Delete() error {
	return common.GetDB().Delete(systemAttachmentCategory).Error
}

func (systemAttachmentCategory *SystemAttachmentCategory) List(rawQuery string, rawOrder string, offset int, limit int) (*[]SystemAttachmentCategory, int, error) {
	systemAttachmentCategorys := []SystemAttachmentCategory{}
	total := 0
	db := common.GetDB().Model(systemAttachmentCategory)
	db, err := buildWhere(rawQuery, db)
	if err != nil {
		return &systemAttachmentCategorys, total, err
	}
	db, err = buildOrder(rawOrder, db)
	if err != nil {
		return &systemAttachmentCategorys, total, err
	}
	db.Offset(offset).
		Limit(limit).
		Find(&systemAttachmentCategorys).
		Count(&total)
	err = db.Error
	return &systemAttachmentCategorys, total, err
}

func (systemAttachmentCategory *SystemAttachmentCategory) Get() (*SystemAttachmentCategory, error) {
	err := common.GetDB().Find(&systemAttachmentCategory).Error
	return systemAttachmentCategory, err
}
