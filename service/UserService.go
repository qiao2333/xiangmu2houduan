package service

import (
	"net/http"
	"project/project3/controller"

	"github.com/gin-gonic/gin"
)

//UserLogin 用户登录
func UserLogin(c *gin.Context) {
	userName := c.PostForm("userName")
	password := c.PostForm("password")
	result, user := controller.Login(userName, password)
	if result == true {
		c.JSON(http.StatusOK, gin.H{"login": true, "datas": user})
	} else {
		c.JSON(http.StatusNotFound, gin.H{"result": false})
	}

}
