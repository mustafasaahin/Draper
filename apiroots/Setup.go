package apiroots

import (
	"github.com/gin-gonic/gin"
	"github.com/mustafasaahin/Draper/apicontrollers"
)

func SetupApiRoot(api *gin.RouterGroup) {
	api.GET("/setup", apicontrollers.GETSetup)
	api.GET("/setup/:id", apicontrollers.GETSetupByID)
	api.POST("/setup", apicontrollers.POSTSetup)
	api.PUT("/setup/:id", apicontrollers.PUTSetup)
	api.DELETE("/setup/:id", apicontrollers.DELETESetup)
}
