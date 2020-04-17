package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/medivhzhan/weapp/v2"

	"goshop/libs/cache"
	"goshop/libs/utils"
	"goshop/restful/common"
	"goshop/restful/config"
	"goshop/restful/models"
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
	code := mdata["code"].(float64)
	spid := mdata["spid"].(float64)

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

	wechatUser.RoutineOpenid = res.OpenID
	wechatUser.Nickname = res.Nickname
	wechatUser.Sex = res.Gender
	wechatUser.Province = res.Province
	wechatUser.Language = res.Language
	wechatUser.Country = res.Country
	wechatUser.City = res.City
	wechatUser.Headimgurl = res.Avatar
	wechatUser.Unionid = res.UnionID
	wechatUser.UserType = "routine"
	wechatUser.SessionKey = sessionKey.(string)

	rWechatUser, err := wechatUser.GetByRoutineOpenid(res.OpenID)
	if err != nil && err.Error() == "record not found" {
		wechatUser.AddTime = int(time.Now().Unix())

		err = wechatUser.Insert()
		if err != nil {
			fmt.Println("插入微信用户失败:", err)
			handleErr(c, err)
			return
		}

		user := &models.User{}

		user.Account = fmt.Sprintf("xw_%d%d", wechatUser.Uid, time.Now().Unix())

		user.Pwd = utils.Md5("lxw123456")
		user.Nickname = wechatUser.Nickname
		user.Avatar = wechatUser.Headimgurl

		user.Uid = wechatUser.Uid
		user.AddTime = wechatUser.AddTime
		user.LastTime = int(time.Now().Unix())
		user.Status = 1

		user.AddIp = c.ClientIP()
		user.LastIp = c.ClientIP()
		user.UserType = wechatUser.UserType

		err = user.Insert()
		if err != nil {
			fmt.Println("插入用户失败:", err)
			handleErr(c, err)
			return
		}

	} else {
		wechatUser.Uid = rWechatUser.Uid

		err = wechatUser.UpdateByUid(wechatUser.Uid)
		if err != nil {
			fmt.Println("更新微信用户失败:", wechatUser)
			handleErr(c, err)
			return
		}

		user := &models.User{}

		user.Uid = wechatUser.Uid
		user.Nickname = wechatUser.Nickname
		user.Avatar = wechatUser.Headimgurl
		user.LastIp = c.ClientIP()
		user.LastTime = int(time.Now().Unix())

		err = user.UpdateByUid(wechatUser.Uid)
		if err != nil {
			handleErr(c, err)
			return
		}
	}

	mdata = make(map[string]interface{}, 0)

	err = json.Unmarshal(b1, &mdata)
	if err != nil {
		handleErr(c, err)
		return
	}

	mdata["session_key"] = sessionKey.(string)

	mdata["uid"] = wechatUser.Uid

	token, err := common.GenerateToken(wechatUser.Uid, wechatUser.RoutineOpenid)
	if err != nil {
		handleErr(c, err)
		return
	}

	user := &models.User{}
	ruser, err := user.Get()
	if err != nil {
		handleErr(c, err)
		return
	}

	mdata["status"] = ruser.Status
	mdata["code"] = code

	mdata["token"] = token
	mdata["spid"] = spid

	fmt.Println("b1:", string(b1))

	//@todo 获取是否有扫码进小程序

	fmt.Println(string(b))
	handleOk(c, mdata)

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
