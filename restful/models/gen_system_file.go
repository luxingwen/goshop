//generate by gen
package models

import (
	"goshop/restful/common"
)

//文件对比表
type SystemFile struct {
	Id       int    `gorm:"column:id"`       //文件对比ID
	Cthash   string `gorm:"column:cthash"`   //文件内容
	Filename string `gorm:"column:filename"` //文价名称
	Atime    string `gorm:"column:atime"`    //上次访问时间
	Mtime    string `gorm:"column:mtime"`    //上次修改时间
	Ctime    string `gorm:"column:ctime"`    //上次改变时间

}

//修改默认表名
func (SystemFile) TableName() string {
	return "eb_system_file"
}

func (systemFile *SystemFile) Insert() error {
	err := common.GetDB().Create(systemFile).Error
	return err
}

func (systemFile *SystemFile) Patch() error {
	err := common.GetDB().Model(systemFile).Updates(systemFile).Error
	return err
}

func (systemFile *SystemFile) Update() error {
	err := common.GetDB().Save(systemFile).Error
	return err
}

func (systemFile *SystemFile) Delete() error {
	return common.GetDB().Delete(systemFile).Error
}

func (systemFile *SystemFile) List(rawQuery string, rawOrder string, offset int, limit int) (*[]SystemFile, int, error) {
	systemFiles := []SystemFile{}
	total := 0
	db := common.GetDB().Model(systemFile)
	db, err := buildWhere(rawQuery, db)
	if err != nil {
		return &systemFiles, total, err
	}
	db, err = buildOrder(rawOrder, db)
	if err != nil {
		return &systemFiles, total, err
	}
	db.Offset(offset).
		Limit(limit).
		Find(&systemFiles).
		Count(&total)
	err = db.Error
	return &systemFiles, total, err
}

func (systemFile *SystemFile) Get() (*SystemFile, error) {
	err := common.GetDB().Find(&systemFile).Error
	return systemFile, err
}
