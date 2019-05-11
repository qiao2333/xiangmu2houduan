package service

import (
	"net/http"
	"project/project3/controller"
	"project/project3/sql/impl"
	"strconv"

	"github.com/gin-gonic/gin"
)

//AddressService 地址Service
type AddressService struct {
}

//GetAreas 获取城区数据
func (a AddressService) GetAreas(c *gin.Context) {
	area := new(impl.AddressareaImpl)
	s := c.Param("code")
	code, _ := strconv.Atoi(s)
	result := area.GetAreas(code)
	c.JSON(http.StatusOK, gin.H{"datas": result})
}

//GetCity 获取城市数据
func (a AddressService) GetCity(c *gin.Context) {
	city := new(impl.AddresscityImpl)
	s := c.Param("code")
	code, _ := strconv.Atoi(s)
	result := city.GetCities(code)
	c.JSON(http.StatusOK, gin.H{"datas": result})
}

//GetState 获取城区数据
func (a AddressService) GetState(c *gin.Context) {
	state := new(impl.AddressstateImpl)
	result := state.GetState()
	c.JSON(http.StatusOK, gin.H{"datas": result})
}

//GetStreet 获取街道数据
func (a AddressService) GetStreet(c *gin.Context) {
	street := new(impl.AddressstreetImpl)
	s := c.Param("code")
	code, _ := strconv.Atoi(s)
	result := street.GetStreet(code)
	c.JSON(http.StatusOK, gin.H{"datas": result})
}

//GetAll 一次性获取所有数据（用于更新操作）
func (a AddressService) GetAll(c *gin.Context) {
	stateCode := c.Param("stateCode")
	cityCode := c.Param("cityCode")
	areaCode := c.Param("areaCode")
	addressController := new(controller.AddressController)
	states, cities, areas, streets := addressController.GetAll(stateCode, cityCode, areaCode)

	c.JSON(http.StatusOK, gin.H{"states": states, "areas": areas, "cities": cities, "streets": streets})
}
