//generate by gen
package controllers

import (
	"github.com/gin-gonic/gin"
	"goshop/restful/models"
	"net/http"
	"strconv"
)

type SystemMenusController struct {
}

// @Summary Create
// @Tags    SystemMenus
// @Param body body models.SystemMenus true "SystemMenus"
// @Success 200 {string} string ""
// @Router /systemMenuss [post]
func (ctl *SystemMenusController) Create(c *gin.Context) {
	systemMenus := models.SystemMenus{}
	if err := ParseRequest(c, &systemMenus); err != nil {
		return
	}
	if err := systemMenus.Insert(); err != nil {
		c.JSON(http.StatusBadGateway, err)
		return
	}
	c.JSON(http.StatusOK, systemMenus)
}

// @Summary  Delete
// @Tags     SystemMenus
// @Param  systemMenusId  path string true "systemMenusId"
// @Success 200 {string} string ""
// @Router /systemMenuss/{systemMenusId} [delete]
func (ctl *SystemMenusController) Delete(c *gin.Context) {
	systemMenus := models.SystemMenus{}
	id := c.Param("systemMenusId")
	var err error
	systemMenus.Id, err = strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	err = systemMenus.Delete()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, nil)
}

// @Summary Put
// @Tags    SystemMenus
// @Param body body models.SystemMenus true "systemMenus"
// @Param  systemMenusId path string true "systemMenusId"
// @Success 200 {string} string ""
// @Router /systemMenuss/{systemMenusId} [put]
func (ctl *SystemMenusController) Put(c *gin.Context) {
	systemMenus := models.SystemMenus{}
	id := c.Param("systemMenusId")
	var err error
	systemMenus.Id, err = strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	if err := ParseRequest(c, &systemMenus); err != nil {
		return
	}
	err = systemMenus.Update()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, nil)
}

// @Summary Patch
// @Tags    SystemMenus
// @Param body body models.SystemMenus true "systemMenus"
// @Param  systemMenusId path string true "systemMenusId"
// @Success 200 {string} string ""
// @Router /systemMenuss/{systemMenusId} [patch]
func (ctl *SystemMenusController) Patch(c *gin.Context) {
	systemMenus := models.SystemMenus{}
	id := c.Param("systemMenusId")
	var err error
	systemMenus.Id, err = strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	if err := ParseRequest(c, &systemMenus); err != nil {
		return
	}
	err = systemMenus.Patch()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, nil)
}

// @Summary List
// @Tags    SystemMenus
// @Param query query string false "query, ?query=age:>:21,name:like:%jason%"
// @Param order query string false "order, ?order=age:desc,created_at:asc"
// @Param page query int false "page"
// @Param pageSize query int false "pageSize"
// @Success 200 {array} models.SystemMenus "systemMenus array"
// @Router /systemMenuss [get]
func (ctl *SystemMenusController) List(c *gin.Context) {
	systemMenus := &models.SystemMenus{}
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
	systemMenuss, total, err := systemMenus.List(rawQuery, rawOrder, offset, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"total": total,
		"data":  systemMenuss,
	})
}

// @Summary Get
// @Tags    SystemMenus
// @Param  systemMenusId path string true "systemMenusId"
// @Success 200 {object} models.SystemMenus "systemMenus object"
// @Router /systemMenuss/{systemMenusId} [get]
func (ctl *SystemMenusController) Get(c *gin.Context) {
	systemMenus := &models.SystemMenus{}
	id := c.Param("systemMenusId")

	var err error
	systemMenus.Id, err = strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	systemMenus, err = systemMenus.Get()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, systemMenus)
}
