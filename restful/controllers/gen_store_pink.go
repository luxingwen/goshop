//generate by gen
package controllers

import (
	"github.com/gin-gonic/gin"
	"goshop/restful/models"
	"net/http"
	"strconv"
)

type StorePinkController struct {
}

// @Summary Create
// @Tags    StorePink
// @Param body body models.StorePink true "StorePink"
// @Success 200 {string} string ""
// @Router /storePinks [post]
func (ctl *StorePinkController) Create(c *gin.Context) {
	storePink := models.StorePink{}
	if err := ParseRequest(c, &storePink); err != nil {
		return
	}
	if err := storePink.Insert(); err != nil {
		c.JSON(http.StatusBadGateway, err)
		return
	}
	c.JSON(http.StatusOK, storePink)
}

// @Summary  Delete
// @Tags     StorePink
// @Param  storePinkId  path string true "storePinkId"
// @Success 200 {string} string ""
// @Router /storePinks/{storePinkId} [delete]
func (ctl *StorePinkController) Delete(c *gin.Context) {
	storePink := models.StorePink{}
	id := c.Param("storePinkId")
	var err error
	storePink.Id, err = strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	err = storePink.Delete()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, nil)
}

// @Summary Put
// @Tags    StorePink
// @Param body body models.StorePink true "storePink"
// @Param  storePinkId path string true "storePinkId"
// @Success 200 {string} string ""
// @Router /storePinks/{storePinkId} [put]
func (ctl *StorePinkController) Put(c *gin.Context) {
	storePink := models.StorePink{}
	id := c.Param("storePinkId")
	var err error
	storePink.Id, err = strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	if err := ParseRequest(c, &storePink); err != nil {
		return
	}
	err = storePink.Update()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, nil)
}

// @Summary Patch
// @Tags    StorePink
// @Param body body models.StorePink true "storePink"
// @Param  storePinkId path string true "storePinkId"
// @Success 200 {string} string ""
// @Router /storePinks/{storePinkId} [patch]
func (ctl *StorePinkController) Patch(c *gin.Context) {
	storePink := models.StorePink{}
	id := c.Param("storePinkId")
	var err error
	storePink.Id, err = strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	if err := ParseRequest(c, &storePink); err != nil {
		return
	}
	err = storePink.Patch()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, nil)
}

// @Summary List
// @Tags    StorePink
// @Param query query string false "query, ?query=age:>:21,name:like:%jason%"
// @Param order query string false "order, ?order=age:desc,created_at:asc"
// @Param page query int false "page"
// @Param pageSize query int false "pageSize"
// @Success 200 {array} models.StorePink "storePink array"
// @Router /storePinks [get]
func (ctl *StorePinkController) List(c *gin.Context) {
	storePink := &models.StorePink{}
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
	storePinks, total, err := storePink.List(rawQuery, rawOrder, offset, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"total": total,
		"data":  storePinks,
	})
}

// @Summary Get
// @Tags    StorePink
// @Param  storePinkId path string true "storePinkId"
// @Success 200 {object} models.StorePink "storePink object"
// @Router /storePinks/{storePinkId} [get]
func (ctl *StorePinkController) Get(c *gin.Context) {
	storePink := &models.StorePink{}
	id := c.Param("storePinkId")

	var err error
	storePink.Id, err = strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	storePink, err = storePink.Get()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, storePink)
}
