//generate by gen
package controllers

import (
	"github.com/gin-gonic/gin"
	"goshop/restful/models"
	"net/http"
	"strconv"
)

type SystemNoticeController struct {
}

// @Summary Create
// @Tags    SystemNotice
// @Param body body models.SystemNotice true "SystemNotice"
// @Success 200 {string} string ""
// @Router /systemNotices [post]
func (ctl *SystemNoticeController) Create(c *gin.Context) {
	systemNotice := models.SystemNotice{}
	if err := ParseRequest(c, &systemNotice); err != nil {
		return
	}
	if err := systemNotice.Insert(); err != nil {
		c.JSON(http.StatusBadGateway, err)
		return
	}
	c.JSON(http.StatusOK, systemNotice)
}

// @Summary  Delete
// @Tags     SystemNotice
// @Param  systemNoticeId  path string true "systemNoticeId"
// @Success 200 {string} string ""
// @Router /systemNotices/{systemNoticeId} [delete]
func (ctl *SystemNoticeController) Delete(c *gin.Context) {
	systemNotice := models.SystemNotice{}
	id := c.Param("systemNoticeId")
	var err error
	systemNotice.Id, err = strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	err = systemNotice.Delete()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, nil)
}

// @Summary Put
// @Tags    SystemNotice
// @Param body body models.SystemNotice true "systemNotice"
// @Param  systemNoticeId path string true "systemNoticeId"
// @Success 200 {string} string ""
// @Router /systemNotices/{systemNoticeId} [put]
func (ctl *SystemNoticeController) Put(c *gin.Context) {
	systemNotice := models.SystemNotice{}
	id := c.Param("systemNoticeId")
	var err error
	systemNotice.Id, err = strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	if err := ParseRequest(c, &systemNotice); err != nil {
		return
	}
	err = systemNotice.Update()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, nil)
}

// @Summary Patch
// @Tags    SystemNotice
// @Param body body models.SystemNotice true "systemNotice"
// @Param  systemNoticeId path string true "systemNoticeId"
// @Success 200 {string} string ""
// @Router /systemNotices/{systemNoticeId} [patch]
func (ctl *SystemNoticeController) Patch(c *gin.Context) {
	systemNotice := models.SystemNotice{}
	id := c.Param("systemNoticeId")
	var err error
	systemNotice.Id, err = strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	if err := ParseRequest(c, &systemNotice); err != nil {
		return
	}
	err = systemNotice.Patch()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, nil)
}

// @Summary List
// @Tags    SystemNotice
// @Param query query string false "query, ?query=age:>:21,name:like:%jason%"
// @Param order query string false "order, ?order=age:desc,created_at:asc"
// @Param page query int false "page"
// @Param pageSize query int false "pageSize"
// @Success 200 {array} models.SystemNotice "systemNotice array"
// @Router /systemNotices [get]
func (ctl *SystemNoticeController) List(c *gin.Context) {
	systemNotice := &models.SystemNotice{}
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
	systemNotices, total, err := systemNotice.List(rawQuery, rawOrder, offset, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"total": total,
		"data":  systemNotices,
	})
}

// @Summary Get
// @Tags    SystemNotice
// @Param  systemNoticeId path string true "systemNoticeId"
// @Success 200 {object} models.SystemNotice "systemNotice object"
// @Router /systemNotices/{systemNoticeId} [get]
func (ctl *SystemNoticeController) Get(c *gin.Context) {
	systemNotice := &models.SystemNotice{}
	id := c.Param("systemNoticeId")

	var err error
	systemNotice.Id, err = strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	systemNotice, err = systemNotice.Get()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, systemNotice)
}
