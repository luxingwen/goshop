//generate by gen
package controllers

import (
	"github.com/gin-gonic/gin"
	"goshop/restful/models"
	"net/http"
	"strconv"
)

type StoreProductController struct {
}

// @Summary Create
// @Tags    StoreProduct
// @Param body body models.StoreProduct true "StoreProduct"
// @Success 200 {string} string ""
// @Router /storeProducts [post]
func (ctl *StoreProductController) Create(c *gin.Context) {
	storeProduct := models.StoreProduct{}
	if err := ParseRequest(c, &storeProduct); err != nil {
		return
	}
	if err := storeProduct.Insert(); err != nil {
		c.JSON(http.StatusBadGateway, err)
		return
	}
	c.JSON(http.StatusOK, storeProduct)
}

// @Summary  Delete
// @Tags     StoreProduct
// @Param  storeProductId  path string true "storeProductId"
// @Success 200 {string} string ""
// @Router /storeProducts/{storeProductId} [delete]
func (ctl *StoreProductController) Delete(c *gin.Context) {
	storeProduct := models.StoreProduct{}
	id := c.Param("storeProductId")
	var err error
	storeProduct.Id, err = strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	err = storeProduct.Delete()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, nil)
}

// @Summary Put
// @Tags    StoreProduct
// @Param body body models.StoreProduct true "storeProduct"
// @Param  storeProductId path string true "storeProductId"
// @Success 200 {string} string ""
// @Router /storeProducts/{storeProductId} [put]
func (ctl *StoreProductController) Put(c *gin.Context) {
	storeProduct := models.StoreProduct{}
	id := c.Param("storeProductId")
	var err error
	storeProduct.Id, err = strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	if err := ParseRequest(c, &storeProduct); err != nil {
		return
	}
	err = storeProduct.Update()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, nil)
}

// @Summary Patch
// @Tags    StoreProduct
// @Param body body models.StoreProduct true "storeProduct"
// @Param  storeProductId path string true "storeProductId"
// @Success 200 {string} string ""
// @Router /storeProducts/{storeProductId} [patch]
func (ctl *StoreProductController) Patch(c *gin.Context) {
	storeProduct := models.StoreProduct{}
	id := c.Param("storeProductId")
	var err error
	storeProduct.Id, err = strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	if err := ParseRequest(c, &storeProduct); err != nil {
		return
	}
	err = storeProduct.Patch()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, nil)
}

// @Summary List
// @Tags    StoreProduct
// @Param query query string false "query, ?query=age:>:21,name:like:%jason%"
// @Param order query string false "order, ?order=age:desc,created_at:asc"
// @Param page query int false "page"
// @Param pageSize query int false "pageSize"
// @Success 200 {array} models.StoreProduct "storeProduct array"
// @Router /storeProducts [get]
func (ctl *StoreProductController) List(c *gin.Context) {
	storeProduct := &models.StoreProduct{}
	var err error
	pageParam := c.DefaultQuery("page", "-1")
	pageSizeParam := c.DefaultQuery("pageSize", "-1")
	rawQuery := c.DefaultQuery("query", "")
	rawOrder := c.DefaultQuery("order", "")
	pageInt, err := strconv.Atoi(pageParam)
	pageSizeInt, err := strconv.Atoi(pageSizeParam)
	offset := pageInt*pageSizeInt - pageSizeInt
	limit := pageSizeInt
	if pageInt < 0 || pageSizeInt < 0 {
		limit = -1
	}
	storeProducts, total, err := storeProduct.List(rawQuery, rawOrder, offset, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"total": total,
		"data":  storeProducts,
	})
}

// @Summary Get
// @Tags    StoreProduct
// @Param  storeProductId path string true "storeProductId"
// @Success 200 {object} models.StoreProduct "storeProduct object"
// @Router /storeProducts/{storeProductId} [get]
func (ctl *StoreProductController) Get(c *gin.Context) {
	storeProduct := &models.StoreProduct{}
	id := c.Param("storeProductId")

	var err error
	storeProduct.Id, err = strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	storeProduct, err = storeProduct.Get()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, storeProduct)
}
