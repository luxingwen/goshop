//generate by gen
package controllers

import (
	"github.com/gin-gonic/gin"
	"goshop/restful/models"
	"net/http"
	"strconv"
)

type SystemNoticeAdminController struct {
}

// @Summary Create
// @Tags    SystemNoticeAdmin
// @Param body body models.SystemNoticeAdmin true "SystemNoticeAdmin"
// @Success 200 {string} string ""
// @Router /systemNoticeAdmins [post]
func (ctl *SystemNoticeAdminController) Create(c *gin.Context) {
	systemNoticeAdmin := models.SystemNoticeAdmin{}
	if err := ParseRequest(c, &systemNoticeAdmin); err != nil {
		return
	}
	if err := systemNoticeAdmin.Insert(); err != nil {
		c.JSON(http.StatusBadGateway, err)
		return
	}
	c.JSON(http.StatusOK, systemNoticeAdmin)
}

// @Summary  Delete
// @Tags     SystemNoticeAdmin
// @Param  systemNoticeAdminId  path string true "systemNoticeAdminId"
// @Success 200 {string} string ""
// @Router /systemNoticeAdmins/{systemNoticeAdminId} [delete]
func (ctl *SystemNoticeAdminController) Delete(c *gin.Context) {
	systemNoticeAdmin := models.SystemNoticeAdmin{}
	id := c.Param("systemNoticeAdminId")
	var err error
	systemNoticeAdmin.Id, err = strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	err = systemNoticeAdmin.Delete()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, nil)
}

// @Summary Put
// @Tags    SystemNoticeAdmin
// @Param body body models.SystemNoticeAdmin true "systemNoticeAdmin"
// @Param  systemNoticeAdminId path string true "systemNoticeAdminId"
// @Success 200 {string} string ""
// @Router /systemNoticeAdmins/{systemNoticeAdminId} [put]
func (ctl *SystemNoticeAdminController) Put(c *gin.Context) {
	systemNoticeAdmin := models.SystemNoticeAdmin{}
	id := c.Param("systemNoticeAdminId")
	var err error
	systemNoticeAdmin.Id, err = strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	if err := ParseRequest(c, &systemNoticeAdmin); err != nil {
		return
	}
	err = systemNoticeAdmin.Update()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, nil)
}

// @Summary Patch
// @Tags    SystemNoticeAdmin
// @Param body body models.SystemNoticeAdmin true "systemNoticeAdmin"
// @Param  systemNoticeAdminId path string true "systemNoticeAdminId"
// @Success 200 {string} string ""
// @Router /systemNoticeAdmins/{systemNoticeAdminId} [patch]
func (ctl *SystemNoticeAdminController) Patch(c *gin.Context) {
	systemNoticeAdmin := models.SystemNoticeAdmin{}
	id := c.Param("systemNoticeAdminId")
	var err error
	systemNoticeAdmin.Id, err = strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	if err := ParseRequest(c, &systemNoticeAdmin); err != nil {
		return
	}
	err = systemNoticeAdmin.Patch()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, nil)
}

// @Summary List
// @Tags    SystemNoticeAdmin
// @Param query query string false "query, ?query=age:>:21,name:like:%jason%"
// @Param order query string false "order, ?order=age:desc,created_at:asc"
// @Param page query int false "page"
// @Param pageSize query int false "pageSize"
// @Success 200 {array} models.SystemNoticeAdmin "systemNoticeAdmin array"
// @Router /systemNoticeAdmins [get]
func (ctl *SystemNoticeAdminController) List(c *gin.Context) {
	systemNoticeAdmin := &models.SystemNoticeAdmin{}
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
	systemNoticeAdmins, total, err := systemNoticeAdmin.List(rawQuery, rawOrder, offset, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"total": total,
		"data":  systemNoticeAdmins,
	})
}

// @Summary Get
// @Tags    SystemNoticeAdmin
// @Param  systemNoticeAdminId path string true "systemNoticeAdminId"
// @Success 200 {object} models.SystemNoticeAdmin "systemNoticeAdmin object"
// @Router /systemNoticeAdmins/{systemNoticeAdminId} [get]
func (ctl *SystemNoticeAdminController) Get(c *gin.Context) {
	systemNoticeAdmin := &models.SystemNoticeAdmin{}
	id := c.Param("systemNoticeAdminId")

	var err error
	systemNoticeAdmin.Id, err = strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	systemNoticeAdmin, err = systemNoticeAdmin.Get()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, systemNoticeAdmin)
}
