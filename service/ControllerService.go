package service

import (
	"fmt"
	"log"
	"net/http"
	"project/project3/controller"
	"project/project3/sql/impl"
	"strconv"

	"github.com/gin-gonic/gin"
)

//ControllerLightSQL 灯饰基本操作
func ControllerLightSQL(c *gin.Context) {
	chioce := c.Param("chioce")
	log.Println("-----------", chioce)
	var result bool
	var s string
	if chioce == "update" {

		log.Println("********修改")
	} else if chioce == "delete" {
		c1 := c.PostForm("id")
		id, _ := strconv.Atoi(c1)
		result, s = controller.LightDelete(id)
		log.Println("********删除")
	} else if chioce == "insert" {
		light := new(impl.LightImpl)
		result, s = controller.LightInsert(light)
		log.Println("********插入")
	}
	c.String(http.StatusOK, "结果%v，%v", result, s)
}

//ControllerSelect 不知道
func ControllerSelect(c *gin.Context) {
	choice := c.Param("chioce")
	fmt.Println(choice)
}

//GetLightMessage 获取灯饰信息返回到修改页面
func GetLightMessage(c *gin.Context) {

}
