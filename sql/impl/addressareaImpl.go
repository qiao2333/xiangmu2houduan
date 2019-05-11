package impl

import (
	"fmt"
	"project/project3/sql/common"
	"project/project3/sql/models"
)

//AddressareaImpl 地址表area Impl
type AddressareaImpl struct {
	MArea models.Addressarea
}

//New conditions接口
func (impl AddressareaImpl) New() error {
	panic("implement me")
}

//Save conditions接口
func (impl AddressareaImpl) Save(conditions models.Addressarea) error {
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

//Insert 地址表Area插入
func (impl *AddressareaImpl) Insert() (bool, string) {
	engine := common.GetDBEngine()
	session := engine.NewSession()
	defer session.Close()
	err := session.Begin()
	success := false
	if err != nil {
		panic(err)
	}
	result, err := session.InsertOne(&impl.MArea)
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

//Update 地址表Area更新
func (impl *AddressareaImpl) Update(id int) (bool, string) {
	engine := common.GetDBEngine()
	session := engine.NewSession()
	defer session.Close()
	err := session.Begin()
	success := false
	if err != nil {
		panic(err)
	}
	result, err := session.ID(id).Update(&impl.MArea)
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

//Delete 地址表Area删除
func (impl *AddressareaImpl) Delete(id int) (bool, string) {
	engine := common.GetDBEngine()
	session := engine.NewSession()
	defer session.Close()
	err := session.Begin()
	success := false
	if err != nil {
		panic(err)
	}
	result, err := session.ID(id).Delete(&impl.MArea)
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

//GetAreas 根据城市code获取城区code
func (impl *AddressareaImpl) GetAreas(code int) []models.Addressarea {
	engine := common.GetDBEngine()
	areas := make([]models.Addressarea, 0)
	err := engine.Where("cityCode=?", code).Find(&areas)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	return areas
}
