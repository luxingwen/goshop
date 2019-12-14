package models

import (
	"github.com/jinzhu/gorm"
)

type Category struct {
	gorm.Model
	Title     string `gorm:"column:title" json:"title"`          // 分类名
	SortOrder int    `gorm:"column:sort_order" json:"sortOrder"` // 排序
	PicUrl    string `gorm:"column:pic_url" json:"picUrl"`       // 分类大封面
	Type      int    `gorm:"column:type" json:"type"`            //类型
	RelateId  int    `gorm:"column:relate_id" json:"relateId"`   // 关联id
	ParentId  int    `gorm:"column:parent_id" json:"parent_id"`  // 父级id
}

// 返回的category 结构体
type ResCategory struct {
	Id        int            `json:"id"`
	Title     string         `json:"title"`     // 分类名
	SortOrder int            `json:"sortOrder"` // 排序
	PicUrl    string         `json:"picUrl"`    // 分类大封面
	Type      int            `json:"type"`      // 类型
	RelateId  int            `json:"relateId"`  // 关联id
	Children  []*ResCategory `json:"children"`  // 孩子
}
