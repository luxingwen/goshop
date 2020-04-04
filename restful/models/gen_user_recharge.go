//generate by gen
package models

import (
	"goshop/restful/common"
)

//用户充值表
type UserRecharge struct {
	Id           int     `gorm:"column:id"`            //
	Uid          int     `gorm:"column:uid"`           //充值用户UID
	OrderId      string  `gorm:"column:order_id"`      //订单号
	Price        float64 `gorm:"column:price"`         //充值金额
	RechargeType string  `gorm:"column:recharge_type"` //充值类型
	Paid         int     `gorm:"column:paid"`          //是否充值
	PayTime      int     `gorm:"column:pay_time"`      //充值支付时间
	AddTime      int     `gorm:"column:add_time"`      //充值时间
	RefundPrice  float64 `gorm:"column:refund_price"`  //退款金额

}

//修改默认表名
func (UserRecharge) TableName() string {
	return "eb_user_recharge"
}

func (userRecharge *UserRecharge) Insert() error {
	err := common.GetDB().Create(userRecharge).Error
	return err
}

func (userRecharge *UserRecharge) Patch() error {
	err := common.GetDB().Model(userRecharge).Updates(userRecharge).Error
	return err
}

func (userRecharge *UserRecharge) Update() error {
	err := common.GetDB().Save(userRecharge).Error
	return err
}

func (userRecharge *UserRecharge) Delete() error {
	return common.GetDB().Delete(userRecharge).Error
}

func (userRecharge *UserRecharge) List(rawQuery string, rawOrder string, offset int, limit int) (*[]UserRecharge, int, error) {
	userRecharges := []UserRecharge{}
	total := 0
	db := common.GetDB().Model(userRecharge)
	db, err := buildWhere(rawQuery, db)
	if err != nil {
		return &userRecharges, total, err
	}
	db, err = buildOrder(rawOrder, db)
	if err != nil {
		return &userRecharges, total, err
	}
	db.Offset(offset).
		Limit(limit).
		Find(&userRecharges).
		Count(&total)
	err = db.Error
	return &userRecharges, total, err
}

func (userRecharge *UserRecharge) Get() (*UserRecharge, error) {
	err := common.GetDB().Find(&userRecharge).Error
	return userRecharge, err
}
