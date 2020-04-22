//generate by gen
package models

import (
	"goshop/restful/common"

	"time"
)

//优惠券发放记录表
type StoreCouponUser struct {
	Id          int     `gorm:"column:id"`            //优惠券发放记录id
	Cid         int     `gorm:"column:cid"`           //兑换的项目id
	Uid         int     `gorm:"column:uid"`           //优惠券所属用户
	CouponTitle string  `gorm:"column:coupon_title"`  //优惠券名称
	CouponPrice float64 `gorm:"column:coupon_price"`  //优惠券的面值
	UseMinPrice float64 `gorm:"column:use_min_price"` //最低消费多少金额可用优惠券
	AddTime     int     `gorm:"column:add_time"`      //优惠券创建时间
	EndTime     int     `gorm:"column:end_time"`      //优惠券结束时间
	UseTime     int     `gorm:"column:use_time"`      //使用时间
	Type        string  `gorm:"column:type"`          //获取方式
	Status      int     `gorm:"column:status"`        //状态（0：未使用，1：已使用, 2:已过期）
	IsFail      int     `gorm:"column:is_fail"`       //是否有效

}

//修改默认表名
func (StoreCouponUser) TableName() string {
	return "eb_store_coupon_user"
}

func (storeCouponUser *StoreCouponUser) Insert() error {
	err := common.GetDB().Create(storeCouponUser).Error
	return err
}

func (storeCouponUser *StoreCouponUser) Patch() error {
	err := common.GetDB().Model(storeCouponUser).Updates(storeCouponUser).Error
	return err
}

func (storeCouponUser *StoreCouponUser) Update() error {
	err := common.GetDB().Save(storeCouponUser).Error
	return err
}

func (storeCouponUser *StoreCouponUser) Delete() error {
	return common.GetDB().Delete(storeCouponUser).Error
}

func (storeCouponUser *StoreCouponUser) List(rawQuery string, rawOrder string, offset int, limit int) (*[]StoreCouponUser, int, error) {
	storeCouponUsers := []StoreCouponUser{}
	total := 0
	db := common.GetDB().Model(storeCouponUser)
	db, err := buildWhere(rawQuery, db)
	if err != nil {
		return &storeCouponUsers, total, err
	}
	db, err = buildOrder(rawOrder, db)
	if err != nil {
		return &storeCouponUsers, total, err
	}
	db.Offset(offset).
		Limit(limit).
		Find(&storeCouponUsers).
		Count(&total)
	err = db.Error
	return &storeCouponUsers, total, err
}

func (storeCouponUser *StoreCouponUser) Get() (*StoreCouponUser, error) {
	err := common.GetDB().Find(&storeCouponUser).Error
	return storeCouponUser, err
}

type MyStoreCoupon struct {
	Id          int     `gorm:"column:id" json:"id"`                     //优惠券发放记录id
	Cid         int     `gorm:"column:cid" json:"cid"`                   //兑换的项目id
	Uid         int     `gorm:"column:uid" json:"uid"`                   //优惠券所属用户
	CouponTitle string  `gorm:"column:coupon_title" json:"coupon_time"`  //优惠券名称
	CouponPrice float64 `gorm:"column:coupon_price" json:"coupon_price"` //优惠券的面值
	AddTime     int     `gorm:"column:add_time" json:"add_time"`         //优惠券创建时间
	EndTime     int     `gorm:"column:end_time" json:"end_time"`         //优惠券结束时间
	Msg         string  `json:"_msg"`
	Type        int     `json:"_type"`
}

func (storeCouponUser *StoreCouponUser) CheckInvalidCoupon() (err error) {
	db := common.GetDB()
	return db.Table(storeCouponUser.TableName()).Update("status", 2).Where("end_time < ? AND status = ?", time.Now().Unix(), 0).Error
}

func (storeCouponUser *StoreCouponUser) MyStoreCoupon(uid int) (r []*MyStoreCoupon, err error) {

	db := common.GetDB()

	err = db.Table(storeCouponUser.TableName()).Where("uid = ? AND status = ?", uid, 0).Order("is_fail ASC, status ASC, add_time DESC").Find(&r).Error
	return

}

