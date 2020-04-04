//generate by gen
package controllers

import (
	"github.com/gin-gonic/gin"
	"goshop/restful/models"
	"net/http"
	"strconv"
)

type UserRechargeController struct {
}

// @Summary Create
// @Tags    UserRecharge
// @Param body body models.UserRecharge true "UserRecharge"
// @Success 200 {string} string ""
// @Router /userRecharges [post]
func (ctl *UserRechargeController) Create(c *gin.Context) {
	userRecharge := models.UserRecharge{}
	if err := ParseRequest(c, &userRecharge); err != nil {
		return
	}
	if err := userRecharge.Insert(); err != nil {
		c.JSON(http.StatusBadGateway, err)
		return
	}
	c.JSON(http.StatusOK, userRecharge)
}

// @Summary  Delete
// @Tags     UserRecharge
// @Param  userRechargeId  path string true "userRechargeId"
// @Success 200 {string} string ""
// @Router /userRecharges/{userRechargeId} [delete]
func (ctl *UserRechargeController) Delete(c *gin.Context) {
	userRecharge := models.UserRecharge{}
	id := c.Param("userRechargeId")
	var err error
	userRecharge.Id, err = strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	err = userRecharge.Delete()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, nil)
}

// @Summary Put
// @Tags    UserRecharge
// @Param body body models.UserRecharge true "userRecharge"
// @Param  userRechargeId path string true "userRechargeId"
// @Success 200 {string} string ""
// @Router /userRecharges/{userRechargeId} [put]
func (ctl *UserRechargeController) Put(c *gin.Context) {
	userRecharge := models.UserRecharge{}
	id := c.Param("userRechargeId")
	var err error
	userRecharge.Id, err = strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	if err := ParseRequest(c, &userRecharge); err != nil {
		return
	}
	err = userRecharge.Update()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, nil)
}

// @Summary Patch
// @Tags    UserRecharge
// @Param body body models.UserRecharge true "userRecharge"
// @Param  userRechargeId path string true "userRechargeId"
// @Success 200 {string} string ""
// @Router /userRecharges/{userRechargeId} [patch]
func (ctl *UserRechargeController) Patch(c *gin.Context) {
	userRecharge := models.UserRecharge{}
	id := c.Param("userRechargeId")
	var err error
	userRecharge.Id, err = strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	if err := ParseRequest(c, &userRecharge); err != nil {
		return
	}
	err = userRecharge.Patch()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, nil)
}

// @Summary List
// @Tags    UserRecharge
// @Param query query string false "query, ?query=age:>:21,name:like:%jason%"
// @Param order query string false "order, ?order=age:desc,created_at:asc"
// @Param page query int false "page"
// @Param pageSize query int false "pageSize"
// @Success 200 {array} models.UserRecharge "userRecharge array"
// @Router /userRecharges [get]
func (ctl *UserRechargeController) List(c *gin.Context) {
	userRecharge := &models.UserRecharge{}
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
	userRecharges, total, err := userRecharge.List(rawQuery, rawOrder, offset, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"total": total,
		"data":  userRecharges,
	})
}

// @Summary Get
// @Tags    UserRecharge
// @Param  userRechargeId path string true "userRechargeId"
// @Success 200 {object} models.UserRecharge "userRecharge object"
// @Router /userRecharges/{userRechargeId} [get]
func (ctl *UserRechargeController) Get(c *gin.Context) {
	userRecharge := &models.UserRecharge{}
	id := c.Param("userRechargeId")

	var err error
	userRecharge.Id, err = strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	userRecharge, err = userRecharge.Get()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, userRecharge)
}
