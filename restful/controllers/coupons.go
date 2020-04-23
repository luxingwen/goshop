package controllers

import (
	"errors"
	"fmt"
	"strconv"

	"goshop/restful/models"

	"github.com/gin-gonic/gin"
)

type CouponsController struct {
}

func (ctl *CouponsController) IssueCouponList(c *gin.Context) {
	req := new(models.Query)
	err := c.ShouldBindQuery(req)
	if err != nil {
		handleErr(c, err)
		return
	}

	var uid int
	uidT, ok := c.Get("uid")
	if ok {
		uid = uidT.(int)

	}

	storeCouponIssue := &models.StoreCouponIssue{}
	list, err := storeCouponIssue.GetIssueCouponList(uid, req)
	if err != nil {
		handleErr(c, err)
		return
	}
	handleOk(c, list)
}

func (ctl *CouponsController) GetUseCoupons(c *gin.Context) {
	uidT, ok := c.Get("uid")
	if !ok {
		handleErr(c, errors.New("无效的uid"))
		return
	}
	uid := uidT.(int)
	fmt.Println("uid ==> ", uid)
	if uid <= 0 {
		handleErr(c, errors.New("无效的uid"))
		return
	}
	typ := c.Param("typ")
	storeCouponUser := &models.StoreCouponUser{}
	if typ == "0" {
		list, err := storeCouponUser.GetUserAllCoupon(uid)
		if err != nil {
			handleErr(c, err)
			return
		}
		handleOk(c, list)
		return
	}
	if typ == "1" {
		list, err := storeCouponUser.GetUserValidCoupon(uid)
		if err != nil {
			handleErr(c, err)
			return
		}
		handleOk(c, list)
		return
	}
	if typ == "2" {
		list, err := storeCouponUser.GetUserAlreadyUsedCoupon(uid)
		if err != nil {
			handleErr(c, err)
			return
		}
		handleOk(c, list)
		return
	}
	list, err := storeCouponUser.GetUserBeOverdueCoupon(uid)
	if err != nil {
		handleErr(c, err)
		return
	}
	handleOk(c, list)
	return
}

//领取优惠券
func (ctl *CouponsController) UserGetCoupon(c *gin.Context) {
	uidT, ok := c.Get("uid")
	if !ok {
		handleErr(c, errors.New("无效的uid"))
		return
	}
	uid := uidT.(int)

	couposeId := c.Param("id")
	id, err := strconv.Atoi(couposeId)
	if err != nil {
		handleErr(c, err)
		return
	}

	storeCouponIssue := &models.StoreCouponIssue{}
	err = storeCouponIssue.IssueUserCoupon(id, uid)
	if err != nil {
		handleErr(c, err)
		return
	}
	handleOk(c, "ok")
}
