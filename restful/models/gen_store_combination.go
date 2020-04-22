//generate by gen
package models

import (
	"goshop/restful/common"

	"encoding/json"
	"time"
)

//拼团产品表
type StoreCombination struct {
	Id          int     `gorm:"column:id" json:"id"`                   //
	ProductId   int     `gorm:"column:product_id" json:"product_id"`   //商品id
	MerId       int     `gorm:"column:mer_id" json:"mer_id"`           //商户id
	Image       string  `gorm:"column:image" json:"image"`             //推荐图
	Images      string  `gorm:"column:images" json:"images"`           //轮播图
	Title       string  `gorm:"column:title" json:"title"`             //活动标题
	Attr        string  `gorm:"column:attr" json:"attr"`               //活动属性
	People      int     `gorm:"column:people" json:"people"`           //参团人数
	Info        string  `gorm:"column:info" json:"info"`               //简介
	Price       float64 `gorm:"column:price" json:"price"`             //价格
	Sort        int     `gorm:"column:sort" json:"sort"`               //排序
	Sales       int     `gorm:"column:sales" json:"sales"`             //销量
	Stock       int     `gorm:"column:stock" json:"stock"`             //库存
	AddTime     string  `gorm:"column:add_time" json:"add_time"`       //添加时间
	IsHost      int     `gorm:"column:is_host" json:"is_host"`         //推荐
	IsShow      int     `gorm:"column:is_show" json:"is_show"`         //产品状态
	IsDel       int     `gorm:"column:is_del" json:"is_del"`           //
	Combination int     `gorm:"column:combination" json:"combination"` //
	MerUse      int     `gorm:"column:mer_use" json:"mer_use"`         //商户是否可用1可用0不可用
	IsPostage   int     `gorm:"column:is_postage" json:"is_postage"`   //是否包邮1是0否
	Postage     float64 `gorm:"column:postage" json:"postage"`         //邮费
	Description string  `gorm:"column:description" json:"description"` //拼团内容
	StartTime   int     `gorm:"column:start_time" json:"start_time"`   //拼团开始时间
	StopTime    int     `gorm:"column:stop_time" json:"stop_time"`     //拼团结束时间
	Cost        int     `gorm:"column:cost" json:"cost"`               //拼图产品成本
	Browse      int     `gorm:"column:browse" json:"browse"`           //浏览量
	UnitName    string  `gorm:"column:unit_name" json:"unit_name"`     //单位名

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

type RStoreCombination struct {
	StoreCombination
	ProductPrice float64 `gorm:"column:product_price" json:"product_price"`
}

type ResStoreCombination struct {
	Id           int      `gorm:"column:id" json:"id"`                   //
	ProductId    int      `gorm:"column:product_id" json:"product_id"`   //商品id
	MerId        int      `gorm:"column:mer_id" json:"mer_id"`           //商户id
	Image        string   `gorm:"column:image" json:"image"`             //推荐图
	Images       []string `gorm:"column:images" json:"images"`           //轮播图
	Title        string   `gorm:"column:title" json:"title"`             //活动标题
	Attr         string   `gorm:"column:attr" json:"attr"`               //活动属性
	People       int      `gorm:"column:people" json:"people"`           //参团人数
	Info         string   `gorm:"column:info" json:"info"`               //简介
	Price        float64  `gorm:"column:price" json:"price"`             //价格
	Sort         int      `gorm:"column:sort" json:"sort"`               //排序
	Sales        int      `gorm:"column:sales" json:"sales"`             //销量
	Stock        int      `gorm:"column:stock" json:"stock"`             //库存
	AddTime      string   `gorm:"column:add_time" json:"add_time"`       //添加时间
	IsHost       int      `gorm:"column:is_host" json:"is_host"`         //推荐
	IsShow       int      `gorm:"column:is_show" json:"is_show"`         //产品状态
	IsDel        int      `gorm:"column:is_del" json:"is_del"`           //
	Combination  int      `gorm:"column:combination" json:"combination"` //
	MerUse       int      `gorm:"column:mer_use" json:"mer_use"`         //商户是否可用1可用0不可用
	IsPostage    int      `gorm:"column:is_postage" json:"is_postage"`   //是否包邮1是0否
	Postage      float64  `gorm:"column:postage" json:"postage"`         //邮费
	Description  string   `gorm:"column:description" json:"description"` //拼团内容
	StartTime    int      `gorm:"column:start_time" json:"start_time"`   //拼团开始时间
	StopTime     int      `gorm:"column:stop_time" json:"stop_time"`     //拼团结束时间
	Cost         int      `gorm:"column:cost" json:"cost"`               //拼图产品成本
	Browse       int      `gorm:"column:browse" json:"browse"`           //浏览量
	UnitName     string   `gorm:"column:unit_name" json:"unit_name"`     //单位名
	ProductPrice float64  `gorm:"column:product_price" json:"product_price"`
}

//获取所有拼团数据
func (storeCombination *StoreCombination) GetAll(req *Query) (r []*ResStoreCombination, err error) {
	storeProduct := &StoreProduct{}
	nowTime := time.Now().Unix()
	db := common.GetDB()
	limit := 10
	page := 0

	if req.Page > 0 {
		page = req.Page - 1
	}
	if req.PageNum > 0 {
		limit = req.PageNum
	}

	offset := limit * page

	list := make([]*RStoreCombination, 0)

	err = db.Raw("SELECT c.*, s.price as product_price FROM "+storeCombination.TableName()+" c LEFET JOIN "+storeProduct.TableName()+" s ON s.id = c.product_id"+
		"WHERE c.is_show = ? AND c.is_del = ? AND c.start_time < ? AND c.stop_time > ? ORDER BY c.sort desc,c.id desc LIMIT ?, ?", 1, 0, nowTime, nowTime, offset, limit).Scan(&list).Error

	if err != nil {
		return
	}

	r = make([]*ResStoreCombination, 0)
	for _, item := range list {
		itemData := &ResStoreCombination{
			Id:           item.Id,
			ProductId:    item.ProductId,
			MerId:        item.MerId,
			Image:        item.Image,
			Title:        item.Title,
			Attr:         item.Attr,
			People:       item.People,
			Info:         item.Info,
			Price:        item.Price,
			Sort:         item.Sort,
			Sales:        item.Sales,
			Stock:        item.Stock,
			AddTime:      item.AddTime,
			IsHost:       item.IsHost,
			IsShow:       item.IsShow,
			IsDel:        item.IsDel,
			Combination:  item.Combination,
			MerUse:       item.MerUse,
			IsPostage:    item.IsPostage,
			Postage:      item.Postage,
			Description:  item.Description,
			StartTime:    item.StartTime,
			StopTime:     item.StopTime,
			Cost:         item.Cost,
			Browse:       item.Browse,
			UnitName:     item.UnitName,
			ProductPrice: item.ProductPrice,
		}

		err = json.Unmarshal([]byte(item.Images), itemData.Images)
		if err != nil {
			return r, err
		}
		r = append(r, itemData)
	}

	return
}

// 获取是否有拼团产品
func (storeCombination *StoreCombination) GetPinkIsOpen() (r bool, err error) {
	nowTime := time.Now().Unix()
	db := common.GetDB()
	var count int
	storeProduct := &StoreProduct{}
	rows, err := db.Raw("SELECT count(*) FROM "+storeCombination.TableName()+" c LEFT JOIN "+storeProduct.TableName()+" s ON "+
		"s.id=c.product_id WHERE c.is_show = ? AND c.is_del = ? AND c.start_time < ? AND c.stop_time > ?", 1, 0, nowTime, nowTime).Rows()
	if err != nil {
		return
	}
	for rows.Next() {
		rows.Scan(&count)
	}
	if count > 0 {
		return true, nil
	}
	return
}
