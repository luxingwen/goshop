package controllers

import (
	"github.com/gin-gonic/gin"
	// "goshop/restful/common"
	"goshop/restful/models"
	// "goshop/libs/errcode"
	"errors"
	"fmt"
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
