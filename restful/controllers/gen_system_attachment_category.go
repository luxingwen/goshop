//generate by gen
package controllers

import (
	"github.com/gin-gonic/gin"
	"goshop/restful/models"
	"net/http"
	"strconv"
)

type SystemAttachmentCategoryController struct {
}

// @Summary Create
// @Tags    SystemAttachmentCategory
// @Param body body models.SystemAttachmentCategory true "SystemAttachmentCategory"
// @Success 200 {string} string ""
// @Router /systemAttachmentCategorys [post]
func (ctl *SystemAttachmentCategoryController) Create(c *gin.Context) {
	systemAttachmentCategory := models.SystemAttachmentCategory{}
	if err := ParseRequest(c, &systemAttachmentCategory); err != nil {
		return
	}
	if err := systemAttachmentCategory.Insert(); err != nil {
		c.JSON(http.StatusBadGateway, err)
		return
	}
	c.JSON(http.StatusOK, systemAttachmentCategory)
}

// @Summary  Delete
// @Tags     SystemAttachmentCategory
// @Param  systemAttachmentCategoryId  path string true "systemAttachmentCategoryId"
// @Success 200 {string} string ""
// @Router /systemAttachmentCategorys/{systemAttachmentCategoryId} [delete]
func (ctl *SystemAttachmentCategoryController) Delete(c *gin.Context) {
	systemAttachmentCategory := models.SystemAttachmentCategory{}
	id := c.Param("systemAttachmentCategoryId")
	var err error
	systemAttachmentCategory.Id, err = strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	err = systemAttachmentCategory.Delete()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, nil)
}

// @Summary Put
// @Tags    SystemAttachmentCategory
// @Param body body models.SystemAttachmentCategory true "systemAttachmentCategory"
// @Param  systemAttachmentCategoryId path string true "systemAttachmentCategoryId"
// @Success 200 {string} string ""
// @Router /systemAttachmentCategorys/{systemAttachmentCategoryId} [put]
func (ctl *SystemAttachmentCategoryController) Put(c *gin.Context) {
	systemAttachmentCategory := models.SystemAttachmentCategory{}
	id := c.Param("systemAttachmentCategoryId")
	var err error
	systemAttachmentCategory.Id, err = strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	if err := ParseRequest(c, &systemAttachmentCategory); err != nil {
		return
	}
	err = systemAttachmentCategory.Update()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, nil)
}

// @Summary Patch
// @Tags    SystemAttachmentCategory
// @Param body body models.SystemAttachmentCategory true "systemAttachmentCategory"
// @Param  systemAttachmentCategoryId path string true "systemAttachmentCategoryId"
// @Success 200 {string} string ""
// @Router /systemAttachmentCategorys/{systemAttachmentCategoryId} [patch]
func (ctl *SystemAttachmentCategoryController) Patch(c *gin.Context) {
	systemAttachmentCategory := models.SystemAttachmentCategory{}
	id := c.Param("systemAttachmentCategoryId")
	var err error
	systemAttachmentCategory.Id, err = strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	if err := ParseRequest(c, &systemAttachmentCategory); err != nil {
		return
	}
	err = systemAttachmentCategory.Patch()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, nil)
}

// @Summary List
// @Tags    SystemAttachmentCategory
// @Param query query string false "query, ?query=age:>:21,name:like:%jason%"
// @Param order query string false "order, ?order=age:desc,created_at:asc"
// @Param page query int false "page"
// @Param pageSize query int false "pageSize"
// @Success 200 {array} models.SystemAttachmentCategory "systemAttachmentCategory array"
// @Router /systemAttachmentCategorys [get]
func (ctl *SystemAttachmentCategoryController) List(c *gin.Context) {
	systemAttachmentCategory := &models.SystemAttachmentCategory{}
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
	systemAttachmentCategorys, total, err := systemAttachmentCategory.List(rawQuery, rawOrder, offset, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"total": total,
		"data":  systemAttachmentCategorys,
	})
}

// @Summary Get
// @Tags    SystemAttachmentCategory
// @Param  systemAttachmentCategoryId path string true "systemAttachmentCategoryId"
// @Success 200 {object} models.SystemAttachmentCategory "systemAttachmentCategory object"
// @Router /systemAttachmentCategorys/{systemAttachmentCategoryId} [get]
func (ctl *SystemAttachmentCategoryController) Get(c *gin.Context) {
	systemAttachmentCategory := &models.SystemAttachmentCategory{}
	id := c.Param("systemAttachmentCategoryId")

	var err error
	systemAttachmentCategory.Id, err = strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	systemAttachmentCategory, err = systemAttachmentCategory.Get()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, systemAttachmentCategory)
}
