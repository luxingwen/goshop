//generate by gen
package models

import (
	"goshop/restful/common"
)

//用户等级记录表
type UserLevel struct {
	Id        int    `gorm:"column:id"`         //
	Uid       int    `gorm:"column:uid"`        //用户uid
	LevelId   int    `gorm:"column:level_id"`   //等级vip
	Grade     int    `gorm:"column:grade"`      //会员等级
	ValidTime int    `gorm:"column:valid_time"` //过期时间
	IsForever int    `gorm:"column:is_forever"` //是否永久
	MerId     int    `gorm:"column:mer_id"`     //商户id
	Status    int    `gorm:"column:status"`     //0:禁止,1:正常
	Mark      string `gorm:"column:mark"`       //备注
	Remind    int    `gorm:"column:remind"`     //是否已通知
	IsDel     int    `gorm:"column:is_del"`     //是否删除,0=未删除,1=删除
	AddTime   int    `gorm:"column:add_time"`   //添加时间
	Discount  int    `gorm:"column:discount"`   //享受折扣

}

//修改默认表名
func (UserLevel) TableName() string {
	return "eb_user_level"
}

func (userLevel *UserLevel) Insert() error {
	err := common.GetDB().Create(userLevel).Error
	return err
}

func (userLevel *UserLevel) Patch() error {
	err := common.GetDB().Model(userLevel).Updates(userLevel).Error
	return err
}

func (userLevel *UserLevel) Update() error {
	err := common.GetDB().Save(userLevel).Error
	return err
}

func (userLevel *UserLevel) Delete() error {
	return common.GetDB().Delete(userLevel).Error
}

func (userLevel *UserLevel) List(rawQuery string, rawOrder string, offset int, limit int) (*[]UserLevel, int, error) {
	userLevels := []UserLevel{}
	total := 0
	db := common.GetDB().Model(userLevel)
	db, err := buildWhere(rawQuery, db)
	if err != nil {
		return &userLevels, total, err
	}
	db, err = buildOrder(rawOrder, db)
	if err != nil {
		return &userLevels, total, err
	}
	db.Offset(offset).
		Limit(limit).
		Find(&userLevels).
		Count(&total)
	err = db.Error
	return &userLevels, total, err
}

func (userLevel *UserLevel) Get() (*UserLevel, error) {
	err := common.GetDB().Find(&userLevel).Error
	return userLevel, err
}
