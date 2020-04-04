//generate by gen
package controllers

import (
	"github.com/gin-gonic/gin"
	"goshop/restful/models"
	"net/http"
	"strconv"
)

type StoreSeckillController struct {
}

// @Summary Create
// @Tags    StoreSeckill
// @Param body body models.StoreSeckill true "StoreSeckill"
// @Success 200 {string} string ""
// @Router /storeSeckills [post]
func (ctl *StoreSeckillController) Create(c *gin.Context) {
	storeSeckill := models.StoreSeckill{}
	if err := ParseRequest(c, &storeSeckill); err != nil {
		return
	}
	if err := storeSeckill.Insert(); err != nil {
		c.JSON(http.StatusBadGateway, err)
		return
	}
	c.JSON(http.StatusOK, storeSeckill)
}

// @Summary  Delete
// @Tags     StoreSeckill
// @Param  storeSeckillId  path string true "storeSeckillId"
// @Success 200 {string} string ""
// @Router /storeSeckills/{storeSeckillId} [delete]
func (ctl *StoreSeckillController) Delete(c *gin.Context) {
	storeSeckill := models.StoreSeckill{}
	id := c.Param("storeSeckillId")
	var err error
	storeSeckill.Id, err = strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	err = storeSeckill.Delete()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, nil)
}

// @Summary Put
// @Tags    StoreSeckill
// @Param body body models.StoreSeckill true "storeSeckill"
// @Param  storeSeckillId path string true "storeSeckillId"
// @Success 200 {string} string ""
// @Router /storeSeckills/{storeSeckillId} [put]
func (ctl *StoreSeckillController) Put(c *gin.Context) {
	storeSeckill := models.StoreSeckill{}
	id := c.Param("storeSeckillId")
	var err error
	storeSeckill.Id, err = strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	if err := ParseRequest(c, &storeSeckill); err != nil {
		return
	}
	err = storeSeckill.Update()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, nil)
}

// @Summary Patch
// @Tags    StoreSeckill
// @Param body body models.StoreSeckill true "storeSeckill"
// @Param  storeSeckillId path string true "storeSeckillId"
// @Success 200 {string} string ""
// @Router /storeSeckills/{storeSeckillId} [patch]
func (ctl *StoreSeckillController) Patch(c *gin.Context) {
	storeSeckill := models.StoreSeckill{}
	id := c.Param("storeSeckillId")
	var err error
	storeSeckill.Id, err = strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	if err := ParseRequest(c, &storeSeckill); err != nil {
		return
	}
	err = storeSeckill.Patch()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, nil)
}

// @Summary List
// @Tags    StoreSeckill
// @Param query query string false "query, ?query=age:>:21,name:like:%jason%"
// @Param order query string false "order, ?order=age:desc,created_at:asc"
// @Param page query int false "page"
// @Param pageSize query int false "pageSize"
// @Success 200 {array} models.StoreSeckill "storeSeckill array"
// @Router /storeSeckills [get]
func (ctl *StoreSeckillController) List(c *gin.Context) {
	storeSeckill := &models.StoreSeckill{}
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
	storeSeckills, total, err := storeSeckill.List(rawQuery, rawOrder, offset, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"total": total,
		"data":  storeSeckills,
	})
}

// @Summary Get
// @Tags    StoreSeckill
// @Param  storeSeckillId path string true "storeSeckillId"
// @Success 200 {object} models.StoreSeckill "storeSeckill object"
// @Router /storeSeckills/{storeSeckillId} [get]
func (ctl *StoreSeckillController) Get(c *gin.Context) {
	storeSeckill := &models.StoreSeckill{}
	id := c.Param("storeSeckillId")

	var err error
	storeSeckill.Id, err = strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	storeSeckill, err = storeSeckill.Get()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, storeSeckill)
}
