//generate by gen
package models

import (
	"goshop/restful/common"
)

//微信模板
type RoutineTemplate struct {
	Id      int    `gorm:"column:id"`       //模板id
	Tempkey string `gorm:"column:tempkey"`  //模板编号
	Name    string `gorm:"column:name"`     //模板名
	Content string `gorm:"column:content"`  //回复内容
	Tempid  string `gorm:"column:tempid"`   //模板ID
	AddTime string `gorm:"column:add_time"` //添加时间
	Status  int    `gorm:"column:status"`   //状态

}

//修改默认表名
func (RoutineTemplate) TableName() string {
	return "eb_routine_template"
}

func (routineTemplate *RoutineTemplate) Insert() error {
	err := common.GetDB().Create(routineTemplate).Error
	return err
}

func (routineTemplate *RoutineTemplate) Patch() error {
	err := common.GetDB().Model(routineTemplate).Updates(routineTemplate).Error
	return err
}

func (routineTemplate *RoutineTemplate) Update() error {
	err := common.GetDB().Save(routineTemplate).Error
	return err
}

func (routineTemplate *RoutineTemplate) Delete() error {
	return common.GetDB().Delete(routineTemplate).Error
}

func (routineTemplate *RoutineTemplate) List(rawQuery string, rawOrder string, offset int, limit int) (*[]RoutineTemplate, int, error) {
	routineTemplates := []RoutineTemplate{}
	total := 0
	db := common.GetDB().Model(routineTemplate)
	db, err := buildWhere(rawQuery, db)
	if err != nil {
		return &routineTemplates, total, err
	}
	db, err = buildOrder(rawOrder, db)
	if err != nil {
		return &routineTemplates, total, err
	}
	db.Offset(offset).
		Limit(limit).
		Find(&routineTemplates).
		Count(&total)
	err = db.Error
	return &routineTemplates, total, err
}

func (routineTemplate *RoutineTemplate) Get() (*RoutineTemplate, error) {
	err := common.GetDB().Find(&routineTemplate).Error
	return routineTemplate, err
}
