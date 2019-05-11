package impl

import (
	"project/project3/sql/common"
	"project/project3/sql/models"
)

//LightbrandImpl 灯饰品牌表Impl
type LightbrandImpl struct {
	MLBrand models.Lightbrand
}

//New conditions接口
func (impl LightbrandImpl) New() error {
	panic("implement me")
}

//Save conditions接口
func (impl LightbrandImpl) Save(conditions models.Lightbrand) error {
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

//Insert 灯饰品牌表插入
func (impl *LightbrandImpl) Insert() (bool, string) {
	engine := common.GetDBEngine()
	session := engine.NewSession()
	defer session.Close()
	err := session.Begin()
	success := false
	if err != nil {
		panic(err)
	}
	result, err := session.InsertOne(&impl.MLBrand)
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

//Update 灯饰品牌表更新
func (impl *LightbrandImpl) Update(id int) (bool, string) {
	engine := common.GetDBEngine()
	session := engine.NewSession()
	defer session.Close()
	err := session.Begin()
	success := false
	if err != nil {
		panic(err)
	}
	result, err := session.ID(id).Update(&impl.MLBrand)
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

//Delete 灯饰品牌表删除
func (impl *LightbrandImpl) Delete(id int) (bool, string) {
	engine := common.GetDBEngine()
	session := engine.NewSession()
	defer session.Close()
	err := session.Begin()
	success := false
	if err != nil {
		panic(err)
	}
	result, err := session.ID(id).Delete(&impl.MLBrand)
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
