//generate by gen
package models

import (
	"goshop/restful/common"
)

//小程序二维码管理表
type RoutineQrcode struct {
	Id        int    `gorm:"column:id"`         //微信二维码ID
	ThirdType string `gorm:"column:third_type"` //二维码类型 spread(用户推广) product_spread(产品推广)
	ThirdId   int    `gorm:"column:third_id"`   //用户id
	Status    int    `gorm:"column:status"`     //状态 0不可用 1可用
	AddTime   string `gorm:"column:add_time"`   //添加时间
	Page      string `gorm:"column:page"`       //小程序页面路径带参数
	QrcodeUrl string `gorm:"column:qrcode_url"` //小程序二维码路径
	UrlTime   int    `gorm:"column:url_time"`   //二维码添加时间

}

//修改默认表名
func (RoutineQrcode) TableName() string {
	return "eb_routine_qrcode"
}

func (routineQrcode *RoutineQrcode) Insert() error {
	err := common.GetDB().Create(routineQrcode).Error
	return err
}

func (routineQrcode *RoutineQrcode) Patch() error {
	err := common.GetDB().Model(routineQrcode).Updates(routineQrcode).Error
	return err
}

func (routineQrcode *RoutineQrcode) Update() error {
	err := common.GetDB().Save(routineQrcode).Error
	return err
}

func (routineQrcode *RoutineQrcode) Delete() error {
	return common.GetDB().Delete(routineQrcode).Error
}

func (routineQrcode *RoutineQrcode) List(rawQuery string, rawOrder string, offset int, limit int) (*[]RoutineQrcode, int, error) {
	routineQrcodes := []RoutineQrcode{}
	total := 0
	db := common.GetDB().Model(routineQrcode)
	db, err := buildWhere(rawQuery, db)
	if err != nil {
		return &routineQrcodes, total, err
	}
	db, err = buildOrder(rawOrder, db)
	if err != nil {
		return &routineQrcodes, total, err
	}
	db.Offset(offset).
		Limit(limit).
		Find(&routineQrcodes).
		Count(&total)
	err = db.Error
	return &routineQrcodes, total, err
}

func (routineQrcode *RoutineQrcode) Get() (*RoutineQrcode, error) {
	err := common.GetDB().Find(&routineQrcode).Error
	return routineQrcode, err
}
