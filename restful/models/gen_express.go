//generate by gen
package models

import (
	"goshop/restful/common"
)

//快递公司表
type Express struct {
	Id     int    `gorm:"column:id"`      //快递公司id
	Code   string `gorm:"column:code"`    //快递公司简称
	Name   string `gorm:"column:name"`    //快递公司全称
	Sort   int    `gorm:"column:sort"`    //排序
	IsShow int    `gorm:"column:is_show"` //是否显示

}

//修改默认表名
func (Express) TableName() string {
	return "eb_express"
}

func (express *Express) Insert() error {
	err := common.GetDB().Create(express).Error
	return err
}

func (express *Express) Patch() error {
	err := common.GetDB().Model(express).Updates(express).Error
	return err
}

func (express *Express) Update() error {
	err := common.GetDB().Save(express).Error
	return err
}

func (express *Express) Delete() error {
	return common.GetDB().Delete(express).Error
}

func (express *Express) List(rawQuery string, rawOrder string, offset int, limit int) (*[]Express, int, error) {
	expresss := []Express{}
	total := 0
	db := common.GetDB().Model(express)
	db, err := buildWhere(rawQuery, db)
	if err != nil {
		return &expresss, total, err
	}
	db, err = buildOrder(rawOrder, db)
	if err != nil {
		return &expresss, total, err
	}
	db.Offset(offset).
		Limit(limit).
		Find(&expresss).
		Count(&total)
	err = db.Error
	return &expresss, total, err
}

func (express *Express) Get() (*Express, error) {
	err := common.GetDB().Find(&express).Error
	return express, err
}
