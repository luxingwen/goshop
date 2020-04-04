//generate by gen
package models

import (
	"goshop/restful/common"
)

//砍价用户帮助表
type StoreBargainUserHelp struct {
	Id            int     `gorm:"column:id"`              //砍价用户帮助表ID
	Uid           int     `gorm:"column:uid"`             //帮助的用户id
	BargainId     int     `gorm:"column:bargain_id"`      //砍价产品ID
	BargainUserId int     `gorm:"column:bargain_user_id"` //用户参与砍价表id
	Price         float64 `gorm:"column:price"`           //帮助砍价多少金额
	AddTime       int     `gorm:"column:add_time"`        //添加时间

}

//修改默认表名
func (StoreBargainUserHelp) TableName() string {
	return "eb_store_bargain_user_help"
}

func (storeBargainUserHelp *StoreBargainUserHelp) Insert() error {
	err := common.GetDB().Create(storeBargainUserHelp).Error
	return err
}

func (storeBargainUserHelp *StoreBargainUserHelp) Patch() error {
	err := common.GetDB().Model(storeBargainUserHelp).Updates(storeBargainUserHelp).Error
	return err
}

func (storeBargainUserHelp *StoreBargainUserHelp) Update() error {
	err := common.GetDB().Save(storeBargainUserHelp).Error
	return err
}

func (storeBargainUserHelp *StoreBargainUserHelp) Delete() error {
	return common.GetDB().Delete(storeBargainUserHelp).Error
}

func (storeBargainUserHelp *StoreBargainUserHelp) List(rawQuery string, rawOrder string, offset int, limit int) (*[]StoreBargainUserHelp, int, error) {
	storeBargainUserHelps := []StoreBargainUserHelp{}
	total := 0
	db := common.GetDB().Model(storeBargainUserHelp)
	db, err := buildWhere(rawQuery, db)
	if err != nil {
		return &storeBargainUserHelps, total, err
	}
	db, err = buildOrder(rawOrder, db)
	if err != nil {
		return &storeBargainUserHelps, total, err
	}
	db.Offset(offset).
		Limit(limit).
		Find(&storeBargainUserHelps).
		Count(&total)
	err = db.Error
	return &storeBargainUserHelps, total, err
}

func (storeBargainUserHelp *StoreBargainUserHelp) Get() (*StoreBargainUserHelp, error) {
	err := common.GetDB().Find(&storeBargainUserHelp).Error
	return storeBargainUserHelp, err
}
