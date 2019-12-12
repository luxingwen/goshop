package controllers

import (
	"github.com/gin-gonic/gin"

	"goshop/restful/common"
	"goshop/restful/models"
)

type ReqLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// @Summary 登录
// @Accept json
// @Produce  json
// @Param param body controllers.ReqLogin true "{}"
// @Success 200 {string} string "{"code":0,"data":{},"msg":"ok"}"
// @Router /api/user/login [post]
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
