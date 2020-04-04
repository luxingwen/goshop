//generate by gen
package controllers

import (
	"github.com/gin-gonic/gin"
	"goshop/restful/models"
	"net/http"
	"strconv"
)

type UserNoticeSeeController struct {
}

// @Summary Create
// @Tags    UserNoticeSee
// @Param body body models.UserNoticeSee true "UserNoticeSee"
// @Success 200 {string} string ""
// @Router /userNoticeSees [post]
func (ctl *UserNoticeSeeController) Create(c *gin.Context) {
	userNoticeSee := models.UserNoticeSee{}
	if err := ParseRequest(c, &userNoticeSee); err != nil {
		return
	}
	if err := userNoticeSee.Insert(); err != nil {
		c.JSON(http.StatusBadGateway, err)
		return
	}
	c.JSON(http.StatusOK, userNoticeSee)
}

// @Summary  Delete
// @Tags     UserNoticeSee
// @Param  userNoticeSeeId  path string true "userNoticeSeeId"
// @Success 200 {string} string ""
// @Router /userNoticeSees/{userNoticeSeeId} [delete]
func (ctl *UserNoticeSeeController) Delete(c *gin.Context) {
	userNoticeSee := models.UserNoticeSee{}
	id := c.Param("userNoticeSeeId")
	var err error
	userNoticeSee.Id, err = strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	err = userNoticeSee.Delete()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, nil)
}

// @Summary Put
// @Tags    UserNoticeSee
// @Param body body models.UserNoticeSee true "userNoticeSee"
// @Param  userNoticeSeeId path string true "userNoticeSeeId"
// @Success 200 {string} string ""
// @Router /userNoticeSees/{userNoticeSeeId} [put]
func (ctl *UserNoticeSeeController) Put(c *gin.Context) {
	userNoticeSee := models.UserNoticeSee{}
	id := c.Param("userNoticeSeeId")
	var err error
	userNoticeSee.Id, err = strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	if err := ParseRequest(c, &userNoticeSee); err != nil {
		return
	}
	err = userNoticeSee.Update()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, nil)
}

// @Summary Patch
// @Tags    UserNoticeSee
// @Param body body models.UserNoticeSee true "userNoticeSee"
// @Param  userNoticeSeeId path string true "userNoticeSeeId"
// @Success 200 {string} string ""
// @Router /userNoticeSees/{userNoticeSeeId} [patch]
func (ctl *UserNoticeSeeController) Patch(c *gin.Context) {
	userNoticeSee := models.UserNoticeSee{}
	id := c.Param("userNoticeSeeId")
	var err error
	userNoticeSee.Id, err = strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	if err := ParseRequest(c, &userNoticeSee); err != nil {
		return
	}
	err = userNoticeSee.Patch()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, nil)
}

// @Summary List
// @Tags    UserNoticeSee
// @Param query query string false "query, ?query=age:>:21,name:like:%jason%"
// @Param order query string false "order, ?order=age:desc,created_at:asc"
// @Param page query int false "page"
// @Param pageSize query int false "pageSize"
// @Success 200 {array} models.UserNoticeSee "userNoticeSee array"
// @Router /userNoticeSees [get]
func (ctl *UserNoticeSeeController) List(c *gin.Context) {
	userNoticeSee := &models.UserNoticeSee{}
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
	userNoticeSees, total, err := userNoticeSee.List(rawQuery, rawOrder, offset, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"total": total,
		"data":  userNoticeSees,
	})
}

// @Summary Get
// @Tags    UserNoticeSee
// @Param  userNoticeSeeId path string true "userNoticeSeeId"
// @Success 200 {object} models.UserNoticeSee "userNoticeSee object"
// @Router /userNoticeSees/{userNoticeSeeId} [get]
func (ctl *UserNoticeSeeController) Get(c *gin.Context) {
	userNoticeSee := &models.UserNoticeSee{}
	id := c.Param("userNoticeSeeId")

	var err error
	userNoticeSee.Id, err = strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	userNoticeSee, err = userNoticeSee.Get()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, userNoticeSee)
}
