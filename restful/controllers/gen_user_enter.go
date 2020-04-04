//generate by gen
package controllers

import (
	"github.com/gin-gonic/gin"
	"goshop/restful/models"
	"net/http"
	"strconv"
)

type UserEnterController struct {
}

// @Summary Create
// @Tags    UserEnter
// @Param body body models.UserEnter true "UserEnter"
// @Success 200 {string} string ""
// @Router /userEnters [post]
func (ctl *UserEnterController) Create(c *gin.Context) {
	userEnter := models.UserEnter{}
	if err := ParseRequest(c, &userEnter); err != nil {
		return
	}
	if err := userEnter.Insert(); err != nil {
		c.JSON(http.StatusBadGateway, err)
		return
	}
	c.JSON(http.StatusOK, userEnter)
}

// @Summary  Delete
// @Tags     UserEnter
// @Param  userEnterId  path string true "userEnterId"
// @Success 200 {string} string ""
// @Router /userEnters/{userEnterId} [delete]
func (ctl *UserEnterController) Delete(c *gin.Context) {
	userEnter := models.UserEnter{}
	id := c.Param("userEnterId")
	var err error
	userEnter.Id, err = strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	err = userEnter.Delete()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, nil)
}

// @Summary Put
// @Tags    UserEnter
// @Param body body models.UserEnter true "userEnter"
// @Param  userEnterId path string true "userEnterId"
// @Success 200 {string} string ""
// @Router /userEnters/{userEnterId} [put]
func (ctl *UserEnterController) Put(c *gin.Context) {
	userEnter := models.UserEnter{}
	id := c.Param("userEnterId")
	var err error
	userEnter.Id, err = strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	if err := ParseRequest(c, &userEnter); err != nil {
		return
	}
	err = userEnter.Update()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, nil)
}

// @Summary Patch
// @Tags    UserEnter
// @Param body body models.UserEnter true "userEnter"
// @Param  userEnterId path string true "userEnterId"
// @Success 200 {string} string ""
// @Router /userEnters/{userEnterId} [patch]
func (ctl *UserEnterController) Patch(c *gin.Context) {
	userEnter := models.UserEnter{}
	id := c.Param("userEnterId")
	var err error
	userEnter.Id, err = strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	if err := ParseRequest(c, &userEnter); err != nil {
		return
	}
	err = userEnter.Patch()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, nil)
}

// @Summary List
// @Tags    UserEnter
// @Param query query string false "query, ?query=age:>:21,name:like:%jason%"
// @Param order query string false "order, ?order=age:desc,created_at:asc"
// @Param page query int false "page"
// @Param pageSize query int false "pageSize"
// @Success 200 {array} models.UserEnter "userEnter array"
// @Router /userEnters [get]
func (ctl *UserEnterController) List(c *gin.Context) {
	userEnter := &models.UserEnter{}
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
	userEnters, total, err := userEnter.List(rawQuery, rawOrder, offset, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"total": total,
		"data":  userEnters,
	})
}

// @Summary Get
// @Tags    UserEnter
// @Param  userEnterId path string true "userEnterId"
// @Success 200 {object} models.UserEnter "userEnter object"
// @Router /userEnters/{userEnterId} [get]
func (ctl *UserEnterController) Get(c *gin.Context) {
	userEnter := &models.UserEnter{}
	id := c.Param("userEnterId")

	var err error
	userEnter.Id, err = strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	userEnter, err = userEnter.Get()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, userEnter)
}
