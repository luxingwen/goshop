//generate by gen
package controllers

import (
	"github.com/gin-gonic/gin"
	"goshop/restful/models"
	"net/http"
	"strconv"
)

type StoreCategoryController struct {
}

// @Summary Create
// @Tags    StoreCategory
// @Param body body models.StoreCategory true "StoreCategory"
// @Success 200 {string} string ""
// @Router /storeCategorys [post]
func (ctl *StoreCategoryController) Create(c *gin.Context) {
	storeCategory := models.StoreCategory{}
	if err := ParseRequest(c, &storeCategory); err != nil {
		return
	}
	if err := storeCategory.Insert(); err != nil {
		c.JSON(http.StatusBadGateway, err)
		return
	}
	c.JSON(http.StatusOK, storeCategory)
}

// @Summary  Delete
// @Tags     StoreCategory
// @Param  storeCategoryId  path string true "storeCategoryId"
// @Success 200 {string} string ""
// @Router /storeCategorys/{storeCategoryId} [delete]
func (ctl *StoreCategoryController) Delete(c *gin.Context) {
	storeCategory := models.StoreCategory{}
	id := c.Param("storeCategoryId")
	var err error
	storeCategory.Id, err = strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	err = storeCategory.Delete()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, nil)
}

// @Summary Put
// @Tags    StoreCategory
// @Param body body models.StoreCategory true "storeCategory"
// @Param  storeCategoryId path string true "storeCategoryId"
// @Success 200 {string} string ""
// @Router /storeCategorys/{storeCategoryId} [put]
func (ctl *StoreCategoryController) Put(c *gin.Context) {
	storeCategory := models.StoreCategory{}
	id := c.Param("storeCategoryId")
	var err error
	storeCategory.Id, err = strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	if err := ParseRequest(c, &storeCategory); err != nil {
		return
	}
	err = storeCategory.Update()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, nil)
}

// @Summary Patch
// @Tags    StoreCategory
// @Param body body models.StoreCategory true "storeCategory"
// @Param  storeCategoryId path string true "storeCategoryId"
// @Success 200 {string} string ""
// @Router /storeCategorys/{storeCategoryId} [patch]
func (ctl *StoreCategoryController) Patch(c *gin.Context) {
	storeCategory := models.StoreCategory{}
	id := c.Param("storeCategoryId")
	var err error
	storeCategory.Id, err = strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	if err := ParseRequest(c, &storeCategory); err != nil {
		return
	}
	err = storeCategory.Patch()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, nil)
}

// @Summary List
// @Tags    StoreCategory
// @Param query query string false "query, ?query=age:>:21,name:like:%jason%"
// @Param order query string false "order, ?order=age:desc,created_at:asc"
// @Param page query int false "page"
// @Param pageSize query int false "pageSize"
// @Success 200 {array} models.StoreCategory "storeCategory array"
// @Router /storeCategorys [get]
func (ctl *StoreCategoryController) List(c *gin.Context) {
	storeCategory := &models.StoreCategory{}
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
	storeCategorys, total, err := storeCategory.List(rawQuery, rawOrder, offset, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"total": total,
		"data":  storeCategorys,
	})
}

// @Summary Get
// @Tags    StoreCategory
// @Param  storeCategoryId path string true "storeCategoryId"
// @Success 200 {object} models.StoreCategory "storeCategory object"
// @Router /storeCategorys/{storeCategoryId} [get]
func (ctl *StoreCategoryController) Get(c *gin.Context) {
	storeCategory := &models.StoreCategory{}
	id := c.Param("storeCategoryId")

	var err error
	storeCategory.Id, err = strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	storeCategory, err = storeCategory.Get()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, storeCategory)
}
