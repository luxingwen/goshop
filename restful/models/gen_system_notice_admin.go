//generate by gen
package models

import (
	"goshop/restful/common"
)

//通知记录表
type SystemNoticeAdmin struct {
	Id         int    `gorm:"column:id"`          //通知记录ID
	NoticeType string `gorm:"column:notice_type"` //通知类型
	AdminId    int    `gorm:"column:admin_id"`    //通知的管理员
	LinkId     int    `gorm:"column:link_id"`     //关联ID
	TableData  string `gorm:"column:table_data"`  //通知的数据
	IsClick    int    `gorm:"column:is_click"`    //点击次数
	IsVisit    int    `gorm:"column:is_visit"`    //访问次数
	VisitTime  int    `gorm:"column:visit_time"`  //访问时间
	AddTime    int    `gorm:"column:add_time"`    //通知时间

}

//修改默认表名
func (SystemNoticeAdmin) TableName() string {
	return "eb_system_notice_admin"
}

func (systemNoticeAdmin *SystemNoticeAdmin) Insert() error {
	err := common.GetDB().Create(systemNoticeAdmin).Error
	return err
}

func (systemNoticeAdmin *SystemNoticeAdmin) Patch() error {
	err := common.GetDB().Model(systemNoticeAdmin).Updates(systemNoticeAdmin).Error
	return err
}

func (systemNoticeAdmin *SystemNoticeAdmin) Update() error {
	err := common.GetDB().Save(systemNoticeAdmin).Error
	return err
}

func (systemNoticeAdmin *SystemNoticeAdmin) Delete() error {
	return common.GetDB().Delete(systemNoticeAdmin).Error
}

func (systemNoticeAdmin *SystemNoticeAdmin) List(rawQuery string, rawOrder string, offset int, limit int) (*[]SystemNoticeAdmin, int, error) {
	systemNoticeAdmins := []SystemNoticeAdmin{}
	total := 0
	db := common.GetDB().Model(systemNoticeAdmin)
	db, err := buildWhere(rawQuery, db)
	if err != nil {
		return &systemNoticeAdmins, total, err
	}
	db, err = buildOrder(rawOrder, db)
	if err != nil {
		return &systemNoticeAdmins, total, err
	}
	db.Offset(offset).
		Limit(limit).
		Find(&systemNoticeAdmins).
		Count(&total)
	err = db.Error
	return &systemNoticeAdmins, total, err
}

func (systemNoticeAdmin *SystemNoticeAdmin) Get() (*SystemNoticeAdmin, error) {
	err := common.GetDB().Find(&systemNoticeAdmin).Error
	return systemNoticeAdmin, err
}
