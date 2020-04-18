//generate by gen
package models

import (
	"goshop/restful/common"
)

//用户提现表
type UserExtract struct {
	Id           int     `gorm:"column:id"`            //
	Uid          int     `gorm:"column:uid"`           //
	RealName     string  `gorm:"column:real_name"`     //名称
	ExtractType  string  `gorm:"column:extract_type"`  //bank = 银行卡 alipay = 支付宝wx=微信
	BankCode     string  `gorm:"column:bank_code"`     //银行卡
	BankAddress  string  `gorm:"column:bank_address"`  //开户地址
	AlipayCode   string  `gorm:"column:alipay_code"`   //支付宝账号
	ExtractPrice float64 `gorm:"column:extract_price"` //提现金额
	Mark         string  `gorm:"column:mark"`          //
	Balance      float64 `gorm:"column:balance"`       //
	FailMsg      string  `gorm:"column:fail_msg"`      //无效原因
	FailTime     int     `gorm:"column:fail_time"`     //
	AddTime      int     `gorm:"column:add_time"`      //添加时间
	Status       int     `gorm:"column:status"`        //-1 未通过 0 审核中 1 已提现
	Wechat       string  `gorm:"column:wechat"`        //微信号

}

//修改默认表名
func (UserExtract) TableName() string {
	return "eb_user_extract"
}

func (userExtract *UserExtract) Insert() error {
	err := common.GetDB().Create(userExtract).Error
	return err
}

func (userExtract *UserExtract) Patch() error {
	err := common.GetDB().Model(userExtract).Updates(userExtract).Error
	return err
}

func (userExtract *UserExtract) Update() error {
	err := common.GetDB().Save(userExtract).Error
	return err
}

func (userExtract *UserExtract) Delete() error {
	return common.GetDB().Delete(userExtract).Error
}

func (userExtract *UserExtract) List(rawQuery string, rawOrder string, offset int, limit int) (*[]UserExtract, int, error) {
	userExtracts := []UserExtract{}
	total := 0
	db := common.GetDB().Model(userExtract)
	db, err := buildWhere(rawQuery, db)
	if err != nil {
		return &userExtracts, total, err
	}
	db, err = buildOrder(rawOrder, db)
	if err != nil {
		return &userExtracts, total, err
	}
	db.Offset(offset).
		Limit(limit).
		Find(&userExtracts).
		Count(&total)
	err = db.Error
	return &userExtracts, total, err
}

func (userExtract *UserExtract) Get() (*UserExtract, error) {
	err := common.GetDB().Find(&userExtract).Error
	return userExtract, err
}

// 累计提现
func (userExtract *UserExtract) UserExtractTotalPrice(uid int) (count float64, err error) {
	db := common.GetDB()
	rows, err := db.Table(userExtract.TableName()).Select("sum(extract_price)").Where("uid = ? AND status = ?", uid, 1).Rows()
	if err != nil {
		return
	}
	for rows.Next() {
		rows.Scan(&count)
	}
	return
}

// 累计提现
func (userExtract *UserExtract) UserExtractTotalPriceByStatus(uid int, status int) (count float64, err error) {
	db := common.GetDB()
	rows, err := db.Table(userExtract.TableName()).Select("sum(extract_price)").Where("uid = ? AND status = ?", uid, status).Rows()
	if err != nil {
		return
	}
	for rows.Next() {
		rows.Scan(&count)
	}
	return
}
