//generate by gen
package controllers

import (
	"github.com/gin-gonic/gin"
	"goshop/restful/models"
	"net/http"
	"strconv"
)

type StoreCouponController struct {
}

// @Summary Create
// @Tags    StoreCoupon
// @Param body body models.StoreCoupon true "StoreCoupon"
// @Success 200 {string} string ""
// @Router /storeCoupons [post]
func (ctl *StoreCouponController) Create(c *gin.Context) {
	storeCoupon := models.StoreCoupon{}
	if err := ParseRequest(c, &storeCoupon); err != nil {
		return
	}
	if err := storeCoupon.Insert(); err != nil {
		c.JSON(http.StatusBadGateway, err)
		return
	}
	c.JSON(http.StatusOK, storeCoupon)
}

// @Summary  Delete
// @Tags     StoreCoupon
// @Param  storeCouponId  path string true "storeCouponId"
// @Success 200 {string} string ""
// @Router /storeCoupons/{storeCouponId} [delete]
func (ctl *StoreCouponController) Delete(c *gin.Context) {
	storeCoupon := models.StoreCoupon{}
	id := c.Param("storeCouponId")
	var err error
	storeCoupon.Id, err = strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	err = storeCoupon.Delete()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, nil)
}

// @Summary Put
// @Tags    StoreCoupon
// @Param body body models.StoreCoupon true "storeCoupon"
// @Param  storeCouponId path string true "storeCouponId"
// @Success 200 {string} string ""
// @Router /storeCoupons/{storeCouponId} [put]
func (ctl *StoreCouponController) Put(c *gin.Context) {
	storeCoupon := models.StoreCoupon{}
	id := c.Param("storeCouponId")
	var err error
	storeCoupon.Id, err = strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	if err := ParseRequest(c, &storeCoupon); err != nil {
		return
	}
	err = storeCoupon.Update()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, nil)
}

// @Summary Patch
// @Tags    StoreCoupon
// @Param body body models.StoreCoupon true "storeCoupon"
// @Param  storeCouponId path string true "storeCouponId"
// @Success 200 {string} string ""
// @Router /storeCoupons/{storeCouponId} [patch]
func (ctl *StoreCouponController) Patch(c *gin.Context) {
	storeCoupon := models.StoreCoupon{}
	id := c.Param("storeCouponId")
	var err error
	storeCoupon.Id, err = strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	if err := ParseRequest(c, &storeCoupon); err != nil {
		return
	}
	err = storeCoupon.Patch()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, nil)
}

// @Summary List
// @Tags    StoreCoupon
// @Param query query string false "query, ?query=age:>:21,name:like:%jason%"
// @Param order query string false "order, ?order=age:desc,created_at:asc"
// @Param page query int false "page"
// @Param pageSize query int false "pageSize"
// @Success 200 {array} models.StoreCoupon "storeCoupon array"
// @Router /storeCoupons [get]
func (ctl *StoreCouponController) List(c *gin.Context) {
	storeCoupon := &models.StoreCoupon{}
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
	storeCoupons, total, err := storeCoupon.List(rawQuery, rawOrder, offset, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"total": total,
		"data":  storeCoupons,
	})
}

// @Summary Get
// @Tags    StoreCoupon
// @Param  storeCouponId path string true "storeCouponId"
// @Success 200 {object} models.StoreCoupon "storeCoupon object"
// @Router /storeCoupons/{storeCouponId} [get]
func (ctl *StoreCouponController) Get(c *gin.Context) {
	storeCoupon := &models.StoreCoupon{}
	id := c.Param("storeCouponId")

	var err error
	storeCoupon.Id, err = strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	storeCoupon, err = storeCoupon.Get()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, storeCoupon)
}
