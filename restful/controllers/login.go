package controllers

import (
	"errors"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/medivhzhan/weapp/v2"

	"goshop/restful/config"

	"goshop/libs/cache"
	"goshop/libs/utils"

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

	sessionKey, ok := cache.Get(mdata["cache_key"])

	if !ok {
		handleErr(c, errors.New("没有找到sessionkey"))
		return
	}

	encryptedData := mdata["encryptedData"].(string)
	rawData := mdata["rawData"].(string)
	sign := mdata["signature"].(string)
	iv := mdata["iv"].(string)

	res, err := weapp.DecryptUserInfo(sessionKey.(string), rawData, encryptedData, sign, iv)
	if err != nil {
		// 处理一般错误信息
		fmt.Println("res err:", err)
		handleErr(c, err)
		return
	}

	b1, err := json.Marshal(res)

	if err != nil {
		handleErr(c, err)
		return
	}

	wechatUser := &models.WechatUser{}

	fmt.Println("b1:", string(b1))

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

	cacheKey := fmt.Sprintf("api_code_%s_%d", req.Code, time.Now().Unix())
	cacheKey = utils.Md5(cacheKey)
	cache.Put(cacheKey, res.SessionKey, 86400)

	mdata := make(map[string]interface{}, 0)
	mdata["cache_key"] = cacheKey

	b, err := json.Marshal(res)
	if err != nil {
		handleErr(c, err)
		return
	}

	fmt.Println("set_code:", string(b))
	handleOk(c, mdata)

}
