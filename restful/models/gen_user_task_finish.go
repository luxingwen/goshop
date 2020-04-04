//generate by gen
package models

import (
	"goshop/restful/common"
)

//用户任务完成记录表
type UserTaskFinish struct {
	Id      int `gorm:"column:id"`       //
	TaskId  int `gorm:"column:task_id"`  //任务id
	Uid     int `gorm:"column:uid"`      //用户id
	Status  int `gorm:"column:status"`   //是否有效
	AddTime int `gorm:"column:add_time"` //添加时间

}

//修改默认表名
func (UserTaskFinish) TableName() string {
	return "eb_user_task_finish"
}

func (userTaskFinish *UserTaskFinish) Insert() error {
	err := common.GetDB().Create(userTaskFinish).Error
	return err
}

func (userTaskFinish *UserTaskFinish) Patch() error {
	err := common.GetDB().Model(userTaskFinish).Updates(userTaskFinish).Error
	return err
}

func (userTaskFinish *UserTaskFinish) Update() error {
	err := common.GetDB().Save(userTaskFinish).Error
	return err
}

func (userTaskFinish *UserTaskFinish) Delete() error {
	return common.GetDB().Delete(userTaskFinish).Error
}

func (userTaskFinish *UserTaskFinish) List(rawQuery string, rawOrder string, offset int, limit int) (*[]UserTaskFinish, int, error) {
	userTaskFinishs := []UserTaskFinish{}
	total := 0
	db := common.GetDB().Model(userTaskFinish)
	db, err := buildWhere(rawQuery, db)
	if err != nil {
		return &userTaskFinishs, total, err
	}
	db, err = buildOrder(rawOrder, db)
	if err != nil {
		return &userTaskFinishs, total, err
	}
	db.Offset(offset).
		Limit(limit).
		Find(&userTaskFinishs).
		Count(&total)
	err = db.Error
	return &userTaskFinishs, total, err
}

func (userTaskFinish *UserTaskFinish) Get() (*UserTaskFinish, error) {
	err := common.GetDB().Find(&userTaskFinish).Error
	return userTaskFinish, err
}
