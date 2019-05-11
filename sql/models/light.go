package models

import (
	"time"
)

//Light 灯饰表
type Light struct {
	Code          int       `xorm:"not null pk autoincr comment('灯饰Code') INT(11)"`
	LightName     string    `xorm:"not null default '' comment('灯饰名称') VARCHAR(45)"`
	LightType     int       `xorm:"not null default 0 comment('灯饰类型') index INT(11)"`
	LightBrand    int       `xorm:"not null default 0 comment('灯饰品牌') index INT(11)"`
	LightStyle    int       `xorm:"not null default 0 comment('灯饰风格') index INT(11)"`
	LightStuff    int       `xorm:"not null default 0 comment('灯饰材质') index INT(11)"`
	LightColor    string    `xorm:"not null default '颜色不详' comment('灯饰颜色') VARCHAR(30)"`
	LightPrice    float64   `xorm:"not null default 0 comment('灯饰价格') DOUBLE"`
	LightDescript string    `xorm:"not null default '' comment('灯饰说明') VARCHAR(250)"`
	Hasdelete     time.Time `xorm:"comment('灯饰时间，如果删除则获取删除的时间') DATETIME deleted"`
	CreateDate    time.Time `xorm:"comment('灯饰数据创建时间') DATETIME created"`
}

//LightInterface 灯饰表接口
type LightInterface interface {
	New() error
	Save(conditions Light) error
}
