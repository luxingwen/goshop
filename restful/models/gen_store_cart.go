//generate by gen
package models

import (
	"goshop/restful/common"

	"errors"
	"time"
)

//购物车表
type StoreCart struct {
	Id                int    `gorm:"column:id"`                  //购物车表ID
	Uid               int    `gorm:"column:uid"`                 //用户ID
	Type              string `gorm:"column:type"`                //类型
	ProductId         int    `gorm:"column:product_id"`          //商品ID
	ProductAttrUnique string `gorm:"column:product_attr_unique"` //商品属性
	CartNum           int    `gorm:"column:cart_num"`            //商品数量
	AddTime           int    `gorm:"column:add_time"`            //添加时间
	IsPay             int    `gorm:"column:is_pay"`              //0 = 未购买 1 = 已购买
	IsDel             int    `gorm:"column:is_del"`              //是否删除
	IsNew             int    `gorm:"column:is_new"`              //是否为立即购买
	CombinationId     int    `gorm:"column:combination_id"`      //拼团id
	SeckillId         int    `gorm:"column:seckill_id"`          //秒杀产品ID
	BargainId         int    `gorm:"column:bargain_id"`          //砍价id

}

//修改默认表名
func (StoreCart) TableName() string {
	return "eb_store_cart"
}

func (storeCart *StoreCart) Insert() error {
	err := common.GetDB().Create(storeCart).Error
	return err
}

func (storeCart *StoreCart) Patch() error {
	err := common.GetDB().Model(storeCart).Updates(storeCart).Error
	return err
}

func (storeCart *StoreCart) Update() error {
	err := common.GetDB().Save(storeCart).Error
	return err
}

func (storeCart *StoreCart) Delete() error {
	return common.GetDB().Delete(storeCart).Error
}

func (storeCart *StoreCart) List(rawQuery string, rawOrder string, offset int, limit int) (*[]StoreCart, int, error) {
	storeCarts := []StoreCart{}
	total := 0
	db := common.GetDB().Model(storeCart)
	db, err := buildWhere(rawQuery, db)
	if err != nil {
		return &storeCarts, total, err
	}
	db, err = buildOrder(rawOrder, db)
	if err != nil {
		return &storeCarts, total, err
	}
	db.Offset(offset).
		Limit(limit).
		Find(&storeCarts).
		Count(&total)
	err = db.Error
	return &storeCarts, total, err
}

func (storeCart *StoreCart) Get() (*StoreCart, error) {
	err := common.GetDB().Find(&storeCart).Error
	return storeCart, err
}

func (storeCart *StoreCart) SetDelete(id int) (err error) {
	db := common.GetDB().Model(storeCart)
	err = db.Where("id = ?", id).Update("is_del", 1).Error
	return
}

func (storeCart *StoreCart) GetUserCartNum(uid int, typ string) (count int, err error) {
	db := common.GetDB().Table(storeCart.TableName())
	err = db.Where("uid = ? AND type = ? AND is_pay = ? AND is_del = ? AND is_new = ?", uid, typ, 0, 0, 0).Count(&count).Error
	return
}

//购物车表
type UserStoreCart struct {
	Id                int             `gorm:"column:id" json:"id"`                                   //购物车表ID
	Uid               int             `gorm:"column:uid" json:"uid"`                                 //用户ID
	Type              string          `gorm:"column:type" json:"type"`                               //类型
	ProductId         int             `gorm:"column:product_id" json:"product_id"`                   //商品ID
	ProductAttrUnique string          `gorm:"column:product_attr_unique" json:"product_attr_unique"` //商品属性
	CartNum           int             `gorm:"column:cart_num" json:"cart_num"`                       //商品数量
	AddTime           int             `gorm:"column:add_time" json:"add_time"`                       //添加时间
	IsPay             int             `gorm:"column:is_pay" json:"is_pay"`                           //0 = 未购买 1 = 已购买
	IsDel             int             `gorm:"column:is_del" json:"is_del"`                           //是否删除
	IsNew             int             `gorm:"column:is_new" json:"is_new"`                           //是否为立即购买
	CombinationId     int             `gorm:"column:combination_id" json:"combination_id"`           //拼团id
	SeckillId         int             `gorm:"column:seckill_id" json:"seckill_id"`                   //秒杀产品ID
	BargainId         int             `gorm:"column:bargain_id" json:"bargain_id"`                   //砍价id
	ProductInfo       *ResProductInfo `json:"productInfo"`                                           // 产品信息
}

