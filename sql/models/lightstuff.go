package models

//Lightstuff 灯饰材质表
type Lightstuff struct {
	Code    int    `xorm:"not null pk autoincr comment('编号') INT(11)"`
	Label string `xorm:"comment('材质名称') VARCHAR(45)"`
}

//LightstuffInterface 灯饰材质表接口
type LightstuffInterface interface {
	New() error
	Save(conditions Lightstuff) error
}
