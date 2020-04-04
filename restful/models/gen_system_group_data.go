//generate by gen
package models

import (
	"goshop/restful/common"
)

//组合数据详情表
type SystemGroupData struct {
	Id      int    `gorm:"column:id"`       //组合数据详情ID
	Gid     int    `gorm:"column:gid"`      //对应的数据组id
	Value   string `gorm:"column:value"`    //数据组对应的数据值（json数据）
	AddTime int    `gorm:"column:add_time"` //添加数据时间
	Sort    int    `gorm:"column:sort"`     //数据排序
	Status  int    `gorm:"column:status"`   //状态（1：开启；2：关闭；）

}

//修改默认表名
func (SystemGroupData) TableName() string {
	return "eb_system_group_data"
}

func (systemGroupData *SystemGroupData) Insert() error {
	err := common.GetDB().Create(systemGroupData).Error
	return err
}

func (systemGroupData *SystemGroupData) Patch() error {
	err := common.GetDB().Model(systemGroupData).Updates(systemGroupData).Error
	return err
}

func (systemGroupData *SystemGroupData) Update() error {
	err := common.GetDB().Save(systemGroupData).Error
	return err
}

func (systemGroupData *SystemGroupData) Delete() error {
	return common.GetDB().Delete(systemGroupData).Error
}

func (systemGroupData *SystemGroupData) List(rawQuery string, rawOrder string, offset int, limit int) (*[]SystemGroupData, int, error) {
	systemGroupDatas := []SystemGroupData{}
	total := 0
	db := common.GetDB().Model(systemGroupData)
	db, err := buildWhere(rawQuery, db)
	if err != nil {
		return &systemGroupDatas, total, err
	}
	db, err = buildOrder(rawOrder, db)
	if err != nil {
		return &systemGroupDatas, total, err
	}
	db.Offset(offset).
		Limit(limit).
		Find(&systemGroupDatas).
		Count(&total)
	err = db.Error
	return &systemGroupDatas, total, err
}

func (systemGroupData *SystemGroupData) Get() (*SystemGroupData, error) {
	err := common.GetDB().Find(&systemGroupData).Error
	return systemGroupData, err
}
