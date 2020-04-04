//generate by gen
package models

import (
	"goshop/restful/common"
)

//微信用户表
type WechatUser struct {
	Uid           int     `gorm:"column:uid"`            //微信用户id
	Unionid       string  `gorm:"column:unionid"`        //只有在用户将公众号绑定到微信开放平台帐号后，才会出现该字段
	Openid        string  `gorm:"column:openid"`         //用户的标识，对当前公众号唯一
	RoutineOpenid string  `gorm:"column:routine_openid"` //小程序唯一身份ID
	Nickname      string  `gorm:"column:nickname"`       //用户的昵称
	Headimgurl    string  `gorm:"column:headimgurl"`     //用户头像
	Sex           int     `gorm:"column:sex"`            //用户的性别，值为1时是男性，值为2时是女性，值为0时是未知
	City          string  `gorm:"column:city"`           //用户所在城市
	Language      string  `gorm:"column:language"`       //用户的语言，简体中文为zh_CN
	Province      string  `gorm:"column:province"`       //用户所在省份
	Country       string  `gorm:"column:country"`        //用户所在国家
	Remark        string  `gorm:"column:remark"`         //公众号运营者对粉丝的备注，公众号运营者可在微信公众平台用户管理界面对粉丝添加备注
	Groupid       int     `gorm:"column:groupid"`        //用户所在的分组ID（兼容旧的用户分组接口）
	TagidList     string  `gorm:"column:tagid_list"`     //用户被打上的标签ID列表
	Subscribe     int     `gorm:"column:subscribe"`      //用户是否订阅该公众号标识
	SubscribeTime int     `gorm:"column:subscribe_time"` //关注公众号时间
	AddTime       int     `gorm:"column:add_time"`       //添加时间
	Stair         int     `gorm:"column:stair"`          //一级推荐人
	Second        int     `gorm:"column:second"`         //二级推荐人
	OrderStair    int     `gorm:"column:order_stair"`    //一级推荐人订单
	OrderSecond   int     `gorm:"column:order_second"`   //二级推荐人订单
	NowMoney      float64 `gorm:"column:now_money"`      //佣金
	SessionKey    string  `gorm:"column:session_key"`    //小程序用户会话密匙
	UserType      string  `gorm:"column:user_type"`      //用户类型

}

//修改默认表名
func (WechatUser) TableName() string {
	return "eb_wechat_user"
}

func (wechatUser *WechatUser) Insert() error {
	err := common.GetDB().Create(wechatUser).Error
	return err
}

func (wechatUser *WechatUser) Patch() error {
	err := common.GetDB().Model(wechatUser).Updates(wechatUser).Error
	return err
}

func (wechatUser *WechatUser) Update() error {
	err := common.GetDB().Save(wechatUser).Error
	return err
}

func (wechatUser *WechatUser) Delete() error {
	return common.GetDB().Delete(wechatUser).Error
}

func (wechatUser *WechatUser) List(rawQuery string, rawOrder string, offset int, limit int) (*[]WechatUser, int, error) {
	wechatUsers := []WechatUser{}
	total := 0
	db := common.GetDB().Model(wechatUser)
	db, err := buildWhere(rawQuery, db)
	if err != nil {
		return &wechatUsers, total, err
	}
	db, err = buildOrder(rawOrder, db)
	if err != nil {
		return &wechatUsers, total, err
	}
	db.Offset(offset).
		Limit(limit).
		Find(&wechatUsers).
		Count(&total)
	err = db.Error
	return &wechatUsers, total, err
}

func (wechatUser *WechatUser) Get() (*WechatUser, error) {
	err := common.GetDB().Find(&wechatUser).Error
	return wechatUser, err
}
