package models

import (
	"time"
)

//UserOrders 用户订单表
type UserOrders struct {
	Code        int       `xorm:"not null pk autoincr comment('编号') INT(11)"`
	UserCode    int       `xorm:"not null comment('用户Code') index INT(11)"`
	LightCode   int       `xorm:"not null comment('灯饰Code') index INT(11)"`
	LightNumber int       `xorm:"not null comment('购买数量') INT(11)"`
	AddressCode int       `xorm:"not null comment('地址表Code') index INT(11)"`
	TotalPrice  float64   `xorm:"not null default 0 comment('总价') DOUBLE"`
	Status      int       `xorm:"not null default 0 comment('状态') INT(11)"`
	PostType    int       `xorm:"not null default 0 comment('邮寄类型1、顺丰 2、圆通 3韵达') INT(11)"`
	PostFee     float64   `xorm:"not null default 0 comment('邮寄费用') DOUBLE"`
	Hasdelete   time.Time `xorm:"comment('删除时间') DATETIME deleted"`
	CreateDate  time.Time `xorm:"comment('创建时间') DATETIME created"`
}

//UserOrdersInterface 用户订单表接口
type UserOrdersInterface interface {
	New() error
	Save(conditions UserOrders) error
}
