package models

import "time"

//UserAddress 用户地址表
type UserAddress struct {
	Code         int       `xorm:"not null pk autoincr comment('编号') INT(11)"`
	UserCode     int       `xorm:"not null comment('用户Code') index INT(11)"`
	Statecode  int       `xorm:"comment('省份编号') index INT(11)"`
	Citycode   int       `xorm:"comment('城市编号') index INT(11)"`
	Areacode   int       `xorm:"comment('城区编号') index INT(11)"`
	Streetcode int       `xorm:"comment('街道编号') index INT(11)"`
	Hasdelete  time.Time `xorm:"comment('删除时间') DATETIME deleted"`
}

//UserAddressInterface 用户地址表接口
type UserAddressInterface interface {
	New() error
	Save(conditions UserAddress) error
}
