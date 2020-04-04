//generate by gen
package controllers

import (
	"github.com/gin-gonic/gin"
	"goshop/restful/models"
	"net/http"
	"strconv"
)

type RoutineTemplateController struct {
}

// @Summary Create
// @Tags    RoutineTemplate
// @Param body body models.RoutineTemplate true "RoutineTemplate"
// @Success 200 {string} string ""
// @Router /routineTemplates [post]
func (ctl *RoutineTemplateController) Create(c *gin.Context) {
	routineTemplate := models.RoutineTemplate{}
	if err := ParseRequest(c, &routineTemplate); err != nil {
		return
	}
	if err := routineTemplate.Insert(); err != nil {
		c.JSON(http.StatusBadGateway, err)
		return
	}
	c.JSON(http.StatusOK, routineTemplate)
}

// @Summary  Delete
// @Tags     RoutineTemplate
// @Param  routineTemplateId  path string true "routineTemplateId"
// @Success 200 {string} string ""
// @Router /routineTemplates/{routineTemplateId} [delete]
func (ctl *RoutineTemplateController) Delete(c *gin.Context) {
	routineTemplate := models.RoutineTemplate{}
	id := c.Param("routineTemplateId")
	var err error
	routineTemplate.Id, err = strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	err = routineTemplate.Delete()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, nil)
}

// @Summary Put
// @Tags    RoutineTemplate
// @Param body body models.RoutineTemplate true "routineTemplate"
// @Param  routineTemplateId path string true "routineTemplateId"
// @Success 200 {string} string ""
// @Router /routineTemplates/{routineTemplateId} [put]
func (ctl *RoutineTemplateController) Put(c *gin.Context) {
	routineTemplate := models.RoutineTemplate{}
	id := c.Param("routineTemplateId")
	var err error
	routineTemplate.Id, err = strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	if err := ParseRequest(c, &routineTemplate); err != nil {
		return
	}
	err = routineTemplate.Update()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, nil)
}

// @Summary Patch
// @Tags    RoutineTemplate
// @Param body body models.RoutineTemplate true "routineTemplate"
// @Param  routineTemplateId path string true "routineTemplateId"
// @Success 200 {string} string ""
// @Router /routineTemplates/{routineTemplateId} [patch]
func (ctl *RoutineTemplateController) Patch(c *gin.Context) {
	routineTemplate := models.RoutineTemplate{}
	id := c.Param("routineTemplateId")
	var err error
	routineTemplate.Id, err = strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	if err := ParseRequest(c, &routineTemplate); err != nil {
		return
	}
	err = routineTemplate.Patch()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, nil)
}

// @Summary List
// @Tags    RoutineTemplate
// @Param query query string false "query, ?query=age:>:21,name:like:%jason%"
// @Param order query string false "order, ?order=age:desc,created_at:asc"
// @Param page query int false "page"
// @Param pageSize query int false "pageSize"
// @Success 200 {array} models.RoutineTemplate "routineTemplate array"
// @Router /routineTemplates [get]
func (ctl *RoutineTemplateController) List(c *gin.Context) {
	routineTemplate := &models.RoutineTemplate{}
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
	routineTemplates, total, err := routineTemplate.List(rawQuery, rawOrder, offset, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"total": total,
		"data":  routineTemplates,
	})
}

// @Summary Get
// @Tags    RoutineTemplate
// @Param  routineTemplateId path string true "routineTemplateId"
// @Success 200 {object} models.RoutineTemplate "routineTemplate object"
// @Router /routineTemplates/{routineTemplateId} [get]
func (ctl *RoutineTemplateController) Get(c *gin.Context) {
	routineTemplate := &models.RoutineTemplate{}
	id := c.Param("routineTemplateId")

	var err error
	routineTemplate.Id, err = strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	routineTemplate, err = routineTemplate.Get()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, routineTemplate)
}
