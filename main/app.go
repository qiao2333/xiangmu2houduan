package main

import (
	"project/project3/service"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	//登录
	r.POST("/login", service.UserLogin)
	//管理员管理灯饰的路由组

	c := r.Group("/controller/light")
	{
		c.GET("/select/:choice", service.ControllerSelect)
		c.POST("/SQL/:chioce", service.ControllerLightSQL)
	}

	//地址路由
	addressService := new(service.AddressService)
	c1 := r.Group("/address")
	{
		c1.GET("/state", addressService.GetState)
		c1.GET("/city/:code", addressService.GetCity)
		c1.GET("/area/:code", addressService.GetAreas)
		c1.GET("/street/:code", addressService.GetStreet)
		c1.GET("/getAll/:stateCode/:cityCode/:areaCode", addressService.GetAll)
	}

	r.Run(":8080")
}
