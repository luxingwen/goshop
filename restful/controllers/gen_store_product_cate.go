//generate by gen
package controllers

import (
	"github.com/gin-gonic/gin"
	"goshop/restful/models"
	"net/http"
	"strconv"
)

type StoreProductCateController struct {
}

// @Summary Create
// @Tags    StoreProductCate
// @Param body body models.StoreProductCate true "StoreProductCate"
// @Success 200 {string} string ""
// @Router /storeProductCates [post]
func (ctl *StoreProductCateController) Create(c *gin.Context) {
	storeProductCate := models.StoreProductCate{}
	if err := ParseRequest(c, &storeProductCate); err != nil {
		return
	}
	if err := storeProductCate.Insert(); err != nil {
		c.JSON(http.StatusBadGateway, err)
		return
	}
	c.JSON(http.StatusOK, storeProductCate)
}

// @Summary  Delete
// @Tags     StoreProductCate
// @Param  storeProductCateId  path string true "storeProductCateId"
// @Success 200 {string} string ""
// @Router /storeProductCates/{storeProductCateId} [delete]
func (ctl *StoreProductCateController) Delete(c *gin.Context) {
	storeProductCate := models.StoreProductCate{}
	id := c.Param("storeProductCateId")
	var err error
	storeProductCate.Id, err = strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	err = storeProductCate.Delete()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, nil)
}

// @Summary Put
// @Tags    StoreProductCate
// @Param body body models.StoreProductCate true "storeProductCate"
// @Param  storeProductCateId path string true "storeProductCateId"
// @Success 200 {string} string ""
// @Router /storeProductCates/{storeProductCateId} [put]
func (ctl *StoreProductCateController) Put(c *gin.Context) {
	storeProductCate := models.StoreProductCate{}
	id := c.Param("storeProductCateId")
	var err error
	storeProductCate.Id, err = strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	if err := ParseRequest(c, &storeProductCate); err != nil {
		return
	}
	err = storeProductCate.Update()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, nil)
}

// @Summary Patch
// @Tags    StoreProductCate
// @Param body body models.StoreProductCate true "storeProductCate"
// @Param  storeProductCateId path string true "storeProductCateId"
// @Success 200 {string} string ""
// @Router /storeProductCates/{storeProductCateId} [patch]
func (ctl *StoreProductCateController) Patch(c *gin.Context) {
	storeProductCate := models.StoreProductCate{}
	id := c.Param("storeProductCateId")
	var err error
	storeProductCate.Id, err = strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	if err := ParseRequest(c, &storeProductCate); err != nil {
		return
	}
	err = storeProductCate.Patch()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, nil)
}

// @Summary List
// @Tags    StoreProductCate
// @Param query query string false "query, ?query=age:>:21,name:like:%jason%"
// @Param order query string false "order, ?order=age:desc,created_at:asc"
// @Param page query int false "page"
// @Param pageSize query int false "pageSize"
// @Success 200 {array} models.StoreProductCate "storeProductCate array"
// @Router /storeProductCates [get]
func (ctl *StoreProductCateController) List(c *gin.Context) {
	storeProductCate := &models.StoreProductCate{}
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
	storeProductCates, total, err := storeProductCate.List(rawQuery, rawOrder, offset, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"total": total,
		"data":  storeProductCates,
	})
}

// @Summary Get
// @Tags    StoreProductCate
// @Param  storeProductCateId path string true "storeProductCateId"
// @Success 200 {object} models.StoreProductCate "storeProductCate object"
// @Router /storeProductCates/{storeProductCateId} [get]
func (ctl *StoreProductCateController) Get(c *gin.Context) {
	storeProductCate := &models.StoreProductCate{}
	id := c.Param("storeProductCateId")

	var err error
	storeProductCate.Id, err = strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	storeProductCate, err = storeProductCate.Get()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, storeProductCate)
}
