//generate by gen
package models

import (
	"goshop/restful/common"
)

//优惠券前台领取表
type StoreCouponIssue struct {
	Id          int `gorm:"column:id"`           //
	Cid         int `gorm:"column:cid"`          //优惠券ID
	StartTime   int `gorm:"column:start_time"`   //优惠券领取开启时间
	EndTime     int `gorm:"column:end_time"`     //优惠券领取结束时间
	TotalCount  int `gorm:"column:total_count"`  //优惠券领取数量
	RemainCount int `gorm:"column:remain_count"` //优惠券剩余领取数量
	IsPermanent int `gorm:"column:is_permanent"` //是否无限张数
	Status      int `gorm:"column:status"`       //1 正常 0 未开启 -1 已无效
	IsDel       int `gorm:"column:is_del"`       //
	AddTime     int `gorm:"column:add_time"`     //优惠券添加时间

}

//修改默认表名
func (StoreCouponIssue) TableName() string {
	return "eb_store_coupon_issue"
}

func (storeCouponIssue *StoreCouponIssue) Insert() error {
	err := common.GetDB().Create(storeCouponIssue).Error
	return err
}

func (storeCouponIssue *StoreCouponIssue) Patch() error {
	err := common.GetDB().Model(storeCouponIssue).Updates(storeCouponIssue).Error
	return err
}

func (storeCouponIssue *StoreCouponIssue) Update() error {
	err := common.GetDB().Save(storeCouponIssue).Error
	return err
}

func (storeCouponIssue *StoreCouponIssue) Delete() error {
	return common.GetDB().Delete(storeCouponIssue).Error
}

func (storeCouponIssue *StoreCouponIssue) List(rawQuery string, rawOrder string, offset int, limit int) (*[]StoreCouponIssue, int, error) {
	storeCouponIssues := []StoreCouponIssue{}
	total := 0
	db := common.GetDB().Model(storeCouponIssue)
	db, err := buildWhere(rawQuery, db)
	if err != nil {
		return &storeCouponIssues, total, err
	}
	db, err = buildOrder(rawOrder, db)
	if err != nil {
		return &storeCouponIssues, total, err
	}
	db.Offset(offset).
		Limit(limit).
		Find(&storeCouponIssues).
		Count(&total)
	err = db.Error
	return &storeCouponIssues, total, err
}

func (storeCouponIssue *StoreCouponIssue) Get() (*StoreCouponIssue, error) {
	err := common.GetDB().Find(&storeCouponIssue).Error
	return storeCouponIssue, err
}
