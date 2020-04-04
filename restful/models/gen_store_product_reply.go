//generate by gen
package models

import (
	"goshop/restful/common"
)

//评论表
type StoreProductReply struct {
	Id                   int    `gorm:"column:id"`                     //评论ID
	Uid                  int    `gorm:"column:uid"`                    //用户ID
	Oid                  int    `gorm:"column:oid"`                    //订单ID
	Unique               string `gorm:"column:unique"`                 //唯一id
	ProductId            int    `gorm:"column:product_id"`             //产品id
	ReplyType            string `gorm:"column:reply_type"`             //某种商品类型(普通商品、秒杀商品）
	ProductScore         int    `gorm:"column:product_score"`          //商品分数
	ServiceScore         int    `gorm:"column:service_score"`          //服务分数
	Comment              string `gorm:"column:comment"`                //评论内容
	Pics                 string `gorm:"column:pics"`                   //评论图片
	AddTime              int    `gorm:"column:add_time"`               //评论时间
	MerchantReplyContent string `gorm:"column:merchant_reply_content"` //管理员回复内容
	MerchantReplyTime    int    `gorm:"column:merchant_reply_time"`    //管理员回复时间
	IsDel                int    `gorm:"column:is_del"`                 //0未删除1已删除
	IsReply              int    `gorm:"column:is_reply"`               //0未回复1已回复

}

//修改默认表名
func (StoreProductReply) TableName() string {
	return "eb_store_product_reply"
}

func (storeProductReply *StoreProductReply) Insert() error {
	err := common.GetDB().Create(storeProductReply).Error
	return err
}

func (storeProductReply *StoreProductReply) Patch() error {
	err := common.GetDB().Model(storeProductReply).Updates(storeProductReply).Error
	return err
}

func (storeProductReply *StoreProductReply) Update() error {
	err := common.GetDB().Save(storeProductReply).Error
	return err
}

func (storeProductReply *StoreProductReply) Delete() error {
	return common.GetDB().Delete(storeProductReply).Error
}

func (storeProductReply *StoreProductReply) List(rawQuery string, rawOrder string, offset int, limit int) (*[]StoreProductReply, int, error) {
	storeProductReplys := []StoreProductReply{}
	total := 0
	db := common.GetDB().Model(storeProductReply)
	db, err := buildWhere(rawQuery, db)
	if err != nil {
		return &storeProductReplys, total, err
	}
	db, err = buildOrder(rawOrder, db)
	if err != nil {
		return &storeProductReplys, total, err
	}
	db.Offset(offset).
		Limit(limit).
		Find(&storeProductReplys).
		Count(&total)
	err = db.Error
	return &storeProductReplys, total, err
}

func (storeProductReply *StoreProductReply) Get() (*StoreProductReply, error) {
	err := common.GetDB().Find(&storeProductReply).Error
	return storeProductReply, err
}
