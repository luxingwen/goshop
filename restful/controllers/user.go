package controllers

import (
	"github.com/gin-gonic/gin"
	"goshop/restful/common"
	"goshop/restful/models"
	"goshop/libs/errcode"
	"strings"
)

type ReqLogin struct {
	Username string `json:"username" faker:"username"`
	Password string `json:"password" faker:"password"`
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
		if strings.Contains(err.Error(), "Duplicate entry") {
			handleErr(c, errcode.NewErrcode(1001, "用户名已存在"))
			return
		}
		handleErr(c, err)
		return
	}
	user.Password = ""
	handleOk(c, user)
}

// @Summary 查询用户列表
// @Produce  json
// @Param page query int true "page"
// @Param limit query int true "limit"
// @Param username query string false "username"
// @Success 200 {string} string "{"code":0,"data":{},"msg":"ok"}"
// @Router /api/user/userlist [get]
func UserList(c *gin.Context) {
	req := new(models.ReqUser)

	err := c.ShouldBindQuery(&req)
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
