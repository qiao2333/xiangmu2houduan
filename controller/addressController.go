package controller

import (
	"project/project3/sql/common"
	"project/project3/sql/models"
)

//AddressController 地址Controller类
type AddressController struct {
}

//GetAll 获取所有地址（用于更新操作）
func (a AddressController) GetAll(stateCode string, cityCode string, areaCode string) ([]models.Addressstate, []models.Addresscity, []models.Addressarea, []models.Addressstreet) {
	engine := common.GetDBEngine()
	states := make([]models.Addressstate, 0)
	cities := make([]models.Addresscity, 0)
	areas := make([]models.Addressarea, 0)
	streets := make([]models.Addressstreet, 0)
	engine.Sql("select * from addressstate").Find(&states)
	engine.Sql("select * from addresscity where stateCode = ?", stateCode).Find(&cities)
	engine.Sql("select * from addressarea where cityCode = ?", cityCode).Find(&areas)
	engine.Sql("select * from addressstreet where areaCode = ?", areaCode).Find(&streets)
	return states, cities, areas, streets
}
