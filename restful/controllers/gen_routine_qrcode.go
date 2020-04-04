//generate by gen
package controllers

import (
	"github.com/gin-gonic/gin"
	"goshop/restful/models"
	"net/http"
	"strconv"
)

type RoutineQrcodeController struct {
}

// @Summary Create
// @Tags    RoutineQrcode
// @Param body body models.RoutineQrcode true "RoutineQrcode"
// @Success 200 {string} string ""
// @Router /routineQrcodes [post]
func (ctl *RoutineQrcodeController) Create(c *gin.Context) {
	routineQrcode := models.RoutineQrcode{}
	if err := ParseRequest(c, &routineQrcode); err != nil {
		return
	}
	if err := routineQrcode.Insert(); err != nil {
		c.JSON(http.StatusBadGateway, err)
		return
	}
	c.JSON(http.StatusOK, routineQrcode)
}

// @Summary  Delete
// @Tags     RoutineQrcode
// @Param  routineQrcodeId  path string true "routineQrcodeId"
// @Success 200 {string} string ""
// @Router /routineQrcodes/{routineQrcodeId} [delete]
func (ctl *RoutineQrcodeController) Delete(c *gin.Context) {
	routineQrcode := models.RoutineQrcode{}
	id := c.Param("routineQrcodeId")
	var err error
	routineQrcode.Id, err = strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	err = routineQrcode.Delete()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, nil)
}

// @Summary Put
// @Tags    RoutineQrcode
// @Param body body models.RoutineQrcode true "routineQrcode"
// @Param  routineQrcodeId path string true "routineQrcodeId"
// @Success 200 {string} string ""
// @Router /routineQrcodes/{routineQrcodeId} [put]
func (ctl *RoutineQrcodeController) Put(c *gin.Context) {
	routineQrcode := models.RoutineQrcode{}
	id := c.Param("routineQrcodeId")
	var err error
	routineQrcode.Id, err = strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	if err := ParseRequest(c, &routineQrcode); err != nil {
		return
	}
	err = routineQrcode.Update()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, nil)
}

// @Summary Patch
// @Tags    RoutineQrcode
// @Param body body models.RoutineQrcode true "routineQrcode"
// @Param  routineQrcodeId path string true "routineQrcodeId"
// @Success 200 {string} string ""
// @Router /routineQrcodes/{routineQrcodeId} [patch]
func (ctl *RoutineQrcodeController) Patch(c *gin.Context) {
	routineQrcode := models.RoutineQrcode{}
	id := c.Param("routineQrcodeId")
	var err error
	routineQrcode.Id, err = strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	if err := ParseRequest(c, &routineQrcode); err != nil {
		return
	}
	err = routineQrcode.Patch()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, nil)
}

// @Summary List
// @Tags    RoutineQrcode
// @Param query query string false "query, ?query=age:>:21,name:like:%jason%"
// @Param order query string false "order, ?order=age:desc,created_at:asc"
// @Param page query int false "page"
// @Param pageSize query int false "pageSize"
// @Success 200 {array} models.RoutineQrcode "routineQrcode array"
// @Router /routineQrcodes [get]
func (ctl *RoutineQrcodeController) List(c *gin.Context) {
	routineQrcode := &models.RoutineQrcode{}
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
	routineQrcodes, total, err := routineQrcode.List(rawQuery, rawOrder, offset, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"total": total,
		"data":  routineQrcodes,
	})
}

// @Summary Get
// @Tags    RoutineQrcode
// @Param  routineQrcodeId path string true "routineQrcodeId"
// @Success 200 {object} models.RoutineQrcode "routineQrcode object"
// @Router /routineQrcodes/{routineQrcodeId} [get]
func (ctl *RoutineQrcodeController) Get(c *gin.Context) {
	routineQrcode := &models.RoutineQrcode{}
	id := c.Param("routineQrcodeId")

	var err error
	routineQrcode.Id, err = strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	routineQrcode, err = routineQrcode.Get()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, routineQrcode)
}
