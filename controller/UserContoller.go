package controller

import (
	"fmt"
	"project/project3/sql/common"
	"project/project3/sql/impl"
	"project/project3/sql/models"
)

//Login 用户登录操作
func Login(userName string, password string) (bool, models.User) {
	user := new(impl.UserImpl)
	engine := common.GetDBEngine()
	result, err := engine.Where("user_name = ?", userName).And("user_password = ?", password).Get(&user.MU)
	if err != nil {
		fmt.Println("错误", err)
		panic(nil)
	}
	return result, user.MU
}
