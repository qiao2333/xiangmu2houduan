package xorm

import (
	"fmt"
	"project/project3/sql/common"
	"project/project3/sql/models"
)

func Insert_light(name string) int64 {
	engine := common.GetDBEngine()
	session := engine.NewSession()
	err := session.Begin()
	if err != nil {
		fmt.Println("事务开启失败")
		panic(err)
	}
	light := &models.Light{
		LightName: name,
	}

	result, err := session.Cols("light_name").InsertOne(light)

	if err != nil {
		fmt.Println("插入失败")
		session.Rollback()
		panic(err)
	}
	err = session.Commit()

	if err != nil {
		fmt.Println("事务提交失败")
		panic(err)
	}
	return result
}
