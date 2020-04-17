package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/medivhzhan/weapp/v2"

	"goshop/restful/config"

	"encoding/json"
)

type LoginController struct {
}

type Wxlogin struct {
	Code      string `json:"code"`
	NickName  string `json:"nickName"`
	AvatarUrl string `json:"avatarUrl"`
	Gender    int    `json:"gender"`
}

func (ctl *LoginController) Login(c *gin.Context) {
	mdata := make(map[string]interface{}, 0)
	err := c.ShouldBind(&mdata)
	if err != nil {
		handleErr(c, err)
		return
	}

	b, err := json.Marshal(mdata)
	if err != nil {
		handleErr(c, err)
		return
	}

	// res, err := weapp.DecryptUserInfo("session-key", "raw-data", "encrypted-date", "signature", "iv")
	// if err != nil {
	// 	// 处理一般错误信息
	// 	return
	// }

	fmt.Println(string(b))
	handleOk(c, "ok")

}

func (ctl *LoginController) SetCode(c *gin.Context) {
	req := new(Wxlogin)

	err := c.ShouldBindJSON(&req)
	if err != nil {
		handleErr(c, err)
		return
	}

	fmt.Printf("req: %#v\n", req)
	fmt.Println("code:", req.Code)

	fmt.Println(config.WxConf.AppId)
	fmt.Println(config.WxConf.AppSecret)

	res, err := weapp.Login(config.WxConf.AppId, config.WxConf.AppSecret, req.Code)
	if err != nil {
		// 处理一般错误信息
		fmt.Println("login err:", err)
		handleErr(c, err)
		return
	}

	if err := res.GetResponseError(); err != nil {
		// 处理微信返回错误信息
		fmt.Println("GetResponseError:", err)
		handleErr(c, err)
		return
	}

	b, err := json.Marshal(res)
	if err != nil {
		handleErr(c, err)
		return
	}

	fmt.Println("set_code:", string(b))
	handleOk(c, "ok")

}
