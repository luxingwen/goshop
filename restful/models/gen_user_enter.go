//generate by gen
package models

import (
	"goshop/restful/common"
)

//商户申请表
type UserEnter struct {
	Id           int    `gorm:"column:id"`            //商户申请ID
	Uid          int    `gorm:"column:uid"`           //用户ID
	Province     string `gorm:"column:province"`      //商户所在省
	City         string `gorm:"column:city"`          //商户所在市
	District     string `gorm:"column:district"`      //商户所在区
	Address      string `gorm:"column:address"`       //商户详细地址
	MerchantName string `gorm:"column:merchant_name"` //商户名称
	LinkUser     string `gorm:"column:link_user"`     //
	LinkTel      string `gorm:"column:link_tel"`      //商户电话
	Charter      string `gorm:"column:charter"`       //商户证书
	AddTime      int    `gorm:"column:add_time"`      //添加时间
	ApplyTime    int    `gorm:"column:apply_time"`    //审核时间
	SuccessTime  int    `gorm:"column:success_time"`  //通过时间
	FailMessage  string `gorm:"column:fail_message"`  //未通过原因
	FailTime     int    `gorm:"column:fail_time"`     //未通过时间
	Status       int    `gorm:"column:status"`        //-1 审核未通过 0未审核 1审核通过
	IsLock       int    `gorm:"column:is_lock"`       //0 = 开启 1= 关闭
	IsDel        int    `gorm:"column:is_del"`        //是否删除

}

//修改默认表名
func (UserEnter) TableName() string {
	return "eb_user_enter"
}

func (userEnter *UserEnter) Insert() error {
	err := common.GetDB().Create(userEnter).Error
	return err
}

func (userEnter *UserEnter) Patch() error {
	err := common.GetDB().Model(userEnter).Updates(userEnter).Error
	return err
}

func (userEnter *UserEnter) Update() error {
	err := common.GetDB().Save(userEnter).Error
	return err
}

func (userEnter *UserEnter) Delete() error {
	return common.GetDB().Delete(userEnter).Error
}

func (userEnter *UserEnter) List(rawQuery string, rawOrder string, offset int, limit int) (*[]UserEnter, int, error) {
	userEnters := []UserEnter{}
	total := 0
	db := common.GetDB().Model(userEnter)
	db, err := buildWhere(rawQuery, db)
	if err != nil {
		return &userEnters, total, err
	}
	db, err = buildOrder(rawOrder, db)
	if err != nil {
		return &userEnters, total, err
	}
	db.Offset(offset).
		Limit(limit).
		Find(&userEnters).
		Count(&total)
	err = db.Error
	return &userEnters, total, err
}

func (userEnter *UserEnter) Get() (*UserEnter, error) {
	err := common.GetDB().Find(&userEnter).Error
	return userEnter, err
}
