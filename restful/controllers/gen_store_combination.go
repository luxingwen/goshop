//generate by gen
package controllers

import (
	"github.com/gin-gonic/gin"
	"goshop/restful/models"
	"net/http"
	"strconv"
)

type StoreCombinationController struct {
}

// @Summary Create
// @Tags    StoreCombination
// @Param body body models.StoreCombination true "StoreCombination"
// @Success 200 {string} string ""
// @Router /storeCombinations [post]
func (ctl *StoreCombinationController) Create(c *gin.Context) {
	storeCombination := models.StoreCombination{}
	if err := ParseRequest(c, &storeCombination); err != nil {
		return
	}
	if err := storeCombination.Insert(); err != nil {
		c.JSON(http.StatusBadGateway, err)
		return
	}
	c.JSON(http.StatusOK, storeCombination)
}

// @Summary  Delete
// @Tags     StoreCombination
// @Param  storeCombinationId  path string true "storeCombinationId"
// @Success 200 {string} string ""
// @Router /storeCombinations/{storeCombinationId} [delete]
func (ctl *StoreCombinationController) Delete(c *gin.Context) {
	storeCombination := models.StoreCombination{}
	id := c.Param("storeCombinationId")
	var err error
	storeCombination.Id, err = strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	err = storeCombination.Delete()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, nil)
}

// @Summary Put
// @Tags    StoreCombination
// @Param body body models.StoreCombination true "storeCombination"
// @Param  storeCombinationId path string true "storeCombinationId"
// @Success 200 {string} string ""
// @Router /storeCombinations/{storeCombinationId} [put]
func (ctl *StoreCombinationController) Put(c *gin.Context) {
	storeCombination := models.StoreCombination{}
	id := c.Param("storeCombinationId")
	var err error
	storeCombination.Id, err = strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	if err := ParseRequest(c, &storeCombination); err != nil {
		return
	}
	err = storeCombination.Update()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, nil)
}

// @Summary Patch
// @Tags    StoreCombination
// @Param body body models.StoreCombination true "storeCombination"
// @Param  storeCombinationId path string true "storeCombinationId"
// @Success 200 {string} string ""
// @Router /storeCombinations/{storeCombinationId} [patch]
func (ctl *StoreCombinationController) Patch(c *gin.Context) {
	storeCombination := models.StoreCombination{}
	id := c.Param("storeCombinationId")
	var err error
	storeCombination.Id, err = strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	if err := ParseRequest(c, &storeCombination); err != nil {
		return
	}
	err = storeCombination.Patch()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, nil)
}

// @Summary List
// @Tags    StoreCombination
// @Param query query string false "query, ?query=age:>:21,name:like:%jason%"
// @Param order query string false "order, ?order=age:desc,created_at:asc"
// @Param page query int false "page"
// @Param pageSize query int false "pageSize"
// @Success 200 {array} models.StoreCombination "storeCombination array"
// @Router /storeCombinations [get]
func (ctl *StoreCombinationController) List(c *gin.Context) {
	storeCombination := &models.StoreCombination{}
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
	storeCombinations, total, err := storeCombination.List(rawQuery, rawOrder, offset, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"total": total,
		"data":  storeCombinations,
	})
}

// @Summary Get
// @Tags    StoreCombination
// @Param  storeCombinationId path string true "storeCombinationId"
// @Success 200 {object} models.StoreCombination "storeCombination object"
// @Router /storeCombinations/{storeCombinationId} [get]
func (ctl *StoreCombinationController) Get(c *gin.Context) {
	storeCombination := &models.StoreCombination{}
	id := c.Param("storeCombinationId")

	var err error
	storeCombination.Id, err = strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	storeCombination, err = storeCombination.Get()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, storeCombination)
}
