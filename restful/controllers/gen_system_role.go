//generate by gen
package controllers

import (
	"github.com/gin-gonic/gin"
	"goshop/restful/models"
	"net/http"
	"strconv"
)

type SystemRoleController struct {
}

// @Summary Create
// @Tags    SystemRole
// @Param body body models.SystemRole true "SystemRole"
// @Success 200 {string} string ""
// @Router /systemRoles [post]
func (ctl *SystemRoleController) Create(c *gin.Context) {
	systemRole := models.SystemRole{}
	if err := ParseRequest(c, &systemRole); err != nil {
		return
	}
	if err := systemRole.Insert(); err != nil {
		c.JSON(http.StatusBadGateway, err)
		return
	}
	c.JSON(http.StatusOK, systemRole)
}

// @Summary  Delete
// @Tags     SystemRole
// @Param  systemRoleId  path string true "systemRoleId"
// @Success 200 {string} string ""
// @Router /systemRoles/{systemRoleId} [delete]
func (ctl *SystemRoleController) Delete(c *gin.Context) {
	systemRole := models.SystemRole{}
	id := c.Param("systemRoleId")
	var err error
	systemRole.Id, err = strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	err = systemRole.Delete()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, nil)
}

// @Summary Put
// @Tags    SystemRole
// @Param body body models.SystemRole true "systemRole"
// @Param  systemRoleId path string true "systemRoleId"
// @Success 200 {string} string ""
// @Router /systemRoles/{systemRoleId} [put]
func (ctl *SystemRoleController) Put(c *gin.Context) {
	systemRole := models.SystemRole{}
	id := c.Param("systemRoleId")
	var err error
	systemRole.Id, err = strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	if err := ParseRequest(c, &systemRole); err != nil {
		return
	}
	err = systemRole.Update()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, nil)
}

// @Summary Patch
// @Tags    SystemRole
// @Param body body models.SystemRole true "systemRole"
// @Param  systemRoleId path string true "systemRoleId"
// @Success 200 {string} string ""
// @Router /systemRoles/{systemRoleId} [patch]
func (ctl *SystemRoleController) Patch(c *gin.Context) {
	systemRole := models.SystemRole{}
	id := c.Param("systemRoleId")
	var err error
	systemRole.Id, err = strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	if err := ParseRequest(c, &systemRole); err != nil {
		return
	}
	err = systemRole.Patch()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, nil)
}

// @Summary List
// @Tags    SystemRole
// @Param query query string false "query, ?query=age:>:21,name:like:%jason%"
// @Param order query string false "order, ?order=age:desc,created_at:asc"
// @Param page query int false "page"
// @Param pageSize query int false "pageSize"
// @Success 200 {array} models.SystemRole "systemRole array"
// @Router /systemRoles [get]
func (ctl *SystemRoleController) List(c *gin.Context) {
	systemRole := &models.SystemRole{}
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
	systemRoles, total, err := systemRole.List(rawQuery, rawOrder, offset, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"total": total,
		"data":  systemRoles,
	})
}

// @Summary Get
// @Tags    SystemRole
// @Param  systemRoleId path string true "systemRoleId"
// @Success 200 {object} models.SystemRole "systemRole object"
// @Router /systemRoles/{systemRoleId} [get]
func (ctl *SystemRoleController) Get(c *gin.Context) {
	systemRole := &models.SystemRole{}
	id := c.Param("systemRoleId")

	var err error
	systemRole.Id, err = strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	systemRole, err = systemRole.Get()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, systemRole)
}
