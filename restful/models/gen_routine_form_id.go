//generate by gen
package models

import (
	"goshop/restful/common"
)

//表单id表记录表
type RoutineFormId struct {
	Id       int    `gorm:"column:id"`        //表单ID表ID
	Uid      int    `gorm:"column:uid"`       //用户uid
	FormId   string `gorm:"column:form_id"`   //表单ID
	StopTime int    `gorm:"column:stop_time"` //表单ID失效时间
	Status   int    `gorm:"column:status"`    //状态1 未使用 2不能使用

}

//修改默认表名
func (RoutineFormId) TableName() string {
	return "eb_routine_form_id"
}

func (routineFormId *RoutineFormId) Insert() error {
	err := common.GetDB().Create(routineFormId).Error
	return err
}

func (routineFormId *RoutineFormId) Patch() error {
	err := common.GetDB().Model(routineFormId).Updates(routineFormId).Error
	return err
}

func (routineFormId *RoutineFormId) Update() error {
	err := common.GetDB().Save(routineFormId).Error
	return err
}

func (routineFormId *RoutineFormId) Delete() error {
	return common.GetDB().Delete(routineFormId).Error
}

func (routineFormId *RoutineFormId) List(rawQuery string, rawOrder string, offset int, limit int) (*[]RoutineFormId, int, error) {
	routineFormIds := []RoutineFormId{}
	total := 0
	db := common.GetDB().Model(routineFormId)
	db, err := buildWhere(rawQuery, db)
	if err != nil {
		return &routineFormIds, total, err
	}
	db, err = buildOrder(rawOrder, db)
	if err != nil {
		return &routineFormIds, total, err
	}
	db.Offset(offset).
		Limit(limit).
		Find(&routineFormIds).
		Count(&total)
	err = db.Error
	return &routineFormIds, total, err
}

func (routineFormId *RoutineFormId) Get() (*RoutineFormId, error) {
	err := common.GetDB().Find(&routineFormId).Error
	return routineFormId, err
}
