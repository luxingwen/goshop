//generate by gen
package models

import (
	"goshop/restful/common"
)

//拼团产品表
type StoreCombination struct {
	Id          int     `gorm:"column:id"`          //
	ProductId   int     `gorm:"column:product_id"`  //商品id
	MerId       int     `gorm:"column:mer_id"`      //商户id
	Image       string  `gorm:"column:image"`       //推荐图
	Images      string  `gorm:"column:images"`      //轮播图
	Title       string  `gorm:"column:title"`       //活动标题
	Attr        string  `gorm:"column:attr"`        //活动属性
	People      int     `gorm:"column:people"`      //参团人数
	Info        string  `gorm:"column:info"`        //简介
	Price       float64 `gorm:"column:price"`       //价格
	Sort        int     `gorm:"column:sort"`        //排序
	Sales       int     `gorm:"column:sales"`       //销量
	Stock       int     `gorm:"column:stock"`       //库存
	AddTime     string  `gorm:"column:add_time"`    //添加时间
	IsHost      int     `gorm:"column:is_host"`     //推荐
	IsShow      int     `gorm:"column:is_show"`     //产品状态
	IsDel       int     `gorm:"column:is_del"`      //
	Combination int     `gorm:"column:combination"` //
	MerUse      int     `gorm:"column:mer_use"`     //商户是否可用1可用0不可用
	IsPostage   int     `gorm:"column:is_postage"`  //是否包邮1是0否
	Postage     float64 `gorm:"column:postage"`     //邮费
	Description string  `gorm:"column:description"` //拼团内容
	StartTime   int     `gorm:"column:start_time"`  //拼团开始时间
	StopTime    int     `gorm:"column:stop_time"`   //拼团结束时间
	Cost        int     `gorm:"column:cost"`        //拼图产品成本
	Browse      int     `gorm:"column:browse"`      //浏览量
	UnitName    string  `gorm:"column:unit_name"`   //单位名

}

//修改默认表名
func (StoreCombination) TableName() string {
	return "eb_store_combination"
}

func (storeCombination *StoreCombination) Insert() error {
	err := common.GetDB().Create(storeCombination).Error
	return err
}

func (storeCombination *StoreCombination) Patch() error {
	err := common.GetDB().Model(storeCombination).Updates(storeCombination).Error
	return err
}

func (storeCombination *StoreCombination) Update() error {
	err := common.GetDB().Save(storeCombination).Error
	return err
}

func (storeCombination *StoreCombination) Delete() error {
	return common.GetDB().Delete(storeCombination).Error
}

func (storeCombination *StoreCombination) List(rawQuery string, rawOrder string, offset int, limit int) (*[]StoreCombination, int, error) {
	storeCombinations := []StoreCombination{}
	total := 0
	db := common.GetDB().Model(storeCombination)
	db, err := buildWhere(rawQuery, db)
	if err != nil {
		return &storeCombinations, total, err
	}
	db, err = buildOrder(rawOrder, db)
	if err != nil {
		return &storeCombinations, total, err
	}
	db.Offset(offset).
		Limit(limit).
		Find(&storeCombinations).
		Count(&total)
	err = db.Error
	return &storeCombinations, total, err
}

func (storeCombination *StoreCombination) Get() (*StoreCombination, error) {
	err := common.GetDB().Find(&storeCombination).Error
	return storeCombination, err
}
