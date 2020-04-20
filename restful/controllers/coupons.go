package controllers

import (
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
