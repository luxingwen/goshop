//generate by gen
package models

import (
	"goshop/restful/common"

	"time"
)

//砍价表
type StoreBargain struct {
	Id              int     `gorm:"column:id"`                //砍价产品ID
	ProductId       int     `gorm:"column:product_id"`        //关联产品ID
	Title           string  `gorm:"column:title"`             //砍价活动名称
	Image           string  `gorm:"column:image"`             //砍价活动图片
	UnitName        string  `gorm:"column:unit_name"`         //单位名称
	Stock           int     `gorm:"column:stock"`             //库存
	Sales           int     `gorm:"column:sales"`             //销量
	Images          string  `gorm:"column:images"`            //砍价产品轮播图
	StartTime       int     `gorm:"column:start_time"`        //砍价开启时间
	StopTime        int     `gorm:"column:stop_time"`         //砍价结束时间
	StoreName       string  `gorm:"column:store_name"`        //砍价产品名称
	Price           float64 `gorm:"column:price"`             //砍价金额
	MinPrice        float64 `gorm:"column:min_price"`         //砍价商品最低价
	Num             int     `gorm:"column:num"`               //每次购买的砍价产品数量
	BargainMaxPrice float64 `gorm:"column:bargain_max_price"` //用户每次砍价的最大金额
	BargainMinPrice float64 `gorm:"column:bargain_min_price"` //用户每次砍价的最小金额
	BargainNum      int     `gorm:"column:bargain_num"`       //用户每次砍价的次数
	Status          int     `gorm:"column:status"`            //砍价状态 0(到砍价时间不自动开启)  1(到砍价时间自动开启时间)
	Description     string  `gorm:"column:description"`       //砍价详情
	GiveIntegral    float64 `gorm:"column:give_integral"`     //反多少积分
	Info            string  `gorm:"column:info"`              //砍价活动简介
	Cost            float64 `gorm:"column:cost"`              //成本价
	Sort            int     `gorm:"column:sort"`              //排序
	IsHot           int     `gorm:"column:is_hot"`            //是否推荐0不推荐1推荐
	IsDel           int     `gorm:"column:is_del"`            //是否删除 0未删除 1删除
	AddTime         int     `gorm:"column:add_time"`          //添加时间
	IsPostage       int     `gorm:"column:is_postage"`        //是否包邮 0不包邮 1包邮
	Postage         float64 `gorm:"column:postage"`           //邮费
	Rule            string  `gorm:"column:rule"`              //砍价规则
	Look            int     `gorm:"column:look"`              //砍价产品浏览量
	Share           int     `gorm:"column:share"`             //砍价产品分享量

}

//修改默认表名
func (StoreBargain) TableName() string {
	return "eb_store_bargain"
}

func (storeBargain *StoreBargain) Insert() error {
	err := common.GetDB().Create(storeBargain).Error
	return err
}

func (storeBargain *StoreBargain) Patch() error {
	err := common.GetDB().Model(storeBargain).Updates(storeBargain).Error
	return err
}

func (storeBargain *StoreBargain) Update() error {
	err := common.GetDB().Save(storeBargain).Error
	return err
}

func (storeBargain *StoreBargain) Delete() error {
	return common.GetDB().Delete(storeBargain).Error
}

func (storeBargain *StoreBargain) List(rawQuery string, rawOrder string, offset int, limit int) (*[]StoreBargain, int, error) {
	storeBargains := []StoreBargain{}
	total := 0
	db := common.GetDB().Model(storeBargain)
	db, err := buildWhere(rawQuery, db)
	if err != nil {
		return &storeBargains, total, err
	}
	db, err = buildOrder(rawOrder, db)
	if err != nil {
		return &storeBargains, total, err
	}
	db.Offset(offset).
		Limit(limit).
		Find(&storeBargains).
		Count(&total)
	err = db.Error
	return &storeBargains, total, err
}

func (storeBargain *StoreBargain) Get() (*StoreBargain, error) {
	err := common.GetDB().Find(&storeBargain).Error
	return storeBargain, err
}

type ResStoreBargain struct {
	Id          int     `gorm:"column:id" json:"id"`                 //砍价产品ID
	ProductId   int     `gorm:"column:product_id" json:"product_id"` //关联产品ID
	Title       string  `gorm:"column:title" json:"title"`           //砍价活动名称
	Image       string  `gorm:"column:image" json:"image"`           //砍价活动图片
	Stock       int     `gorm:"column:stock" json:"stock"`           //库存
	StoreName   string  `gorm:"column:store_name" json:"store_name"` //砍价产品名称
	Price       float64 `gorm:"column:price" json:"price"`           //砍价金额
	MinPrice    float64 `gorm:"column:min_price" json:"min_price"`   //砍价商品最低价
	PeopleCount int     `json:"people"`
}

func (storeBargain *StoreBargain) GetList(req *Query) (r []*ResStoreBargain, err error) {
	nowTime := time.Now().Unix()
	db := common.GetDB()

	err = db.Select("id,product_id,title,price,min_price,image").Where("is_del = ? AND status = ？ AND start_time < ? AND stop_time > ?", 0, 1, nowTime, nowTime).Scan(&r).Error
	if err != nil {
		return
	}
	// ids := make([]int, 0)
	// for _, item := range r {
	// 	ids = append(ids, item.Id)
	// }
	return

}
