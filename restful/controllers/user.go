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

type ResUserList struct {
	Count    int            `json:"count"`
	UserList []*models.User `json:"user_list"`
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
	err = user.IsUserExist()
	if err != nil {
		handleErr(c, err)
		return
	}

	token, err := common.GenerateToken(user.Username, user.Password)
	if err != nil {
		handleErr(c, err)
		return
	}

	m := make(map[string]interface{}, 0)
	m["token"] = token
	handleOk(c, m)
}

// @Summary 注册
// @Accept json
// @Produce  json
// @Param param body controllers.ReqLogin true "{}"
// @Success 200 {string} string "{"code":0,"data":{},"msg":"ok"}"
// @Router /api/user/register [post]
func Register(c *gin.Context) {
	user := new(models.User)

	err := c.ShouldBindJSON(&user)
	if err != nil {
		handleErr(c, err)
		return
	}

	err = user.RegisterUser()
	if err != nil {
		handleErr(c, err)
		return
	}
	handleOk(c, "")
}

// @Summary 查询用户列表
// @Accept json
// @Produce  json
// @Param param body models.ReqUser true "{}"
// @Success 200 {string} string "{"code":0,"data":{},"msg":"ok"}"
// @Router /api/user/userlist [get]
func UserList(c *gin.Context) {
	req := new(models.ReqUser)

	err := c.ShouldBindJSON(&req)
	if err != nil {
		handleErr(c, err)
		return
	}

	userList, count, err := req.UserList()
	if err != nil {
		handleErr(c, err)
		return
	}

	res := &ResUserList{
		Count:    count,
		UserList: userList,
	}

	handleOk(c, res)
}
