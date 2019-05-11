package models

//Addresscity 地址城市表
type Addresscity struct {
	Code      int    `xorm:"not null pk comment('城市表编号') INT(11)" json:"value"`
	Label     string `xorm:"comment('城市名称') VARCHAR(45)" json:"label"`
	Statecode int    `xorm:"comment('省份表编号') index INT(11)"`
}

//AddresscityInterface 地址城市表接口
type AddresscityInterface interface {
	New() error
	Save(conditions Addresscity) error
}
