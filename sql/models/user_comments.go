package models

import (
	"time"
)

//UserComments 用户评论表
type UserComments struct {
	Code       int       `xorm:"not null pk autoincr comment('编号') INT(11)"`
	UserCode   int       `xorm:"not null comment('用户Code') index INT(11)"`
	LightCode  int       `xorm:"not null comment('灯饰Code') index INT(11)"`
	Comments   string    `xorm:"comment('评论') VARCHAR(200)"`
	Hasdelete  time.Time `xorm:"comment('删除时间') DATETIME created"`
	CreateDate time.Time `xorm:"comment('创建时间') DATETIME deleted"`
}

//UserCommentsInterface 用户评论表
type UserCommentsInterface interface {
	New() error
	Save(conditions UserComments) error
}
