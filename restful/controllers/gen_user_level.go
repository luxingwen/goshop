//generate by gen
package controllers

import (
	"github.com/gin-gonic/gin"
	"goshop/restful/models"
	"net/http"
	"strconv"
)

type UserLevelController struct {
}

// @Summary Create
// @Tags    UserLevel
// @Param body body models.UserLevel true "UserLevel"
// @Success 200 {string} string ""
// @Router /userLevels [post]
func (ctl *UserLevelController) Create(c *gin.Context) {
	userLevel := models.UserLevel{}
	if err := ParseRequest(c, &userLevel); err != nil {
		return
	}
	if err := userLevel.Insert(); err != nil {
		c.JSON(http.StatusBadGateway, err)
		return
	}
	c.JSON(http.StatusOK, userLevel)
}

// @Summary  Delete
// @Tags     UserLevel
// @Param  userLevelId  path string true "userLevelId"
// @Success 200 {string} string ""
// @Router /userLevels/{userLevelId} [delete]
func (ctl *UserLevelController) Delete(c *gin.Context) {
	userLevel := models.UserLevel{}
	id := c.Param("userLevelId")
	var err error
	userLevel.Id, err = strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	err = userLevel.Delete()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, nil)
}

// @Summary Put
// @Tags    UserLevel
// @Param body body models.UserLevel true "userLevel"
// @Param  userLevelId path string true "userLevelId"
// @Success 200 {string} string ""
// @Router /userLevels/{userLevelId} [put]
func (ctl *UserLevelController) Put(c *gin.Context) {
	userLevel := models.UserLevel{}
	id := c.Param("userLevelId")
	var err error
	userLevel.Id, err = strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	if err := ParseRequest(c, &userLevel); err != nil {
		return
	}
	err = userLevel.Update()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, nil)
}

// @Summary Patch
// @Tags    UserLevel
// @Param body body models.UserLevel true "userLevel"
// @Param  userLevelId path string true "userLevelId"
// @Success 200 {string} string ""
// @Router /userLevels/{userLevelId} [patch]
func (ctl *UserLevelController) Patch(c *gin.Context) {
	userLevel := models.UserLevel{}
	id := c.Param("userLevelId")
	var err error
	userLevel.Id, err = strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	if err := ParseRequest(c, &userLevel); err != nil {
		return
	}
	err = userLevel.Patch()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, nil)
}

// @Summary List
// @Tags    UserLevel
// @Param query query string false "query, ?query=age:>:21,name:like:%jason%"
// @Param order query string false "order, ?order=age:desc,created_at:asc"
// @Param page query int false "page"
// @Param pageSize query int false "pageSize"
// @Success 200 {array} models.UserLevel "userLevel array"
// @Router /userLevels [get]
func (ctl *UserLevelController) List(c *gin.Context) {
	userLevel := &models.UserLevel{}
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
	userLevels, total, err := userLevel.List(rawQuery, rawOrder, offset, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"total": total,
		"data":  userLevels,
	})
}

// @Summary Get
// @Tags    UserLevel
// @Param  userLevelId path string true "userLevelId"
// @Success 200 {object} models.UserLevel "userLevel object"
// @Router /userLevels/{userLevelId} [get]
func (ctl *UserLevelController) Get(c *gin.Context) {
	userLevel := &models.UserLevel{}
	id := c.Param("userLevelId")

	var err error
	userLevel.Id, err = strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	userLevel, err = userLevel.Get()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, userLevel)
}
