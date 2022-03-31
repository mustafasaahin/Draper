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
		models.User{},
	); err != nil {
		fmt.Println(err.Error())
	}
	if err := config.DB.AutoMigrate(
		models.Setup{},
		models.Product{},
		models.Yarns{},
		models.User{},
		models.DressingPriceList{}); err != nil {
		fmt.Println(err.Error())
	}
	InitData.InitDressingPriceList()
	InitData.InitYarns()
	InitData.InitSetup()
	InitData.InitUser()
	InitData.InitProduct()

}
