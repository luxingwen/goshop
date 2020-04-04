//generate by gen
package models

import (
	"goshop/restful/common"
)

//优惠券表
type StoreCoupon struct {
	Id          int     `gorm:"column:id"`            //优惠券表ID
	Title       string  `gorm:"column:title"`         //优惠券名称
	Integral    int     `gorm:"column:integral"`      //兑换消耗积分值
	CouponPrice float64 `gorm:"column:coupon_price"`  //兑换的优惠券面值
	UseMinPrice float64 `gorm:"column:use_min_price"` //最低消费多少金额可用优惠券
	CouponTime  int     `gorm:"column:coupon_time"`   //优惠券有效期限（单位：天）
	Sort        int     `gorm:"column:sort"`          //排序
	Status      int     `gorm:"column:status"`        //状态（0：关闭，1：开启）
	AddTime     int     `gorm:"column:add_time"`      //兑换项目添加时间
	IsDel       int     `gorm:"column:is_del"`        //是否删除

}

//修改默认表名
func (StoreCoupon) TableName() string {
	return "eb_store_coupon"
}

func (storeCoupon *StoreCoupon) Insert() error {
	err := common.GetDB().Create(storeCoupon).Error
	return err
}

func (storeCoupon *StoreCoupon) Patch() error {
	err := common.GetDB().Model(storeCoupon).Updates(storeCoupon).Error
	return err
}

func (storeCoupon *StoreCoupon) Update() error {
	err := common.GetDB().Save(storeCoupon).Error
	return err
}

func (storeCoupon *StoreCoupon) Delete() error {
	return common.GetDB().Delete(storeCoupon).Error
}

func (storeCoupon *StoreCoupon) List(rawQuery string, rawOrder string, offset int, limit int) (*[]StoreCoupon, int, error) {
	storeCoupons := []StoreCoupon{}
	total := 0
	db := common.GetDB().Model(storeCoupon)
	db, err := buildWhere(rawQuery, db)
	if err != nil {
		return &storeCoupons, total, err
	}
	db, err = buildOrder(rawOrder, db)
	if err != nil {
		return &storeCoupons, total, err
	}
	db.Offset(offset).
		Limit(limit).
		Find(&storeCoupons).
		Count(&total)
	err = db.Error
	return &storeCoupons, total, err
}

func (storeCoupon *StoreCoupon) Get() (*StoreCoupon, error) {
	err := common.GetDB().Find(&storeCoupon).Error
	return storeCoupon, err
}
