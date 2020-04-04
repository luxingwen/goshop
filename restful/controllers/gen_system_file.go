//generate by gen
package controllers

import (
	"github.com/gin-gonic/gin"
	"goshop/restful/models"
	"net/http"
	"strconv"
)

type SystemFileController struct {
}

// @Summary Create
// @Tags    SystemFile
// @Param body body models.SystemFile true "SystemFile"
// @Success 200 {string} string ""
// @Router /systemFiles [post]
func (ctl *SystemFileController) Create(c *gin.Context) {
	systemFile := models.SystemFile{}
	if err := ParseRequest(c, &systemFile); err != nil {
		return
	}
	if err := systemFile.Insert(); err != nil {
		c.JSON(http.StatusBadGateway, err)
		return
	}
	c.JSON(http.StatusOK, systemFile)
}

// @Summary  Delete
// @Tags     SystemFile
// @Param  systemFileId  path string true "systemFileId"
// @Success 200 {string} string ""
// @Router /systemFiles/{systemFileId} [delete]
func (ctl *SystemFileController) Delete(c *gin.Context) {
	systemFile := models.SystemFile{}
	id := c.Param("systemFileId")
	var err error
	systemFile.Id, err = strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	err = systemFile.Delete()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, nil)
}

// @Summary Put
// @Tags    SystemFile
// @Param body body models.SystemFile true "systemFile"
// @Param  systemFileId path string true "systemFileId"
// @Success 200 {string} string ""
// @Router /systemFiles/{systemFileId} [put]
func (ctl *SystemFileController) Put(c *gin.Context) {
	systemFile := models.SystemFile{}
	id := c.Param("systemFileId")
	var err error
	systemFile.Id, err = strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	if err := ParseRequest(c, &systemFile); err != nil {
		return
	}
	err = systemFile.Update()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, nil)
}

// @Summary Patch
// @Tags    SystemFile
// @Param body body models.SystemFile true "systemFile"
// @Param  systemFileId path string true "systemFileId"
// @Success 200 {string} string ""
// @Router /systemFiles/{systemFileId} [patch]
func (ctl *SystemFileController) Patch(c *gin.Context) {
	systemFile := models.SystemFile{}
	id := c.Param("systemFileId")
	var err error
	systemFile.Id, err = strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	if err := ParseRequest(c, &systemFile); err != nil {
		return
	}
	err = systemFile.Patch()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, nil)
}

// @Summary List
// @Tags    SystemFile
// @Param query query string false "query, ?query=age:>:21,name:like:%jason%"
// @Param order query string false "order, ?order=age:desc,created_at:asc"
// @Param page query int false "page"
// @Param pageSize query int false "pageSize"
// @Success 200 {array} models.SystemFile "systemFile array"
// @Router /systemFiles [get]
func (ctl *SystemFileController) List(c *gin.Context) {
	systemFile := &models.SystemFile{}
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
	systemFiles, total, err := systemFile.List(rawQuery, rawOrder, offset, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"total": total,
		"data":  systemFiles,
	})
}

// @Summary Get
// @Tags    SystemFile
// @Param  systemFileId path string true "systemFileId"
// @Success 200 {object} models.SystemFile "systemFile object"
// @Router /systemFiles/{systemFileId} [get]
func (ctl *SystemFileController) Get(c *gin.Context) {
	systemFile := &models.SystemFile{}
	id := c.Param("systemFileId")

	var err error
	systemFile.Id, err = strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	systemFile, err = systemFile.Get()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, systemFile)
}
