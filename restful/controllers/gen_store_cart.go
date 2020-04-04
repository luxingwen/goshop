//generate by gen
package controllers

import (
	"github.com/gin-gonic/gin"
	"goshop/restful/models"
	"net/http"
	"strconv"
)

type StoreCartController struct {
}

// @Summary Create
// @Tags    StoreCart
// @Param body body models.StoreCart true "StoreCart"
// @Success 200 {string} string ""
// @Router /storeCarts [post]
func (ctl *StoreCartController) Create(c *gin.Context) {
	storeCart := models.StoreCart{}
	if err := ParseRequest(c, &storeCart); err != nil {
		return
	}
	if err := storeCart.Insert(); err != nil {
		c.JSON(http.StatusBadGateway, err)
		return
	}
	c.JSON(http.StatusOK, storeCart)
}

// @Summary  Delete
// @Tags     StoreCart
// @Param  storeCartId  path string true "storeCartId"
// @Success 200 {string} string ""
// @Router /storeCarts/{storeCartId} [delete]
func (ctl *StoreCartController) Delete(c *gin.Context) {
	storeCart := models.StoreCart{}
	id := c.Param("storeCartId")
	var err error
	storeCart.Id, err = strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	err = storeCart.Delete()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, nil)
}

// @Summary Put
// @Tags    StoreCart
// @Param body body models.StoreCart true "storeCart"
// @Param  storeCartId path string true "storeCartId"
// @Success 200 {string} string ""
// @Router /storeCarts/{storeCartId} [put]
func (ctl *StoreCartController) Put(c *gin.Context) {
	storeCart := models.StoreCart{}
	id := c.Param("storeCartId")
	var err error
	storeCart.Id, err = strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	if err := ParseRequest(c, &storeCart); err != nil {
		return
	}
	err = storeCart.Update()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, nil)
}

// @Summary Patch
// @Tags    StoreCart
// @Param body body models.StoreCart true "storeCart"
// @Param  storeCartId path string true "storeCartId"
// @Success 200 {string} string ""
// @Router /storeCarts/{storeCartId} [patch]
func (ctl *StoreCartController) Patch(c *gin.Context) {
	storeCart := models.StoreCart{}
	id := c.Param("storeCartId")
	var err error
	storeCart.Id, err = strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	if err := ParseRequest(c, &storeCart); err != nil {
		return
	}
	err = storeCart.Patch()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, nil)
}

// @Summary List
// @Tags    StoreCart
// @Param query query string false "query, ?query=age:>:21,name:like:%jason%"
// @Param order query string false "order, ?order=age:desc,created_at:asc"
// @Param page query int false "page"
// @Param pageSize query int false "pageSize"
// @Success 200 {array} models.StoreCart "storeCart array"
// @Router /storeCarts [get]
func (ctl *StoreCartController) List(c *gin.Context) {
	storeCart := &models.StoreCart{}
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
	storeCarts, total, err := storeCart.List(rawQuery, rawOrder, offset, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"total": total,
		"data":  storeCarts,
	})
}

// @Summary Get
// @Tags    StoreCart
// @Param  storeCartId path string true "storeCartId"
// @Success 200 {object} models.StoreCart "storeCart object"
// @Router /storeCarts/{storeCartId} [get]
func (ctl *StoreCartController) Get(c *gin.Context) {
	storeCart := &models.StoreCart{}
	id := c.Param("storeCartId")

	var err error
	storeCart.Id, err = strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	storeCart, err = storeCart.Get()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, storeCart)
}
