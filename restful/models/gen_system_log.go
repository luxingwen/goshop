//generate by gen
package models

import (
	"goshop/restful/common"
)

//管理员操作记录表
type SystemLog struct {
	Id         int    `gorm:"column:id"`          //管理员操作记录ID
	AdminId    int    `gorm:"column:admin_id"`    //管理员id
	AdminName  string `gorm:"column:admin_name"`  //管理员姓名
	Path       string `gorm:"column:path"`        //链接
	Page       string `gorm:"column:page"`        //行为
	Method     string `gorm:"column:method"`      //访问类型
	Ip         string `gorm:"column:ip"`          //登录IP
	Type       string `gorm:"column:type"`        //类型
	AddTime    int    `gorm:"column:add_time"`    //操作时间
	MerchantId int    `gorm:"column:merchant_id"` //商户id

}

//修改默认表名
func (SystemLog) TableName() string {
	return "eb_system_log"
}

func (systemLog *SystemLog) Insert() error {
	err := common.GetDB().Create(systemLog).Error
	return err
}

func (systemLog *SystemLog) Patch() error {
	err := common.GetDB().Model(systemLog).Updates(systemLog).Error
	return err
}

func (systemLog *SystemLog) Update() error {
	err := common.GetDB().Save(systemLog).Error
	return err
}

func (systemLog *SystemLog) Delete() error {
	return common.GetDB().Delete(systemLog).Error
}

func (systemLog *SystemLog) List(rawQuery string, rawOrder string, offset int, limit int) (*[]SystemLog, int, error) {
	systemLogs := []SystemLog{}
	total := 0
	db := common.GetDB().Model(systemLog)
	db, err := buildWhere(rawQuery, db)
	if err != nil {
		return &systemLogs, total, err
	}
	db, err = buildOrder(rawOrder, db)
	if err != nil {
		return &systemLogs, total, err
	}
	db.Offset(offset).
		Limit(limit).
		Find(&systemLogs).
		Count(&total)
	err = db.Error
	return &systemLogs, total, err
}

func (systemLog *SystemLog) Get() (*SystemLog, error) {
	err := common.GetDB().Find(&systemLog).Error
	return systemLog, err
}
