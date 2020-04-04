//generate by gen
package controllers

import (
	"github.com/gin-gonic/gin"
	"goshop/restful/models"
	"net/http"
	"strconv"
)

type StoreCouponUserController struct {
}

// @Summary Create
// @Tags    StoreCouponUser
// @Param body body models.StoreCouponUser true "StoreCouponUser"
// @Success 200 {string} string ""
// @Router /storeCouponUsers [post]
func (ctl *StoreCouponUserController) Create(c *gin.Context) {
	storeCouponUser := models.StoreCouponUser{}
	if err := ParseRequest(c, &storeCouponUser); err != nil {
		return
	}
	if err := storeCouponUser.Insert(); err != nil {
		c.JSON(http.StatusBadGateway, err)
		return
	}
	c.JSON(http.StatusOK, storeCouponUser)
}

// @Summary  Delete
// @Tags     StoreCouponUser
// @Param  storeCouponUserId  path string true "storeCouponUserId"
// @Success 200 {string} string ""
// @Router /storeCouponUsers/{storeCouponUserId} [delete]
func (ctl *StoreCouponUserController) Delete(c *gin.Context) {
	storeCouponUser := models.StoreCouponUser{}
	id := c.Param("storeCouponUserId")
	var err error
	storeCouponUser.Id, err = strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	err = storeCouponUser.Delete()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, nil)
}

// @Summary Put
// @Tags    StoreCouponUser
// @Param body body models.StoreCouponUser true "storeCouponUser"
// @Param  storeCouponUserId path string true "storeCouponUserId"
// @Success 200 {string} string ""
// @Router /storeCouponUsers/{storeCouponUserId} [put]
func (ctl *StoreCouponUserController) Put(c *gin.Context) {
	storeCouponUser := models.StoreCouponUser{}
	id := c.Param("storeCouponUserId")
	var err error
	storeCouponUser.Id, err = strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	if err := ParseRequest(c, &storeCouponUser); err != nil {
		return
	}
	err = storeCouponUser.Update()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, nil)
}

// @Summary Patch
// @Tags    StoreCouponUser
// @Param body body models.StoreCouponUser true "storeCouponUser"
// @Param  storeCouponUserId path string true "storeCouponUserId"
// @Success 200 {string} string ""
// @Router /storeCouponUsers/{storeCouponUserId} [patch]
func (ctl *StoreCouponUserController) Patch(c *gin.Context) {
	storeCouponUser := models.StoreCouponUser{}
	id := c.Param("storeCouponUserId")
	var err error
	storeCouponUser.Id, err = strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	if err := ParseRequest(c, &storeCouponUser); err != nil {
		return
	}
	err = storeCouponUser.Patch()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, nil)
}

// @Summary List
// @Tags    StoreCouponUser
// @Param query query string false "query, ?query=age:>:21,name:like:%jason%"
// @Param order query string false "order, ?order=age:desc,created_at:asc"
// @Param page query int false "page"
// @Param pageSize query int false "pageSize"
// @Success 200 {array} models.StoreCouponUser "storeCouponUser array"
// @Router /storeCouponUsers [get]
func (ctl *StoreCouponUserController) List(c *gin.Context) {
	storeCouponUser := &models.StoreCouponUser{}
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
	storeCouponUsers, total, err := storeCouponUser.List(rawQuery, rawOrder, offset, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"total": total,
		"data":  storeCouponUsers,
	})
}

// @Summary Get
// @Tags    StoreCouponUser
// @Param  storeCouponUserId path string true "storeCouponUserId"
// @Success 200 {object} models.StoreCouponUser "storeCouponUser object"
// @Router /storeCouponUsers/{storeCouponUserId} [get]
func (ctl *StoreCouponUserController) Get(c *gin.Context) {
	storeCouponUser := &models.StoreCouponUser{}
	id := c.Param("storeCouponUserId")

	var err error
	storeCouponUser.Id, err = strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	storeCouponUser, err = storeCouponUser.Get()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, storeCouponUser)
}
