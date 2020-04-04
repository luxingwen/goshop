//generate by gen
package controllers

import (
	"github.com/gin-gonic/gin"
	"goshop/restful/models"
	"net/http"
	"strconv"
)

type StoreBargainUserHelpController struct {
}

// @Summary Create
// @Tags    StoreBargainUserHelp
// @Param body body models.StoreBargainUserHelp true "StoreBargainUserHelp"
// @Success 200 {string} string ""
// @Router /storeBargainUserHelps [post]
func (ctl *StoreBargainUserHelpController) Create(c *gin.Context) {
	storeBargainUserHelp := models.StoreBargainUserHelp{}
	if err := ParseRequest(c, &storeBargainUserHelp); err != nil {
		return
	}
	if err := storeBargainUserHelp.Insert(); err != nil {
		c.JSON(http.StatusBadGateway, err)
		return
	}
	c.JSON(http.StatusOK, storeBargainUserHelp)
}

// @Summary  Delete
// @Tags     StoreBargainUserHelp
// @Param  storeBargainUserHelpId  path string true "storeBargainUserHelpId"
// @Success 200 {string} string ""
// @Router /storeBargainUserHelps/{storeBargainUserHelpId} [delete]
func (ctl *StoreBargainUserHelpController) Delete(c *gin.Context) {
	storeBargainUserHelp := models.StoreBargainUserHelp{}
	id := c.Param("storeBargainUserHelpId")
	var err error
	storeBargainUserHelp.Id, err = strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	err = storeBargainUserHelp.Delete()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, nil)
}

// @Summary Put
// @Tags    StoreBargainUserHelp
// @Param body body models.StoreBargainUserHelp true "storeBargainUserHelp"
// @Param  storeBargainUserHelpId path string true "storeBargainUserHelpId"
// @Success 200 {string} string ""
// @Router /storeBargainUserHelps/{storeBargainUserHelpId} [put]
func (ctl *StoreBargainUserHelpController) Put(c *gin.Context) {
	storeBargainUserHelp := models.StoreBargainUserHelp{}
	id := c.Param("storeBargainUserHelpId")
	var err error
	storeBargainUserHelp.Id, err = strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	if err := ParseRequest(c, &storeBargainUserHelp); err != nil {
		return
	}
	err = storeBargainUserHelp.Update()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, nil)
}

// @Summary Patch
// @Tags    StoreBargainUserHelp
// @Param body body models.StoreBargainUserHelp true "storeBargainUserHelp"
// @Param  storeBargainUserHelpId path string true "storeBargainUserHelpId"
// @Success 200 {string} string ""
// @Router /storeBargainUserHelps/{storeBargainUserHelpId} [patch]
func (ctl *StoreBargainUserHelpController) Patch(c *gin.Context) {
	storeBargainUserHelp := models.StoreBargainUserHelp{}
	id := c.Param("storeBargainUserHelpId")
	var err error
	storeBargainUserHelp.Id, err = strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	if err := ParseRequest(c, &storeBargainUserHelp); err != nil {
		return
	}
	err = storeBargainUserHelp.Patch()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, nil)
}

// @Summary List
// @Tags    StoreBargainUserHelp
// @Param query query string false "query, ?query=age:>:21,name:like:%jason%"
// @Param order query string false "order, ?order=age:desc,created_at:asc"
// @Param page query int false "page"
// @Param pageSize query int false "pageSize"
// @Success 200 {array} models.StoreBargainUserHelp "storeBargainUserHelp array"
// @Router /storeBargainUserHelps [get]
func (ctl *StoreBargainUserHelpController) List(c *gin.Context) {
	storeBargainUserHelp := &models.StoreBargainUserHelp{}
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
	storeBargainUserHelps, total, err := storeBargainUserHelp.List(rawQuery, rawOrder, offset, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"total": total,
		"data":  storeBargainUserHelps,
	})
}

// @Summary Get
// @Tags    StoreBargainUserHelp
// @Param  storeBargainUserHelpId path string true "storeBargainUserHelpId"
// @Success 200 {object} models.StoreBargainUserHelp "storeBargainUserHelp object"
// @Router /storeBargainUserHelps/{storeBargainUserHelpId} [get]
func (ctl *StoreBargainUserHelpController) Get(c *gin.Context) {
	storeBargainUserHelp := &models.StoreBargainUserHelp{}
	id := c.Param("storeBargainUserHelpId")

	var err error
	storeBargainUserHelp.Id, err = strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	storeBargainUserHelp, err = storeBargainUserHelp.Get()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, storeBargainUserHelp)
}
