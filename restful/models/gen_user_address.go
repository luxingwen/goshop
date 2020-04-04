//generate by gen
package models

import (
	"goshop/restful/common"
)

//用户地址表
type UserAddress struct {
	Id        int    `gorm:"column:id"`         //用户地址id
	Uid       int    `gorm:"column:uid"`        //用户id
	RealName  string `gorm:"column:real_name"`  //收货人姓名
	Phone     string `gorm:"column:phone"`      //收货人电话
	Province  string `gorm:"column:province"`   //收货人所在省
	City      string `gorm:"column:city"`       //收货人所在市
	District  string `gorm:"column:district"`   //收货人所在区
	Detail    string `gorm:"column:detail"`     //收货人详细地址
	PostCode  int    `gorm:"column:post_code"`  //邮编
	Longitude string `gorm:"column:longitude"`  //经度
	Latitude  string `gorm:"column:latitude"`   //纬度
	IsDefault int    `gorm:"column:is_default"` //是否默认
	IsDel     int    `gorm:"column:is_del"`     //是否删除
	AddTime   int    `gorm:"column:add_time"`   //添加时间

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
