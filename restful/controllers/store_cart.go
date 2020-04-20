package controllers

import (
	"goshop/restful/models"

	"github.com/gin-gonic/gin"
)

//获取购物车数量
func (ctl *StoreCartController) GetCartNum(c *gin.Context) {
	var uid int
	uidT, ok := c.Get("uid")
	if ok {
		uid = uidT.(int)
	}

	storeCart := &models.StoreCart{}
	count, err := storeCart.GetUserCartNum(uid, "product")
	if err != nil {
		handleErr(c, err)
		return
	}
	mdata := make(map[string]interface{}, 0)
	mdata["count"] = count
	handleOk(c, mdata)
}
