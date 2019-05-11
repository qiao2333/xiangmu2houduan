package models

import (
	"time"
)

//User 用户表
type User struct {
	Code           int       `xorm:"not null pk autoincr comment('编号') INT(11)"`
	UserName     string    `xorm:"not null comment('用户名') unique VARCHAR(20)"`
	UserPassword string    `xorm:"not null comment('密码') VARCHAR(20)"`
	UserType     int       `xorm:"not null default 0 comment('用户类型0：用户、1：雇员、2：管理员') INT(11)"`
	Hasdelete    time.Time `xorm:"comment('删除时间') DATETIME deleted"`
	CreateDate   time.Time `xorm:"comment('创建时间') DATETIME created "`
}

//UserInterface 用户表接口
type UserInterface interface {
	New() error
	Save(conditions User) error
}
