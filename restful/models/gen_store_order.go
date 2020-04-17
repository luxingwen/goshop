//generate by gen
package models

import (
	"goshop/restful/common"
)

//订单表
type StoreOrder struct {
	Id                     int     `gorm:"column:id"`                        //订单ID
	OrderId                string  `gorm:"column:order_id"`                  //订单号
	Uid                    int     `gorm:"column:uid"`                       //用户id
	RealName               string  `gorm:"column:real_name"`                 //用户姓名
	UserPhone              string  `gorm:"column:user_phone"`                //用户电话
	UserAddress            string  `gorm:"column:user_address"`              //详细地址
	CartId                 string  `gorm:"column:cart_id"`                   //购物车id
	TotalNum               int     `gorm:"column:total_num"`                 //订单商品总数
	TotalPrice             float64 `gorm:"column:total_price"`               //订单总价
	TotalPostage           float64 `gorm:"column:total_postage"`             //邮费
	PayPrice               float64 `gorm:"column:pay_price"`                 //实际支付金额
	PayPostage             float64 `gorm:"column:pay_postage"`               //支付邮费
	DeductionPrice         float64 `gorm:"column:deduction_price"`           //抵扣金额
	CouponId               int     `gorm:"column:coupon_id"`                 //优惠券id
	CouponPrice            float64 `gorm:"column:coupon_price"`              //优惠券金额
	Paid                   int     `gorm:"column:paid"`                      //支付状态
	PayTime                int     `gorm:"column:pay_time"`                  //支付时间
	PayType                string  `gorm:"column:pay_type"`                  //支付方式
	AddTime                int     `gorm:"column:add_time"`                  //创建时间
	Status                 int     `gorm:"column:status"`                    //订单状态（-1 : 申请退款 -2 : 退货成功 0：待发货；1：待收货；2：已收货；3：待评价；-1：已退款）
	RefundStatus           int     `gorm:"column:refund_status"`             //0 未退款 1 申请中 2 已退款
	RefundReasonWapImg     string  `gorm:"column:refund_reason_wap_img"`     //退款图片
	RefundReasonWapExplain string  `gorm:"column:refund_reason_wap_explain"` //退款用户说明
	RefundReasonTime       int     `gorm:"column:refund_reason_time"`        //退款时间
	RefundReasonWap        string  `gorm:"column:refund_reason_wap"`         //前台退款原因
	RefundReason           string  `gorm:"column:refund_reason"`             //不退款的理由
	RefundPrice            float64 `gorm:"column:refund_price"`              //退款金额
	DeliveryName           string  `gorm:"column:delivery_name"`             //快递名称/送货人姓名
	DeliveryType           string  `gorm:"column:delivery_type"`             //发货类型
	DeliveryId             string  `gorm:"column:delivery_id"`               //快递单号/手机号
	GainIntegral           float64 `gorm:"column:gain_integral"`             //消费赚取积分
	UseIntegral            float64 `gorm:"column:use_integral"`              //使用积分
	BackIntegral           float64 `gorm:"column:back_integral"`             //给用户退了多少积分
	Mark                   string  `gorm:"column:mark"`                      //备注
	IsDel                  int     `gorm:"column:is_del"`                    //是否删除
	Unique                 string  `gorm:"column:unique"`                    //唯一id(md5加密)类似id
	Remark                 string  `gorm:"column:remark"`                    //管理员备注
	MerId                  int     `gorm:"column:mer_id"`                    //商户ID
	IsMerCheck             int     `gorm:"column:is_mer_check"`              //
	CombinationId          int     `gorm:"column:combination_id"`            //拼团产品id0一般产品
	PinkId                 int     `gorm:"column:pink_id"`                   //拼团id 0没有拼团
	Cost                   float64 `gorm:"column:cost"`                      //成本价
	SeckillId              int     `gorm:"column:seckill_id"`                //秒杀产品ID
	BargainId              int     `gorm:"column:bargain_id"`                //砍价id
	IsChannel              int     `gorm:"column:is_channel"`                //支付渠道(0微信公众号1微信小程序)
	IsSystemDel            int     `gorm:"column:is_system_del"`             //后台管理员删除

}

//修改默认表名
func (StoreOrder) TableName() string {
	return "eb_store_order"
}

func (storeOrder *StoreOrder) Insert() error {
	err := common.GetDB().Create(storeOrder).Error
	return err
}

