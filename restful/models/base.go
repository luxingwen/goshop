package models

type Query struct {
	Page    int `form:"page" json:"page"`
	PageNum int `form:"limit" json:"pageNum"`
}
