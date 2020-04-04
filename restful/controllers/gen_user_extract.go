//generate by gen
package controllers

import (
	"github.com/gin-gonic/gin"
	"goshop/restful/models"
	"net/http"
	"strconv"
)

type UserExtractController struct {
}

// @Summary Create
// @Tags    UserExtract
// @Param body body models.UserExtract true "UserExtract"
// @Success 200 {string} string ""
// @Router /userExtracts [post]
func (ctl *UserExtractController) Create(c *gin.Context) {
	userExtract := models.UserExtract{}
	if err := ParseRequest(c, &userExtract); err != nil {
		return
	}
	if err := userExtract.Insert(); err != nil {
		c.JSON(http.StatusBadGateway, err)
		return
	}
	c.JSON(http.StatusOK, userExtract)
}

// @Summary  Delete
// @Tags     UserExtract
// @Param  userExtractId  path string true "userExtractId"
// @Success 200 {string} string ""
// @Router /userExtracts/{userExtractId} [delete]
func (ctl *UserExtractController) Delete(c *gin.Context) {
	userExtract := models.UserExtract{}
	id := c.Param("userExtractId")
	var err error
	userExtract.Id, err = strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	err = userExtract.Delete()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, nil)
}

// @Summary Put
// @Tags    UserExtract
// @Param body body models.UserExtract true "userExtract"
// @Param  userExtractId path string true "userExtractId"
// @Success 200 {string} string ""
// @Router /userExtracts/{userExtractId} [put]
func (ctl *UserExtractController) Put(c *gin.Context) {
	userExtract := models.UserExtract{}
	id := c.Param("userExtractId")
	var err error
	userExtract.Id, err = strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	if err := ParseRequest(c, &userExtract); err != nil {
		return
	}
	err = userExtract.Update()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, nil)
}

// @Summary Patch
// @Tags    UserExtract
// @Param body body models.UserExtract true "userExtract"
// @Param  userExtractId path string true "userExtractId"
// @Success 200 {string} string ""
// @Router /userExtracts/{userExtractId} [patch]
func (ctl *UserExtractController) Patch(c *gin.Context) {
	userExtract := models.UserExtract{}
	id := c.Param("userExtractId")
	var err error
	userExtract.Id, err = strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	if err := ParseRequest(c, &userExtract); err != nil {
		return
	}
	err = userExtract.Patch()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, nil)
}

// @Summary List
// @Tags    UserExtract
// @Param query query string false "query, ?query=age:>:21,name:like:%jason%"
// @Param order query string false "order, ?order=age:desc,created_at:asc"
// @Param page query int false "page"
// @Param pageSize query int false "pageSize"
// @Success 200 {array} models.UserExtract "userExtract array"
// @Router /userExtracts [get]
func (ctl *UserExtractController) List(c *gin.Context) {
	userExtract := &models.UserExtract{}
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
	userExtracts, total, err := userExtract.List(rawQuery, rawOrder, offset, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"total": total,
		"data":  userExtracts,
	})
}

// @Summary Get
// @Tags    UserExtract
// @Param  userExtractId path string true "userExtractId"
// @Success 200 {object} models.UserExtract "userExtract object"
// @Router /userExtracts/{userExtractId} [get]
func (ctl *UserExtractController) Get(c *gin.Context) {
	userExtract := &models.UserExtract{}
	id := c.Param("userExtractId")

	var err error
	userExtract.Id, err = strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	userExtract, err = userExtract.Get()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, userExtract)
}
