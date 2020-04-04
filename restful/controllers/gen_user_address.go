//generate by gen
package controllers

import (
	"github.com/gin-gonic/gin"
	"goshop/restful/models"
	"net/http"
	"strconv"
)

type UserAddressController struct {
}

// @Summary Create
// @Tags    UserAddress
// @Param body body models.UserAddress true "UserAddress"
// @Success 200 {string} string ""
// @Router /userAddresss [post]
func (ctl *UserAddressController) Create(c *gin.Context) {
	userAddress := models.UserAddress{}
	if err := ParseRequest(c, &userAddress); err != nil {
		return
	}
	if err := userAddress.Insert(); err != nil {
		c.JSON(http.StatusBadGateway, err)
		return
	}
	c.JSON(http.StatusOK, userAddress)
}

// @Summary  Delete
// @Tags     UserAddress
// @Param  userAddressId  path string true "userAddressId"
// @Success 200 {string} string ""
// @Router /userAddresss/{userAddressId} [delete]
func (ctl *UserAddressController) Delete(c *gin.Context) {
	userAddress := models.UserAddress{}
	id := c.Param("userAddressId")
	var err error
	userAddress.Id, err = strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	err = userAddress.Delete()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, nil)
}

// @Summary Put
// @Tags    UserAddress
// @Param body body models.UserAddress true "userAddress"
// @Param  userAddressId path string true "userAddressId"
// @Success 200 {string} string ""
// @Router /userAddresss/{userAddressId} [put]
func (ctl *UserAddressController) Put(c *gin.Context) {
	userAddress := models.UserAddress{}
	id := c.Param("userAddressId")
	var err error
	userAddress.Id, err = strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	if err := ParseRequest(c, &userAddress); err != nil {
		return
	}
	err = userAddress.Update()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, nil)
}

// @Summary Patch
// @Tags    UserAddress
// @Param body body models.UserAddress true "userAddress"
// @Param  userAddressId path string true "userAddressId"
// @Success 200 {string} string ""
// @Router /userAddresss/{userAddressId} [patch]
func (ctl *UserAddressController) Patch(c *gin.Context) {
	userAddress := models.UserAddress{}
	id := c.Param("userAddressId")
	var err error
	userAddress.Id, err = strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	if err := ParseRequest(c, &userAddress); err != nil {
		return
	}
	err = userAddress.Patch()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, nil)
}

// @Summary List
// @Tags    UserAddress
// @Param query query string false "query, ?query=age:>:21,name:like:%jason%"
// @Param order query string false "order, ?order=age:desc,created_at:asc"
// @Param page query int false "page"
// @Param pageSize query int false "pageSize"
// @Success 200 {array} models.UserAddress "userAddress array"
// @Router /userAddresss [get]
func (ctl *UserAddressController) List(c *gin.Context) {
	userAddress := &models.UserAddress{}
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
	userAddresss, total, err := userAddress.List(rawQuery, rawOrder, offset, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"total": total,
		"data":  userAddresss,
	})
}

// @Summary Get
// @Tags    UserAddress
// @Param  userAddressId path string true "userAddressId"
// @Success 200 {object} models.UserAddress "userAddress object"
// @Router /userAddresss/{userAddressId} [get]
func (ctl *UserAddressController) Get(c *gin.Context) {
	userAddress := &models.UserAddress{}
	id := c.Param("userAddressId")

	var err error
	userAddress.Id, err = strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	userAddress, err = userAddress.Get()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, userAddress)
}
