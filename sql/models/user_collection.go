package models

import (
	"time"
)

//UserCollection 用户收藏夹表
type UserCollection struct {
	Code           int       `xorm:"not null pk autoincr comment('编号') INT(11)"`
	UserCode       int       `xorm:"not null comment('用户Code') index INT(11)"`
	LightCode      int       `xorm:"not null comment('灯饰Code') index INT(11)"`
	Hasdelete      time.Time `xorm:"comment('删除时间') DATETIME deleted"`
	CollectionDate time.Time `xorm:"comment('收藏时间') DATETIME created"`
}

//UserCollectionInterface 用户收藏夹表接口
type UserCollectionInterface interface {
	New() error
	Save(conditions UserCollection) error
}
