//generate by gen
package controllers

import (
	"github.com/gin-gonic/gin"
	"goshop/restful/models"
	"net/http"
	"strconv"
)

type SystemConfigTabController struct {
}

// @Summary Create
// @Tags    SystemConfigTab
// @Param body body models.SystemConfigTab true "SystemConfigTab"
// @Success 200 {string} string ""
// @Router /systemConfigTabs [post]
func (ctl *SystemConfigTabController) Create(c *gin.Context) {
	systemConfigTab := models.SystemConfigTab{}
	if err := ParseRequest(c, &systemConfigTab); err != nil {
		return
	}
	if err := systemConfigTab.Insert(); err != nil {
		c.JSON(http.StatusBadGateway, err)
		return
	}
	c.JSON(http.StatusOK, systemConfigTab)
}

// @Summary  Delete
// @Tags     SystemConfigTab
// @Param  systemConfigTabId  path string true "systemConfigTabId"
// @Success 200 {string} string ""
// @Router /systemConfigTabs/{systemConfigTabId} [delete]
func (ctl *SystemConfigTabController) Delete(c *gin.Context) {
	systemConfigTab := models.SystemConfigTab{}
	id := c.Param("systemConfigTabId")
	var err error
	systemConfigTab.Id, err = strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	err = systemConfigTab.Delete()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, nil)
}

// @Summary Put
// @Tags    SystemConfigTab
// @Param body body models.SystemConfigTab true "systemConfigTab"
// @Param  systemConfigTabId path string true "systemConfigTabId"
// @Success 200 {string} string ""
// @Router /systemConfigTabs/{systemConfigTabId} [put]
func (ctl *SystemConfigTabController) Put(c *gin.Context) {
	systemConfigTab := models.SystemConfigTab{}
	id := c.Param("systemConfigTabId")
	var err error
	systemConfigTab.Id, err = strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	if err := ParseRequest(c, &systemConfigTab); err != nil {
		return
	}
	err = systemConfigTab.Update()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, nil)
}

// @Summary Patch
// @Tags    SystemConfigTab
// @Param body body models.SystemConfigTab true "systemConfigTab"
// @Param  systemConfigTabId path string true "systemConfigTabId"
// @Success 200 {string} string ""
// @Router /systemConfigTabs/{systemConfigTabId} [patch]
func (ctl *SystemConfigTabController) Patch(c *gin.Context) {
	systemConfigTab := models.SystemConfigTab{}
	id := c.Param("systemConfigTabId")
	var err error
	systemConfigTab.Id, err = strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	if err := ParseRequest(c, &systemConfigTab); err != nil {
		return
	}
	err = systemConfigTab.Patch()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, nil)
}

// @Summary List
// @Tags    SystemConfigTab
// @Param query query string false "query, ?query=age:>:21,name:like:%jason%"
// @Param order query string false "order, ?order=age:desc,created_at:asc"
// @Param page query int false "page"
// @Param pageSize query int false "pageSize"
// @Success 200 {array} models.SystemConfigTab "systemConfigTab array"
// @Router /systemConfigTabs [get]
func (ctl *SystemConfigTabController) List(c *gin.Context) {
	systemConfigTab := &models.SystemConfigTab{}
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
	systemConfigTabs, total, err := systemConfigTab.List(rawQuery, rawOrder, offset, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"total": total,
		"data":  systemConfigTabs,
	})
}

// @Summary Get
// @Tags    SystemConfigTab
// @Param  systemConfigTabId path string true "systemConfigTabId"
// @Success 200 {object} models.SystemConfigTab "systemConfigTab object"
// @Router /systemConfigTabs/{systemConfigTabId} [get]
func (ctl *SystemConfigTabController) Get(c *gin.Context) {
	systemConfigTab := &models.SystemConfigTab{}
	id := c.Param("systemConfigTabId")

	var err error
	systemConfigTab.Id, err = strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	systemConfigTab, err = systemConfigTab.Get()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, systemConfigTab)
}
