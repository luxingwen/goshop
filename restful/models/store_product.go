package models

import (
	"strings"

	"goshop/restful/common"
)

// $info['bastList'] = StoreProduct::getBestProduct('id,image,store_name,cate_id,price,ot_price,IFNULL(sales,0) + IFNULL(ficti,0) as sales,unit_name,sort',$bastNumber,$this->uid);//TODO 精品推荐个数
type BestStoreProduct struct {
	Id        int     `gorm:"column:id" json:"id"`                 //商品id
	Image     string  `gorm:"column:image" json:"image"`           //商品图片
	StoreName string  `gorm:"column:store_name" json:"store_name"` //商品名称
	CateId    string  `gorm:"column:cate_id" json:"cate_id"`       //分类id
	Price     float64 `gorm:"column:price" json:"price"`           //商品价格
	OtPrice   float64 `gorm:"column:ot_price" json:"ot_price"`     //市场价
	UnitName  string  `gorm:"column:unit_name" json:"unit_name"`   //单位名
	Stock     int     `gorm:"column:stock" json:"stock"`           //库存
	Sort      int     `gorm:"column:sort" json:"sort"`             //排序
	Sales     float64 `gorm:"column:sales" json:"sales"`
}

func (storeProduct *StoreProduct) GetBestProduct(limit, uid int) (r []*BestStoreProduct, err error) {
	db := common.GetDB()

	r = make([]*BestStoreProduct, 0)
	err = db.Raw("SELECT id, image, store_name, cate_id, price, ot_price,IFNULL(sales,0) + IFNULL(ficti,0) as sales, unit_name, sort FROM "+
		storeProduct.TableName()+" WHERE is_best = ? AND is_del = ? AND mer_id = ? AND stock > ? AND is_show = ? ORDER BY `sort` DESC, `id` DESC LIMIT ?", 1, 0, 0, 0, 1, limit).Scan(&r).Error

	// @Todo 设置会员价格

	return
}

//id,image,store_name,cate_id,price,unit_name,sort
func (storeProduct *StoreProduct) GetNewProduct(limit int) (r []*BestStoreProduct, err error) {
	db := common.GetDB()

	r = make([]*BestStoreProduct, 0)

	err = db.Raw("SELECT id,image,store_name,cate_id,price,unit_name,sort FROM "+storeProduct.TableName()+
		" WHERE is_new = ? AND is_del = ? AND mer_id = ? AND stock > ? AND is_show > ? ORDER BY `sort` DESC, `id` DESC LIMIT ?", 1, 0, 0, 0, 1, limit).Scan(&r).Error
	return
}

// 优惠产品
// id,image,store_name,cate_id,price,ot_price,stock,unit_name,sort
func (storeProduct *StoreProduct) GetBenefitProduct(limit int) (r []*BestStoreProduct, err error) {
	db := common.GetDB()
	r = make([]*BestStoreProduct, 0)

	err = db.Raw("SELECT id,image,store_name,cate_id,price,ot_price,stock,unit_name,sort FROM "+storeProduct.TableName()+
		" WHERE is_benefit = ? AND is_del = ? AND mer_id = ? AND stock > ? AND is_show = ? ORDER BY `sort` DESC, id DESC LIMIT ?", 1, 0, 0, 0, 1, limit).Scan(&r).Error
	return
}

// 热卖产品
//id,image,store_name,cate_id,price,unit_name,sort
func (storeProduct *StoreProduct) GetHotProduct(limit, uid int) (r []*BestStoreProduct, err error) {
	db := common.GetDB()
	r = make([]*BestStoreProduct, 0)

	err = db.Raw("SELECT id,image,store_name,cate_id,price,unit_name,sort FROM "+storeProduct.TableName()+
		" WHERE is_hot = ? AND is_del = ? AND mer_id = ? AND stock > ? AND is_show = ? ORDER BY `sort` DESC, `id` DESC LIMIT ?", 1, 0, 0, 0, 1, limit).Scan(&r).Error
	// @Todo 设置会员价格
	return
}

//add_time,browse,cate_id,code_path,cost,description,ficti,give_integral,id,image,is_bargain,is_benefit,is_best,is_del,is_hot,is_new,is_postage,is_seckill,is_show,keyword,mer_id,mer_use,ot_price,postage,price,sales,slider_image,sort,stock,store_info,store_name,unit_name,vip_price,IFNULL(sales,0) + IFNULL(ficti,0) as fsales
type DetailStoreProduct struct {
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
	Fsales       float64 `gorm:"column:fsales"`
}

func (storeProduct *StoreProduct) GetValidProduct(id int) (r *DetailStoreProduct, err error) {
	db := common.GetDB()

	err = db.Raw("SELECT add_time,browse,cate_id,code_path,cost,description,ficti,give_integral,id,image,is_bargain,is_benefit,is_best,"+
		"is_del,is_hot,is_new,is_postage,is_seckill,is_show,keyword,mer_id,mer_use,ot_price,postage,price,sales,slider_image,sort,"+
		"stock,store_info,store_name,unit_name,vip_price,IFNULL(sales,0) + IFNULL(ficti,0) as fsales FROM "+storeProduct.TableName()+
		" WHERE is_del = ? AND is_show = ? AND id = ?", 0, 1, id).Scan(&r).Error
	return
}

