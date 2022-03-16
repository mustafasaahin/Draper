package main

import (
	"fmt"
	"github.com/mustafasaahin/Draper/InitData"
	"github.com/mustafasaahin/Draper/config"
	"github.com/mustafasaahin/Draper/models"
)

func InitRecords() {
	if err :=
		config.DB.Migrator().DropTable(
			models.Setup{},
			models.Yarns{},
			models.Product{},
			models.DressingPriceList{},
		); err != nil {
		fmt.Println(err.Error())
	}
	if err := config.DB.AutoMigrate(
		models.Setup{},
		models.Yarns{},
		models.Product{},
		models.DressingPriceList{}); err != nil {
		fmt.Println(err.Error())
	}
	InitData.InitDressingPriceList()
	InitData.InitDressingPriceList()
	InitData.InitYarns()
}
