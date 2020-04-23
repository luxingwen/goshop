package controllers

import (
	"github.com/gin-gonic/gin"
	// "goshop/restful/common"
	"goshop/restful/models"
	// "goshop/libs/errcode"
	//"encoding/json"
	"errors"
	"fmt"
	"strconv"
)

type ReqLogin struct {
	Username string `json:"username" faker:"username"`
	Password string `json:"password" faker:"password"`
}

// type UserController struct {
// }

func (crtl *UserController) My(c *gin.Context) {
	uidT, ok := c.Get("uid")
	if !ok {
		handleErr(c, errors.New("无效的uid"))
		return
	}
	uid := uidT.(int)

	storeCouponUser := &models.StoreCouponUser{}
	couponCount, err := storeCouponUser.GetUserValidCouponCount(uid)
	if err != nil {
		handleErr(c, err)
		return
	}
	storeProductRelation := &models.StoreProductRelation{}
	like, err := storeProductRelation.GetUserIdCollect(uid)

	if err != nil {
		handleErr(c, err)
		return
	}

	storeOrder := &models.StoreOrder{}
	orderStatusNum, err := storeOrder.GetOrderStatusNum(uid)
	if err != nil {
		handleErr(c, err)
		return
	}

	userBill := &models.UserBill{}
	brokerage, err := userBill.GetBrokerage(uid)
	if err != nil {
		handleErr(c, err)
		return
	}
	recharge, err := userBill.GetRecharge(uid)
	if err != nil {
		handleErr(c, err)
		return
	}

	orderStatusSum, err := storeOrder.GetOrderStatusSum(uid)

	if err != nil {
		handleErr(c, err)
		return
	}

	userExtract := &models.UserExtract{}
	extractTotalPrice, err := userExtract.UserExtractTotalPrice(uid)

	if err != nil {
		handleErr(c, err)
		return
	}

	if brokerage > extractTotalPrice {
		orderYuePrice, err := storeOrder.GetOrderStatusYueSum(uid)
		if err != nil {
			handleErr(c, err)
			return
		}
		systemAdd, err := userBill.GetSystemAdd(uid)
		if err != nil {
			handleErr(c, err)
			return
		}
		yueCount := recharge + systemAdd
		if yueCount > orderYuePrice {
			orderYuePrice = 0
		} else {
			orderYuePrice = orderYuePrice - yueCount
		}
		brokerage = brokerage - extractTotalPrice
		extract_price, err := userExtract.UserExtractTotalPriceByStatus(uid, 0)
		if err != nil {
			handleErr(c, err)
			return
		}

		if extract_price < brokerage {
			brokerage = brokerage - extract_price
		} else {
			brokerage = 0
		}

		if brokerage > orderYuePrice {
			brokerage = brokerage - orderYuePrice
		} else {
			brokerage = 0
		}

	} else {
		brokerage = 0
	}
	var extractPrice float64
	if brokerage-extractTotalPrice > 0 {
		extractPrice = brokerage - extractTotalPrice
	} else {
		extractPrice = 0
	}

	userLevel := &models.UserLevel{}
	err = userLevel.GetLevelInfo(uid, 0)
	if err != nil {
		if err.Error() != "record not found" {
			handleErr(c, err)
			return
		}

	}

	fmt.Println("userLevel:", userLevel)

	user := &models.User{Uid: uid}
	ruser, err := user.GetByUid(uid)
	if err != nil {
		handleErr(c, err)
		return
	}

	mdata := make(map[string]interface{}, 0)
	mdata["couponCount"] = couponCount
	mdata["extractPrice"] = extractPrice

	mdata["like"] = like
	mdata["orderStatusNum"] = orderStatusNum
	mdata["brokerage"] = brokerage
	mdata["recharge"] = recharge
	mdata["orderStatusSum"] = orderStatusSum
	mdata["extractTotalPrice"] = extractTotalPrice
	mdata["nickname"] = ruser.Nickname
	mdata["avatar"] = ruser.Avatar
	mdata["now_money"] = ruser.NowMoney

	mdata["uid"] = ruser.Uid
	handleOk(c, mdata)

}

func (ctl *UserController) MyUserInfo(c *gin.Context) {
	uidT, ok := c.Get("uid")
	if !ok {
		handleErr(c, errors.New("无效的uid"))
		return
	}
	uid := uidT.(int)

	user := &models.User{}
	ruser, err := user.GetByUid(uid)
	if err != nil {
		handleErr(c, err)
		return
	}

	// @Todo 是否统计签到
	// @Todo 是否统计积分使用情况
	handleOk(c, ruser)
	//_ = ruser
	// @todo
}

