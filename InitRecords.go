package main

import (
	"fmt"
	"github.com/mustafasaahin/Draper/InitData"
	"github.com/mustafasaahin/Draper/config"
	"github.com/mustafasaahin/Draper/models"
)

func InitRecords() {
	if err := config.DB.Migrator().DropTable(
		models.DressingPriceList{},
		models.Yarns{},
		models.Product{},
		models.Setup{},
	); err != nil {
		fmt.Println(err.Error())
	}
	if err := config.DB.AutoMigrate(
		models.Setup{},
		models.Product{},
		models.Yarns{},
		models.DressingPriceList{}); err != nil {
		fmt.Println(err.Error())
	}
	InitData.InitDressingPriceList()
	InitData.InitProduct()
	InitData.InitYarns()
	InitData.InitSetup()

}
