//generate by gen
package models

import (
	"goshop/restful/common"

	"encoding/json"
	"strings"
)

//商品表
type StoreProduct struct {
	Id           int     `gorm:"column:id"`            //商品id
	MerId        int     `gorm:"column:mer_id"`        //商户Id(0为总后台管理员创建,不为0的时候是商户后台创建)
	Image        string  `gorm:"column:image"`         //商品图片
	SliderImage  string  `gorm:"column:slider_image"`  //轮播图
	StoreName    string  `gorm:"column:store_name"`    //商品名称
	StoreInfo    string  `gorm:"column:store_info"`    //商品简介
	Keyword      string  `gorm:"column:keyword"`       //关键字
	CateId       string  `gorm:"column:cate_id"`       //分类id
	Price        float64 `gorm:"column:price"`         //商品价格
	VipPrice     float64 `gorm:"column:vip_price"`     //会员价格
	OtPrice      float64 `gorm:"column:ot_price"`      //市场价
	Postage      float64 `gorm:"column:postage"`       //邮费
	UnitName     string  `gorm:"column:unit_name"`     //单位名
	Sort         int     `gorm:"column:sort"`          //排序
	Sales        int     `gorm:"column:sales"`         //销量
	Stock        int     `gorm:"column:stock"`         //库存
	IsShow       int     `gorm:"column:is_show"`       //状态（0：未上架，1：上架）
	IsHot        int     `gorm:"column:is_hot"`        //是否热卖
	IsBenefit    int     `gorm:"column:is_benefit"`    //是否优惠
	IsBest       int     `gorm:"column:is_best"`       //是否精品
	IsNew        int     `gorm:"column:is_new"`        //是否新品
	Description  string  `gorm:"column:description"`   //产品描述
	AddTime      int     `gorm:"column:add_time"`      //添加时间
	IsPostage    int     `gorm:"column:is_postage"`    //是否包邮
	IsDel        int     `gorm:"column:is_del"`        //是否删除
	MerUse       int     `gorm:"column:mer_use"`       //商户是否代理 0不可代理1可代理
	GiveIntegral float64 `gorm:"column:give_integral"` //获得积分
	Cost         float64 `gorm:"column:cost"`          //成本价
	IsSeckill    int     `gorm:"column:is_seckill"`    //秒杀状态 0 未开启 1已开启
	IsBargain    int     `gorm:"column:is_bargain"`    //砍价状态 0未开启 1开启
	Ficti        int     `gorm:"column:ficti"`         //虚拟销量
	Browse       int     `gorm:"column:browse"`        //浏览量
	CodePath     string  `gorm:"column:code_path"`     //产品二维码地址(用户小程序海报)
	StoreCode    string  `gorm:"column:store_code"`    //产品编码
	SoureLink    string  `gorm:"column:soure_link"`    //淘宝、天猫、1688商品保存标识，避免商品重复入库
	Brand        int     `gorm:"column:brand"`         //品牌id

}

//修改默认表名
func (StoreProduct) TableName() string {
	return "eb_store_product"
}

func (storeProduct *StoreProduct) Insert() error {
	err := common.GetDB().Create(storeProduct).Error
	return err
}

func (storeProduct *StoreProduct) Patch() error {
	err := common.GetDB().Model(storeProduct).Updates(storeProduct).Error
	return err
}

func (storeProduct *StoreProduct) Update() error {
	err := common.GetDB().Save(storeProduct).Error
	return err
}

func (storeProduct *StoreProduct) Delete() error {
	return common.GetDB().Delete(storeProduct).Error
}

func (storeProduct *StoreProduct) List(rawQuery string, rawOrder string, offset int, limit int) (*[]StoreProduct, int, error) {
	storeProducts := []StoreProduct{}
	total := 0
	db := common.GetDB().Model(storeProduct)
	db, err := buildWhere(rawQuery, db)
	if err != nil {
		return &storeProducts, total, err
	}
	db, err = buildOrder(rawOrder, db)
	if err != nil {
		return &storeProducts, total, err
	}
	db.Offset(offset).
		Limit(limit).
		Find(&storeProducts).
		Count(&total)
	err = db.Error
	return &storeProducts, total, err
}

func (storeProduct *StoreProduct) Get() (*StoreProduct, error) {
	err := common.GetDB().Find(&storeProduct).Error
	return storeProduct, err
}

type ResProductInfo struct {
	MerId        int                      `json:"mer_id"` //商户Id(0为总后台管理员创建,不为0的时候是商户后台创建)
	PriceName    string                   `json:"priceName"`
	ProductAttr  []map[string]interface{} `json:"productAttr"`
	ProductValue map[string]interface{}   `json:"productValue"`
	Reply        interface{}              `json:"reply"`
	ReplyChance  int                      `json:"replyChance"`
	ReplyCount   int                      `json:"replyCount"`
	StoreInfo    *StoreInfo               `json:"storeInfo"`
}

