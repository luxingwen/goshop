//generate by gen
package controllers

import (
	"github.com/gin-gonic/gin"
	"goshop/restful/models"
	"net/http"
	"strconv"
)

type SystemUserTaskController struct {
}

// @Summary Create
// @Tags    SystemUserTask
// @Param body body models.SystemUserTask true "SystemUserTask"
// @Success 200 {string} string ""
// @Router /systemUserTasks [post]
func (ctl *SystemUserTaskController) Create(c *gin.Context) {
	systemUserTask := models.SystemUserTask{}
	if err := ParseRequest(c, &systemUserTask); err != nil {
		return
	}
	if err := systemUserTask.Insert(); err != nil {
		c.JSON(http.StatusBadGateway, err)
		return
	}
	c.JSON(http.StatusOK, systemUserTask)
}

// @Summary  Delete
// @Tags     SystemUserTask
// @Param  systemUserTaskId  path string true "systemUserTaskId"
// @Success 200 {string} string ""
// @Router /systemUserTasks/{systemUserTaskId} [delete]
func (ctl *SystemUserTaskController) Delete(c *gin.Context) {
	systemUserTask := models.SystemUserTask{}
	id := c.Param("systemUserTaskId")
	var err error
	systemUserTask.Id, err = strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	err = systemUserTask.Delete()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, nil)
}

// @Summary Put
// @Tags    SystemUserTask
// @Param body body models.SystemUserTask true "systemUserTask"
// @Param  systemUserTaskId path string true "systemUserTaskId"
// @Success 200 {string} string ""
// @Router /systemUserTasks/{systemUserTaskId} [put]
func (ctl *SystemUserTaskController) Put(c *gin.Context) {
	systemUserTask := models.SystemUserTask{}
	id := c.Param("systemUserTaskId")
	var err error
	systemUserTask.Id, err = strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	if err := ParseRequest(c, &systemUserTask); err != nil {
		return
	}
	err = systemUserTask.Update()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, nil)
}

// @Summary Patch
// @Tags    SystemUserTask
// @Param body body models.SystemUserTask true "systemUserTask"
// @Param  systemUserTaskId path string true "systemUserTaskId"
// @Success 200 {string} string ""
// @Router /systemUserTasks/{systemUserTaskId} [patch]
func (ctl *SystemUserTaskController) Patch(c *gin.Context) {
	systemUserTask := models.SystemUserTask{}
	id := c.Param("systemUserTaskId")
	var err error
	systemUserTask.Id, err = strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	if err := ParseRequest(c, &systemUserTask); err != nil {
		return
	}
	err = systemUserTask.Patch()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, nil)
}

// @Summary List
// @Tags    SystemUserTask
// @Param query query string false "query, ?query=age:>:21,name:like:%jason%"
// @Param order query string false "order, ?order=age:desc,created_at:asc"
// @Param page query int false "page"
// @Param pageSize query int false "pageSize"
// @Success 200 {array} models.SystemUserTask "systemUserTask array"
// @Router /systemUserTasks [get]
func (ctl *SystemUserTaskController) List(c *gin.Context) {
	systemUserTask := &models.SystemUserTask{}
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
	systemUserTasks, total, err := systemUserTask.List(rawQuery, rawOrder, offset, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"total": total,
		"data":  systemUserTasks,
	})
}

// @Summary Get
// @Tags    SystemUserTask
// @Param  systemUserTaskId path string true "systemUserTaskId"
// @Success 200 {object} models.SystemUserTask "systemUserTask object"
// @Router /systemUserTasks/{systemUserTaskId} [get]
func (ctl *SystemUserTaskController) Get(c *gin.Context) {
	systemUserTask := &models.SystemUserTask{}
	id := c.Param("systemUserTaskId")

	var err error
	systemUserTask.Id, err = strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	systemUserTask, err = systemUserTask.Get()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, systemUserTask)
}
