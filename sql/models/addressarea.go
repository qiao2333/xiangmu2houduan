package models

//Addressarea 地址城区表
type Addressarea struct {
	Code      int    `xorm:"not null pk comment('编号') INT(11)" json:"value"`
	Label     string `xorm:"comment('名称') VARCHAR(45)" json:"label"`
	Citycode  int    `xorm:"comment('城市表编号') index INT(11)"`
	Statecode int    `xorm:"comment('省份表编号') index INT(11)"`
}

//AddressareaInterface 地址城区表接口
type AddressareaInterface interface {
	New() error
	Save(conditions Addressarea) error
}
