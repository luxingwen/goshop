//generate by gen
package models

import (
	"goshop/restful/common"
)

//用户地址表
type UserAddress struct {
	Id        int    `gorm:"column:id" json:"id"`                 //用户地址id
	Uid       int    `gorm:"column:uid" json:"uid"`               //用户id
	RealName  string `gorm:"column:real_name" json:"real_name"`   //收货人姓名
	Phone     string `gorm:"column:phone" json:"phone"`           //收货人电话
	Province  string `gorm:"column:province" json:"province"`     //收货人所在省
	City      string `gorm:"column:city" json:"city"`             //收货人所在市
	District  string `gorm:"column:district" json:"district"`     //收货人所在区
	Detail    string `gorm:"column:detail" json:"detail"`         //收货人详细地址
	PostCode  int    `gorm:"column:post_code" json:"post_code"`   //邮编
	Longitude string `gorm:"column:longitude" json:"longitude"`   //经度
	Latitude  string `gorm:"column:latitude" json:"latitude"`     //纬度
	IsDefault int    `gorm:"column:is_default" json:"is_default"` //是否默认
	IsDel     int    `gorm:"column:is_del" json:"is_del"`         //是否删除
	AddTime   int    `gorm:"column:add_time" json:"add_time"`     //添加时间

}

//修改默认表名
func (UserAddress) TableName() string {
	return "eb_user_address"
}

func (userAddress *UserAddress) Insert() error {
	err := common.GetDB().Create(userAddress).Error
	return err
}

func (userAddress *UserAddress) Patch() error {
	err := common.GetDB().Model(userAddress).Updates(userAddress).Error
	return err
}

func (userAddress *UserAddress) Update() error {
	err := common.GetDB().Save(userAddress).Error
	return err
}

func (userAddress *UserAddress) Delete() error {
	return common.GetDB().Delete(userAddress).Error
}

func (userAddress *UserAddress) List(rawQuery string, rawOrder string, offset int, limit int) (*[]UserAddress, int, error) {
	userAddresss := []UserAddress{}
	total := 0
	db := common.GetDB().Model(userAddress)
	db, err := buildWhere(rawQuery, db)
	if err != nil {
		return &userAddresss, total, err
	}
	db, err = buildOrder(rawOrder, db)
	if err != nil {
		return &userAddresss, total, err
	}
	db.Offset(offset).
		Limit(limit).
		Find(&userAddresss).
		Count(&total)
	err = db.Error
	return &userAddresss, total, err
}

func (userAddress *UserAddress) Get() (*UserAddress, error) {
	err := common.GetDB().Find(&userAddress).Error
	return userAddress, err
}

func (userAddress *UserAddress) GetUserValidAddressList(uid int, req *Query) (r []*UserAddress, err error) {
	db := common.GetDB().Model(userAddress)
	limit := 10
	page := 0

	if req.Page > 0 {
		page = req.Page - 1
	}
	if req.PageNum > 0 {
		limit = req.PageNum
	}

	offset := limit * page

	err = db.Select("id,real_name,phone,province,city,district,detail,is_default").Where("uid = ?", uid).Order("add_time DESC").Offset(offset).Limit(limit).Find(&r).Error
	return
}
