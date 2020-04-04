//generate by gen
package models

import (
	"goshop/restful/common"
)

//等级任务设置
type SystemUserTask struct {
	Id         int    `gorm:"column:id"`         //
	Name       string `gorm:"column:name"`       //任务名称
	RealName   string `gorm:"column:real_name"`  //配置原名
	TaskType   string `gorm:"column:task_type"`  //任务类型
	Number     int    `gorm:"column:number"`     //限定数
	LevelId    int    `gorm:"column:level_id"`   //等级id
	Sort       int    `gorm:"column:sort"`       //排序
	IsShow     int    `gorm:"column:is_show"`    //是否显示
	IsMust     int    `gorm:"column:is_must"`    //是否务必达成任务,1务必达成,0=满足其一
	Illustrate string `gorm:"column:illustrate"` //任务说明
	AddTime    int    `gorm:"column:add_time"`   //新增时间

}

//修改默认表名
func (SystemUserTask) TableName() string {
	return "eb_system_user_task"
}

func (systemUserTask *SystemUserTask) Insert() error {
	err := common.GetDB().Create(systemUserTask).Error
	return err
}

func (systemUserTask *SystemUserTask) Patch() error {
	err := common.GetDB().Model(systemUserTask).Updates(systemUserTask).Error
	return err
}

func (systemUserTask *SystemUserTask) Update() error {
	err := common.GetDB().Save(systemUserTask).Error
	return err
}

func (systemUserTask *SystemUserTask) Delete() error {
	return common.GetDB().Delete(systemUserTask).Error
}

func (systemUserTask *SystemUserTask) List(rawQuery string, rawOrder string, offset int, limit int) (*[]SystemUserTask, int, error) {
	systemUserTasks := []SystemUserTask{}
	total := 0
	db := common.GetDB().Model(systemUserTask)
	db, err := buildWhere(rawQuery, db)
	if err != nil {
		return &systemUserTasks, total, err
	}
	db, err = buildOrder(rawOrder, db)
	if err != nil {
		return &systemUserTasks, total, err
	}
	db.Offset(offset).
		Limit(limit).
		Find(&systemUserTasks).
		Count(&total)
	err = db.Error
	return &systemUserTasks, total, err
}

func (systemUserTask *SystemUserTask) Get() (*SystemUserTask, error) {
	err := common.GetDB().Find(&systemUserTask).Error
	return systemUserTask, err
}
