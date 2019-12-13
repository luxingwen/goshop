package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"goshop/libs/errcode"
)

func handleOk(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "success",
		"data":    data,
	})
}

func handleErr(c *gin.Context, err error) {
	c.JSON(http.StatusOK, gin.H{
		"code":    errcode.ErrCode(err),
		"message": err.Error(),
	})
}
