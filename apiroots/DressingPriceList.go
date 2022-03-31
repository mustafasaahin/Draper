package apiroots

import (
	"github.com/gin-gonic/gin"
	"github.com/mustafasaahin/Draper/apicontrollers"
)

func DressingPriceListApiRoot(api *gin.RouterGroup) {
	api.GET("/dressingPriceList", apicontrollers.GETDressingPriceList)
	api.GET("/dressingPriceList/:id", apicontrollers.GETDressingPriceListByID)
	api.POST("/dressingPriceList", apicontrollers.POSTDressingPriceList)
	api.PUT("/dressingPriceList/:id", apicontrollers.PUTDressingPriceList)
	api.DELETE("/dressingPriceList/:id", apicontrollers.DELETEDressingPriceList)
}
