//generate by gen
package controllers

import (
	"github.com/gin-gonic/gin"
	"goshop/restful/models"
	"net/http"
	"strconv"
)

type TokenController struct {
}

// @Summary Create
// @Tags    Token
// @Param body body models.Token true "Token"
// @Success 200 {string} string ""
// @Router /tokens [post]
func (ctl *TokenController) Create(c *gin.Context) {
	token := models.Token{}
	if err := ParseRequest(c, &token); err != nil {
		return
	}
	if err := token.Insert(); err != nil {
		c.JSON(http.StatusBadGateway, err)
		return
	}
	c.JSON(http.StatusOK, token)
}

// @Summary  Delete
// @Tags     Token
// @Param  tokenId  path string true "tokenId"
// @Success 200 {string} string ""
// @Router /tokens/{tokenId} [delete]
func (ctl *TokenController) Delete(c *gin.Context) {
	token := models.Token{}
	id := c.Param("tokenId")
	var err error
	token.Id, err = strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	err = token.Delete()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, nil)
}

// @Summary Put
// @Tags    Token
// @Param body body models.Token true "token"
// @Param  tokenId path string true "tokenId"
// @Success 200 {string} string ""
// @Router /tokens/{tokenId} [put]
func (ctl *TokenController) Put(c *gin.Context) {
	token := models.Token{}
	id := c.Param("tokenId")
	var err error
	token.Id, err = strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	if err := ParseRequest(c, &token); err != nil {
		return
	}
	err = token.Update()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, nil)
}

// @Summary Patch
// @Tags    Token
// @Param body body models.Token true "token"
// @Param  tokenId path string true "tokenId"
// @Success 200 {string} string ""
// @Router /tokens/{tokenId} [patch]
func (ctl *TokenController) Patch(c *gin.Context) {
	token := models.Token{}
	id := c.Param("tokenId")
	var err error
	token.Id, err = strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	if err := ParseRequest(c, &token); err != nil {
		return
	}
	err = token.Patch()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, nil)
}

// @Summary List
// @Tags    Token
// @Param query query string false "query, ?query=age:>:21,name:like:%jason%"
// @Param order query string false "order, ?order=age:desc,created_at:asc"
// @Param page query int false "page"
// @Param pageSize query int false "pageSize"
// @Success 200 {array} models.Token "token array"
// @Router /tokens [get]
func (ctl *TokenController) List(c *gin.Context) {
	token := &models.Token{}
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
	tokens, total, err := token.List(rawQuery, rawOrder, offset, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"total": total,
		"data":  tokens,
	})
}

// @Summary Get
// @Tags    Token
// @Param  tokenId path string true "tokenId"
// @Success 200 {object} models.Token "token object"
// @Router /tokens/{tokenId} [get]
func (ctl *TokenController) Get(c *gin.Context) {
	token := &models.Token{}
	id := c.Param("tokenId")

	var err error
	token.Id, err = strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	token, err = token.Get()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, token)
}
