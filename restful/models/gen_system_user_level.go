//generate by gen
package models

import (
	"goshop/restful/common"
)

//设置用户等级表
type SystemUserLevel struct {
	Id        int     `gorm:"column:id"`         //
	MerId     int     `gorm:"column:mer_id"`     //商户id
	Name      string  `gorm:"column:name"`       //会员名称
	Money     float64 `gorm:"column:money"`      //购买金额
	ValidDate int     `gorm:"column:valid_date"` //有效时间
	IsForever int     `gorm:"column:is_forever"` //是否为永久会员
	IsPay     int     `gorm:"column:is_pay"`     //是否购买,1=购买,0=不购买
	IsShow    int     `gorm:"column:is_show"`    //是否显示 1=显示,0=隐藏
	Grade     int     `gorm:"column:grade"`      //会员等级
	Discount  float64 `gorm:"column:discount"`   //享受折扣
	Image     string  `gorm:"column:image"`      //会员卡背景
	Icon      string  `gorm:"column:icon"`       //会员图标
	Explain   string  `gorm:"column:explain"`    //说明
	AddTime   int     `gorm:"column:add_time"`   //添加时间
	IsDel     int     `gorm:"column:is_del"`     //是否删除.1=删除,0=未删除

}

//修改默认表名
func (SystemUserLevel) TableName() string {
	return "eb_system_user_level"
}

func (systemUserLevel *SystemUserLevel) Insert() error {
	err := common.GetDB().Create(systemUserLevel).Error
	return err
}

func (systemUserLevel *SystemUserLevel) Patch() error {
	err := common.GetDB().Model(systemUserLevel).Updates(systemUserLevel).Error
	return err
}

func (systemUserLevel *SystemUserLevel) Update() error {
	err := common.GetDB().Save(systemUserLevel).Error
	return err
}

func (systemUserLevel *SystemUserLevel) Delete() error {
	return common.GetDB().Delete(systemUserLevel).Error
}

func (systemUserLevel *SystemUserLevel) List(rawQuery string, rawOrder string, offset int, limit int) (*[]SystemUserLevel, int, error) {
	systemUserLevels := []SystemUserLevel{}
	total := 0
	db := common.GetDB().Model(systemUserLevel)
	db, err := buildWhere(rawQuery, db)
	if err != nil {
		return &systemUserLevels, total, err
	}
	db, err = buildOrder(rawOrder, db)
	if err != nil {
		return &systemUserLevels, total, err
	}
	db.Offset(offset).
		Limit(limit).
		Find(&systemUserLevels).
		Count(&total)
	err = db.Error
	return &systemUserLevels, total, err
}

func (systemUserLevel *SystemUserLevel) Get() (*SystemUserLevel, error) {
	err := common.GetDB().Find(&systemUserLevel).Error
	return systemUserLevel, err
}
