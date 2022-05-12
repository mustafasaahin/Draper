package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/mustafasaahin/Draper/apicontrollers"
	"github.com/mustafasaahin/Draper/apiroots"
	"github.com/mustafasaahin/Draper/config"
	"github.com/mustafasaahin/Draper/libs"
	"net/http"
	"time"
)

func main() {
	config.InitDB()
	InitRecords()
	app := gin.Default()

	app.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"PUT", "PATCH"},
		AllowHeaders:     []string{"Origin", "Accept", "Content-Type","Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	api := app.Group("/api", libs.BiletKontrol)
	apiroots.YarnsApiRoot(api)
	apiroots.ProductApiRoot(api)
	apiroots.SetupApiRoot(api)
	apiroots.DressingPriceListApiRoot(api)
	app.POST("/register", apicontrollers.POSTRegister)
	app.POST("/login", apicontrollers.POSTLogin)

	app.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello from Mustafa.")
	})

	err := app.Run(config.GetOutboundIP() + ":8014")
	if err != nil {
		return
	}
}
