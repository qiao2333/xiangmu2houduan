package impl

import (
	"project/project3/sql/common"
	"project/project3/sql/models"
)

//UserCollectionImpl 用户收藏表Impl
type UserCollectionImpl struct {
	MUC models.UserCollection
}

//New conditions接口
func (impl UserCollectionImpl) New() error {
	panic("implement me")
}

//Save conditions接口
func (impl UserCollectionImpl) Save(conditions models.UserCollection) error {
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

//Insert 用户收藏表插入
func (impl *UserCollectionImpl) Insert() (bool, string) {
	engine := common.GetDBEngine()
	session := engine.NewSession()
	defer session.Close()
	err := session.Begin()
	success := false
	if err != nil {
		panic(err)
	}
	result, err := session.InsertOne(&impl.MUC)
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

//Update 用户收藏表更新
func (impl *UserCollectionImpl) Update(id int) (bool, string) {
	engine := common.GetDBEngine()
	session := engine.NewSession()
	defer session.Close()
	err := session.Begin()
	success := false
	if err != nil {
		panic(err)
	}
	result, err := session.ID(id).Update(&impl.MUC)
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

//Delete 用户收藏表删除
func (impl *UserCollectionImpl) Delete(id int) (bool, string) {
	engine := common.GetDBEngine()
	session := engine.NewSession()
	defer session.Close()
	err := session.Begin()
	success := false
	if err != nil {
		panic(err)
	}
	result, err := session.ID(id).Delete(&impl.MUC)
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
