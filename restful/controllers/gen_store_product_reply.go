//generate by gen
package controllers

import (
	"github.com/gin-gonic/gin"
	"goshop/restful/models"
	"net/http"
	"strconv"
)

type StoreProductReplyController struct {
}

// @Summary Create
// @Tags    StoreProductReply
// @Param body body models.StoreProductReply true "StoreProductReply"
// @Success 200 {string} string ""
// @Router /storeProductReplys [post]
func (ctl *StoreProductReplyController) Create(c *gin.Context) {
	storeProductReply := models.StoreProductReply{}
	if err := ParseRequest(c, &storeProductReply); err != nil {
		return
	}
	if err := storeProductReply.Insert(); err != nil {
		c.JSON(http.StatusBadGateway, err)
		return
	}
	c.JSON(http.StatusOK, storeProductReply)
}

// @Summary  Delete
// @Tags     StoreProductReply
// @Param  storeProductReplyId  path string true "storeProductReplyId"
// @Success 200 {string} string ""
// @Router /storeProductReplys/{storeProductReplyId} [delete]
func (ctl *StoreProductReplyController) Delete(c *gin.Context) {
	storeProductReply := models.StoreProductReply{}
	id := c.Param("storeProductReplyId")
	var err error
	storeProductReply.Id, err = strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	err = storeProductReply.Delete()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, nil)
}

// @Summary Put
// @Tags    StoreProductReply
// @Param body body models.StoreProductReply true "storeProductReply"
// @Param  storeProductReplyId path string true "storeProductReplyId"
// @Success 200 {string} string ""
// @Router /storeProductReplys/{storeProductReplyId} [put]
func (ctl *StoreProductReplyController) Put(c *gin.Context) {
	storeProductReply := models.StoreProductReply{}
	id := c.Param("storeProductReplyId")
	var err error
	storeProductReply.Id, err = strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	if err := ParseRequest(c, &storeProductReply); err != nil {
		return
	}
	err = storeProductReply.Update()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, nil)
}

// @Summary Patch
// @Tags    StoreProductReply
// @Param body body models.StoreProductReply true "storeProductReply"
// @Param  storeProductReplyId path string true "storeProductReplyId"
// @Success 200 {string} string ""
// @Router /storeProductReplys/{storeProductReplyId} [patch]
func (ctl *StoreProductReplyController) Patch(c *gin.Context) {
	storeProductReply := models.StoreProductReply{}
	id := c.Param("storeProductReplyId")
	var err error
	storeProductReply.Id, err = strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	if err := ParseRequest(c, &storeProductReply); err != nil {
		return
	}
	err = storeProductReply.Patch()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, nil)
}

// @Summary List
// @Tags    StoreProductReply
// @Param query query string false "query, ?query=age:>:21,name:like:%jason%"
// @Param order query string false "order, ?order=age:desc,created_at:asc"
// @Param page query int false "page"
// @Param pageSize query int false "pageSize"
// @Success 200 {array} models.StoreProductReply "storeProductReply array"
// @Router /storeProductReplys [get]
func (ctl *StoreProductReplyController) List(c *gin.Context) {
	storeProductReply := &models.StoreProductReply{}
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
	storeProductReplys, total, err := storeProductReply.List(rawQuery, rawOrder, offset, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"total": total,
		"data":  storeProductReplys,
	})
}

// @Summary Get
// @Tags    StoreProductReply
// @Param  storeProductReplyId path string true "storeProductReplyId"
// @Success 200 {object} models.StoreProductReply "storeProductReply object"
// @Router /storeProductReplys/{storeProductReplyId} [get]
func (ctl *StoreProductReplyController) Get(c *gin.Context) {
	storeProductReply := &models.StoreProductReply{}
	id := c.Param("storeProductReplyId")

	var err error
	storeProductReply.Id, err = strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	storeProductReply, err = storeProductReply.Get()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, storeProductReply)
}
