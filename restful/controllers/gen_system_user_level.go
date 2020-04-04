//generate by gen
package controllers

import (
	"github.com/gin-gonic/gin"
	"goshop/restful/models"
	"net/http"
	"strconv"
)

type SystemUserLevelController struct {
}

// @Summary Create
// @Tags    SystemUserLevel
// @Param body body models.SystemUserLevel true "SystemUserLevel"
// @Success 200 {string} string ""
// @Router /systemUserLevels [post]
func (ctl *SystemUserLevelController) Create(c *gin.Context) {
	systemUserLevel := models.SystemUserLevel{}
	if err := ParseRequest(c, &systemUserLevel); err != nil {
		return
	}
	if err := systemUserLevel.Insert(); err != nil {
		c.JSON(http.StatusBadGateway, err)
		return
	}
	c.JSON(http.StatusOK, systemUserLevel)
}

// @Summary  Delete
// @Tags     SystemUserLevel
// @Param  systemUserLevelId  path string true "systemUserLevelId"
// @Success 200 {string} string ""
// @Router /systemUserLevels/{systemUserLevelId} [delete]
func (ctl *SystemUserLevelController) Delete(c *gin.Context) {
	systemUserLevel := models.SystemUserLevel{}
	id := c.Param("systemUserLevelId")
	var err error
	systemUserLevel.Id, err = strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	err = systemUserLevel.Delete()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, nil)
}

// @Summary Put
// @Tags    SystemUserLevel
// @Param body body models.SystemUserLevel true "systemUserLevel"
// @Param  systemUserLevelId path string true "systemUserLevelId"
// @Success 200 {string} string ""
// @Router /systemUserLevels/{systemUserLevelId} [put]
func (ctl *SystemUserLevelController) Put(c *gin.Context) {
	systemUserLevel := models.SystemUserLevel{}
	id := c.Param("systemUserLevelId")
	var err error
	systemUserLevel.Id, err = strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	if err := ParseRequest(c, &systemUserLevel); err != nil {
		return
	}
	err = systemUserLevel.Update()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, nil)
}

// @Summary Patch
// @Tags    SystemUserLevel
// @Param body body models.SystemUserLevel true "systemUserLevel"
// @Param  systemUserLevelId path string true "systemUserLevelId"
// @Success 200 {string} string ""
// @Router /systemUserLevels/{systemUserLevelId} [patch]
func (ctl *SystemUserLevelController) Patch(c *gin.Context) {
	systemUserLevel := models.SystemUserLevel{}
	id := c.Param("systemUserLevelId")
	var err error
	systemUserLevel.Id, err = strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	if err := ParseRequest(c, &systemUserLevel); err != nil {
		return
	}
	err = systemUserLevel.Patch()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, nil)
}

// @Summary List
// @Tags    SystemUserLevel
// @Param query query string false "query, ?query=age:>:21,name:like:%jason%"
// @Param order query string false "order, ?order=age:desc,created_at:asc"
// @Param page query int false "page"
// @Param pageSize query int false "pageSize"
// @Success 200 {array} models.SystemUserLevel "systemUserLevel array"
// @Router /systemUserLevels [get]
func (ctl *SystemUserLevelController) List(c *gin.Context) {
	systemUserLevel := &models.SystemUserLevel{}
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
	systemUserLevels, total, err := systemUserLevel.List(rawQuery, rawOrder, offset, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"total": total,
		"data":  systemUserLevels,
	})
}

// @Summary Get
// @Tags    SystemUserLevel
// @Param  systemUserLevelId path string true "systemUserLevelId"
// @Success 200 {object} models.SystemUserLevel "systemUserLevel object"
// @Router /systemUserLevels/{systemUserLevelId} [get]
func (ctl *SystemUserLevelController) Get(c *gin.Context) {
	systemUserLevel := &models.SystemUserLevel{}
	id := c.Param("systemUserLevelId")

	var err error
	systemUserLevel.Id, err = strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	systemUserLevel, err = systemUserLevel.Get()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, systemUserLevel)
}
