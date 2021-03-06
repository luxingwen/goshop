//generate by gen
package models

import (
	"goshop/restful/common"
)

//用户表
type User struct {
	Uid         int     `gorm:"column:uid" json:"uid"`                   //用户id
	Account     string  `gorm:"column:account" json:"account"`           //用户账号
	Pwd         string  `gorm:"column:pwd" json:"pwd"`                   //用户密码
	RealName    string  `gorm:"column:real_name" json:"real_name"`       //真实姓名
	Birthday    int     `gorm:"column:birthday" json:"birthday"`         //生日
	CardId      string  `gorm:"column:card_id" json:"card_id"`           //身份证号码
	Mark        string  `gorm:"column:mark" json:"mark"`                 //用户备注
	GroupId     int     `gorm:"column:group_id" json:"group_id"`         //用户分组id
	PartnerId   int     `gorm:"column:partner_id" json:"partner_id"`     //合伙人id
	Addres      string  `gorm:"column:addres" json:"addres"`             //详细地址
	Nickname    string  `gorm:"column:nickname" json:"nickname"`         //用户昵称
	Avatar      string  `gorm:"column:avatar" json:"avatar"`             //用户头像
	Phone       string  `gorm:"column:phone" json:"phone"`               //手机号码
	AddTime     int     `gorm:"column:add_time" json:"add_time"`         //添加时间
	AddIp       string  `gorm:"column:add_ip" json:"add_ip"`             //添加ip
	LastTime    int     `gorm:"column:last_time" json:"last_time"`       //最后一次登录时间
	LastIp      string  `gorm:"column:last_ip" json:"last_ip"`           //最后一次登录ip
	NowMoney    float64 `gorm:"column:now_money" json:"now_money"`       //用户余额
	Integral    float64 `gorm:"column:integral" json:"integral"`         //用户剩余积分
	SignNum     int     `gorm:"column:sign_num" json:"sign_num"`         //连续签到天数
	Status      int     `gorm:"column:status" json:"status"`             //1为正常，0为禁止
	Level       int     `gorm:"column:level" json:"level"`               //等级
	SpreadUid   int     `gorm:"column:spread_uid" json:"spread_uid"`     //推广元id
	SpreadTime  int     `gorm:"column:spread_time" json:"spread_time"`   //推广员关联时间
	UserType    string  `gorm:"column:user_type" json:"user_type"`       //用户类型
	IsPromoter  int     `gorm:"column:is_promoter" json:"is_promoter"`   //是否为推广员
	PayCount    int     `gorm:"column:pay_count" json:"pay_count"`       //用户购买次数
	SpreadCount int     `gorm:"column:spread_count" json:"spread_count"` //下级人数
	CleanTime   int     `gorm:"column:clean_time" json:"clean_time"`     //等级清除时间为0没有清除过

}

//修改默认表名
func (User) TableName() string {
	return "eb_user"
}

func (user *User) Insert() error {
	err := common.GetDB().Create(user).Error
	return err
}

func (user *User) Patch() error {
	err := common.GetDB().Model(user).Updates(user).Error
	return err
}

func (user *User) Update() error {
	err := common.GetDB().Save(user).Error
	return err
}

func (user *User) Delete() error {
	return common.GetDB().Delete(user).Error
}

func (user *User) List(rawQuery string, rawOrder string, offset int, limit int) (*[]User, int, error) {
	users := []User{}
	total := 0
	db := common.GetDB().Model(user)
	db, err := buildWhere(rawQuery, db)
	if err != nil {
		return &users, total, err
	}
	db, err = buildOrder(rawOrder, db)
	if err != nil {
		return &users, total, err
	}
	db.Offset(offset).
		Limit(limit).
		Find(&users).
		Count(&total)
	err = db.Error
	return &users, total, err
}

func (user *User) Get() (*User, error) {
	err := common.GetDB().Find(&user).Error
	return user, err
}

func (user *User) GetByUid(uid int) (r *User, err error) {
	r = new(User)
	err = common.GetDB().Model(user).Where("uid = ?", uid).First(&r).Error
	return
}

func (user *User) UpdateByUid(uid int) (err error) {
	err = common.GetDB().Model(user).Where("uid = ?", uid).Update(user).Error
	return err
}
