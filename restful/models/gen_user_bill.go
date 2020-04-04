//generate by gen
package models

import (
	"goshop/restful/common"
)

//用户账单表
type UserBill struct {
	Id       int     `gorm:"column:id"`       //用户账单id
	Uid      int     `gorm:"column:uid"`      //用户uid
	LinkId   string  `gorm:"column:link_id"`  //关联id
	Pm       int     `gorm:"column:pm"`       //0 = 支出 1 = 获得
	Title    string  `gorm:"column:title"`    //账单标题
	Category string  `gorm:"column:category"` //明细种类
	Type     string  `gorm:"column:type"`     //明细类型
	Number   float64 `gorm:"column:number"`   //明细数字
	Balance  float64 `gorm:"column:balance"`  //剩余
	Mark     string  `gorm:"column:mark"`     //备注
	AddTime  int     `gorm:"column:add_time"` //添加时间
	Status   int     `gorm:"column:status"`   //0 = 带确定 1 = 有效 -1 = 无效

}

//修改默认表名
func (UserBill) TableName() string {
	return "eb_user_bill"
}

func (userBill *UserBill) Insert() error {
	err := common.GetDB().Create(userBill).Error
	return err
}

func (userBill *UserBill) Patch() error {
	err := common.GetDB().Model(userBill).Updates(userBill).Error
	return err
}

func (userBill *UserBill) Update() error {
	err := common.GetDB().Save(userBill).Error
	return err
}

func (userBill *UserBill) Delete() error {
	return common.GetDB().Delete(userBill).Error
}

func (userBill *UserBill) List(rawQuery string, rawOrder string, offset int, limit int) (*[]UserBill, int, error) {
	userBills := []UserBill{}
	total := 0
	db := common.GetDB().Model(userBill)
	db, err := buildWhere(rawQuery, db)
	if err != nil {
		return &userBills, total, err
	}
	db, err = buildOrder(rawOrder, db)
	if err != nil {
		return &userBills, total, err
	}
	db.Offset(offset).
		Limit(limit).
		Find(&userBills).
		Count(&total)
	err = db.Error
	return &userBills, total, err
}

func (userBill *UserBill) Get() (*UserBill, error) {
	err := common.GetDB().Find(&userBill).Error
	return userBill, err
}
