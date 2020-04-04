//generate by gen
package controllers

import (
	"github.com/gin-gonic/gin"
	"goshop/restful/models"
	"net/http"
	"strconv"
)

type UserSignController struct {
}

// @Summary Create
// @Tags    UserSign
// @Param body body models.UserSign true "UserSign"
// @Success 200 {string} string ""
// @Router /userSigns [post]
func (ctl *UserSignController) Create(c *gin.Context) {
	userSign := models.UserSign{}
	if err := ParseRequest(c, &userSign); err != nil {
		return
	}
	if err := userSign.Insert(); err != nil {
		c.JSON(http.StatusBadGateway, err)
		return
	}
	c.JSON(http.StatusOK, userSign)
}

// @Summary  Delete
// @Tags     UserSign
// @Param  userSignId  path string true "userSignId"
// @Success 200 {string} string ""
// @Router /userSigns/{userSignId} [delete]
func (ctl *UserSignController) Delete(c *gin.Context) {
	userSign := models.UserSign{}
	id := c.Param("userSignId")
	var err error
	userSign.Id, err = strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	err = userSign.Delete()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, nil)
}

// @Summary Put
// @Tags    UserSign
// @Param body body models.UserSign true "userSign"
// @Param  userSignId path string true "userSignId"
// @Success 200 {string} string ""
// @Router /userSigns/{userSignId} [put]
func (ctl *UserSignController) Put(c *gin.Context) {
	userSign := models.UserSign{}
	id := c.Param("userSignId")
	var err error
	userSign.Id, err = strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	if err := ParseRequest(c, &userSign); err != nil {
		return
	}
	err = userSign.Update()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, nil)
}

// @Summary Patch
// @Tags    UserSign
// @Param body body models.UserSign true "userSign"
// @Param  userSignId path string true "userSignId"
// @Success 200 {string} string ""
// @Router /userSigns/{userSignId} [patch]
func (ctl *UserSignController) Patch(c *gin.Context) {
	userSign := models.UserSign{}
	id := c.Param("userSignId")
	var err error
	userSign.Id, err = strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	if err := ParseRequest(c, &userSign); err != nil {
		return
	}
	err = userSign.Patch()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, nil)
}

// @Summary List
// @Tags    UserSign
// @Param query query string false "query, ?query=age:>:21,name:like:%jason%"
// @Param order query string false "order, ?order=age:desc,created_at:asc"
// @Param page query int false "page"
// @Param pageSize query int false "pageSize"
// @Success 200 {array} models.UserSign "userSign array"
// @Router /userSigns [get]
func (ctl *UserSignController) List(c *gin.Context) {
	userSign := &models.UserSign{}
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
	userSigns, total, err := userSign.List(rawQuery, rawOrder, offset, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"total": total,
		"data":  userSigns,
	})
}

// @Summary Get
// @Tags    UserSign
// @Param  userSignId path string true "userSignId"
// @Success 200 {object} models.UserSign "userSign object"
// @Router /userSigns/{userSignId} [get]
func (ctl *UserSignController) Get(c *gin.Context) {
	userSign := &models.UserSign{}
	id := c.Param("userSignId")

	var err error
	userSign.Id, err = strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	userSign, err = userSign.Get()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, userSign)
}
