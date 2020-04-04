//generate by gen
package controllers

import (
	"github.com/gin-gonic/gin"
	"goshop/restful/models"
	"net/http"
	"strconv"
)

type StoreOrderController struct {
}

// @Summary Create
// @Tags    StoreOrder
// @Param body body models.StoreOrder true "StoreOrder"
// @Success 200 {string} string ""
// @Router /storeOrders [post]
func (ctl *StoreOrderController) Create(c *gin.Context) {
	storeOrder := models.StoreOrder{}
	if err := ParseRequest(c, &storeOrder); err != nil {
		return
	}
	if err := storeOrder.Insert(); err != nil {
		c.JSON(http.StatusBadGateway, err)
		return
	}
	c.JSON(http.StatusOK, storeOrder)
}

// @Summary  Delete
// @Tags     StoreOrder
// @Param  storeOrderId  path string true "storeOrderId"
// @Success 200 {string} string ""
// @Router /storeOrders/{storeOrderId} [delete]
func (ctl *StoreOrderController) Delete(c *gin.Context) {
	storeOrder := models.StoreOrder{}
	id := c.Param("storeOrderId")
	var err error
	storeOrder.Id, err = strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	err = storeOrder.Delete()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, nil)
}

// @Summary Put
// @Tags    StoreOrder
// @Param body body models.StoreOrder true "storeOrder"
// @Param  storeOrderId path string true "storeOrderId"
// @Success 200 {string} string ""
// @Router /storeOrders/{storeOrderId} [put]
func (ctl *StoreOrderController) Put(c *gin.Context) {
	storeOrder := models.StoreOrder{}
	id := c.Param("storeOrderId")
	var err error
	storeOrder.Id, err = strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	if err := ParseRequest(c, &storeOrder); err != nil {
		return
	}
	err = storeOrder.Update()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, nil)
}

// @Summary Patch
// @Tags    StoreOrder
// @Param body body models.StoreOrder true "storeOrder"
// @Param  storeOrderId path string true "storeOrderId"
// @Success 200 {string} string ""
// @Router /storeOrders/{storeOrderId} [patch]
func (ctl *StoreOrderController) Patch(c *gin.Context) {
	storeOrder := models.StoreOrder{}
	id := c.Param("storeOrderId")
	var err error
	storeOrder.Id, err = strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	if err := ParseRequest(c, &storeOrder); err != nil {
		return
	}
	err = storeOrder.Patch()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, nil)
}

// @Summary List
// @Tags    StoreOrder
// @Param query query string false "query, ?query=age:>:21,name:like:%jason%"
// @Param order query string false "order, ?order=age:desc,created_at:asc"
// @Param page query int false "page"
// @Param pageSize query int false "pageSize"
// @Success 200 {array} models.StoreOrder "storeOrder array"
// @Router /storeOrders [get]
func (ctl *StoreOrderController) List(c *gin.Context) {
	storeOrder := &models.StoreOrder{}
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
	storeOrders, total, err := storeOrder.List(rawQuery, rawOrder, offset, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"total": total,
		"data":  storeOrders,
	})
}

// @Summary Get
// @Tags    StoreOrder
// @Param  storeOrderId path string true "storeOrderId"
// @Success 200 {object} models.StoreOrder "storeOrder object"
// @Router /storeOrders/{storeOrderId} [get]
func (ctl *StoreOrderController) Get(c *gin.Context) {
	storeOrder := &models.StoreOrder{}
	id := c.Param("storeOrderId")

	var err error
	storeOrder.Id, err = strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	storeOrder, err = storeOrder.Get()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, storeOrder)
}
