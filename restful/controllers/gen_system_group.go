//generate by gen
package controllers

import (
	"github.com/gin-gonic/gin"
	"goshop/restful/models"
	"net/http"
	"strconv"
)

type SystemGroupController struct {
}

// @Summary Create
// @Tags    SystemGroup
// @Param body body models.SystemGroup true "SystemGroup"
// @Success 200 {string} string ""
// @Router /systemGroups [post]
func (ctl *SystemGroupController) Create(c *gin.Context) {
	systemGroup := models.SystemGroup{}
	if err := ParseRequest(c, &systemGroup); err != nil {
		return
	}
	if err := systemGroup.Insert(); err != nil {
		c.JSON(http.StatusBadGateway, err)
		return
	}
	c.JSON(http.StatusOK, systemGroup)
}

// @Summary  Delete
// @Tags     SystemGroup
// @Param  systemGroupId  path string true "systemGroupId"
// @Success 200 {string} string ""
// @Router /systemGroups/{systemGroupId} [delete]
func (ctl *SystemGroupController) Delete(c *gin.Context) {
	systemGroup := models.SystemGroup{}
	id := c.Param("systemGroupId")
	var err error
	systemGroup.Id, err = strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	err = systemGroup.Delete()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, nil)
}

// @Summary Put
// @Tags    SystemGroup
// @Param body body models.SystemGroup true "systemGroup"
// @Param  systemGroupId path string true "systemGroupId"
// @Success 200 {string} string ""
// @Router /systemGroups/{systemGroupId} [put]
func (ctl *SystemGroupController) Put(c *gin.Context) {
	systemGroup := models.SystemGroup{}
	id := c.Param("systemGroupId")
	var err error
	systemGroup.Id, err = strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	if err := ParseRequest(c, &systemGroup); err != nil {
		return
	}
	err = systemGroup.Update()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, nil)
}

// @Summary Patch
// @Tags    SystemGroup
// @Param body body models.SystemGroup true "systemGroup"
// @Param  systemGroupId path string true "systemGroupId"
// @Success 200 {string} string ""
// @Router /systemGroups/{systemGroupId} [patch]
func (ctl *SystemGroupController) Patch(c *gin.Context) {
	systemGroup := models.SystemGroup{}
	id := c.Param("systemGroupId")
	var err error
	systemGroup.Id, err = strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	if err := ParseRequest(c, &systemGroup); err != nil {
		return
	}
	err = systemGroup.Patch()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, nil)
}

// @Summary List
// @Tags    SystemGroup
// @Param query query string false "query, ?query=age:>:21,name:like:%jason%"
// @Param order query string false "order, ?order=age:desc,created_at:asc"
// @Param page query int false "page"
// @Param pageSize query int false "pageSize"
// @Success 200 {array} models.SystemGroup "systemGroup array"
// @Router /systemGroups [get]
func (ctl *SystemGroupController) List(c *gin.Context) {
	systemGroup := &models.SystemGroup{}
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
	systemGroups, total, err := systemGroup.List(rawQuery, rawOrder, offset, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"total": total,
		"data":  systemGroups,
	})
}

// @Summary Get
// @Tags    SystemGroup
// @Param  systemGroupId path string true "systemGroupId"
// @Success 200 {object} models.SystemGroup "systemGroup object"
// @Router /systemGroups/{systemGroupId} [get]
func (ctl *SystemGroupController) Get(c *gin.Context) {
	systemGroup := &models.SystemGroup{}
	id := c.Param("systemGroupId")

	var err error
	systemGroup.Id, err = strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	systemGroup, err = systemGroup.Get()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, systemGroup)
}
