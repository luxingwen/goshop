//generate by gen
package controllers

import (
	"github.com/gin-gonic/gin"
	"goshop/restful/models"
	"net/http"
	"strconv"
)

type ArticleCategoryController struct {
}

// @Summary Create
// @Tags    ArticleCategory
// @Param body body models.ArticleCategory true "ArticleCategory"
// @Success 200 {string} string ""
// @Router /articleCategorys [post]
func (ctl *ArticleCategoryController) Create(c *gin.Context) {
	articleCategory := models.ArticleCategory{}
	if err := ParseRequest(c, &articleCategory); err != nil {
		return
	}
	if err := articleCategory.Insert(); err != nil {
		c.JSON(http.StatusBadGateway, err)
		return
	}
	c.JSON(http.StatusOK, articleCategory)
}

// @Summary  Delete
// @Tags     ArticleCategory
// @Param  articleCategoryId  path string true "articleCategoryId"
// @Success 200 {string} string ""
// @Router /articleCategorys/{articleCategoryId} [delete]
func (ctl *ArticleCategoryController) Delete(c *gin.Context) {
	articleCategory := models.ArticleCategory{}
	id := c.Param("articleCategoryId")
	var err error
	articleCategory.Id, err = strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	err = articleCategory.Delete()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, nil)
}

// @Summary Put
// @Tags    ArticleCategory
// @Param body body models.ArticleCategory true "articleCategory"
// @Param  articleCategoryId path string true "articleCategoryId"
// @Success 200 {string} string ""
// @Router /articleCategorys/{articleCategoryId} [put]
func (ctl *ArticleCategoryController) Put(c *gin.Context) {
	articleCategory := models.ArticleCategory{}
	id := c.Param("articleCategoryId")
	var err error
	articleCategory.Id, err = strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	if err := ParseRequest(c, &articleCategory); err != nil {
		return
	}
	err = articleCategory.Update()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, nil)
}

// @Summary Patch
// @Tags    ArticleCategory
// @Param body body models.ArticleCategory true "articleCategory"
// @Param  articleCategoryId path string true "articleCategoryId"
// @Success 200 {string} string ""
// @Router /articleCategorys/{articleCategoryId} [patch]
func (ctl *ArticleCategoryController) Patch(c *gin.Context) {
	articleCategory := models.ArticleCategory{}
	id := c.Param("articleCategoryId")
	var err error
	articleCategory.Id, err = strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	if err := ParseRequest(c, &articleCategory); err != nil {
		return
	}
	err = articleCategory.Patch()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, nil)
}

// @Summary List
// @Tags    ArticleCategory
// @Param query query string false "query, ?query=age:>:21,name:like:%jason%"
// @Param order query string false "order, ?order=age:desc,created_at:asc"
// @Param page query int false "page"
// @Param pageSize query int false "pageSize"
// @Success 200 {array} models.ArticleCategory "articleCategory array"
// @Router /articleCategorys [get]
func (ctl *ArticleCategoryController) List(c *gin.Context) {
	articleCategory := &models.ArticleCategory{}
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
	articleCategorys, total, err := articleCategory.List(rawQuery, rawOrder, offset, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"total": total,
		"data":  articleCategorys,
	})
}

// @Summary Get
// @Tags    ArticleCategory
// @Param  articleCategoryId path string true "articleCategoryId"
// @Success 200 {object} models.ArticleCategory "articleCategory object"
// @Router /articleCategorys/{articleCategoryId} [get]
func (ctl *ArticleCategoryController) Get(c *gin.Context) {
	articleCategory := &models.ArticleCategory{}
	id := c.Param("articleCategoryId")

	var err error
	articleCategory.Id, err = strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	articleCategory, err = articleCategory.Get()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, articleCategory)
}
