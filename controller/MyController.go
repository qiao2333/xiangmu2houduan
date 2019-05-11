package controller

import (
	"project/project3/sql/impl"
)

//LightDelete 灯饰删除
func LightDelete(id int) (bool, string) {
	light := new(impl.LightImpl)
	result, s := light.Delete(id)
	return result, s
}

//LightInsert 灯饰插入
func LightInsert(light *impl.LightImpl) (bool, string) {
	result, s := light.Insert()
	return result, s
}
