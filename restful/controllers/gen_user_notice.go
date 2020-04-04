//generate by gen
package controllers

import (
	"github.com/gin-gonic/gin"
	"goshop/restful/models"
	"net/http"
	"strconv"
)

type UserNoticeController struct {
}

// @Summary Create
// @Tags    UserNotice
// @Param body body models.UserNotice true "UserNotice"
// @Success 200 {string} string ""
// @Router /userNotices [post]
func (ctl *UserNoticeController) Create(c *gin.Context) {
	userNotice := models.UserNotice{}
	if err := ParseRequest(c, &userNotice); err != nil {
		return
	}
	if err := userNotice.Insert(); err != nil {
		c.JSON(http.StatusBadGateway, err)
		return
	}
	c.JSON(http.StatusOK, userNotice)
}

// @Summary  Delete
// @Tags     UserNotice
// @Param  userNoticeId  path string true "userNoticeId"
// @Success 200 {string} string ""
// @Router /userNotices/{userNoticeId} [delete]
func (ctl *UserNoticeController) Delete(c *gin.Context) {
	userNotice := models.UserNotice{}
	id := c.Param("userNoticeId")
	var err error
	userNotice.Id, err = strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	err = userNotice.Delete()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, nil)
}

// @Summary Put
// @Tags    UserNotice
// @Param body body models.UserNotice true "userNotice"
// @Param  userNoticeId path string true "userNoticeId"
// @Success 200 {string} string ""
// @Router /userNotices/{userNoticeId} [put]
func (ctl *UserNoticeController) Put(c *gin.Context) {
	userNotice := models.UserNotice{}
	id := c.Param("userNoticeId")
	var err error
	userNotice.Id, err = strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	if err := ParseRequest(c, &userNotice); err != nil {
		return
	}
	err = userNotice.Update()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, nil)
}

// @Summary Patch
// @Tags    UserNotice
// @Param body body models.UserNotice true "userNotice"
// @Param  userNoticeId path string true "userNoticeId"
// @Success 200 {string} string ""
// @Router /userNotices/{userNoticeId} [patch]
func (ctl *UserNoticeController) Patch(c *gin.Context) {
	userNotice := models.UserNotice{}
	id := c.Param("userNoticeId")
	var err error
	userNotice.Id, err = strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	if err := ParseRequest(c, &userNotice); err != nil {
		return
	}
	err = userNotice.Patch()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, nil)
}

// @Summary List
// @Tags    UserNotice
// @Param query query string false "query, ?query=age:>:21,name:like:%jason%"
// @Param order query string false "order, ?order=age:desc,created_at:asc"
// @Param page query int false "page"
// @Param pageSize query int false "pageSize"
// @Success 200 {array} models.UserNotice "userNotice array"
// @Router /userNotices [get]
func (ctl *UserNoticeController) List(c *gin.Context) {
	userNotice := &models.UserNotice{}
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
	userNotices, total, err := userNotice.List(rawQuery, rawOrder, offset, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"total": total,
		"data":  userNotices,
	})
}

// @Summary Get
// @Tags    UserNotice
// @Param  userNoticeId path string true "userNoticeId"
// @Success 200 {object} models.UserNotice "userNotice object"
// @Router /userNotices/{userNoticeId} [get]
func (ctl *UserNoticeController) Get(c *gin.Context) {
	userNotice := &models.UserNotice{}
	id := c.Param("userNoticeId")

	var err error
	userNotice.Id, err = strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	userNotice, err = userNotice.Get()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, userNotice)
}
