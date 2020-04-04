//generate by gen
package controllers

import (
	"github.com/gin-gonic/gin"
	"goshop/restful/models"
	"net/http"
	"strconv"
)

type UserController struct {
}

// @Summary Create
// @Tags    User
// @Param body body models.User true "User"
// @Success 200 {string} string ""
// @Router /users [post]
func (ctl *UserController) Create(c *gin.Context) {
	user := models.User{}
	if err := ParseRequest(c, &user); err != nil {
		return
	}
	if err := user.Insert(); err != nil {
		c.JSON(http.StatusBadGateway, err)
		return
	}
	c.JSON(http.StatusOK, user)
}

// @Summary  Delete
// @Tags     User
// @Param  userId  path string true "userId"
// @Success 200 {string} string ""
// @Router /users/{userId} [delete]
func (ctl *UserController) Delete(c *gin.Context) {
	user := models.User{}
	id := c.Param("userId")
	var err error
	user.Uid, err = strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	err = user.Delete()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, nil)
}

// @Summary Put
// @Tags    User
// @Param body body models.User true "user"
// @Param  userId path string true "userId"
// @Success 200 {string} string ""
// @Router /users/{userId} [put]
func (ctl *UserController) Put(c *gin.Context) {
	user := models.User{}
	id := c.Param("userId")
	var err error
	user.Uid, err = strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	if err := ParseRequest(c, &user); err != nil {
		return
	}
	err = user.Update()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, nil)
}

// @Summary Patch
// @Tags    User
// @Param body body models.User true "user"
// @Param  userId path string true "userId"
// @Success 200 {string} string ""
// @Router /users/{userId} [patch]
func (ctl *UserController) Patch(c *gin.Context) {
	user := models.User{}
	id := c.Param("userId")
	var err error
	user.Uid, err = strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	if err := ParseRequest(c, &user); err != nil {
		return
	}
	err = user.Patch()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, nil)
}

// @Summary List
// @Tags    User
// @Param query query string false "query, ?query=age:>:21,name:like:%jason%"
// @Param order query string false "order, ?order=age:desc,created_at:asc"
// @Param page query int false "page"
// @Param pageSize query int false "pageSize"
// @Success 200 {array} models.User "user array"
// @Router /users [get]
func (ctl *UserController) List(c *gin.Context) {
	user := &models.User{}
	id := c.Param("userId")
	var err error
	user.Uid, err = strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

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
	users, total, err := user.List(rawQuery, rawOrder, offset, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"total": total,
		"data":  users,
	})
}

// @Summary Get
// @Tags    User
// @Param  userId path string true "userId"
// @Success 200 {object} models.User "user object"
// @Router /users/{userId} [get]
func (ctl *UserController) Get(c *gin.Context) {
	user := &models.User{}
	id := c.Param("userId")

	var err error
	user.Uid, err = strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	user, err = user.Get()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, user)
}
