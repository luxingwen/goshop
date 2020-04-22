package controllers

import (
	"errors"
	"fmt"
	"strconv"

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

func (crtl *StoreProductController) Details(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		handleErr(c, err)
		return
	}

	storeProduct := &models.StoreProduct{}
	data, err := storeProduct.GetById(id)
	if err != nil {
		handleErr(c, err)
		return
	}
	handleOk(c, data)
}

func (ctl *StoreProductController) HotProductList(c *gin.Context) {
	req := new(models.ReqGoodsSearch)
	err := c.ShouldBindQuery(req)
	if err != nil {
		handleErr(c, err)
		return
	}

	storeProduct := &models.StoreProduct{}
	list, count, err := storeProduct.GetHotProductLoading(req)
	if err != nil {
		handleErr(c, err)
		return
	}

	mdata := make(map[string]interface{}, 0)
	mdata["list"] = list
	mdata["count"] = count
	handleOk(c, mdata)
}

func (ctl *StoreProductController) ProductCollect(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		handleErr(c, err)
		return
	}

	var uid int
	uidT, ok := c.Get("uid")
	if ok {
		uid = uidT.(int)

	}

	storeProductRelation := &models.StoreProductRelation{}
	relation, err := storeProductRelation.IsProductRelation(id, uid, "collect")
	if err != nil {
		handleErr(c, err)
		return
	}

	mdata := make(map[string]interface{}, 0)
	mdata["userCollect"] = relation
	handleOk(c, mdata)

}

// 获取收藏产品
func (ctl *StoreProductController) GetUserCollectProduct(c *gin.Context) {
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

	storeProductRelation := &models.StoreProductRelation{}
	list, err := storeProductRelation.GetUserCollectProduct(uid, req)
	if err != nil {
		handleErr(c, err)
		return
	}

	handleOk(c, list)

}

// 取消收藏
func (ctl *StoreProductController) UncollectProduct(c *gin.Context) {
	uidT, ok := c.Get("uid")
	if !ok {
		handleErr(c, errors.New("无效的uid"))
		return
	}
	uid := uidT.(int)
	if uid <= 0 {
		handleErr(c, errors.New("无效的uid"))
		return
	}

	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		handleErr(c, err)
		return
	}

	storeProductRelation := &models.StoreProductRelation{}

	err = storeProductRelation.UnProductRelation(uid, id, "collect", "product")
	if err != nil {
		handleErr(c, err)
		return
	}
	handleOk(c, "success")

}

// 添加收藏
func (ctl *StoreProductController) CollectProduct(c *gin.Context) {
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

	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		handleErr(c, err)
		return
	}

	storeProductRelation := &models.StoreProductRelation{}

	err = storeProductRelation.ProductRelation(uid, id, "collect", "product")
	if err != nil {
		handleErr(c, err)
		return
	}
	handleOk(c, "收藏成功")

}