type StoreInfo struct {
	Id           int      `json:"id"`            //商品id
	MerId        int      `json:"mer_id"`        //商户Id(0为总后台管理员创建,不为0的时候是商户后台创建)
	Image        string   `json:"image"`         //商品图片
	SliderImage  []string `json:"slider_image"`  //轮播图
	StoreName    string   `json:"store_name"`    //商品名称
	StoreInfo    string   `json:"store_info"`    //商品简介
	Keyword      string   `json:"keyword"`       //关键字
	CateId       string   `json:"cate_id"`       //分类id
	Price        float64  `json:"price"`         //商品价格
	VipPrice     float64  `json:"vip_price"`     //会员价格
	OtPrice      float64  `json:"ot_price"`      //市场价
	Postage      float64  `json:"postage"`       //邮费
	UnitName     string   `json:"unit_name"`     //单位名
	Sort         int      `json:"sort"`          //排序
	Sales        int      `json:"sales"`         //销量
	Stock        int      `json:"stock"`         //库存
	IsShow       int      `json:"is_show"`       //状态（0：未上架，1：上架）
	IsHot        int      `json:"is_hot"`        //是否热卖
	IsBenefit    int      `json:"is_benefit"`    //是否优惠
	IsBest       int      `json:"is_best"`       //是否精品
	IsNew        int      `json:"is_new"`        //是否新品
	Description  string   `json:"description"`   //产品描述
	AddTime      int      `json:"add_time"`      //添加时间
	IsPostage    int      `json:"is_postage"`    //是否包邮
	IsDel        int      `json:"is_del"`        //是否删除
	MerUse       int      `json:"mer_use"`       //商户是否代理 0不可代理1可代理
	GiveIntegral float64  `json:"give_integral"` //获得积分
	Cost         float64  `json:"cost"`          //成本价
	IsSeckill    int      `json:"is_seckill"`    //秒杀状态 0 未开启 1已开启
	IsBargain    int      `json:"is_bargain"`    //砍价状态 0未开启 1开启
	Ficti        int      `json:"ficti"`         //虚拟销量
	Browse       int      `json:"browse"`        //浏览量
	CodePath     string   `json:"code_path"`     //产品二维码地址(用户小程序海报)
	StoreCode    string   `json:"store_code"`    //产品编码
	Fscales      int      `json:"fscales"`       // 虚拟销量+ 真实销量
}

func (storeProduct *StoreProduct) GetById(id int) (r *ResProductInfo, err error) {
	db := common.GetDB()
	err = db.Table(storeProduct.TableName()).Select("id, add_time, browse, cate_id, code_path, cost, description, ficti, sales, give_integral, image, is_bargain, "+
		"is_benefit, is_best, is_del, is_hot, is_new, is_postage, is_seckill, is_show, keyword, mer_id, mer_use, ot_price, postage, price, slider_image, sort, stock, "+
		"store_info, store_name, unit_name, vip_price").Where("id = ?", id).First(&storeProduct).Error
	if err != nil {
		return
	}

	storeProductAttr := &StoreProductAttr{}
	err = storeProductAttr.GetByProductId(id)
	if err != nil && err.Error() != "record not found" {
		return
	}
	// storeProductAttrResult := &StoreProductAttrResult{}
	// err = storeProductAttrResult.GetByProductId(id)
	// if err != nil {
	// 	return
	// }
	storeProductAttrValue := &StoreProductAttrValue{}
	attrValues, err := storeProductAttrValue.ListByProductId(id)
	if err != nil && err.Error() != "record not found" {
		return
	}

	mAttrs := make([]map[string]interface{}, 0)
	if storeProductAttr != nil {
		for _, item := range strings.Split(storeProductAttr.AttrValues, ",") {
			itemdata := make(map[string]interface{}, 0)
			itemdata["attr"] = item
			itemdata["check"] = false
			mAttrs = append(mAttrs, itemdata)
		}
	}

	mValues := make(map[string]interface{}, 0)
	if attrValues != nil {
		for _, item := range attrValues {
			mValues[item.Suk] = item
		}
	}

	storeInfo := &StoreInfo{
		Id:           storeProduct.Id,
		MerId:        storeProduct.MerId,
		Image:        storeProduct.Image,
		StoreName:    storeProduct.StoreName,
		StoreInfo:    storeProduct.StoreInfo,
		Keyword:      storeProduct.Keyword,
		CateId:       storeProduct.CateId,
		Price:        storeProduct.Price,
		VipPrice:     storeProduct.VipPrice,
		OtPrice:      storeProduct.OtPrice,
		Postage:      storeProduct.Postage,
		UnitName:     storeProduct.UnitName,
		Sort:         storeProduct.Sort,
		Sales:        storeProduct.Sales,
		Stock:        storeProduct.Stock,
		IsShow:       storeProduct.IsShow,
		IsHot:        storeProduct.IsHot,
		IsBenefit:    storeProduct.IsBenefit,
		IsBest:       storeProduct.IsBest,
		IsNew:        storeProduct.IsNew,
		Description:  storeProduct.Description,
		AddTime:      storeProduct.AddTime,
		IsPostage:    storeProduct.IsPostage,
		IsDel:        storeProduct.IsDel,
		MerUse:       storeProduct.MerUse,
		GiveIntegral: storeProduct.GiveIntegral,
		Cost:         storeProduct.Cost,
		IsSeckill:    storeProduct.IsSeckill,
		IsBargain:    storeProduct.IsBargain,
		Ficti:        storeProduct.Ficti,
		Browse:       storeProduct.Browse,
		CodePath:     storeProduct.CodePath,
		StoreCode:    storeProduct.StoreCode,
	}

	err = json.Unmarshal([]byte(storeProduct.SliderImage), &storeInfo.SliderImage)
	if err != nil {
		return
	}

	storeInfo.Fscales = storeProduct.Sales + storeProduct.Ficti

	r = &ResProductInfo{
		StoreInfo:    storeInfo,
		ProductAttr:  mAttrs,
		ProductValue: mValues,
	}

	return
	// @Todo 设置会员的价格
}
