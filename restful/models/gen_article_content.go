//generate by gen
package models

import (
	"goshop/restful/common"
)

//文章内容表
type ArticleContent struct {
	Nid     int    `gorm:"column:nid"`     //文章id
	Content string `gorm:"column:content"` //文章内容

}

//修改默认表名
func (ArticleContent) TableName() string {
	return "eb_article_content"
}

func (articleContent *ArticleContent) Insert() error {
	err := common.GetDB().Create(articleContent).Error
	return err
}

func (articleContent *ArticleContent) Patch() error {
	err := common.GetDB().Model(articleContent).Updates(articleContent).Error
	return err
}

func (articleContent *ArticleContent) Update() error {
	err := common.GetDB().Save(articleContent).Error
	return err
}

func (articleContent *ArticleContent) Delete() error {
	return common.GetDB().Delete(articleContent).Error
}

func (articleContent *ArticleContent) List(rawQuery string, rawOrder string, offset int, limit int) (*[]ArticleContent, int, error) {
	articleContents := []ArticleContent{}
	total := 0
	db := common.GetDB().Model(articleContent)
	db, err := buildWhere(rawQuery, db)
	if err != nil {
		return &articleContents, total, err
	}
	db, err = buildOrder(rawOrder, db)
	if err != nil {
		return &articleContents, total, err
	}
	db.Offset(offset).
		Limit(limit).
		Find(&articleContents).
		Count(&total)
	err = db.Error
	return &articleContents, total, err
}

func (articleContent *ArticleContent) Get() (*ArticleContent, error) {
	err := common.GetDB().Find(&articleContent).Error
	return articleContent, err
}