type ReqStoreProductQuery struct {
	Query
	Sid        int    `form:"sid" json:"sid"`
	Cid        int    `form:"cid" json:"cid"`
	Keyword    string `form:"keyword" json:"keyword"`
	PriceOrder string `form:"priceOrder" json:"priceOrder"`
	SalesOrder string `form:"salesOrder" json:"salesOrder"`
	News       int    `form:"news" json:"news"`
}

type ResProduct struct {
	Id        int     `gorm:"column:id" json:"id"`                 //商品id
	StoreName string  `gorm:"column:store_name" json:"store_name"` //商品名称
	CateId    string  `gorm:"column:cate_id" json:"cate_id"`       //分类id
	Image     string  `gorm:"column:image" json:"image"`           //商品图片
	Sales     float64 `gorm:"column:sales" json:"sales"`
	Price     float64 `gorm:"column:price" json:"price"` //商品价格
	Stock     int     `gorm:"column:stock" json:"stock"` //库存
}

func (storeProduct *StoreProduct) GetProductList(req *ReqStoreProductQuery) (r []*ResProduct, count int, err error) {
	db := common.GetDB().Table(storeProduct.TableName())
	db = db.Select("id,store_name,cate_id,image,IFNULL(sales,0) + IFNULL(ficti,0) as sales,price,stock")
	db = db.Where("is_del = ? AND is_show = ? AND mer_id = ?", 0, 1, 0)
	if req.Sid > 0 {
		storeProductCate := &StoreProductCate{}
		ids, err := storeProductCate.GetProductIdsByCateId(req.Sid)
		if err != nil {
			return nil, 0, err
		}
		if len(ids) > 0 {
			db = db.Where("id in(?)", ids)
		} else {
			db = db.Where("cate_id = ?", -1)
		}

	} else if req.Cid > 0 {
		storeCategory := &StoreCategory{}
		categorys, err := storeCategory.PidBySidList(req.Cid)
		if err != nil {
			return nil, 0, err
		}
		ids := make([]int, 0)
		for _, item := range categorys {
			ids = append(ids, item.Id)
		}
		if len(ids) > 0 {
			db = db.Where("cate_id in(?)", ids)
		}
	}
	req.Keyword = strings.TrimSpace(req.Keyword)
	if req.Keyword != "" {
		db = db.Where("(keyword LIKE %%?%% OR store_name LIKE %%?%%", req.Keyword, req.Keyword)
	}
	if req.News != 0 {
		db = db.Where("is_new = ?", 1)
	}
	if req.PriceOrder != "" {
		if req.PriceOrder == "desc" {
			db = db.Order("price DESC")
		} else {
			db = db.Order("price ASC")
		}
	}

	if req.SalesOrder != "" {
		if req.SalesOrder == "desc" {
			db = db.Order("sales DESC")
		} else {
			db = db.Order("sales ASC")
		}
	}

	limit := 10
	page := 0

	if req.Page > 0 {
		page = req.Page - 1
	}
	if req.PageNum > 0 {
		limit = req.PageNum
	}

	offset := limit * page

	err = db.Offset(offset).Limit(limit).Find(&r).Count(&count).Error
	return

	// @todo 设置会员的价格

}

type ReqGoodsSearch struct {
	Query
	Keyword string `form:"keyword" json:"keyword"`
}

// 分类搜索
func (storeProduct *StoreProduct) GetSearchStorePage(req *ReqGoodsSearch, uid int) (r []*ResProduct, count int, err error) {
	db := common.GetDB().Table(storeProduct.TableName())

	req.Keyword = strings.TrimSpace(req.Keyword)
	if req.Keyword != "" {
		db = db.Where("(keyword LIKE %%?%% OR store_name LIKE %%?%%", req.Keyword, req.Keyword)
	}

	db = db.Select("id,store_name,cate_id,image,ficti as sales,price,stock")
	db = db.Where("is_del = ? AND is_show = ? AND mer_id = ?", 0, 1, 0)

	limit := 10
	page := 0

	if req.Page > 0 {
		page = req.Page - 1
	}
	if req.PageNum > 0 {
		limit = req.PageNum
	}

	offset := limit * page

	err = db.Offset(offset).Limit(limit).Find(&r).Count(&count).Error
	return

	// @Todo 设置会员的价格

}

// 热卖产品
func (storeProduct *StoreProduct) GetHotProductLoading(req *ReqGoodsSearch) (r []*ResProduct, count int, err error) {
	db := common.GetDB().Table(storeProduct.TableName())
	db = db.Select("id,image,store_name,cate_id,price,unit_name,sort")
	limit := 10
	page := 0

	if req.Page > 0 {
		page = req.Page - 1
	}
	if req.PageNum > 0 {
		limit = req.PageNum
	}

	offset := limit * page

	err = db.Where("is_hot = ? AND is_del = ? AND mer_id = ? AND stock > ? AND is_show = ?", 1, 0, 0, 0, 1).Order("sort DESC, id DESC").
		Offset(offset).Limit(limit).Find(&r).Count(&count).Error
	if err != nil {
		return
	}
	return

}
