//generate by gen
package models

import (
	"goshop/restful/common"
)

//文章分类表
type ArticleCategory struct {
	Id      int    `gorm:"column:id"`       //文章分类id
	Pid     int    `gorm:"column:pid"`      //父级ID
	Title   string `gorm:"column:title"`    //文章分类标题
	Intr    string `gorm:"column:intr"`     //文章分类简介
	Image   string `gorm:"column:image"`    //文章分类图片
	Status  int    `gorm:"column:status"`   //状态
	Sort    int    `gorm:"column:sort"`     //排序
	IsDel   int    `gorm:"column:is_del"`   //1删除0未删除
	AddTime string `gorm:"column:add_time"` //添加时间
	Hidden  int    `gorm:"column:hidden"`   //是否隐藏

}

//修改默认表名
func (ArticleCategory) TableName() string {
	return "eb_article_category"
}

func (articleCategory *ArticleCategory) Insert() error {
	err := common.GetDB().Create(articleCategory).Error
	return err
}

func (articleCategory *ArticleCategory) Patch() error {
	err := common.GetDB().Model(articleCategory).Updates(articleCategory).Error
	return err
}

func (articleCategory *ArticleCategory) Update() error {
	err := common.GetDB().Save(articleCategory).Error
	return err
}

func (articleCategory *ArticleCategory) Delete() error {
	return common.GetDB().Delete(articleCategory).Error
}

func (articleCategory *ArticleCategory) List(rawQuery string, rawOrder string, offset int, limit int) (*[]ArticleCategory, int, error) {
	articleCategorys := []ArticleCategory{}
	total := 0
	db := common.GetDB().Model(articleCategory)
	db, err := buildWhere(rawQuery, db)
	if err != nil {
		return &articleCategorys, total, err
	}
	db, err = buildOrder(rawOrder, db)
	if err != nil {
		return &articleCategorys, total, err
	}
	db.Offset(offset).
		Limit(limit).
		Find(&articleCategorys).
		Count(&total)
	err = db.Error
	return &articleCategorys, total, err
}

func (articleCategory *ArticleCategory) Get() (*ArticleCategory, error) {
	err := common.GetDB().Find(&articleCategory).Error
	return articleCategory, err
}

func (articleCategory *ArticleCategory) GetArticleCategory() (r []*ArticleCategory, err error) {
	db := common.GetDB()

	db = db.Select("id, title").Where("hidden = ? AND is_del = ? AND status = ? AND pid = ?", 0, 0, 1, 0)
	err = db.Find(&r).Error
	return
}
