package InitData

import (
	"github.com/mustafasaahin/Draper/config"
	"github.com/mustafasaahin/Draper/models"
)

func InitSetup() {
	var setup models.Setup
	config.DB.
		Where("id<>0").
		Delete(&models.Setup{})
	setup = models.Setup{
		WeavingWarpPrice:    0,
		ExchangeUSD:         0,
		ExchangeEUR:         0,
		DressingExchaneType: "",
	}
	config.DB.Create(&setup)
}
