//generate by gen
package models

import (
	"goshop/restful/common"
)

//附件管理表
type SystemAttachment struct {
	AttId      int    `gorm:"column:att_id"`      //
	Name       string `gorm:"column:name"`        //附件名称
	AttDir     string `gorm:"column:att_dir"`     //附件路径
	SattDir    string `gorm:"column:satt_dir"`    //压缩图片路径
	AttSize    string `gorm:"column:att_size"`    //附件大小
	AttType    string `gorm:"column:att_type"`    //附件类型
	Pid        int    `gorm:"column:pid"`         //分类ID0编辑器,1产品图片,2拼团图片,3砍价图片,4秒杀图片,5文章图片,6组合数据图
	Time       int    `gorm:"column:time"`        //上传时间
	ImageType  int    `gorm:"column:image_type"`  //图片上传类型 1本地 2七牛云 3OSS 4COS
	ModuleType int    `gorm:"column:module_type"` //图片上传模块类型 1 后台上传 2 用户生成

}

//修改默认表名
func (SystemAttachment) TableName() string {
	return "eb_system_attachment"
}

func (systemAttachment *SystemAttachment) Insert() error {
	err := common.GetDB().Create(systemAttachment).Error
	return err
}

func (systemAttachment *SystemAttachment) Patch() error {
	err := common.GetDB().Model(systemAttachment).Updates(systemAttachment).Error
	return err
}

func (systemAttachment *SystemAttachment) Update() error {
	err := common.GetDB().Save(systemAttachment).Error
	return err
}

func (systemAttachment *SystemAttachment) Delete() error {
	return common.GetDB().Delete(systemAttachment).Error
}

func (systemAttachment *SystemAttachment) List(rawQuery string, rawOrder string, offset int, limit int) (*[]SystemAttachment, int, error) {
	systemAttachments := []SystemAttachment{}
	total := 0
	db := common.GetDB().Model(systemAttachment)
	db, err := buildWhere(rawQuery, db)
	if err != nil {
		return &systemAttachments, total, err
	}
	db, err = buildOrder(rawOrder, db)
	if err != nil {
		return &systemAttachments, total, err
	}
	db.Offset(offset).
		Limit(limit).
		Find(&systemAttachments).
		Count(&total)
	err = db.Error
	return &systemAttachments, total, err
}

func (systemAttachment *SystemAttachment) Get() (*SystemAttachment, error) {
	err := common.GetDB().Find(&systemAttachment).Error
	return systemAttachment, err
}
