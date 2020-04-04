//generate by gen
package controllers

import (
	"github.com/gin-gonic/gin"
	"goshop/restful/models"
	"net/http"
	"strconv"
)

type UserBillController struct {
}

// @Summary Create
// @Tags    UserBill
// @Param body body models.UserBill true "UserBill"
// @Success 200 {string} string ""
// @Router /userBills [post]
func (ctl *UserBillController) Create(c *gin.Context) {
	userBill := models.UserBill{}
	if err := ParseRequest(c, &userBill); err != nil {
		return
	}
	if err := userBill.Insert(); err != nil {
		c.JSON(http.StatusBadGateway, err)
		return
	}
	c.JSON(http.StatusOK, userBill)
}

// @Summary  Delete
// @Tags     UserBill
// @Param  userBillId  path string true "userBillId"
// @Success 200 {string} string ""
// @Router /userBills/{userBillId} [delete]
func (ctl *UserBillController) Delete(c *gin.Context) {
	userBill := models.UserBill{}
	id := c.Param("userBillId")
	var err error
	userBill.Id, err = strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	err = userBill.Delete()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, nil)
}

// @Summary Put
// @Tags    UserBill
// @Param body body models.UserBill true "userBill"
// @Param  userBillId path string true "userBillId"
// @Success 200 {string} string ""
// @Router /userBills/{userBillId} [put]
func (ctl *UserBillController) Put(c *gin.Context) {
	userBill := models.UserBill{}
	id := c.Param("userBillId")
	var err error
	userBill.Id, err = strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	if err := ParseRequest(c, &userBill); err != nil {
		return
	}
	err = userBill.Update()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, nil)
}

// @Summary Patch
// @Tags    UserBill
// @Param body body models.UserBill true "userBill"
// @Param  userBillId path string true "userBillId"
// @Success 200 {string} string ""
// @Router /userBills/{userBillId} [patch]
func (ctl *UserBillController) Patch(c *gin.Context) {
	userBill := models.UserBill{}
	id := c.Param("userBillId")
	var err error
	userBill.Id, err = strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	if err := ParseRequest(c, &userBill); err != nil {
		return
	}
	err = userBill.Patch()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, nil)
}

// @Summary List
// @Tags    UserBill
// @Param query query string false "query, ?query=age:>:21,name:like:%jason%"
// @Param order query string false "order, ?order=age:desc,created_at:asc"
// @Param page query int false "page"
// @Param pageSize query int false "pageSize"
// @Success 200 {array} models.UserBill "userBill array"
// @Router /userBills [get]
func (ctl *UserBillController) List(c *gin.Context) {
	userBill := &models.UserBill{}
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
	userBills, total, err := userBill.List(rawQuery, rawOrder, offset, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"total": total,
		"data":  userBills,
	})
}

// @Summary Get
// @Tags    UserBill
// @Param  userBillId path string true "userBillId"
// @Success 200 {object} models.UserBill "userBill object"
// @Router /userBills/{userBillId} [get]
func (ctl *UserBillController) Get(c *gin.Context) {
	userBill := &models.UserBill{}
	id := c.Param("userBillId")

	var err error
	userBill.Id, err = strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	userBill, err = userBill.Get()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, userBill)
}
