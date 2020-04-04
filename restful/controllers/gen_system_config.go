//generate by gen
package controllers

import (
	"github.com/gin-gonic/gin"
	"goshop/restful/models"
	"net/http"
	"strconv"
)

type SystemConfigController struct {
}

// @Summary Create
// @Tags    SystemConfig
// @Param body body models.SystemConfig true "SystemConfig"
// @Success 200 {string} string ""
// @Router /systemConfigs [post]
func (ctl *SystemConfigController) Create(c *gin.Context) {
	systemConfig := models.SystemConfig{}
	if err := ParseRequest(c, &systemConfig); err != nil {
		return
	}
	if err := systemConfig.Insert(); err != nil {
		c.JSON(http.StatusBadGateway, err)
		return
	}
	c.JSON(http.StatusOK, systemConfig)
}

// @Summary  Delete
// @Tags     SystemConfig
// @Param  systemConfigId  path string true "systemConfigId"
// @Success 200 {string} string ""
// @Router /systemConfigs/{systemConfigId} [delete]
func (ctl *SystemConfigController) Delete(c *gin.Context) {
	systemConfig := models.SystemConfig{}
	id := c.Param("systemConfigId")
	var err error
	systemConfig.Id, err = strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	err = systemConfig.Delete()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, nil)
}

// @Summary Put
// @Tags    SystemConfig
// @Param body body models.SystemConfig true "systemConfig"
// @Param  systemConfigId path string true "systemConfigId"
// @Success 200 {string} string ""
// @Router /systemConfigs/{systemConfigId} [put]
func (ctl *SystemConfigController) Put(c *gin.Context) {
	systemConfig := models.SystemConfig{}
	id := c.Param("systemConfigId")
	var err error
	systemConfig.Id, err = strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	if err := ParseRequest(c, &systemConfig); err != nil {
		return
	}
	err = systemConfig.Update()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, nil)
}

// @Summary Patch
// @Tags    SystemConfig
// @Param body body models.SystemConfig true "systemConfig"
// @Param  systemConfigId path string true "systemConfigId"
// @Success 200 {string} string ""
// @Router /systemConfigs/{systemConfigId} [patch]
func (ctl *SystemConfigController) Patch(c *gin.Context) {
	systemConfig := models.SystemConfig{}
	id := c.Param("systemConfigId")
	var err error
	systemConfig.Id, err = strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	if err := ParseRequest(c, &systemConfig); err != nil {
		return
	}
	err = systemConfig.Patch()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, nil)
}

// @Summary List
// @Tags    SystemConfig
// @Param query query string false "query, ?query=age:>:21,name:like:%jason%"
// @Param order query string false "order, ?order=age:desc,created_at:asc"
// @Param page query int false "page"
// @Param pageSize query int false "pageSize"
// @Success 200 {array} models.SystemConfig "systemConfig array"
// @Router /systemConfigs [get]
func (ctl *SystemConfigController) List(c *gin.Context) {
	systemConfig := &models.SystemConfig{}
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
	systemConfigs, total, err := systemConfig.List(rawQuery, rawOrder, offset, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"total": total,
		"data":  systemConfigs,
	})
}

// @Summary Get
// @Tags    SystemConfig
// @Param  systemConfigId path string true "systemConfigId"
// @Success 200 {object} models.SystemConfig "systemConfig object"
// @Router /systemConfigs/{systemConfigId} [get]
func (ctl *SystemConfigController) Get(c *gin.Context) {
	systemConfig := &models.SystemConfig{}
	id := c.Param("systemConfigId")

	var err error
	systemConfig.Id, err = strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	systemConfig, err = systemConfig.Get()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, systemConfig)
}
