package models

import (
	"goshop/restful/common"
)

type SystemUserTaskLevel struct {
	SystemUserTask
	LevelName string
}

type SystemUserLevelGrade struct {
	SystemUserLevel
	IsClear  bool
	TaskList []*SystemUserTaskLevel
}

func (systemUserLevel *SystemUserLevel) GetLevelListAndGrade(levelId int, isTaskList bool) (list []*SystemUserLevelGrade, err error) {
	db := common.GetDB()
	gradle := 0

	list = make([]*SystemUserLevelGrade, 0)
	db = db.Select("name, discount, image, icon, explain, id, grade")
	err = db.Order("grade asc").Find(&list).Error
	if err != nil {
		return
	}

	mSystemUserLevel := make(map[int]*SystemUserLevelGrade, 0)

	for _, item := range list {
		mSystemUserLevel[item.Id] = item
	}

	tasks := make([]*SystemUserTaskLevel, 0)
	if isTaskList {
		err = db.Order("sort desc, add_time desc").Find(&tasks).Error
		if err != nil {
			return
		}
	}

	mTasks := make(map[int][]*SystemUserTaskLevel, 0)

	for _, item := range tasks {
		item.LevelName = mSystemUserLevel[item.LevelId].Name
		if v, ok := mTasks[item.LevelId]; ok {
			mTasks[item.LevelId] = append(v, item)
			continue
		}
		mTasks[item.LevelId] = []*SystemUserTaskLevel{item}
	}

	for _, item := range list {
		if item.Id == levelId {
			gradle = item.Grade
		}

	}

	for _, item := range list {
		if gradle < item.Grade {
			item.IsClear = true
		}
		item.TaskList = mTasks[item.Id]
	}
	return
}
