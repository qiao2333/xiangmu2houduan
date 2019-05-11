package models

import (
	"time"
)

//UserShopcar 用户购物车表
type UserShopcar struct {
	Code          int       `xorm:"not null pk autoincr comment('表') INT(11)"`
	UserCode      int       `xorm:"not null comment('用户Code') index INT(11)"`
	LightCode     int       `xorm:"not null comment('灯饰Code') index INT(11)"`
	LightNumber int       `xorm:"not null default 1 comment('购买数量') INT(11)"`
	Hasdelete   time.Time `xorm:"comment('删除时间') DATETIME deleted"`
	CreateDate  time.Time `xorm:"comment('创建时间') DATETIME created"`
}

//UserShopcarInterface 用户购物车接口
type UserShopcarInterface interface {
	New() error
	Save(conditions UserShopcar) error
}
