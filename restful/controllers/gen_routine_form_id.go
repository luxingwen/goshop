//generate by gen
package controllers

import (
	"github.com/gin-gonic/gin"
	"goshop/restful/models"
	"net/http"
	"strconv"
)

type RoutineFormIdController struct {
}

// @Summary Create
// @Tags    RoutineFormId
// @Param body body models.RoutineFormId true "RoutineFormId"
// @Success 200 {string} string ""
// @Router /routineFormIds [post]
func (ctl *RoutineFormIdController) Create(c *gin.Context) {
	routineFormId := models.RoutineFormId{}
	if err := ParseRequest(c, &routineFormId); err != nil {
		return
	}
	if err := routineFormId.Insert(); err != nil {
		c.JSON(http.StatusBadGateway, err)
		return
	}
	c.JSON(http.StatusOK, routineFormId)
}

// @Summary  Delete
// @Tags     RoutineFormId
// @Param  routineFormIdId  path string true "routineFormIdId"
// @Success 200 {string} string ""
// @Router /routineFormIds/{routineFormIdId} [delete]
func (ctl *RoutineFormIdController) Delete(c *gin.Context) {
	routineFormId := models.RoutineFormId{}
	id := c.Param("routineFormIdId")
	var err error
	routineFormId.Id, err = strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	err = routineFormId.Delete()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, nil)
}

// @Summary Put
// @Tags    RoutineFormId
// @Param body body models.RoutineFormId true "routineFormId"
// @Param  routineFormIdId path string true "routineFormIdId"
// @Success 200 {string} string ""
// @Router /routineFormIds/{routineFormIdId} [put]
func (ctl *RoutineFormIdController) Put(c *gin.Context) {
	routineFormId := models.RoutineFormId{}
	id := c.Param("routineFormIdId")
	var err error
	routineFormId.Id, err = strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	if err := ParseRequest(c, &routineFormId); err != nil {
		return
	}
	err = routineFormId.Update()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, nil)
}

// @Summary Patch
// @Tags    RoutineFormId
// @Param body body models.RoutineFormId true "routineFormId"
// @Param  routineFormIdId path string true "routineFormIdId"
// @Success 200 {string} string ""
// @Router /routineFormIds/{routineFormIdId} [patch]
func (ctl *RoutineFormIdController) Patch(c *gin.Context) {
	routineFormId := models.RoutineFormId{}
	id := c.Param("routineFormIdId")
	var err error
	routineFormId.Id, err = strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	if err := ParseRequest(c, &routineFormId); err != nil {
		return
	}
	err = routineFormId.Patch()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, nil)
}

// @Summary List
// @Tags    RoutineFormId
// @Param query query string false "query, ?query=age:>:21,name:like:%jason%"
// @Param order query string false "order, ?order=age:desc,created_at:asc"
// @Param page query int false "page"
// @Param pageSize query int false "pageSize"
// @Success 200 {array} models.RoutineFormId "routineFormId array"
// @Router /routineFormIds [get]
func (ctl *RoutineFormIdController) List(c *gin.Context) {
	routineFormId := &models.RoutineFormId{}
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
	routineFormIds, total, err := routineFormId.List(rawQuery, rawOrder, offset, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"total": total,
		"data":  routineFormIds,
	})
}

// @Summary Get
// @Tags    RoutineFormId
// @Param  routineFormIdId path string true "routineFormIdId"
// @Success 200 {object} models.RoutineFormId "routineFormId object"
// @Router /routineFormIds/{routineFormIdId} [get]
func (ctl *RoutineFormIdController) Get(c *gin.Context) {
	routineFormId := &models.RoutineFormId{}
	id := c.Param("routineFormIdId")

	var err error
	routineFormId.Id, err = strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	routineFormId, err = routineFormId.Get()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, routineFormId)
}
