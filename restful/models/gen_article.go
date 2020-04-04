//generate by gen
package models

import (
	"goshop/restful/common"
)

//文章管理表
type Article struct {
	Id            int    `gorm:"column:id"`             //文章管理ID
	Cid           string `gorm:"column:cid"`            //分类id
	Title         string `gorm:"column:title"`          //文章标题
	Author        string `gorm:"column:author"`         //文章作者
	ImageInput    string `gorm:"column:image_input"`    //文章图片
	Synopsis      string `gorm:"column:synopsis"`       //文章简介
	ShareTitle    string `gorm:"column:share_title"`    //文章分享标题
	ShareSynopsis string `gorm:"column:share_synopsis"` //文章分享简介
	Visit         string `gorm:"column:visit"`          //浏览次数
	Sort          int    `gorm:"column:sort"`           //排序
	Url           string `gorm:"column:url"`            //原文链接
	Status        int    `gorm:"column:status"`         //状态
	AddTime       string `gorm:"column:add_time"`       //添加时间
	Hide          int    `gorm:"column:hide"`           //是否隐藏
	AdminId       int    `gorm:"column:admin_id"`       //管理员id
	MerId         int    `gorm:"column:mer_id"`         //商户id
	IsHot         int    `gorm:"column:is_hot"`         //是否热门(小程序)
	IsBanner      int    `gorm:"column:is_banner"`      //是否轮播图(小程序)

}

//修改默认表名
func (Article) TableName() string {
	return "eb_article"
}

func (article *Article) Insert() error {
	err := common.GetDB().Create(article).Error
	return err
}

func (article *Article) Patch() error {
	err := common.GetDB().Model(article).Updates(article).Error
	return err
}

func (article *Article) Update() error {
	err := common.GetDB().Save(article).Error
	return err
}

func (article *Article) Delete() error {
	return common.GetDB().Delete(article).Error
}

func (article *Article) List(rawQuery string, rawOrder string, offset int, limit int) (*[]Article, int, error) {
	articles := []Article{}
	total := 0
	db := common.GetDB().Model(article)
	db, err := buildWhere(rawQuery, db)
	if err != nil {
		return &articles, total, err
	}
	db, err = buildOrder(rawOrder, db)
	if err != nil {
		return &articles, total, err
	}
	db.Offset(offset).
		Limit(limit).
		Find(&articles).
		Count(&total)
	err = db.Error
	return &articles, total, err
}

func (article *Article) Get() (*Article, error) {
	err := common.GetDB().Find(&article).Error
	return article, err
}
