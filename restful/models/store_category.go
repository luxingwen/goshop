package models

import (
	"goshop/restful/common"
)

func (storeCategory *StoreCategory) ListByIndex(limit int) (r []*StoreCategory, err error) {
	db := common.GetDB()
	err = db.Select("id, cate_name, pid, pic").Where("pid > ? and is_show = ?", 0, 1).Order("sort DESC").Limit(limit).Find(&r).Error
	return
}

func (storeCategory *StoreCategory) PidByCategory(pid int) (r []*StoreCategory, err error) {
	db := common.GetDB()
	err = db.Select("id,cate_name").Where("pid = ? AND is_show = ?", pid, 1).Order("sort DESC, id DESC").Find(&r).Error
	return
}

// 获取一级和二级分类
func (storeCategory *StoreCategory) GetProductCategory() (r []map[string]interface{}, err error) {
	db := common.GetDB()

	list := make([]*StoreCategory, 0)
	err = db.Find(&list).Error
	if err != nil {
		return
	}

	mdata := make(map[int]*StoreCategory, 0)
	for _, item := range list {
		mdata[item.Id] = item
	}

	r = make([]map[string]interface{}, 0)

	mList := make(map[int][]map[string]interface{}, 0)

	for _, item := range list {
		if item.Pid == 0 {
			continue
		}
		valMap := make(map[string]interface{}, 0)
		valMap["cate_name"] = item.CateName
		valMap["id"] = item.Id
		valMap["pic"] = item.Pic
		if v, ok := mList[item.Pid]; ok {
			mList[item.Pid] = append(v, valMap)
		} else {
			mList[item.Pid] = []map[string]interface{}{valMap}
		}
	}
	for _, item := range list {
		if item.Pid == 0 {
			mMap := make(map[string]interface{}, 0)
			mMap["cate_name"] = item.CateName
			mMap["id"] = item.Id
			mMap["child"] = mList[item.Id]
			r = append(r, mMap)
		}
	}
	return

}
