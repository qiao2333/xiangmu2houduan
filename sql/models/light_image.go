package models

import "time"

//LightImage 灯饰图片表
type LightImage struct {
	Code        int       `xorm:"not null pk autoincr comment('编号') INT(11)"`
	LightCode   int       `xorm:"not null comment('灯饰Code') index INT(11)"`
	Cover     int       `xorm:"default 0 comment('0：不是封面 1：是封面') TINYINT(4)"`
	ImagePath string    `xorm:"comment('图片存储路径') VARCHAR(100)"`
	Hasdelete time.Time `xorm:"comment('是否删除') DATETIME deleted"`
}

//LightImageInterface 灯饰图片表接口
type LightImageInterface interface {
	New() error
	Save(conditions LightImage) error
}
