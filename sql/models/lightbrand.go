package models

//Lightbrand 灯饰品牌表
type Lightbrand struct {
	Code    int    `xorm:"not null pk autoincr comment('编号') INT(11)"`
	Label string `xorm:"comment('品牌名称') VARCHAR(45)"`
}

//LightbrandInterface 灯饰品牌表接口
type LightbrandInterface interface {
	New() error
	Save(conditions Lightbrand) error
}
