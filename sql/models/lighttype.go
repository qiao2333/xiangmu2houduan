package models

//Lighttype 灯饰类型表
type Lighttype struct {
	Code    int    `xorm:"not null pk autoincr comment('编号') INT(11)"`
	Label string `xorm:"comment('类型名称') VARCHAR(45)"`
}

//LightttypeInterface 灯饰类型表接口
type LightttypeInterface interface {
	New() error
	Save(conditions Lighttype) error
}
