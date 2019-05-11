package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//InsertLight 灯饰数据插入
func InsertLight(c *gin.Context) {
	c.HTML(http.StatusOK, "template1.tmpl", gin.H{})
}
