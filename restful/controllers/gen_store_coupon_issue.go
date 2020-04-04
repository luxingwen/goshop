//generate by gen
package controllers

import (
	"github.com/gin-gonic/gin"
	"goshop/restful/models"
	"net/http"
	"strconv"
)

type StoreCouponIssueController struct {
}

// @Summary Create
// @Tags    StoreCouponIssue
// @Param body body models.StoreCouponIssue true "StoreCouponIssue"
// @Success 200 {string} string ""
// @Router /storeCouponIssues [post]
func (ctl *StoreCouponIssueController) Create(c *gin.Context) {
	storeCouponIssue := models.StoreCouponIssue{}
	if err := ParseRequest(c, &storeCouponIssue); err != nil {
		return
	}
	if err := storeCouponIssue.Insert(); err != nil {
		c.JSON(http.StatusBadGateway, err)
		return
	}
	c.JSON(http.StatusOK, storeCouponIssue)
}

// @Summary  Delete
// @Tags     StoreCouponIssue
// @Param  storeCouponIssueId  path string true "storeCouponIssueId"
// @Success 200 {string} string ""
// @Router /storeCouponIssues/{storeCouponIssueId} [delete]
func (ctl *StoreCouponIssueController) Delete(c *gin.Context) {
	storeCouponIssue := models.StoreCouponIssue{}
	id := c.Param("storeCouponIssueId")
	var err error
	storeCouponIssue.Id, err = strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	err = storeCouponIssue.Delete()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, nil)
}

// @Summary Put
// @Tags    StoreCouponIssue
// @Param body body models.StoreCouponIssue true "storeCouponIssue"
// @Param  storeCouponIssueId path string true "storeCouponIssueId"
// @Success 200 {string} string ""
// @Router /storeCouponIssues/{storeCouponIssueId} [put]
func (ctl *StoreCouponIssueController) Put(c *gin.Context) {
	storeCouponIssue := models.StoreCouponIssue{}
	id := c.Param("storeCouponIssueId")
	var err error
	storeCouponIssue.Id, err = strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	if err := ParseRequest(c, &storeCouponIssue); err != nil {
		return
	}
	err = storeCouponIssue.Update()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, nil)
}

// @Summary Patch
// @Tags    StoreCouponIssue
// @Param body body models.StoreCouponIssue true "storeCouponIssue"
// @Param  storeCouponIssueId path string true "storeCouponIssueId"
// @Success 200 {string} string ""
// @Router /storeCouponIssues/{storeCouponIssueId} [patch]
func (ctl *StoreCouponIssueController) Patch(c *gin.Context) {
	storeCouponIssue := models.StoreCouponIssue{}
	id := c.Param("storeCouponIssueId")
	var err error
	storeCouponIssue.Id, err = strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	if err := ParseRequest(c, &storeCouponIssue); err != nil {
		return
	}
	err = storeCouponIssue.Patch()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, nil)
}

// @Summary List
// @Tags    StoreCouponIssue
// @Param query query string false "query, ?query=age:>:21,name:like:%jason%"
// @Param order query string false "order, ?order=age:desc,created_at:asc"
// @Param page query int false "page"
// @Param pageSize query int false "pageSize"
// @Success 200 {array} models.StoreCouponIssue "storeCouponIssue array"
// @Router /storeCouponIssues [get]
func (ctl *StoreCouponIssueController) List(c *gin.Context) {
	storeCouponIssue := &models.StoreCouponIssue{}
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
	storeCouponIssues, total, err := storeCouponIssue.List(rawQuery, rawOrder, offset, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"total": total,
		"data":  storeCouponIssues,
	})
}

// @Summary Get
// @Tags    StoreCouponIssue
// @Param  storeCouponIssueId path string true "storeCouponIssueId"
// @Success 200 {object} models.StoreCouponIssue "storeCouponIssue object"
// @Router /storeCouponIssues/{storeCouponIssueId} [get]
func (ctl *StoreCouponIssueController) Get(c *gin.Context) {
	storeCouponIssue := &models.StoreCouponIssue{}
	id := c.Param("storeCouponIssueId")

	var err error
	storeCouponIssue.Id, err = strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	storeCouponIssue, err = storeCouponIssue.Get()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, storeCouponIssue)
}
