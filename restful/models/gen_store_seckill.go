//generate by gen
package models

import (
	"goshop/restful/common"
)

//商品秒杀产品表
type StoreSeckill struct {
	Id           int     `gorm:"column:id"`            //商品秒杀产品表id
	ProductId    int     `gorm:"column:product_id"`    //商品id
	Image        string  `gorm:"column:image"`         //推荐图
	Images       string  `gorm:"column:images"`        //轮播图
	Title        string  `gorm:"column:title"`         //活动标题
	Info         string  `gorm:"column:info"`          //简介
	Price        float64 `gorm:"column:price"`         //价格
	Cost         float64 `gorm:"column:cost"`          //成本
	OtPrice      float64 `gorm:"column:ot_price"`      //原价
	GiveIntegral float64 `gorm:"column:give_integral"` //返多少积分
	Sort         int     `gorm:"column:sort"`          //排序
	Stock        int     `gorm:"column:stock"`         //库存
	Sales        int     `gorm:"column:sales"`         //销量
	UnitName     string  `gorm:"column:unit_name"`     //单位名
	Postage      float64 `gorm:"column:postage"`       //邮费
	Description  string  `gorm:"column:description"`   //内容
	StartTime    string  `gorm:"column:start_time"`    //开始时间
	StopTime     string  `gorm:"column:stop_time"`     //结束时间
	AddTime      string  `gorm:"column:add_time"`      //添加时间
	Status       int     `gorm:"column:status"`        //产品状态
	IsPostage    int     `gorm:"column:is_postage"`    //是否包邮
	IsHot        int     `gorm:"column:is_hot"`        //热门推荐
	IsDel        int     `gorm:"column:is_del"`        //删除 0未删除1已删除
	Num          int     `gorm:"column:num"`           //最多秒杀几个
	IsShow       int     `gorm:"column:is_show"`       //显示

}

//修改默认表名
func (StoreSeckill) TableName() string {
	return "eb_store_seckill"
}

func (storeSeckill *StoreSeckill) Insert() error {
	err := common.GetDB().Create(storeSeckill).Error
	return err
}

func (storeSeckill *StoreSeckill) Patch() error {
	err := common.GetDB().Model(storeSeckill).Updates(storeSeckill).Error
	return err
}

func (storeSeckill *StoreSeckill) Update() error {
	err := common.GetDB().Save(storeSeckill).Error
	return err
}

func (storeSeckill *StoreSeckill) Delete() error {
	return common.GetDB().Delete(storeSeckill).Error
}

func (storeSeckill *StoreSeckill) List(rawQuery string, rawOrder string, offset int, limit int) (*[]StoreSeckill, int, error) {
	storeSeckills := []StoreSeckill{}
	total := 0
	db := common.GetDB().Model(storeSeckill)
	db, err := buildWhere(rawQuery, db)
	if err != nil {
		return &storeSeckills, total, err
	}
	db, err = buildOrder(rawOrder, db)
	if err != nil {
		return &storeSeckills, total, err
	}
	db.Offset(offset).
		Limit(limit).
		Find(&storeSeckills).
		Count(&total)
	err = db.Error
	return &storeSeckills, total, err
}

func (storeSeckill *StoreSeckill) Get() (*StoreSeckill, error) {
	err := common.GetDB().Find(&storeSeckill).Error
	return storeSeckill, err
}
