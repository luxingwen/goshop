//generate by gen
package models

import (
	"goshop/restful/common"
)

//小程序access_token表
type RoutineAccessToken struct {
	Id          int    `gorm:"column:id"`           //小程序access_token表ID
	AccessToken string `gorm:"column:access_token"` //openid
	StopTime    int    `gorm:"column:stop_time"`    //添加时间

}

//修改默认表名
func (RoutineAccessToken) TableName() string {
	return "eb_routine_access_token"
}

func (routineAccessToken *RoutineAccessToken) Insert() error {
	err := common.GetDB().Create(routineAccessToken).Error
	return err
}

func (routineAccessToken *RoutineAccessToken) Patch() error {
	err := common.GetDB().Model(routineAccessToken).Updates(routineAccessToken).Error
	return err
}

func (routineAccessToken *RoutineAccessToken) Update() error {
	err := common.GetDB().Save(routineAccessToken).Error
	return err
}

func (routineAccessToken *RoutineAccessToken) Delete() error {
	return common.GetDB().Delete(routineAccessToken).Error
}

func (routineAccessToken *RoutineAccessToken) List(rawQuery string, rawOrder string, offset int, limit int) (*[]RoutineAccessToken, int, error) {
	routineAccessTokens := []RoutineAccessToken{}
	total := 0
	db := common.GetDB().Model(routineAccessToken)
	db, err := buildWhere(rawQuery, db)
	if err != nil {
		return &routineAccessTokens, total, err
	}
	db, err = buildOrder(rawOrder, db)
	if err != nil {
		return &routineAccessTokens, total, err
	}
	db.Offset(offset).
		Limit(limit).
		Find(&routineAccessTokens).
		Count(&total)
	err = db.Error
	return &routineAccessTokens, total, err
}

func (routineAccessToken *RoutineAccessToken) Get() (*RoutineAccessToken, error) {
	err := common.GetDB().Find(&routineAccessToken).Error
	return routineAccessToken, err
}
