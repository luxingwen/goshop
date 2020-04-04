//generate by gen
package controllers

import (
	"github.com/gin-gonic/gin"
	"goshop/restful/models"
	"net/http"
	"strconv"
)

type WechatMessageController struct {
}

// @Summary Create
// @Tags    WechatMessage
// @Param body body models.WechatMessage true "WechatMessage"
// @Success 200 {string} string ""
// @Router /wechatMessages [post]
func (ctl *WechatMessageController) Create(c *gin.Context) {
	wechatMessage := models.WechatMessage{}
	if err := ParseRequest(c, &wechatMessage); err != nil {
		return
	}
	if err := wechatMessage.Insert(); err != nil {
		c.JSON(http.StatusBadGateway, err)
		return
	}
	c.JSON(http.StatusOK, wechatMessage)
}

// @Summary  Delete
// @Tags     WechatMessage
// @Param  wechatMessageId  path string true "wechatMessageId"
// @Success 200 {string} string ""
// @Router /wechatMessages/{wechatMessageId} [delete]
func (ctl *WechatMessageController) Delete(c *gin.Context) {
	wechatMessage := models.WechatMessage{}
	id := c.Param("wechatMessageId")
	var err error
	wechatMessage.Id, err = strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	err = wechatMessage.Delete()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, nil)
}

// @Summary Put
// @Tags    WechatMessage
// @Param body body models.WechatMessage true "wechatMessage"
// @Param  wechatMessageId path string true "wechatMessageId"
// @Success 200 {string} string ""
// @Router /wechatMessages/{wechatMessageId} [put]
func (ctl *WechatMessageController) Put(c *gin.Context) {
	wechatMessage := models.WechatMessage{}
	id := c.Param("wechatMessageId")
	var err error
	wechatMessage.Id, err = strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	if err := ParseRequest(c, &wechatMessage); err != nil {
		return
	}
	err = wechatMessage.Update()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, nil)
}

// @Summary Patch
// @Tags    WechatMessage
// @Param body body models.WechatMessage true "wechatMessage"
// @Param  wechatMessageId path string true "wechatMessageId"
// @Success 200 {string} string ""
// @Router /wechatMessages/{wechatMessageId} [patch]
func (ctl *WechatMessageController) Patch(c *gin.Context) {
	wechatMessage := models.WechatMessage{}
	id := c.Param("wechatMessageId")
	var err error
	wechatMessage.Id, err = strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	if err := ParseRequest(c, &wechatMessage); err != nil {
		return
	}
	err = wechatMessage.Patch()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, nil)
}

// @Summary List
// @Tags    WechatMessage
// @Param query query string false "query, ?query=age:>:21,name:like:%jason%"
// @Param order query string false "order, ?order=age:desc,created_at:asc"
// @Param page query int false "page"
// @Param pageSize query int false "pageSize"
// @Success 200 {array} models.WechatMessage "wechatMessage array"
// @Router /wechatMessages [get]
func (ctl *WechatMessageController) List(c *gin.Context) {
	wechatMessage := &models.WechatMessage{}
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
	wechatMessages, total, err := wechatMessage.List(rawQuery, rawOrder, offset, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"total": total,
		"data":  wechatMessages,
	})
}

// @Summary Get
// @Tags    WechatMessage
// @Param  wechatMessageId path string true "wechatMessageId"
// @Success 200 {object} models.WechatMessage "wechatMessage object"
// @Router /wechatMessages/{wechatMessageId} [get]
func (ctl *WechatMessageController) Get(c *gin.Context) {
	wechatMessage := &models.WechatMessage{}
	id := c.Param("wechatMessageId")

	var err error
	wechatMessage.Id, err = strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	wechatMessage, err = wechatMessage.Get()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, wechatMessage)
}
