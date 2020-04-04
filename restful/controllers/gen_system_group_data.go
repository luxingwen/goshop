//generate by gen
package controllers

import (
	"github.com/gin-gonic/gin"
	"goshop/restful/models"
	"net/http"
	"strconv"
)

type SystemGroupDataController struct {
}

// @Summary Create
// @Tags    SystemGroupData
// @Param body body models.SystemGroupData true "SystemGroupData"
// @Success 200 {string} string ""
// @Router /systemGroupDatas [post]
func (ctl *SystemGroupDataController) Create(c *gin.Context) {
	systemGroupData := models.SystemGroupData{}
	if err := ParseRequest(c, &systemGroupData); err != nil {
		return
	}
	if err := systemGroupData.Insert(); err != nil {
		c.JSON(http.StatusBadGateway, err)
		return
	}
	c.JSON(http.StatusOK, systemGroupData)
}

// @Summary  Delete
// @Tags     SystemGroupData
// @Param  systemGroupDataId  path string true "systemGroupDataId"
// @Success 200 {string} string ""
// @Router /systemGroupDatas/{systemGroupDataId} [delete]
func (ctl *SystemGroupDataController) Delete(c *gin.Context) {
	systemGroupData := models.SystemGroupData{}
	id := c.Param("systemGroupDataId")
	var err error
	systemGroupData.Id, err = strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	err = systemGroupData.Delete()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, nil)
}

// @Summary Put
// @Tags    SystemGroupData
// @Param body body models.SystemGroupData true "systemGroupData"
// @Param  systemGroupDataId path string true "systemGroupDataId"
// @Success 200 {string} string ""
// @Router /systemGroupDatas/{systemGroupDataId} [put]
func (ctl *SystemGroupDataController) Put(c *gin.Context) {
	systemGroupData := models.SystemGroupData{}
	id := c.Param("systemGroupDataId")
	var err error
	systemGroupData.Id, err = strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	if err := ParseRequest(c, &systemGroupData); err != nil {
		return
	}
	err = systemGroupData.Update()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, nil)
}

// @Summary Patch
// @Tags    SystemGroupData
// @Param body body models.SystemGroupData true "systemGroupData"
// @Param  systemGroupDataId path string true "systemGroupDataId"
// @Success 200 {string} string ""
// @Router /systemGroupDatas/{systemGroupDataId} [patch]
func (ctl *SystemGroupDataController) Patch(c *gin.Context) {
	systemGroupData := models.SystemGroupData{}
	id := c.Param("systemGroupDataId")
	var err error
	systemGroupData.Id, err = strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	if err := ParseRequest(c, &systemGroupData); err != nil {
		return
	}
	err = systemGroupData.Patch()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, nil)
}

// @Summary List
// @Tags    SystemGroupData
// @Param query query string false "query, ?query=age:>:21,name:like:%jason%"
// @Param order query string false "order, ?order=age:desc,created_at:asc"
// @Param page query int false "page"
// @Param pageSize query int false "pageSize"
// @Success 200 {array} models.SystemGroupData "systemGroupData array"
// @Router /systemGroupDatas [get]
func (ctl *SystemGroupDataController) List(c *gin.Context) {
	systemGroupData := &models.SystemGroupData{}
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
	systemGroupDatas, total, err := systemGroupData.List(rawQuery, rawOrder, offset, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"total": total,
		"data":  systemGroupDatas,
	})
}

// @Summary Get
// @Tags    SystemGroupData
// @Param  systemGroupDataId path string true "systemGroupDataId"
// @Success 200 {object} models.SystemGroupData "systemGroupData object"
// @Router /systemGroupDatas/{systemGroupDataId} [get]
func (ctl *SystemGroupDataController) Get(c *gin.Context) {
	systemGroupData := &models.SystemGroupData{}
	id := c.Param("systemGroupDataId")

	var err error
	systemGroupData.Id, err = strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	systemGroupData, err = systemGroupData.Get()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, systemGroupData)
}
