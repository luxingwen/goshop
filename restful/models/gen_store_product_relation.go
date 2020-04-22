//generate by gen
package models

import (
	"goshop/restful/common"

	"time"
)

//商品点赞和收藏表
type StoreProductRelation struct {
	Uid       int    `gorm:"column:uid"`        //用户ID
	ProductId int    `gorm:"column:product_id"` //商品ID
	Type      string `gorm:"column:type"`       //类型(收藏(collect）、点赞(like))
	Category  string `gorm:"column:category"`   //某种类型的商品(普通商品、秒杀商品)
	AddTime   int    `gorm:"column:add_time"`   //添加时间

}

//修改默认表名
func (StoreProductRelation) TableName() string {
	return "eb_store_product_relation"
}

func (storeProductRelation *StoreProductRelation) Insert() error {
	err := common.GetDB().Create(storeProductRelation).Error
	return err
}

func (storeProductRelation *StoreProductRelation) Patch() error {
	err := common.GetDB().Model(storeProductRelation).Updates(storeProductRelation).Error
	return err
}

func (storeProductRelation *StoreProductRelation) Update() error {
	err := common.GetDB().Save(storeProductRelation).Error
	return err
}

func (storeProductRelation *StoreProductRelation) Delete() error {
	return common.GetDB().Delete(storeProductRelation).Error
}

func (storeProductRelation *StoreProductRelation) List(rawQuery string, rawOrder string, offset int, limit int) (*[]StoreProductRelation, int, error) {
	storeProductRelations := []StoreProductRelation{}
	total := 0
	db := common.GetDB().Model(storeProductRelation)
	db, err := buildWhere(rawQuery, db)
	if err != nil {
		return &storeProductRelations, total, err
	}
	db, err = buildOrder(rawOrder, db)
	if err != nil {
		return &storeProductRelations, total, err
	}
	db.Offset(offset).
		Limit(limit).
		Find(&storeProductRelations).
		Count(&total)
	err = db.Error
	return &storeProductRelations, total, err
}

func (storeProductRelation *StoreProductRelation) Get() (*StoreProductRelation, error) {
	err := common.GetDB().Find(&storeProductRelation).Error
	return storeProductRelation, err
}

// 获取用户收藏的所有产品的个数
func (storeProductRelation *StoreProductRelation) GetUserIdCollect(uid int) (count int, err error) {
	db := common.GetDB()
	err = db.Table(storeProductRelation.TableName()).Where("uid = ? AND type = ?", uid, "collect").Count(&count).Error
	return
}

// 获取用户点赞所有产品的个数
func (storeProductRelation *StoreProductRelation) GetUserIdLike(uid int) (count int, err error) {
	db := common.GetDB()
	err = db.Table(storeProductRelation.TableName()).Where("uid = ? AND type = ?", uid, "like").Count(&count).Error
	return
}

// 是否有关联
func (storeProductRelation *StoreProductRelation) IsProductRelation(productId int, uid int, typ string) (r bool, err error) {
	db := common.GetDB()
	rows, err := db.Table(storeProductRelation.TableName()).Select("count(*)").Where("uid = ? AND product_id = ? AND type = ?", uid, productId, typ).Rows()
	if err != nil {
		return
	}
	count := 0
	for rows.Next() {
		rows.Scan(&count)
	}
	if count > 0 {
		return true, nil
	}
	return
}

// 取消 点赞 收藏
func (storeProductRelation *StoreProductRelation) UnProductRelation(uid, id int, relationType, category string) (err error) {
	if id <= 0 {
		return
	}
	db := common.GetDB()
	err = db.Table(storeProductRelation.TableName()).Where("uid = ? AND product_id = ? AND type = ? AND category = ?", uid, id, relationType, category).Delete(StoreProductRelation{}).Error
	return
}

// 添加收藏
func (storeProductRelation *StoreProductRelation) ProductRelation(uid, id int, relationType, category string) (err error) {
	storeProductRelation.Uid = uid
	storeProductRelation.ProductId = id
	storeProductRelation.Type = relationType
	storeProductRelation.Category = category
	storeProductRelation.AddTime = int(time.Now().Unix())

	db := common.GetDB()
	err = db.Create(storeProductRelation).Error

	return
}

//
type UserCollectProduct struct {
	Id        int     `gorm:"column:id" json:"id"`                 //商品id
	Image     string  `gorm:"column:image" json:"image"`           //商品图片
	StoreName string  `gorm:"column:store_name" json:"store_name"` //商品名称
	Price     float64 `gorm:"column:price" json:"price"`           //商品价格
	OtPrice   float64 `gorm:"column:ot_price" json:"ot_price"`     //市场价
	Sales     int     `gorm:"column:sales" json:"sales"`           //销量
	IsShow    int     `gorm:"column:is_show" json:"is_show"`       //状态（0：未上架，1：上架）
	IsFail    int     `json:"is_fail"`
}

// 获取某个用户收藏的产品
func (storeProductRelation *StoreProductRelation) GetUserCollectProduct(uid int, req *Query) (r []*UserCollectProduct, err error) {
	if uid <= 0 {
		return
	}

	storeProduct := &StoreProduct{}
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

	list := make([]*UserCollectProduct, 0)

	err = db.Raw("SELECT B.id,B.store_name,B.price,B.ot_price,B.sales,B.image,B.is_del,B.is_show FROM "+storeProductRelation.TableName()+
		" A LEFT JOIN "+storeProduct.TableName()+" B ON A.product_id = B.id WHERE A.uid = ? AND A.type = ?"+
		" AND A.category = ? ORDER BY A.add_time DESC LIMIT ?, ? ", uid, "collect", "product", offset, limit).Scan(&list).Error
	if err != nil {
		return
	}

	r = make([]*UserCollectProduct, 0)

	for _, item := range list {
		if item.Id > 0 {
			if item.IsShow > 0 && item.IsShow > 0 {
				item.IsFail = 1
			}
			r = append(r, item)
		}
	}
	return
}

// 收藏产品删除
func (storeProductRelation *StoreProductRelation) UserCollectProductDel(uid, productId int) (err error) {
	db := common.GetDB()
	err = db.Table(storeProductRelation.TableName()).Where("uid = ? AND product_id = ?", uid, productId).Delete(storeProductRelation).Error
	if err != nil {
		return
	}
	return
	//return storeProductRelation.GetUserCollectProduct(uid, &Query{Page: 1, PageNum: 20})
}
