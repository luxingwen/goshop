package models

import (
	"goshop/restful/common"
)

func (storeProductCate *StoreProductCate) GetProductIdsByCateId(cid int) (r []int, err error) {
	db := common.GetDB()

	list := make([]*StoreProductCate, 0)
	err = db.Select("product_id").Where("cate_id = ?", cid).Find(&list).Error
	if err != nil {
		return
	}
	for _, item := range list {
		r = append(r, item.ProductId)
	}
	return
}
