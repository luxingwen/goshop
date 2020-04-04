//generate by gen
package models

import (
	"goshop/restful/common"
)

//拼团表
type StorePink struct {
	Id         int     `gorm:"column:id"`           //
	Uid        int     `gorm:"column:uid"`          //用户id
	OrderId    string  `gorm:"column:order_id"`     //订单id 生成
	OrderIdKey int     `gorm:"column:order_id_key"` //订单id  数据库
	TotalNum   int     `gorm:"column:total_num"`    //购买商品个数
	TotalPrice float64 `gorm:"column:total_price"`  //购买总金额
	Cid        int     `gorm:"column:cid"`          //拼团产品id
	Pid        int     `gorm:"column:pid"`          //产品id
	People     int     `gorm:"column:people"`       //拼图总人数
	Price      float64 `gorm:"column:price"`        //拼团产品单价
	AddTime    string  `gorm:"column:add_time"`     //开始时间
	StopTime   string  `gorm:"column:stop_time"`    //
	KId        int     `gorm:"column:k_id"`         //团长id 0为团长
	IsTpl      int     `gorm:"column:is_tpl"`       //是否发送模板消息0未发送1已发送
	IsRefund   int     `gorm:"column:is_refund"`    //是否退款 0未退款 1已退款
	Status     int     `gorm:"column:status"`       //状态1进行中2已完成3未完成

}

//修改默认表名
func (StorePink) TableName() string {
	return "eb_store_pink"
}

func (storePink *StorePink) Insert() error {
	err := common.GetDB().Create(storePink).Error
	return err
}

func (storePink *StorePink) Patch() error {
	err := common.GetDB().Model(storePink).Updates(storePink).Error
	return err
}

func (storePink *StorePink) Update() error {
	err := common.GetDB().Save(storePink).Error
	return err
}

func (storePink *StorePink) Delete() error {
	return common.GetDB().Delete(storePink).Error
}

func (storePink *StorePink) List(rawQuery string, rawOrder string, offset int, limit int) (*[]StorePink, int, error) {
	storePinks := []StorePink{}
	total := 0
	db := common.GetDB().Model(storePink)
	db, err := buildWhere(rawQuery, db)
	if err != nil {
		return &storePinks, total, err
	}
	db, err = buildOrder(rawOrder, db)
	if err != nil {
		return &storePinks, total, err
	}
	db.Offset(offset).
		Limit(limit).
		Find(&storePinks).
		Count(&total)
	err = db.Error
	return &storePinks, total, err
}

func (storePink *StorePink) Get() (*StorePink, error) {
	err := common.GetDB().Find(&storePink).Error
	return storePink, err
}