func (storeOrder *StoreOrder) Patch() error {
	err := common.GetDB().Model(storeOrder).Updates(storeOrder).Error
	return err
}

func (storeOrder *StoreOrder) Update() error {
	err := common.GetDB().Save(storeOrder).Error
	return err
}

func (storeOrder *StoreOrder) Delete() error {
	return common.GetDB().Delete(storeOrder).Error
}

func (storeOrder *StoreOrder) List(rawQuery string, rawOrder string, offset int, limit int) (*[]StoreOrder, int, error) {
	storeOrders := []StoreOrder{}
	total := 0
	db := common.GetDB().Model(storeOrder)
	db, err := buildWhere(rawQuery, db)
	if err != nil {
		return &storeOrders, total, err
	}
	db, err = buildOrder(rawOrder, db)
	if err != nil {
		return &storeOrders, total, err
	}
	db.Offset(offset).
		Limit(limit).
		Find(&storeOrders).
		Count(&total)
	err = db.Error
	return &storeOrders, total, err
}

func (storeOrder *StoreOrder) Get() (*StoreOrder, error) {
	err := common.GetDB().Find(&storeOrder).Error
	return storeOrder, err
}

func (storeOrder *StoreOrder) GetOrderStatusNum(uid int) (r map[string]int, err error) {
	var (
		noBuy            int
		noPostageNoPink  int
		noPostageYesPink int
		noPostage        int
		noTake           int
		noReply          int
		noPink           int
		noRefund         int
	)

	db := common.GetDB()

	err = db.Table(storeOrder.TableName()).Where("uid = ? AND paid = ? AND is_del = ? pay_type <> ?", uid, 0, 0, "offline").Count(&noBuy).Error
	if err != nil {
		return
	}

	err = db.Table(storeOrder.TableName()).Where("uid = ? AND paid = ? AND pink_id = ? AND is_del = ? status = ? AND pay_type <> ?", uid, 1, 0, 0, 0, "offline").Count(&noPostageNoPink).Error

	if err != nil {
		return
	}

	noPostage = noPostageNoPink + noPostageYesPink

	storePink := &StorePink{}
	err = db.Raw("SELECT count(*) FROM "+storeOrder.TableName()+" o LEFT JOIN "+storePink.TableName()+" p ON o.pink_id = p.id WHERE p.status = ? AND o.paid = ? AND o.is_del = ? AND o.status = ? AND o.pay_type = ?", 2, 1, 0, 0, "offline").Scan(&noPostageYesPink).Error
	if err != nil {
		return
	}

	err = db.Table(storeOrder.TableName()).Where("uid = ? AND paid = ? AND is_del = ? AND status = ? AND pay_type <> ?", uid, 1, 0, 1, "offline").Count(&noTake).Error
	if err != nil {
		return
	}

	err = db.Table(storeOrder.TableName()).Where("uid = ? AND paid = ? AND is_del = ? AND status = ?", uid, 1, 0, 0, 2).Count(&noReply).Error
	if err != nil {
		return
	}

	err = db.Raw("SELECT count(*) FROM "+storeOrder.TableName()+" o LEFT JOIN "+storePink.TableName()+" p ON p.id = o.pink_id WHERE p.status = ? AND o.paid = ? AND o.is_del = ? AND o.status = ? AND o.pay_type <> ?", 1, 1, 0, 0, "offline").Scan(&noPink).Error
	if err != nil {
		return
	}

	err = db.Table(storeOrder.TableName()).Where("uid = ? AND paid = ? AND is_del = ? AND refund_status IN(?)", uid, 1, 0, []int{1, 2}).Count(&noRefund).Error
	if err != nil {
		return
	}

	r = make(map[string]int, 0)

	r["noBuy"] = noBuy
	r["noPostageNoPink"] = noPostageNoPink
	r["noPostageYesPink"] = noPostageYesPink
	r["noPostage"] = noPostage
	r["noTake"] = noTake
	r["noReply"] = noReply
	r["noPink"] = noPink
	r["noRefund"] = noRefund

	return
}

// 累计消费
func (storeOrder *StoreOrder) GetOrderStatusSum(uid int) (count int, err error) {
	db := common.GetDB()
	err = db.Table(storeOrder.TableName()).Select("sum(pay_price)").Where("uid = ? AND is_del = ? AND paid = ? ", uid, 0, 1).Scan(&count).Error
	return
}
