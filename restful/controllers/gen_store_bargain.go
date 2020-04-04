//generate by gen
package controllers

import (
	"github.com/gin-gonic/gin"
	"goshop/restful/models"
	"net/http"
	"strconv"
)

type StoreBargainController struct {
}

// @Summary Create
// @Tags    StoreBargain
// @Param body body models.StoreBargain true "StoreBargain"
// @Success 200 {string} string ""
// @Router /storeBargains [post]
func (ctl *StoreBargainController) Create(c *gin.Context) {
	storeBargain := models.StoreBargain{}
	if err := ParseRequest(c, &storeBargain); err != nil {
		return
	}
	if err := storeBargain.Insert(); err != nil {
		c.JSON(http.StatusBadGateway, err)
		return
	}
	c.JSON(http.StatusOK, storeBargain)
}

// @Summary  Delete
// @Tags     StoreBargain
// @Param  storeBargainId  path string true "storeBargainId"
// @Success 200 {string} string ""
// @Router /storeBargains/{storeBargainId} [delete]
func (ctl *StoreBargainController) Delete(c *gin.Context) {
	storeBargain := models.StoreBargain{}
	id := c.Param("storeBargainId")
	var err error
	storeBargain.Id, err = strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	err = storeBargain.Delete()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, nil)
}

// @Summary Put
// @Tags    StoreBargain
// @Param body body models.StoreBargain true "storeBargain"
// @Param  storeBargainId path string true "storeBargainId"
// @Success 200 {string} string ""
// @Router /storeBargains/{storeBargainId} [put]
func (ctl *StoreBargainController) Put(c *gin.Context) {
	storeBargain := models.StoreBargain{}
	id := c.Param("storeBargainId")
	var err error
	storeBargain.Id, err = strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	if err := ParseRequest(c, &storeBargain); err != nil {
		return
	}
	err = storeBargain.Update()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, nil)
}

// @Summary Patch
// @Tags    StoreBargain
// @Param body body models.StoreBargain true "storeBargain"
// @Param  storeBargainId path string true "storeBargainId"
// @Success 200 {string} string ""
// @Router /storeBargains/{storeBargainId} [patch]
func (ctl *StoreBargainController) Patch(c *gin.Context) {
	storeBargain := models.StoreBargain{}
	id := c.Param("storeBargainId")
	var err error
	storeBargain.Id, err = strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	if err := ParseRequest(c, &storeBargain); err != nil {
		return
	}
	err = storeBargain.Patch()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, nil)
}

// @Summary List
// @Tags    StoreBargain
// @Param query query string false "query, ?query=age:>:21,name:like:%jason%"
// @Param order query string false "order, ?order=age:desc,created_at:asc"
// @Param page query int false "page"
// @Param pageSize query int false "pageSize"
// @Success 200 {array} models.StoreBargain "storeBargain array"
// @Router /storeBargains [get]
func (ctl *StoreBargainController) List(c *gin.Context) {
	storeBargain := &models.StoreBargain{}
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
	storeBargains, total, err := storeBargain.List(rawQuery, rawOrder, offset, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"total": total,
		"data":  storeBargains,
	})
}

// @Summary Get
// @Tags    StoreBargain
// @Param  storeBargainId path string true "storeBargainId"
// @Success 200 {object} models.StoreBargain "storeBargain object"
// @Router /storeBargains/{storeBargainId} [get]
func (ctl *StoreBargainController) Get(c *gin.Context) {
	storeBargain := &models.StoreBargain{}
	id := c.Param("storeBargainId")

	var err error
	storeBargain.Id, err = strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	storeBargain, err = storeBargain.Get()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, storeBargain)
}
