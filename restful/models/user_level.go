package models

import (
	"time"

	"goshop/restful/common"
)

type VipInfo struct {
	Id       int     `gorm:"column:id"`       //
	AddTime  int     `gorm:"column:add_time"` //添加时间
	Discount float64 `gorm:"column:discount"` //享受折扣
	LevelId  int     `gorm:"column:level_id"` //等级vip
	Name     string  `gorm:"column:name"`     //会员名称
	Icon     string  `gorm:"column:icon"`     //会员图标
	IsPay    int     `gorm:"column:is_pay"`   //是否购买,1=购买,0=不购买
	Grade    int     `gorm:"column:grade"`    //会员等级
}

func (userLevel *UserLevel) GetUserLevelInfo(id int) (r *VipInfo, err error) {
	db := common.GetDB()
	systemUserLevel := &SystemUserLevel{}
	err = db.Raw("SELECT l.id,a.add_time,l.discount,a.level_id,l.name,l.money,l.icon,l.is_pay,l.grade FROM "+userLevel.TableName()+
		" a LEFT JOIN "+systemUserLevel.TableName()+" l ON l.id = a.level_id WHERE a.id = ?", id).Scan(&r).Error

	return

}

func (userLevel *UserLevel) GetLevelInfo(uid int, grade int) (err error) {
	db := common.GetDB()

	db = db.Select("level_id,is_forever,valid_time,id,status,grade")
	db = db.Where("status = ? AND is_del = ? AND uid = ?", 1, 0, uid)

	if grade > 0 {
		db = db.Where("grade < ?", grade)
	}
	err = db.Order("grade desc").Find(&userLevel).Error
	if err != nil {
		return
	}

	if userLevel.IsForever > 0 {
		return
	}

	// 会员已经过期
	if int(time.Now().Unix()) > userLevel.ValidTime {
		if userLevel.Status == 1 {
			userLevel.Status = 0
			// @Todo update
		}
		return userLevel.GetLevelInfo(uid, userLevel.Grade)
	}
	return
}
