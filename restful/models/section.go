package models

import (
	"github.com/jinzhu/gorm"
)

// 板块
type Section struct {
	gorm.Model
	Title     string `gorm:"column:title" json:"title"`           // 标题
	Type      int    `gorm:"column:type" json:"type"`             // 板块类型   0 轮播  1 频道图标  2 横向商品板块  3 板块类型 两张板块  4 大图商品带边框板块  5 大图商品无边框板块
	SortOrder int    `gorm:"column:sort_order" json:"sort_order"` // 板块排序
	RelateId  int    `gorm:"column:relate_id" json:"relate_id"`   // 关联id
}

// 板块物品
type SectionItem struct {
	gorm.Model
	SectionId uint    `gorm:"column:section_id" json:"sectionId"`
	Title     string  `gorm:"column:title" json:"title"`         // 标题
	Type      int     `gorm:"column:type" json:"type"`           // 类型
	RelateId  int     `gorm:"column:relate_id" json:"relate_id"` // 关联id
	Price     float64 `gorm:"column:price" json:"price"`         // 价格
	PicUrl    string  `gorm:"pic_url" json:"picUrl"`             // 图片地址
	Tag       string  `gorm:"tag" json:"tag"`                    // 标签
}

// 返回板块信息
type ResSection struct {
	Section
	List []*SectionItem `json:"list"`
}
