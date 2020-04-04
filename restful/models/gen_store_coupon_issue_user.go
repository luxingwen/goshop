//generate by gen
package models

import (
	"goshop/restful/common"
)

//优惠券前台用户领取记录表
type StoreCouponIssueUser struct {
	Uid           int `gorm:"column:uid"`             //领取优惠券用户ID
	IssueCouponId int `gorm:"column:issue_coupon_id"` //优惠券前台领取ID
	AddTime       int `gorm:"column:add_time"`        //领取时间

}

//修改默认表名
func (StoreCouponIssueUser) TableName() string {
	return "eb_store_coupon_issue_user"
}

func (storeCouponIssueUser *StoreCouponIssueUser) Insert() error {
	err := common.GetDB().Create(storeCouponIssueUser).Error
	return err
}

func (storeCouponIssueUser *StoreCouponIssueUser) Patch() error {
	err := common.GetDB().Model(storeCouponIssueUser).Updates(storeCouponIssueUser).Error
	return err
}

func (storeCouponIssueUser *StoreCouponIssueUser) Update() error {
	err := common.GetDB().Save(storeCouponIssueUser).Error
	return err
}

func (storeCouponIssueUser *StoreCouponIssueUser) Delete() error {
	return common.GetDB().Delete(storeCouponIssueUser).Error
}

func (storeCouponIssueUser *StoreCouponIssueUser) List(rawQuery string, rawOrder string, offset int, limit int) (*[]StoreCouponIssueUser, int, error) {
	storeCouponIssueUsers := []StoreCouponIssueUser{}
	total := 0
	db := common.GetDB().Model(storeCouponIssueUser)
	db, err := buildWhere(rawQuery, db)
	if err != nil {
		return &storeCouponIssueUsers, total, err
	}
	db, err = buildOrder(rawOrder, db)
	if err != nil {
		return &storeCouponIssueUsers, total, err
	}
	db.Offset(offset).
		Limit(limit).
		Find(&storeCouponIssueUsers).
		Count(&total)
	err = db.Error
	return &storeCouponIssueUsers, total, err
}

func (storeCouponIssueUser *StoreCouponIssueUser) Get() (*StoreCouponIssueUser, error) {
	err := common.GetDB().Find(&storeCouponIssueUser).Error
	return storeCouponIssueUser, err
}
