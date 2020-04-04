//generate by gen
package controllers

import (
	"github.com/gin-gonic/gin"
	"goshop/restful/models"
	"net/http"
	"strconv"
)

type StoreServiceController struct {
}

// @Summary Create
// @Tags    StoreService
// @Param body body models.StoreService true "StoreService"
// @Success 200 {string} string ""
// @Router /storeServices [post]
func (ctl *StoreServiceController) Create(c *gin.Context) {
	storeService := models.StoreService{}
	if err := ParseRequest(c, &storeService); err != nil {
		return
	}
	if err := storeService.Insert(); err != nil {
		c.JSON(http.StatusBadGateway, err)
		return
	}
	c.JSON(http.StatusOK, storeService)
}

// @Summary  Delete
// @Tags     StoreService
// @Param  storeServiceId  path string true "storeServiceId"
// @Success 200 {string} string ""
// @Router /storeServices/{storeServiceId} [delete]
func (ctl *StoreServiceController) Delete(c *gin.Context) {
	storeService := models.StoreService{}
	id := c.Param("storeServiceId")
	var err error
	storeService.Id, err = strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	err = storeService.Delete()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, nil)
}

// @Summary Put
// @Tags    StoreService
// @Param body body models.StoreService true "storeService"
// @Param  storeServiceId path string true "storeServiceId"
// @Success 200 {string} string ""
// @Router /storeServices/{storeServiceId} [put]
func (ctl *StoreServiceController) Put(c *gin.Context) {
	storeService := models.StoreService{}
	id := c.Param("storeServiceId")
	var err error
	storeService.Id, err = strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	if err := ParseRequest(c, &storeService); err != nil {
		return
	}
	err = storeService.Update()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, nil)
}

// @Summary Patch
// @Tags    StoreService
// @Param body body models.StoreService true "storeService"
// @Param  storeServiceId path string true "storeServiceId"
// @Success 200 {string} string ""
// @Router /storeServices/{storeServiceId} [patch]
func (ctl *StoreServiceController) Patch(c *gin.Context) {
	storeService := models.StoreService{}
	id := c.Param("storeServiceId")
	var err error
	storeService.Id, err = strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	if err := ParseRequest(c, &storeService); err != nil {
		return
	}
	err = storeService.Patch()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, nil)
}

// @Summary List
// @Tags    StoreService
// @Param query query string false "query, ?query=age:>:21,name:like:%jason%"
// @Param order query string false "order, ?order=age:desc,created_at:asc"
// @Param page query int false "page"
// @Param pageSize query int false "pageSize"
// @Success 200 {array} models.StoreService "storeService array"
// @Router /storeServices [get]
func (ctl *StoreServiceController) List(c *gin.Context) {
	storeService := &models.StoreService{}
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
	storeServices, total, err := storeService.List(rawQuery, rawOrder, offset, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"total": total,
		"data":  storeServices,
	})
}

// @Summary Get
// @Tags    StoreService
// @Param  storeServiceId path string true "storeServiceId"
// @Success 200 {object} models.StoreService "storeService object"
// @Router /storeServices/{storeServiceId} [get]
func (ctl *StoreServiceController) Get(c *gin.Context) {
	storeService := &models.StoreService{}
	id := c.Param("storeServiceId")

	var err error
	storeService.Id, err = strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	storeService, err = storeService.Get()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, storeService)
}