// 获取活动是否存在
func (ctl *UserController) GetActivity(c *gin.Context) {
	storeBargain := &models.StoreBargain{}
	isBargin, err := storeBargain.ValidBargain(0)
	if err != nil {
		handleErr(c, err)
		return
	}
	storeCombination := &models.StoreCombination{}
	isPink, err := storeCombination.GetPinkIsOpen()
	if err != nil {
		handleErr(c, err)
		return
	}
	mdata := make(map[string]interface{}, 0)
	mdata["is_bargin"] = isBargin
	mdata["is_pink"] = isPink
	mdata["is_seckill"] = false
	handleOk(c, mdata)
}

// 用户地址列表
func (ctl *UserController) UserAddressList(c *gin.Context) {
	uidT, ok := c.Get("uid")
	if !ok {
		handleErr(c, errors.New("无效的uid"))
		return
	}
	uid := uidT.(int)

	req := new(models.Query)
	err := c.ShouldBindQuery(&req)
	if err != nil {
		return
	}

	userAddress := &models.UserAddress{}
	list, err := userAddress.GetUserValidAddressList(uid, req)
	if err != nil {
		return
	}
	handleOk(c, list)
}

// 修改收货地址
func (ctl *UserController) EditUserAddress(c *gin.Context) {
	uidT, ok := c.Get("uid")
	if !ok {
		handleErr(c, errors.New("无效的uid"))
		return
	}
	uid := uidT.(int)
	mReq := make(map[string]interface{}, 0)
	err := c.ShouldBind(&mReq)
	if err != nil {
		return
	}

	address := mReq["address"].(map[string]interface{})
	detail := mReq["detail"].(string)
	var id int
	if v, ok := mReq["id"].(float64); ok {
		id = int(v)
	}
	if v, ok := mReq["id"].(string); ok {
		id, _ = strconv.Atoi(v)
	}
	iSDefault := int(mReq["is_default"].(float64))
	phone := mReq["phone"].(string)
	realName := mReq["real_name"].(string)

	city := address["city"].(string)
	district := address["district"].(string)
	province := address["province"].(string)

	userAddress := &models.UserAddress{
		Uid:       uid,
		RealName:  realName,
		Phone:     phone,
		Province:  province,
		City:      city,
		District:  district,
		Detail:    detail,
		IsDefault: iSDefault,
		Id:        id,
	}
	if id > 0 {
		_, err := userAddress.GetById(id)
		if err != nil {
			userAddress.Id = 0
			err = userAddress.Insert()
			if err != nil {
				handleErr(c, err)
				return
			}
		} else {
			err = userAddress.UpdateById(id)
			if err != nil {
				handleErr(c, err)
				return
			}
		}
		return
	}

	err = userAddress.Insert()
	if err != nil {
		handleErr(c, err)
		return
	}

	handleOk(c, "ok")
}

// 设置默认地址
func (ctl *UserController) SetUserDefaultAddress(c *gin.Context) {
	uidT, ok := c.Get("uid")
	if !ok {
		handleErr(c, errors.New("无效的uid"))
		return
	}
	uid := uidT.(int)

	addressId := c.Param("id")
	id, err := strconv.Atoi(addressId)
	if err != nil {
		handleErr(c, err)
		return
	}

	userAddress := &models.UserAddress{}
	err = userAddress.SetDefaultAddress(id, uid)
	if err != nil {
		return
	}
	handleOk(c, "ok")
}

// 删除地址
func (ctl *UserController) RemoveUserAddress(c *gin.Context) {
	uidT, ok := c.Get("uid")
	if !ok {
		handleErr(c, errors.New("无效的uid"))
		return
	}
	uid := uidT.(int)

	addressId := c.Param("id")
	id, err := strconv.Atoi(addressId)
	if err != nil {
		handleErr(c, err)
		return
	}

	userAddress := &models.UserAddress{}
	err = userAddress.RemoveUserAddress(id, uid)
	if err != nil {
		return
	}
	handleOk(c, "ok")
}

// 获取一条用户地址
func (ctl *UserController) GetUserAddress(c *gin.Context) {
	uidT, ok := c.Get("uid")
	if !ok {
		handleErr(c, errors.New("无效的uid"))
		return
	}
	uid := uidT.(int)

	addressId := c.Param("id")
	id, err := strconv.Atoi(addressId)
	if err != nil {
		handleErr(c, err)
		return
	}
	userAddress := &models.UserAddress{}
	r, err := userAddress.GetUserAddress(id, uid)
	if err != nil {
		handleErr(c, err)
		return
	}
	handleOk(c, r)

}
