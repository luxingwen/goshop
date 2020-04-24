package controllers

import (
	"encoding/json"
	"errors"
	"fmt"

	"goshop/restful/models"

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
	if uid <= 0 {
		handleErr(c, errors.New("无效的uid"))
		return
	}

	storeCart := &models.StoreCart{}
	count, err := storeCart.GetUserCartNum(uid, "product")
	if err != nil {
		handleErr(c, err)
		return
	}
	handleOk(c, count)
}

func (ctl *AuthApiController) GetCartList(c *gin.Context) {
	uidT, ok := c.Get("uid")
	if !ok {
		handleErr(c, errors.New("无效的uid"))
		return
	}
	uid := uidT.(int)

	if uid <= 0 {
		handleErr(c, errors.New("无效的uid"))
		return
	}

	storeCart := &models.StoreCart{}
	r, err := storeCart.GetUserProductCartList(uid)
	if err != nil {
		handleErr(c, err)
		return
	}
	handleOk(c, r)
}

type ReqRemoveCart struct {
	Ids []int `json:"ids"`
}

func (ctl *AuthApiController) RemoveCart(c *gin.Context) {
	uidT, ok := c.Get("uid")
	if !ok {
		handleErr(c, errors.New("无效的uid"))
		return
	}
	uid := uidT.(int)

	if uid <= 0 {
		handleErr(c, errors.New("无效的uid"))
		return
	}
	req := new(ReqRemoveCart)

	err := c.ShouldBind(&req)
	if err != nil {
		return
	}

	storeCart := &models.StoreCart{}
	err = storeCart.RemoveUserCart(uid, req.Ids)
	if err != nil {
		handleErr(c, err)
		return
	}
	handleOk(c, "移除ok")
}

type UserCartNum struct {
	CartId  int `json:"cartId"`
	CartNum int `json:"cartNum"`
}

// 修改购物车库存
func (ctl *AuthApiController) ChangeUserCartNum(c *gin.Context) {
	uidT, ok := c.Get("uid")
	if !ok {
		handleErr(c, errors.New("无效的uid"))
		return
	}
	uid := uidT.(int)

	if uid <= 0 {
		handleErr(c, errors.New("无效的uid"))
		return
	}
	req := new(UserCartNum)
	err := c.ShouldBind(&req)
	if err != nil {
		handleErr(c, err)
		return
	}
	storeCart := &models.StoreCart{}
	err = storeCart.ChangeUserCartNum(req.CartId, req.CartNum, uid)
	if err != nil {
		return
	}
	handleOk(c, "移除OK")
}

type ReqSetCart struct {
	ProductId int    `json:"ProductId"`
	CartNum   int    `json:"cartNum"`
	UniqueId  string `json:"uniqueId"`
}

// 设置购物车
func (ctl *AuthApiController) SetCart(c *gin.Context) {
	uidT, ok := c.Get("uid")
	if !ok {
		handleErr(c, errors.New("无效的uid"))
		return
	}
	uid := uidT.(int)

	if uid <= 0 {
		handleErr(c, errors.New("无效的uid"))
		return
	}

	req := new(ReqSetCart)

	//mdata := make(map[string]interface{}, 0)
	err := c.ShouldBind(&req)
	if err != nil {
		handleErr(c, err)
		return
	}
	b, err := json.MarshalIndent(req, "", "\t")
	if err != nil {
		handleErr(c, err)
		return
	}
	fmt.Println(string(b))

	storeCart := &models.StoreCart{
		Uid:               uid,
		ProductId:         req.ProductId,
		CartNum:           req.CartNum,
		ProductAttrUnique: req.UniqueId,
		Type:              "product",
	}

	err = storeCart.SetCart()
	if err != nil {
		handleErr(c, err)
		return
	}

	mdata := make(map[string]interface{}, 0)
	mdata["cartId"] = storeCart.Id
	handleOk(c, mdata)
}

//
type ReqNowBuy struct {
	ProductId     int    `json:"ProductId"`
	CartNum       int    `json:"cartNum"`
	UniqueId      string `json:"uniqueId"`
	CombinationId int    `json:"combinationId"`
	SecKillId     int    `json:"SecKillId"`
	BargainId     int    `json:"bargainId"`
}

func (ctl *AuthApiController) NowBuy(c *gin.Context) {
	uidT, ok := c.Get("uid")
	if !ok {
		handleErr(c, errors.New("无效的uid"))
		return
	}
	uid := uidT.(int)

	if uid <= 0 {
		handleErr(c, errors.New("无效的uid"))
		return
	}

	req := new(ReqNowBuy)

	storeCart := &models.StoreCart{
		Uid:               uid,
		ProductId:         req.ProductId,
		CartNum:           req.CartNum,
		ProductAttrUnique: req.UniqueId,
		CombinationId:     req.CombinationId,
		SeckillId:         req.SecKillId,
		BargainId:         req.BargainId,
		Type:              "product",
	}

	err := storeCart.SetCart()
	if err != nil {
		handleErr(c, err)
		return
	}

	mdata := make(map[string]interface{}, 0)
	mdata["cartId"] = storeCart.Id
	handleOk(c, mdata)
}
