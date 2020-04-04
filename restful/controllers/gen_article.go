//generate by gen
package controllers

import (
	"github.com/gin-gonic/gin"
	"goshop/restful/models"
	"net/http"
	"strconv"
)

type ArticleController struct {
}

// @Summary Create
// @Tags    Article
// @Param body body models.Article true "Article"
// @Success 200 {string} string ""
// @Router /articles [post]
func (ctl *ArticleController) Create(c *gin.Context) {
	article := models.Article{}
	if err := ParseRequest(c, &article); err != nil {
		return
	}
	if err := article.Insert(); err != nil {
		c.JSON(http.StatusBadGateway, err)
		return
	}
	c.JSON(http.StatusOK, article)
}

// @Summary  Delete
// @Tags     Article
// @Param  articleId  path string true "articleId"
// @Success 200 {string} string ""
// @Router /articles/{articleId} [delete]
func (ctl *ArticleController) Delete(c *gin.Context) {
	article := models.Article{}
	id := c.Param("articleId")
	var err error
	article.Id, err = strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	err = article.Delete()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, nil)
}

// @Summary Put
// @Tags    Article
// @Param body body models.Article true "article"
// @Param  articleId path string true "articleId"
// @Success 200 {string} string ""
// @Router /articles/{articleId} [put]
func (ctl *ArticleController) Put(c *gin.Context) {
	article := models.Article{}
	id := c.Param("articleId")
	var err error
	article.Id, err = strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	if err := ParseRequest(c, &article); err != nil {
		return
	}
	err = article.Update()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, nil)
}

// @Summary Patch
// @Tags    Article
// @Param body body models.Article true "article"
// @Param  articleId path string true "articleId"
// @Success 200 {string} string ""
// @Router /articles/{articleId} [patch]
func (ctl *ArticleController) Patch(c *gin.Context) {
	article := models.Article{}
	id := c.Param("articleId")
	var err error
	article.Id, err = strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	if err := ParseRequest(c, &article); err != nil {
		return
	}
	err = article.Patch()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, nil)
}

// @Summary List
// @Tags    Article
// @Param query query string false "query, ?query=age:>:21,name:like:%jason%"
// @Param order query string false "order, ?order=age:desc,created_at:asc"
// @Param page query int false "page"
// @Param pageSize query int false "pageSize"
// @Success 200 {array} models.Article "article array"
// @Router /articles [get]
func (ctl *ArticleController) List(c *gin.Context) {
	article := &models.Article{}
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
	articles, total, err := article.List(rawQuery, rawOrder, offset, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"total": total,
		"data":  articles,
	})
}

// @Summary Get
// @Tags    Article
// @Param  articleId path string true "articleId"
// @Success 200 {object} models.Article "article object"
// @Router /articles/{articleId} [get]
func (ctl *ArticleController) Get(c *gin.Context) {
	article := &models.Article{}
	id := c.Param("articleId")

	var err error
	article.Id, err = strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	article, err = article.Get()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, article)
}
