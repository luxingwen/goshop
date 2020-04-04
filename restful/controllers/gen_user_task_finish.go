//generate by gen
package controllers

import (
	"github.com/gin-gonic/gin"
	"goshop/restful/models"
	"net/http"
	"strconv"
)

type UserTaskFinishController struct {
}

// @Summary Create
// @Tags    UserTaskFinish
// @Param body body models.UserTaskFinish true "UserTaskFinish"
// @Success 200 {string} string ""
// @Router /userTaskFinishs [post]
func (ctl *UserTaskFinishController) Create(c *gin.Context) {
	userTaskFinish := models.UserTaskFinish{}
	if err := ParseRequest(c, &userTaskFinish); err != nil {
		return
	}
	if err := userTaskFinish.Insert(); err != nil {
		c.JSON(http.StatusBadGateway, err)
		return
	}
	c.JSON(http.StatusOK, userTaskFinish)
}

// @Summary  Delete
// @Tags     UserTaskFinish
// @Param  userTaskFinishId  path string true "userTaskFinishId"
// @Success 200 {string} string ""
// @Router /userTaskFinishs/{userTaskFinishId} [delete]
func (ctl *UserTaskFinishController) Delete(c *gin.Context) {
	userTaskFinish := models.UserTaskFinish{}
	id := c.Param("userTaskFinishId")
	var err error
	userTaskFinish.Id, err = strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	err = userTaskFinish.Delete()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, nil)
}

// @Summary Put
// @Tags    UserTaskFinish
// @Param body body models.UserTaskFinish true "userTaskFinish"
// @Param  userTaskFinishId path string true "userTaskFinishId"
// @Success 200 {string} string ""
// @Router /userTaskFinishs/{userTaskFinishId} [put]
func (ctl *UserTaskFinishController) Put(c *gin.Context) {
	userTaskFinish := models.UserTaskFinish{}
	id := c.Param("userTaskFinishId")
	var err error
	userTaskFinish.Id, err = strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	if err := ParseRequest(c, &userTaskFinish); err != nil {
		return
	}
	err = userTaskFinish.Update()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, nil)
}

// @Summary Patch
// @Tags    UserTaskFinish
// @Param body body models.UserTaskFinish true "userTaskFinish"
// @Param  userTaskFinishId path string true "userTaskFinishId"
// @Success 200 {string} string ""
// @Router /userTaskFinishs/{userTaskFinishId} [patch]
func (ctl *UserTaskFinishController) Patch(c *gin.Context) {
	userTaskFinish := models.UserTaskFinish{}
	id := c.Param("userTaskFinishId")
	var err error
	userTaskFinish.Id, err = strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	if err := ParseRequest(c, &userTaskFinish); err != nil {
		return
	}
	err = userTaskFinish.Patch()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, nil)
}

// @Summary List
// @Tags    UserTaskFinish
// @Param query query string false "query, ?query=age:>:21,name:like:%jason%"
// @Param order query string false "order, ?order=age:desc,created_at:asc"
// @Param page query int false "page"
// @Param pageSize query int false "pageSize"
// @Success 200 {array} models.UserTaskFinish "userTaskFinish array"
// @Router /userTaskFinishs [get]
func (ctl *UserTaskFinishController) List(c *gin.Context) {
	userTaskFinish := &models.UserTaskFinish{}
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
	userTaskFinishs, total, err := userTaskFinish.List(rawQuery, rawOrder, offset, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"total": total,
		"data":  userTaskFinishs,
	})
}

// @Summary Get
// @Tags    UserTaskFinish
// @Param  userTaskFinishId path string true "userTaskFinishId"
// @Success 200 {object} models.UserTaskFinish "userTaskFinish object"
// @Router /userTaskFinishs/{userTaskFinishId} [get]
func (ctl *UserTaskFinishController) Get(c *gin.Context) {
	userTaskFinish := &models.UserTaskFinish{}
	id := c.Param("userTaskFinishId")

	var err error
	userTaskFinish.Id, err = strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	userTaskFinish, err = userTaskFinish.Get()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, userTaskFinish)
}
