package controllers

import (
	"github.com/gin-gonic/gin"

	"goshop/restful/common"
	"goshop/restful/models"
)

func Login(c *gin.Context) {
	user := new(models.User)

	err := c.ShouldBindJSON(&user)
	if err != nil {
		handleErr(c, err)
		return
	}

	// @todo login

	token, err := common.GenerateToken(user.Username, user.Password)
	if err != nil {
		handleErr(c, err)
		return
	}

	m := make(map[string]interface{}, 0)
	m["token"] = token
	handleOk(c, m)
}
