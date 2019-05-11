package models

//UserHeader 用户头像图片表
type UserHeader struct {
	Code   int    `xorm:"not null pk comment('编号') INT(11)"`
	Path string `xorm:"comment('头像文件路径') VARCHAR(100)"`
}

//UserHeaderInterface 用户头像图片表接口
type UserHeaderInterface interface {
	New() error
	Save(conditions UserHeader) error
}
