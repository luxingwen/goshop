package models

import (
	"goshop/restful/common"
)

// $info['bastList'] = StoreProduct::getBestProduct('id,image,store_name,cate_id,price,ot_price,IFNULL(sales,0) + IFNULL(ficti,0) as sales,unit_name,sort',$bastNumber,$this->uid);//TODO 精品推荐个数
type BestStoreProduct struct {
	Id        int     `gorm:"column:id" json:"id"`                //商品id
	Image     string  `gorm:"column:image" json:"image"`          //商品图片
	StoreName string  `gorm:"column:store_name" json:"storeName"` //商品名称
	CateId    string  `gorm:"column:cate_id" json:"cateId"`       //分类id
	Price     float64 `gorm:"column:price" json:"price"`          //商品价格
	OtPrice   float64 `gorm:"column:ot_price" json:"otPrice"`     //市场价
	UnitName  string  `gorm:"column:unit_name" json:"unitName"`   //单位名
	Stock     int     `gorm:"column:stock" json:"stock"`          //库存
	Sort      int     `gorm:"column:sort" json:"sort"`            //排序
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
