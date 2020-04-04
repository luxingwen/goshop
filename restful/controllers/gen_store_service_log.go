//generate by gen
package controllers

import (
	"github.com/gin-gonic/gin"
	"goshop/restful/models"
	"net/http"
	"strconv"
)

type StoreServiceLogController struct {
}

// @Summary Create
// @Tags    StoreServiceLog
// @Param body body models.StoreServiceLog true "StoreServiceLog"
// @Success 200 {string} string ""
// @Router /storeServiceLogs [post]
func (ctl *StoreServiceLogController) Create(c *gin.Context) {
	storeServiceLog := models.StoreServiceLog{}
	if err := ParseRequest(c, &storeServiceLog); err != nil {
		return
	}
	if err := storeServiceLog.Insert(); err != nil {
		c.JSON(http.StatusBadGateway, err)
		return
	}
	c.JSON(http.StatusOK, storeServiceLog)
}

// @Summary  Delete
// @Tags     StoreServiceLog
// @Param  storeServiceLogId  path string true "storeServiceLogId"
// @Success 200 {string} string ""
// @Router /storeServiceLogs/{storeServiceLogId} [delete]
func (ctl *StoreServiceLogController) Delete(c *gin.Context) {
	storeServiceLog := models.StoreServiceLog{}
	id := c.Param("storeServiceLogId")
	var err error
	storeServiceLog.Id, err = strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	err = storeServiceLog.Delete()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, nil)
}

// @Summary Put
// @Tags    StoreServiceLog
// @Param body body models.StoreServiceLog true "storeServiceLog"
// @Param  storeServiceLogId path string true "storeServiceLogId"
// @Success 200 {string} string ""
// @Router /storeServiceLogs/{storeServiceLogId} [put]
func (ctl *StoreServiceLogController) Put(c *gin.Context) {
	storeServiceLog := models.StoreServiceLog{}
	id := c.Param("storeServiceLogId")
	var err error
	storeServiceLog.Id, err = strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	if err := ParseRequest(c, &storeServiceLog); err != nil {
		return
	}
	err = storeServiceLog.Update()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, nil)
}

// @Summary Patch
// @Tags    StoreServiceLog
// @Param body body models.StoreServiceLog true "storeServiceLog"
// @Param  storeServiceLogId path string true "storeServiceLogId"
// @Success 200 {string} string ""
// @Router /storeServiceLogs/{storeServiceLogId} [patch]
func (ctl *StoreServiceLogController) Patch(c *gin.Context) {
	storeServiceLog := models.StoreServiceLog{}
	id := c.Param("storeServiceLogId")
	var err error
	storeServiceLog.Id, err = strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	if err := ParseRequest(c, &storeServiceLog); err != nil {
		return
	}
	err = storeServiceLog.Patch()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, nil)
}

// @Summary List
// @Tags    StoreServiceLog
// @Param query query string false "query, ?query=age:>:21,name:like:%jason%"
// @Param order query string false "order, ?order=age:desc,created_at:asc"
// @Param page query int false "page"
// @Param pageSize query int false "pageSize"
// @Success 200 {array} models.StoreServiceLog "storeServiceLog array"
// @Router /storeServiceLogs [get]
func (ctl *StoreServiceLogController) List(c *gin.Context) {
	storeServiceLog := &models.StoreServiceLog{}
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
	storeServiceLogs, total, err := storeServiceLog.List(rawQuery, rawOrder, offset, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"total": total,
		"data":  storeServiceLogs,
	})
}

// @Summary Get
// @Tags    StoreServiceLog
// @Param  storeServiceLogId path string true "storeServiceLogId"
// @Success 200 {object} models.StoreServiceLog "storeServiceLog object"
// @Router /storeServiceLogs/{storeServiceLogId} [get]
func (ctl *StoreServiceLogController) Get(c *gin.Context) {
	storeServiceLog := &models.StoreServiceLog{}
	id := c.Param("storeServiceLogId")

	var err error
	storeServiceLog.Id, err = strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	storeServiceLog, err = storeServiceLog.Get()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, storeServiceLog)
}
