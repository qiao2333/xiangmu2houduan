package models

//UserInfo 用户信息表
type UserInfo struct {
	Code      int    `xorm:"not null pk comment('用户Code') INT(11)"`
	UserSex   int    `xorm:"default 0 comment('性别') INT(11)"`
	UserEmail string `xorm:"comment('邮箱') VARCHAR(45)"`
	UserPhone string `xorm:"comment('手机') VARCHAR(20)"`
	UserCode  string `xorm:"comment('邮编地址') VARCHAR(45)"`
}

//UserInfoInterface 用户信息表接口
type UserInfoInterface interface {
	New() error
	Save(conditions UserInfo) error
}