// 获取用户可用优惠券数量
func (storeCouponUser *StoreCouponUser) GetUserValidCouponCount(uid int) (count int, err error) {
	err = storeCouponUser.CheckInvalidCoupon()
	if err != nil {
		return
	}
	db := common.GetDB()

	err = db.Table(storeCouponUser.TableName()).Where("uid = ? AND status = ?", uid, 0).Count(&count).Error
	return
}

//获取用户优惠券（全部）
func (storeCouponUser *StoreCouponUser) GetUserAllCoupon(uid int) (r []*MyStoreCoupon, err error) {
	storeCouponUser.CheckInvalidCoupon()

	db := common.GetDB()
	list := make([]*StoreCouponUser, 0)
	err = db.Table(storeCouponUser.TableName()).Where("uid = ?", uid).Order("is_fail ASC,status ASC,add_time DESC").Find(&list).Error
	if err != nil {
		return
	}
	return storeCouponUser.tidyCouponList(list), nil
}

// 获取用户优惠券（未使用）
func (storeCouponUser *StoreCouponUser) GetUserValidCoupon(uid int) (r []*MyStoreCoupon, err error) {
	storeCouponUser.CheckInvalidCoupon()

	db := common.GetDB()
	list := make([]*StoreCouponUser, 0)
	err = db.Table(storeCouponUser.TableName()).Where("uid = ? AND status = ?", uid, 0).Order("is_fail ASC,status ASC,add_time DESC").Find(&list).Error
	if err != nil {
		return
	}
	return storeCouponUser.tidyCouponList(list), nil
}

// 获取用户优惠券（已经使用)
func (storeCouponUser *StoreCouponUser) GetUserAlreadyUsedCoupon(uid int) (r []*MyStoreCoupon, err error) {
	storeCouponUser.CheckInvalidCoupon()

	db := common.GetDB()
	list := make([]*StoreCouponUser, 0)
	err = db.Table(storeCouponUser.TableName()).Where("uid = ? AND status = ?", uid, 1).Order("is_fail ASC,status ASC,add_time DESC").Find(&list).Error
	if err != nil {
		return
	}
	return storeCouponUser.tidyCouponList(list), nil
}

// 获取用户优惠券（已过期）
func (storeCouponUser *StoreCouponUser) GetUserBeOverdueCoupon(uid int) (r []*MyStoreCoupon, err error) {
	storeCouponUser.CheckInvalidCoupon()

	db := common.GetDB()
	list := make([]*StoreCouponUser, 0)
	err = db.Table(storeCouponUser.TableName()).Where("uid = ? AND status = ?", uid, 2).Order("is_fail ASC,status ASC,add_time DESC").Find(&list).Error
	if err != nil {
		return
	}
	return storeCouponUser.tidyCouponList(list), nil
}

//
func (storeCouponUser *StoreCouponUser) tidyCouponList(list []*StoreCouponUser) (r []*MyStoreCoupon) {
	nowTime := int(time.Now().Unix())
	for _, item := range list {
		itemData := &MyStoreCoupon{
			Id:          item.Id,
			Cid:         item.Cid,
			Uid:         item.Uid,
			CouponTitle: item.CouponTitle,
			CouponPrice: item.CouponPrice,
			AddTime:     item.AddTime,
			EndTime:     item.EndTime,
		}
		if item.IsFail > 0 {
			itemData.Type = 0
			itemData.Msg = "已失效"
		} else if item.Status == 1 {
			itemData.Type = 0
			itemData.Msg = "已使用"
		} else if item.Status == 2 {
			itemData.Type = 0
			itemData.Msg = "已过期"
		} else if item.AddTime > nowTime || item.EndTime < nowTime {
			itemData.Type = 0
			itemData.Msg = "已过期"
		} else {
			if item.AddTime+86400 > nowTime {
				itemData.Type = 2
				itemData.Msg = "可使用"
			} else {
				itemData.Type = 2
				itemData.Msg = "可使用"
			}
		}
		r = append(r, itemData)

	}
	return
}
