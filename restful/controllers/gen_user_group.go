//generate by gen
package controllers

import (
	"github.com/gin-gonic/gin"
	"goshop/restful/models"
	"net/http"
	"strconv"
)

type UserGroupController struct {
}

// @Summary Create
// @Tags    UserGroup
// @Param body body models.UserGroup true "UserGroup"
// @Success 200 {string} string ""
// @Router /userGroups [post]
func (ctl *UserGroupController) Create(c *gin.Context) {
	userGroup := models.UserGroup{}
	if err := ParseRequest(c, &userGroup); err != nil {
		return
	}
	if err := userGroup.Insert(); err != nil {
		c.JSON(http.StatusBadGateway, err)
		return
	}
	c.JSON(http.StatusOK, userGroup)
}

// @Summary  Delete
// @Tags     UserGroup
// @Param  userGroupId  path string true "userGroupId"
// @Success 200 {string} string ""
// @Router /userGroups/{userGroupId} [delete]
func (ctl *UserGroupController) Delete(c *gin.Context) {
	userGroup := models.UserGroup{}
	id := c.Param("userGroupId")
	var err error
	userGroup.Id, err = strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	err = userGroup.Delete()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, nil)
}

// @Summary Put
// @Tags    UserGroup
// @Param body body models.UserGroup true "userGroup"
// @Param  userGroupId path string true "userGroupId"
// @Success 200 {string} string ""
// @Router /userGroups/{userGroupId} [put]
func (ctl *UserGroupController) Put(c *gin.Context) {
	userGroup := models.UserGroup{}
	id := c.Param("userGroupId")
	var err error
	userGroup.Id, err = strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	if err := ParseRequest(c, &userGroup); err != nil {
		return
	}
	err = userGroup.Update()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, nil)
}

// @Summary Patch
// @Tags    UserGroup
// @Param body body models.UserGroup true "userGroup"
// @Param  userGroupId path string true "userGroupId"
// @Success 200 {string} string ""
// @Router /userGroups/{userGroupId} [patch]
func (ctl *UserGroupController) Patch(c *gin.Context) {
	userGroup := models.UserGroup{}
	id := c.Param("userGroupId")
	var err error
	userGroup.Id, err = strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	if err := ParseRequest(c, &userGroup); err != nil {
		return
	}
	err = userGroup.Patch()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, nil)
}

// @Summary List
// @Tags    UserGroup
// @Param query query string false "query, ?query=age:>:21,name:like:%jason%"
// @Param order query string false "order, ?order=age:desc,created_at:asc"
// @Param page query int false "page"
// @Param pageSize query int false "pageSize"
// @Success 200 {array} models.UserGroup "userGroup array"
// @Router /userGroups [get]
func (ctl *UserGroupController) List(c *gin.Context) {
	userGroup := &models.UserGroup{}
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
	userGroups, total, err := userGroup.List(rawQuery, rawOrder, offset, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"total": total,
		"data":  userGroups,
	})
}

// @Summary Get
// @Tags    UserGroup
// @Param  userGroupId path string true "userGroupId"
// @Success 200 {object} models.UserGroup "userGroup object"
// @Router /userGroups/{userGroupId} [get]
func (ctl *UserGroupController) Get(c *gin.Context) {
	userGroup := &models.UserGroup{}
	id := c.Param("userGroupId")

	var err error
	userGroup.Id, err = strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	userGroup, err = userGroup.Get()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, userGroup)
}
