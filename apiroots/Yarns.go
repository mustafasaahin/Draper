package apiroots

import (
	"github.com/gin-gonic/gin"
	"github.com/mustafasaahin/Draper/apicontrollers"
)

func YarnsApiRoot(api *gin.RouterGroup) {
	api.GET("/yarns", apicontrollers.GETYarns)
	api.GET("/yarn/:id", apicontrollers.GETYarnsByID)
	api.POST("/yarn", apicontrollers.POSTYarns)
	api.PUT("/yarn/:id", apicontrollers.PUTYarns)
	api.DELETE("/yarn/:id", apicontrollers.DELETEYarns)

}
