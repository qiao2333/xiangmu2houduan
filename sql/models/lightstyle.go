package models

//Lightstyle 灯饰风格表
type Lightstyle struct {
	Code    int    `xorm:"not null pk autoincr comment('编号') INT(11)"`
	Label string `xorm:"comment('风格名称') VARCHAR(45)"`
}

//LightstyleInterface 灯饰风格表接口
type LightstyleInterface interface {
	New() error
	Save(conditions Lightstyle) error
}
