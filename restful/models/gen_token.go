//generate by gen
package models

import (
	"goshop/restful/common"
)

//保存token随机字符串
type Token struct {
	Id         int    `gorm:"column:id"`          //
	Uid        int    `gorm:"column:uid"`         //用户uid
	RandString string `gorm:"column:rand_string"` //10位随机字符串
	AddTime    int    `gorm:"column:add_time"`    //添加时间

}

//修改默认表名
func (Token) TableName() string {
	return "eb_token"
}

func (token *Token) Insert() error {
	err := common.GetDB().Create(token).Error
	return err
}

func (token *Token) Patch() error {
	err := common.GetDB().Model(token).Updates(token).Error
	return err
}

func (token *Token) Update() error {
	err := common.GetDB().Save(token).Error
	return err
}

func (token *Token) Delete() error {
	return common.GetDB().Delete(token).Error
}

func (token *Token) List(rawQuery string, rawOrder string, offset int, limit int) (*[]Token, int, error) {
	tokens := []Token{}
	total := 0
	db := common.GetDB().Model(token)
	db, err := buildWhere(rawQuery, db)
	if err != nil {
		return &tokens, total, err
	}
	db, err = buildOrder(rawOrder, db)
	if err != nil {
		return &tokens, total, err
	}
	db.Offset(offset).
		Limit(limit).
		Find(&tokens).
		Count(&total)
	err = db.Error
	return &tokens, total, err
}

func (token *Token) Get() (*Token, error) {
	err := common.GetDB().Find(&token).Error
	return token, err
}
