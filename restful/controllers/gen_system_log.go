//generate by gen
package controllers

import (
	"github.com/gin-gonic/gin"
	"goshop/restful/models"
	"net/http"
	"strconv"
)

type SystemLogController struct {
}

// @Summary Create
// @Tags    SystemLog
// @Param body body models.SystemLog true "SystemLog"
// @Success 200 {string} string ""
// @Router /systemLogs [post]
func (ctl *SystemLogController) Create(c *gin.Context) {
	systemLog := models.SystemLog{}
	if err := ParseRequest(c, &systemLog); err != nil {
		return
	}
	if err := systemLog.Insert(); err != nil {
		c.JSON(http.StatusBadGateway, err)
		return
	}
	c.JSON(http.StatusOK, systemLog)
}

// @Summary  Delete
// @Tags     SystemLog
// @Param  systemLogId  path string true "systemLogId"
// @Success 200 {string} string ""
// @Router /systemLogs/{systemLogId} [delete]
func (ctl *SystemLogController) Delete(c *gin.Context) {
	systemLog := models.SystemLog{}
	id := c.Param("systemLogId")
	var err error
	systemLog.Id, err = strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	err = systemLog.Delete()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, nil)
}

// @Summary Put
// @Tags    SystemLog
// @Param body body models.SystemLog true "systemLog"
// @Param  systemLogId path string true "systemLogId"
// @Success 200 {string} string ""
// @Router /systemLogs/{systemLogId} [put]
func (ctl *SystemLogController) Put(c *gin.Context) {
	systemLog := models.SystemLog{}
	id := c.Param("systemLogId")
	var err error
	systemLog.Id, err = strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	if err := ParseRequest(c, &systemLog); err != nil {
		return
	}
	err = systemLog.Update()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, nil)
}

// @Summary Patch
// @Tags    SystemLog
// @Param body body models.SystemLog true "systemLog"
// @Param  systemLogId path string true "systemLogId"
// @Success 200 {string} string ""
// @Router /systemLogs/{systemLogId} [patch]
func (ctl *SystemLogController) Patch(c *gin.Context) {
	systemLog := models.SystemLog{}
	id := c.Param("systemLogId")
	var err error
	systemLog.Id, err = strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	if err := ParseRequest(c, &systemLog); err != nil {
		return
	}
	err = systemLog.Patch()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, nil)
}

// @Summary List
// @Tags    SystemLog
// @Param query query string false "query, ?query=age:>:21,name:like:%jason%"
// @Param order query string false "order, ?order=age:desc,created_at:asc"
// @Param page query int false "page"
// @Param pageSize query int false "pageSize"
// @Success 200 {array} models.SystemLog "systemLog array"
// @Router /systemLogs [get]
func (ctl *SystemLogController) List(c *gin.Context) {
	systemLog := &models.SystemLog{}
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
	systemLogs, total, err := systemLog.List(rawQuery, rawOrder, offset, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"total": total,
		"data":  systemLogs,
	})
}

// @Summary Get
// @Tags    SystemLog
// @Param  systemLogId path string true "systemLogId"
// @Success 200 {object} models.SystemLog "systemLog object"
// @Router /systemLogs/{systemLogId} [get]
func (ctl *SystemLogController) Get(c *gin.Context) {
	systemLog := &models.SystemLog{}
	id := c.Param("systemLogId")

	var err error
	systemLog.Id, err = strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	systemLog, err = systemLog.Get()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, systemLog)
}
