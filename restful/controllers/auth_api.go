package controllers

import (
	//"goshop/restful/models"
	"errors"

	"github.com/gin-gonic/gin"
)

type AuthApiController struct {
}

func (ctl *AuthApiController) GetCartNum(c *gin.Context) {
	uidT, ok := c.Get("uid")
	if !ok {
		handleErr(c, errors.New("无效的uid"))
		return
	}
	uid := uidT.(int)
	_ = uid
}
