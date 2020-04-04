//generate by gen
package controllers

import (
	"github.com/gin-gonic/gin"
	"goshop/restful/models"
	"net/http"
	"strconv"
)

type SystemAdminController struct {
}

// @Summary Create
// @Tags    SystemAdmin
// @Param body body models.SystemAdmin true "SystemAdmin"
// @Success 200 {string} string ""
// @Router /systemAdmins [post]
func (ctl *SystemAdminController) Create(c *gin.Context) {
	systemAdmin := models.SystemAdmin{}
	if err := ParseRequest(c, &systemAdmin); err != nil {
		return
	}
	if err := systemAdmin.Insert(); err != nil {
		c.JSON(http.StatusBadGateway, err)
		return
	}
	c.JSON(http.StatusOK, systemAdmin)
}

// @Summary  Delete
// @Tags     SystemAdmin
// @Param  systemAdminId  path string true "systemAdminId"
// @Success 200 {string} string ""
// @Router /systemAdmins/{systemAdminId} [delete]
func (ctl *SystemAdminController) Delete(c *gin.Context) {
	systemAdmin := models.SystemAdmin{}
	id := c.Param("systemAdminId")
	var err error
	systemAdmin.Id, err = strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	err = systemAdmin.Delete()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, nil)
}

// @Summary Put
// @Tags    SystemAdmin
// @Param body body models.SystemAdmin true "systemAdmin"
// @Param  systemAdminId path string true "systemAdminId"
// @Success 200 {string} string ""
// @Router /systemAdmins/{systemAdminId} [put]
func (ctl *SystemAdminController) Put(c *gin.Context) {
	systemAdmin := models.SystemAdmin{}
	id := c.Param("systemAdminId")
	var err error
	systemAdmin.Id, err = strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	if err := ParseRequest(c, &systemAdmin); err != nil {
		return
	}
	err = systemAdmin.Update()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, nil)
}

// @Summary Patch
// @Tags    SystemAdmin
// @Param body body models.SystemAdmin true "systemAdmin"
// @Param  systemAdminId path string true "systemAdminId"
// @Success 200 {string} string ""
// @Router /systemAdmins/{systemAdminId} [patch]
func (ctl *SystemAdminController) Patch(c *gin.Context) {
	systemAdmin := models.SystemAdmin{}
	id := c.Param("systemAdminId")
	var err error
	systemAdmin.Id, err = strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	if err := ParseRequest(c, &systemAdmin); err != nil {
		return
	}
	err = systemAdmin.Patch()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, nil)
}

// @Summary List
// @Tags    SystemAdmin
// @Param query query string false "query, ?query=age:>:21,name:like:%jason%"
// @Param order query string false "order, ?order=age:desc,created_at:asc"
// @Param page query int false "page"
// @Param pageSize query int false "pageSize"
// @Success 200 {array} models.SystemAdmin "systemAdmin array"
// @Router /systemAdmins [get]
func (ctl *SystemAdminController) List(c *gin.Context) {
	systemAdmin := &models.SystemAdmin{}
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
	systemAdmins, total, err := systemAdmin.List(rawQuery, rawOrder, offset, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"total": total,
		"data":  systemAdmins,
	})
}

// @Summary Get
// @Tags    SystemAdmin
// @Param  systemAdminId path string true "systemAdminId"
// @Success 200 {object} models.SystemAdmin "systemAdmin object"
// @Router /systemAdmins/{systemAdminId} [get]
func (ctl *SystemAdminController) Get(c *gin.Context) {
	systemAdmin := &models.SystemAdmin{}
	id := c.Param("systemAdminId")

	var err error
	systemAdmin.Id, err = strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	systemAdmin, err = systemAdmin.Get()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, systemAdmin)
}
