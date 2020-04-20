//generate by gen
package models

import (
	"goshop/restful/common"

	"time"
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

type IssueCoupon struct {
	Id          int     `gorm:"column:id" json:"id"`                       //
	Cid         int     `gorm:"column:cid" json:"cid"`                     //优惠券ID
	StartTime   int     `gorm:"column:start_time" json:"start_time"`       //优惠券领取开启时间
	EndTime     int     `gorm:"column:end_time" json:"end_time"`           //优惠券领取结束时间
	TotalCount  int     `gorm:"column:total_count" json:"total_count"`     //优惠券领取数量
	RemainCount int     `gorm:"column:remain_count" json:"remain_count"`   //优惠券剩余领取数量
	IsPermanent int     `gorm:"column:is_permanent" json:"is_permanent"`   //是否无限张数
	Status      int     `gorm:"column:status" json:"status"`               //1 正常 0 未开启 -1 已无效
	IsDel       int     `gorm:"column:is_del" json:"is_del"`               //
	AddTime     int     `gorm:"column:add_time" json:"add_time"`           //优惠券添加时间
	CouponPrice float64 `gorm:"column:coupon_price" json:"coupon_price"`   //兑换的优惠券面值
	UseMinPrice float64 `gorm:"column:use_min_price" json:"use_min_price"` //最低消费多少金额可用优惠券
	IsUse       int     `json:"is_use"`                                    // 是否可以使用
}

func (storeCouponIssue *StoreCouponIssue) GetIssueCouponList(uid int, req *Query) (r []*IssueCoupon, err error) {

	storeCoupon := &StoreCoupon{}
	db := common.GetDB()
	nowTime := time.Now().Unix()
	list := make([]*IssueCoupon, 0)
	err = db.Raw("SELECT A.*, B.coupon_price, B.use_min_price FROM "+storeCouponIssue.TableName()+" A LEFT JOIN "+
		storeCoupon.TableName()+" B ON A.cid = B.id WHERE A.status = ? AND ((A.start_time < ? AND A.end_time > ?) OR (A.start_time = ? AND A.end_time = ?))"+
		"AND A.is_del = ?", 1, nowTime, nowTime, 0, 0, 0).Scan(&list).Error
	if err != nil {
		return
	}

	ids := make([]int, 0)

	for _, item := range list {
		ids = append(ids, item.Id)
	}
	storeCouponIssueUser := &StoreCouponIssueUser{}

	issueCoupons, err := storeCouponIssueUser.GetByUidIssueCouponIds(uid, ids)
	if err != nil {
		return
	}
	mIssuCoupons := make(map[int]*StoreCouponIssueUser, 0)
	for _, item := range issueCoupons {
		mIssuCoupons[item.IssueCouponId] = item
	}

	for _, item := range list {
		if _, ok := mIssuCoupons[item.Id]; ok {
			item.IsUse = 1
		} else {
			if item.RemainCount <= 0 && item.IsPermanent != 0 {
				item.IsUse = 2
			}
		}
	}
	r = list
	return
}
