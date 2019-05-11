package impl

import (
	"fmt"
	"project/project3/sql/common"
	"project/project3/sql/models"
)

//AddressstreetImpl 地址街道表Impl
type AddressstreetImpl struct {
	MStreet models.Addressstreet
}

//New conditions接口
func (impl AddressstreetImpl) New() error {
	panic("implement me")
}

//Save conditions接口
func (impl AddressstreetImpl) Save(conditions models.Addressstreet) error {
	panic("implement me")
}

//Insert 地址街道表插入
func (impl *AddressstreetImpl) Insert() (bool, string) {
	engine := common.GetDBEngine()
	session := engine.NewSession()
	defer session.Close()
	err := session.Begin()
	success := false
	if err != nil {
		panic(err)
	}
	result, err := session.InsertOne(&impl.MStreet)
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

//Update 地址街道表更新
func (impl *AddressstreetImpl) Update(id int) (bool, string) {
	engine := common.GetDBEngine()
	session := engine.NewSession()
	defer session.Close()
	err := session.Begin()
	success := false
	if err != nil {
		panic(err)
	}
	result, err := session.ID(id).Update(&impl.MStreet)
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

//Delete 地址街道表删除
func (impl *AddressstreetImpl) Delete(id int) (bool, string) {
	engine := common.GetDBEngine()
	session := engine.NewSession()
	defer session.Close()
	err := session.Begin()
	success := false
	if err != nil {
		panic(err)
	}
	result, err := session.ID(id).Delete(&impl.MStreet)
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

//GetStreet 根据城区code获取街道code
func (impl *AddressstreetImpl) GetStreet(code int) []models.Addressstreet {
	engine := common.GetDBEngine()
	streets := make([]models.Addressstreet, 0)
	err := engine.Where("areaCode=?", code).Find(&streets)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	return streets
}
