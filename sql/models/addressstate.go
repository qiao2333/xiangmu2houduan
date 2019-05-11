package models

//Addressstate 地址省份表
type Addressstate struct {
	Code  int    `xorm:"not null pk comment('省份编号') INT(11)" json:"value"`
	Label string `xorm:"comment('省份名称') VARCHAR(45)" json:"label"`
}

//AddressstateInterface 地址省份表接口
type AddressstateInterface interface {
	New() error
	Save(conditions Addressstate) error
}
