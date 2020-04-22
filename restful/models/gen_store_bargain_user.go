//generate by gen
package models

import (
	"goshop/restful/common"
)

//用户参与砍价表
type StoreBargainUser struct {
	Id              int     `gorm:"column:id"`                //用户参与砍价表ID
	Uid             int     `gorm:"column:uid"`               //用户ID
	BargainId       int     `gorm:"column:bargain_id"`        //砍价产品id
	BargainPriceMin float64 `gorm:"column:bargain_price_min"` //砍价的最低价
	BargainPrice    float64 `gorm:"column:bargain_price"`     //砍价金额
	Price           float64 `gorm:"column:price"`             //砍掉的价格
	Status          int     `gorm:"column:status"`            //状态 1参与中 2 活动结束参与失败 3活动结束参与成功
	AddTime         int     `gorm:"column:add_time"`          //参与时间
	IsDel           int     `gorm:"column:is_del"`            //是否取消

}

//修改默认表名
func (StoreBargainUser) TableName() string {
	return "eb_store_bargain_user"
}

func (storeBargainUser *StoreBargainUser) Insert() error {
	err := common.GetDB().Create(storeBargainUser).Error
	return err
}

func (storeBargainUser *StoreBargainUser) Patch() error {
	err := common.GetDB().Model(storeBargainUser).Updates(storeBargainUser).Error
	return err
}

func (storeBargainUser *StoreBargainUser) Update() error {
	err := common.GetDB().Save(storeBargainUser).Error
	return err
}

func (storeBargainUser *StoreBargainUser) Delete() error {
	return common.GetDB().Delete(storeBargainUser).Error
}

func (storeBargainUser *StoreBargainUser) List(rawQuery string, rawOrder string, offset int, limit int) (*[]StoreBargainUser, int, error) {
	storeBargainUsers := []StoreBargainUser{}
	total := 0
	db := common.GetDB().Model(storeBargainUser)
	db, err := buildWhere(rawQuery, db)
	if err != nil {
		return &storeBargainUsers, total, err
	}
	db, err = buildOrder(rawOrder, db)
	if err != nil {
		return &storeBargainUsers, total, err
	}
	db.Offset(offset).
		Limit(limit).
		Find(&storeBargainUsers).
		Count(&total)
	err = db.Error
	return &storeBargainUsers, total, err
}

func (storeBargainUser *StoreBargainUser) Get() (*StoreBargainUser, error) {
	err := common.GetDB().Find(&storeBargainUser).Error
	return storeBargainUser, err
}

// 根据砍价产品编号获取正在参与人的编号
func (storeBargainUser *StoreBargainUser) GetUserIdList(bargainId int) (r []*StoreBargainUser, err error) {
	if bargainId <= 0 {
		return
	}
	db := common.GetDB()
	err = db.Select("uid, id").Where("status = ? AND bargain_id = ?", 1, bargainId).Scan(&r).Error
	if err != nil {
		return
	}
	return
}
