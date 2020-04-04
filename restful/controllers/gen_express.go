//generate by gen
package controllers

import (
	"github.com/gin-gonic/gin"
	"goshop/restful/models"
	"net/http"
	"strconv"
)

type ExpressController struct {
}

// @Summary Create
// @Tags    Express
// @Param body body models.Express true "Express"
// @Success 200 {string} string ""
// @Router /expresss [post]
func (ctl *ExpressController) Create(c *gin.Context) {
	express := models.Express{}
	if err := ParseRequest(c, &express); err != nil {
		return
	}
	if err := express.Insert(); err != nil {
		c.JSON(http.StatusBadGateway, err)
		return
	}
	c.JSON(http.StatusOK, express)
}

// @Summary  Delete
// @Tags     Express
// @Param  expressId  path string true "expressId"
// @Success 200 {string} string ""
// @Router /expresss/{expressId} [delete]
func (ctl *ExpressController) Delete(c *gin.Context) {
	express := models.Express{}
	id := c.Param("expressId")
	var err error
	express.Id, err = strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	err = express.Delete()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, nil)
}

// @Summary Put
// @Tags    Express
// @Param body body models.Express true "express"
// @Param  expressId path string true "expressId"
// @Success 200 {string} string ""
// @Router /expresss/{expressId} [put]
func (ctl *ExpressController) Put(c *gin.Context) {
	express := models.Express{}
	id := c.Param("expressId")
	var err error
	express.Id, err = strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	if err := ParseRequest(c, &express); err != nil {
		return
	}
	err = express.Update()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, nil)
}

// @Summary Patch
// @Tags    Express
// @Param body body models.Express true "express"
// @Param  expressId path string true "expressId"
// @Success 200 {string} string ""
// @Router /expresss/{expressId} [patch]
func (ctl *ExpressController) Patch(c *gin.Context) {
	express := models.Express{}
	id := c.Param("expressId")
	var err error
	express.Id, err = strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	if err := ParseRequest(c, &express); err != nil {
		return
	}
	err = express.Patch()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, nil)
}

// @Summary List
// @Tags    Express
// @Param query query string false "query, ?query=age:>:21,name:like:%jason%"
// @Param order query string false "order, ?order=age:desc,created_at:asc"
// @Param page query int false "page"
// @Param pageSize query int false "pageSize"
// @Success 200 {array} models.Express "express array"
// @Router /expresss [get]
func (ctl *ExpressController) List(c *gin.Context) {
	express := &models.Express{}
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
	expresss, total, err := express.List(rawQuery, rawOrder, offset, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"total": total,
		"data":  expresss,
	})
}

// @Summary Get
// @Tags    Express
// @Param  expressId path string true "expressId"
// @Success 200 {object} models.Express "express object"
// @Router /expresss/{expressId} [get]
func (ctl *ExpressController) Get(c *gin.Context) {
	express := &models.Express{}
	id := c.Param("expressId")

	var err error
	express.Id, err = strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	express, err = express.Get()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, express)
}
