//generate by gen
package controllers

import (
	"github.com/gin-gonic/gin"
	"goshop/restful/models"
	"net/http"
	"strconv"
)

type StoreVisitController struct {
}

// @Summary Create
// @Tags    StoreVisit
// @Param body body models.StoreVisit true "StoreVisit"
// @Success 200 {string} string ""
// @Router /storeVisits [post]
func (ctl *StoreVisitController) Create(c *gin.Context) {
	storeVisit := models.StoreVisit{}
	if err := ParseRequest(c, &storeVisit); err != nil {
		return
	}
	if err := storeVisit.Insert(); err != nil {
		c.JSON(http.StatusBadGateway, err)
		return
	}
	c.JSON(http.StatusOK, storeVisit)
}

// @Summary  Delete
// @Tags     StoreVisit
// @Param  storeVisitId  path string true "storeVisitId"
// @Success 200 {string} string ""
// @Router /storeVisits/{storeVisitId} [delete]
func (ctl *StoreVisitController) Delete(c *gin.Context) {
	storeVisit := models.StoreVisit{}
	id := c.Param("storeVisitId")
	var err error
	storeVisit.Id, err = strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	err = storeVisit.Delete()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, nil)
}

// @Summary Put
// @Tags    StoreVisit
// @Param body body models.StoreVisit true "storeVisit"
// @Param  storeVisitId path string true "storeVisitId"
// @Success 200 {string} string ""
// @Router /storeVisits/{storeVisitId} [put]
func (ctl *StoreVisitController) Put(c *gin.Context) {
	storeVisit := models.StoreVisit{}
	id := c.Param("storeVisitId")
	var err error
	storeVisit.Id, err = strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	if err := ParseRequest(c, &storeVisit); err != nil {
		return
	}
	err = storeVisit.Update()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, nil)
}

// @Summary Patch
// @Tags    StoreVisit
// @Param body body models.StoreVisit true "storeVisit"
// @Param  storeVisitId path string true "storeVisitId"
// @Success 200 {string} string ""
// @Router /storeVisits/{storeVisitId} [patch]
func (ctl *StoreVisitController) Patch(c *gin.Context) {
	storeVisit := models.StoreVisit{}
	id := c.Param("storeVisitId")
	var err error
	storeVisit.Id, err = strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	if err := ParseRequest(c, &storeVisit); err != nil {
		return
	}
	err = storeVisit.Patch()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, nil)
}

// @Summary List
// @Tags    StoreVisit
// @Param query query string false "query, ?query=age:>:21,name:like:%jason%"
// @Param order query string false "order, ?order=age:desc,created_at:asc"
// @Param page query int false "page"
// @Param pageSize query int false "pageSize"
// @Success 200 {array} models.StoreVisit "storeVisit array"
// @Router /storeVisits [get]
func (ctl *StoreVisitController) List(c *gin.Context) {
	storeVisit := &models.StoreVisit{}
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
	storeVisits, total, err := storeVisit.List(rawQuery, rawOrder, offset, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"total": total,
		"data":  storeVisits,
	})
}

// @Summary Get
// @Tags    StoreVisit
// @Param  storeVisitId path string true "storeVisitId"
// @Success 200 {object} models.StoreVisit "storeVisit object"
// @Router /storeVisits/{storeVisitId} [get]
func (ctl *StoreVisitController) Get(c *gin.Context) {
	storeVisit := &models.StoreVisit{}
	id := c.Param("storeVisitId")

	var err error
	storeVisit.Id, err = strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	storeVisit, err = storeVisit.Get()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, storeVisit)
}
