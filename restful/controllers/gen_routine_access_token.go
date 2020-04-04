//generate by gen
package controllers

import (
	"github.com/gin-gonic/gin"
	"goshop/restful/models"
	"net/http"
	"strconv"
)

type RoutineAccessTokenController struct {
}

// @Summary Create
// @Tags    RoutineAccessToken
// @Param body body models.RoutineAccessToken true "RoutineAccessToken"
// @Success 200 {string} string ""
// @Router /routineAccessTokens [post]
func (ctl *RoutineAccessTokenController) Create(c *gin.Context) {
	routineAccessToken := models.RoutineAccessToken{}
	if err := ParseRequest(c, &routineAccessToken); err != nil {
		return
	}
	if err := routineAccessToken.Insert(); err != nil {
		c.JSON(http.StatusBadGateway, err)
		return
	}
	c.JSON(http.StatusOK, routineAccessToken)
}

// @Summary  Delete
// @Tags     RoutineAccessToken
// @Param  routineAccessTokenId  path string true "routineAccessTokenId"
// @Success 200 {string} string ""
// @Router /routineAccessTokens/{routineAccessTokenId} [delete]
func (ctl *RoutineAccessTokenController) Delete(c *gin.Context) {
	routineAccessToken := models.RoutineAccessToken{}
	id := c.Param("routineAccessTokenId")
	var err error
	routineAccessToken.Id, err = strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	err = routineAccessToken.Delete()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, nil)
}

// @Summary Put
// @Tags    RoutineAccessToken
// @Param body body models.RoutineAccessToken true "routineAccessToken"
// @Param  routineAccessTokenId path string true "routineAccessTokenId"
// @Success 200 {string} string ""
// @Router /routineAccessTokens/{routineAccessTokenId} [put]
func (ctl *RoutineAccessTokenController) Put(c *gin.Context) {
	routineAccessToken := models.RoutineAccessToken{}
	id := c.Param("routineAccessTokenId")
	var err error
	routineAccessToken.Id, err = strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	if err := ParseRequest(c, &routineAccessToken); err != nil {
		return
	}
	err = routineAccessToken.Update()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, nil)
}

// @Summary Patch
// @Tags    RoutineAccessToken
// @Param body body models.RoutineAccessToken true "routineAccessToken"
// @Param  routineAccessTokenId path string true "routineAccessTokenId"
// @Success 200 {string} string ""
// @Router /routineAccessTokens/{routineAccessTokenId} [patch]
func (ctl *RoutineAccessTokenController) Patch(c *gin.Context) {
	routineAccessToken := models.RoutineAccessToken{}
	id := c.Param("routineAccessTokenId")
	var err error
	routineAccessToken.Id, err = strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	if err := ParseRequest(c, &routineAccessToken); err != nil {
		return
	}
	err = routineAccessToken.Patch()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, nil)
}

// @Summary List
// @Tags    RoutineAccessToken
// @Param query query string false "query, ?query=age:>:21,name:like:%jason%"
// @Param order query string false "order, ?order=age:desc,created_at:asc"
// @Param page query int false "page"
// @Param pageSize query int false "pageSize"
// @Success 200 {array} models.RoutineAccessToken "routineAccessToken array"
// @Router /routineAccessTokens [get]
func (ctl *RoutineAccessTokenController) List(c *gin.Context) {
	routineAccessToken := &models.RoutineAccessToken{}
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
	routineAccessTokens, total, err := routineAccessToken.List(rawQuery, rawOrder, offset, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"total": total,
		"data":  routineAccessTokens,
	})
}

// @Summary Get
// @Tags    RoutineAccessToken
// @Param  routineAccessTokenId path string true "routineAccessTokenId"
// @Success 200 {object} models.RoutineAccessToken "routineAccessToken object"
// @Router /routineAccessTokens/{routineAccessTokenId} [get]
func (ctl *RoutineAccessTokenController) Get(c *gin.Context) {
	routineAccessToken := &models.RoutineAccessToken{}
	id := c.Param("routineAccessTokenId")

	var err error
	routineAccessToken.Id, err = strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	routineAccessToken, err = routineAccessToken.Get()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, routineAccessToken)
}
