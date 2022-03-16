package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mustafasaahin/Draper/config"
)

func main() {
	config.InitDB()
	InitRecords()
	app := gin.Default()

	app.Run(config.GetOutboundIP() + ":8014")

}
