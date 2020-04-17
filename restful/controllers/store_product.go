package controllers

import (
	"goshop/restful/models"

	"github.com/gin-gonic/gin"
)

func (ctl *StoreProductController) ProductList(c *gin.Context) {
	req := new(models.ReqStoreProductQuery)
	err := c.ShouldBindQuery(req)
	if err != nil {
		handleErr(c, err)
		return
	}

	storeProduct := &models.StoreProduct{}
	list, count, err := storeProduct.GetProductList(req)
	if err != nil {
		handleErr(c, err)
		return
	}

	mdata := make(map[string]interface{}, 0)
	mdata["list"] = list
	mdata["count"] = count
	handleOk(c, mdata)
}

func (crtl *StoreProductController) GoodsSearch(c *gin.Context) {
	req := new(models.ReqGoodsSearch)
	err := c.ShouldBindQuery(req)
	if err != nil {
		handleErr(c, err)
		return
	}
	storeProduct := &models.StoreProduct{}
	list, count, err := storeProduct.GetSearchStorePage(req, 0)
	if err != nil {
		handleErr(c, err)
		return
	}

	mdata := make(map[string]interface{}, 0)
	mdata["list"] = list
	mdata["count"] = count
	handleOk(c, mdata)
}