// 获取用户购物车列表
func (storeCart *StoreCart) GetUserProductCartList(uid int) (r map[string]interface{}, err error) {
	db := common.GetDB().Model(storeCart)
	list := make([]*StoreCart, 0)
	err = db.Where("uid = ? AND type = ? AND is_pay = ? AND is_del = ?", uid, "product", 0, 0).Find(&list).Error
	if err != nil {
		return
	}
	storeProduct := &StoreProduct{}

	valid := make([]*UserStoreCart, 0)
	inValid := make([]*UserStoreCart, 0)

	for _, item := range list {
		itemData := &UserStoreCart{
			Id:                item.Id,
			Uid:               item.Uid,
			Type:              item.Type,
			ProductId:         item.ProductId,
			ProductAttrUnique: item.ProductAttrUnique,
			CartNum:           item.CartNum,
			AddTime:           item.AddTime,
			IsPay:             item.IsPay,
			IsDel:             item.IsDel,
			IsNew:             item.IsNew,
			CombinationId:     item.CombinationId,
			SeckillId:         item.SeckillId,
			BargainId:         item.BargainId,
		}
		if item.SeckillId > 0 {

		} else if item.BargainId > 0 {

		} else if item.CombinationId > 0 {

		} else {
			productInfo, err := storeProduct.GetById(item.ProductId)
			if err == nil && productInfo != nil {
				itemData.ProductInfo = productInfo
			}
		}

		if itemData.ProductInfo == nil {
			item.SetDelete(item.Id)

		}
		valid = append(valid, itemData)

	}

	r = make(map[string]interface{}, 0)
	r["valid"] = valid
	r["invalid"] = inValid
	return

}

// 设置购物车
func (storeCart *StoreCart) SetCart() (err error) {
	if storeCart.CartNum < 1 {
		storeCart.CartNum = 1
	}
	if storeCart.SeckillId > 0 { // 秒杀

	} else if storeCart.BargainId > 0 { //砍价

	} else if storeCart.CombinationId > 0 { // 拼团

	}

	db := common.GetDB().Model(storeCart)
	storeCart.AddTime = int(time.Now().Unix())

	rstoreCart := new(StoreCart)
	err = db.Where("type = ? AND uid = ? AND product_id = ? AND product_attr_unique = ? AND is_new = ? AND is_pay = ? AND is_del = ? AND combination_id = ?",
		storeCart.Type, storeCart.Uid, storeCart.ProductId, storeCart.ProductAttrUnique, storeCart.IsNew, storeCart.IsPay, storeCart.IsDel, storeCart.CombinationId).Scan(&rstoreCart).Error
	if err != nil {
		storeCart.AddTime = int(time.Now().Unix())
		err = db.Create(&storeCart).Error
		if err != nil {
			return
		}
		return
	}

	err = db.Where("id = ?", rstoreCart.Id).Update("cart_num", storeCart.AddTime).Error
	return

}

// 移除购物车
func (storeCart *StoreCart) RemoveUserCart(uid int, ids []int) (err error) {
	db := common.GetDB().Model(storeCart)
	err = db.Where("uid = ? AND id IN(?)", uid, ids).Update("is_del", 1).Error
	return
}

// 修改购物车库存
func (storeCart *StoreCart) ChangeUserCartNum(cartId, cardNum, uid int) (err error) {
	db := common.GetDB().Model(storeCart)

	rstoreCart := new(StoreCart)
	err = db.Where("uid = ? AND id = ?", uid, cartId).Scan(&rstoreCart).Error
	if err != nil {
		return
	}
	var stock int

	if rstoreCart.BargainId > 0 {
		// @Todo 获取砍价产品的库存
	} else if rstoreCart.SeckillId > 0 {
		// @Todo 获取秒杀产品的库存
	} else if rstoreCart.CombinationId > 0 {
		// @Todo 获取拼团产品的库存
	} else if rstoreCart.ProductId > 0 {
		storeProduct := &StoreProduct{}
		count, err := storeProduct.GetProductStock(rstoreCart.ProductId, rstoreCart.ProductAttrUnique)
		if err != nil {
			return err
		}
		stock = count
	}
	if stock < 0 {
		return errors.New("暂无库存")
	}
	if cardNum <= 0 {
		return errors.New("库存错误")
	}
	if stock < cardNum {
		return errors.New("库存不足")
	}
	err = db.Where("uid = ? AND id = ?", uid, cartId).Update("cart_num", cardNum).Error

	return
}

// func (storeCart *StoreCart) GetUserCartNum(uid int, typ string) (count int, err error) {
// 	db := common.GetDB().Model(storeCart)
// 	err = db.Where("uid = ? AND type = ? AND is_pay = ? AND is_del = ? AND is_new = ?", uid, typ, 0, 0, 0).Count(&count).Error
// 	return
// }
