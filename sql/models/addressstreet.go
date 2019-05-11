package models

//Addressstreet 地址街道表
type Addressstreet struct {
	Code      int    `xorm:"not null pk comment('街道表编号') INT(11)" json:"value"`
	Label     string `xorm:"comment('街道名称') VARCHAR(45)" json:"label"`
	Statecode int    `xorm:"comment('省份编号') index INT(11)"`
	Citycode  int    `xorm:"comment('城市编号') index INT(11)"`
	Areacode  int    `xorm:"comment('城市区域') index INT(11)"`
}

//AddressstreetInterface 地址街道表接口
type AddressstreetInterface interface {
	New() error
	Save(conditions Addressstreet) error
}
