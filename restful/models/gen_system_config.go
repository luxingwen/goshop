//generate by gen
package models

import (
	"goshop/restful/common"
)

//配置表
type SystemConfig struct {
	Id          int    `gorm:"column:id"`            //配置id
	MenuName    string `gorm:"column:menu_name"`     //字段名称
	Type        string `gorm:"column:type"`          //类型(文本框,单选按钮...)
	ConfigTabId int    `gorm:"column:config_tab_id"` //配置分类id
	Parameter   string `gorm:"column:parameter"`     //规则 单选框和多选框
	UploadType  int    `gorm:"column:upload_type"`   //上传文件格式1单图2多图3文件
	Required    string `gorm:"column:required"`      //规则
	Width       int    `gorm:"column:width"`         //多行文本框的宽度
	High        int    `gorm:"column:high"`          //多行文框的高度
	Value       string `gorm:"column:value"`         //默认值
	Info        string `gorm:"column:info"`          //配置名称
	Desc        string `gorm:"column:desc"`          //配置简介
	Sort        int    `gorm:"column:sort"`          //排序
	Status      int    `gorm:"column:status"`        //是否隐藏

}

//修改默认表名
func (SystemConfig) TableName() string {
	return "eb_system_config"
}

func (systemConfig *SystemConfig) Insert() error {
	err := common.GetDB().Create(systemConfig).Error
	return err
}

func (systemConfig *SystemConfig) Patch() error {
	err := common.GetDB().Model(systemConfig).Updates(systemConfig).Error
	return err
}

func (systemConfig *SystemConfig) Update() error {
	err := common.GetDB().Save(systemConfig).Error
	return err
}

func (systemConfig *SystemConfig) Delete() error {
	return common.GetDB().Delete(systemConfig).Error
}

func (systemConfig *SystemConfig) List(rawQuery string, rawOrder string, offset int, limit int) (*[]SystemConfig, int, error) {
	systemConfigs := []SystemConfig{}
	total := 0
	db := common.GetDB().Model(systemConfig)
	db, err := buildWhere(rawQuery, db)
	if err != nil {
		return &systemConfigs, total, err
	}
	db, err = buildOrder(rawOrder, db)
	if err != nil {
		return &systemConfigs, total, err
	}
	db.Offset(offset).
		Limit(limit).
		Find(&systemConfigs).
		Count(&total)
	err = db.Error
	return &systemConfigs, total, err
}

func (systemConfig *SystemConfig) Get() (*SystemConfig, error) {
	err := common.GetDB().Find(&systemConfig).Error
	return systemConfig, err
}
