package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"

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

func ParseRequest(c *gin.Context, request interface{}) error {
	err := c.ShouldBindWith(request, binding.JSON)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "parse Request Error",
			"error":   err.Error(),
		})

		log.Println("ParseRequest Result", request)
		log.Println("ParseRequest Error", err.Error())
		return err
	}
	return nil
}

func SuccessResponse(c *gin.Context, response interface{}) {
	handleOk(c, response)
}

func CheckErr(c *gin.Context, err error) {
	if err != nil {
		c.JSON(500, err.Error())
		return
	}
}
