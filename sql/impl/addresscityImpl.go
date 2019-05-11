package impl

import (
	"fmt"
	"project/project3/sql/common"
	"project/project3/sql/models"
)

//AddresscityImpl 地址表城市类Impl
type AddresscityImpl struct {
	MCity models.Addresscity
}

//New conditions接口
func (impl AddresscityImpl) New() error {
	panic("implement me")
}

//Save conditions接口
func (impl AddresscityImpl) Save(conditions models.Addresscity) error {
	panic("implement me")
}

//session的基本流程
// engine := common.GetDBEngine()
// 	session := Engine.NewSession()
// 	defer session.Close()
// 	err := session.Begin()
// 	if err != nil {
// 		panic(err)
// 	}

// 	err = session.Commit()
// 	if err != nil {
// 		panic(err)
// 	}

//Insert 地址城市表插入
func (impl *AddresscityImpl) Insert() (bool, string) {
	engine := common.GetDBEngine()
	session := engine.NewSession()
	defer session.Close()
	err := session.Begin()
	success := false
	if err != nil {
		panic(err)
	}
	result, err := session.InsertOne(&impl.MCity)
	if err == nil && result > 0 {
		success = true
	} else {
		session.Rollback()
		panic(err)
	}
	err = session.Commit()
	if err != nil {
		panic(err)
	}
	if success {
		return success, "插入成功"
	}
	return false, "插入失败"

}

//Update 地址城市表更新
func (impl *AddresscityImpl) Update(id int) (bool, string) {
	engine := common.GetDBEngine()
	session := engine.NewSession()
	defer session.Close()
	err := session.Begin()
	success := false
	if err != nil {
		panic(err)
	}
	result, err := session.ID(id).Update(&impl.MCity)
	if err == nil && result > 0 {
		success = true
	} else {
		session.Rollback()
	}

	err = session.Commit()
	if err != nil {
		panic(err)
	}

	if success {
		return true, "更新成功"
	}
	return false, "更新失败"
}

//Delete 地址城市表删除
func (impl *AddresscityImpl) Delete(id int) (bool, string) {
	engine := common.GetDBEngine()
	session := engine.NewSession()
	defer session.Close()
	err := session.Begin()
	success := false
	if err != nil {
		panic(err)
	}
	result, err := session.ID(id).Delete(&impl.MCity)
	if err == nil && result > 0 {
		success = true
	} else {
		session.Rollback()
	}
	err = session.Commit()
	if err != nil {
		panic(err)
	}
	if success {
		return true, "删除成功"
	}
	return false, "删除失败"

}

//GetCities 根据省份code获取城市code
func (impl *AddresscityImpl) GetCities(code int) []models.Addresscity {
	engine := common.GetDBEngine()
	cities := make([]models.Addresscity, 0)
	err := engine.Where("stateCode=?", code).Find(&cities)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	return cities
}
