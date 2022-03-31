package InitData

import (
	"github.com/mustafasaahin/Draper/config"
	"github.com/mustafasaahin/Draper/models"
	"time"
)

func InitUser() {
	config.DB.Where("id<>0").Delete(&models.User{})
	var user = models.User{
		Name:               "Mustafa",
		Surname:            "Şahin",
		Mail:               "shnmustafa28@gmail.com",
		Phone:              "05314246047",
		PersonalIdentityNo: "10895301298",
		BirthDate:          time.Date(1998, 02, 07, 0, 0, 0, 0, time.UTC),
		CompanyName:        "Icitech Technology",
		Gender:             "Erkek",
		Password:           config.HashPassword("ms1998ms"),
	}
	config.DB.Create(&user)
	user = models.User{
		Name:               "Elif",
		Surname:            "Aslım",
		Mail:               "elifaslim34@gmail.com",
		Phone:              "05439731085",
		PersonalIdentityNo: "10511521680",
		BirthDate:          time.Date(2005, 10, 14, 0, 0, 0, 0, time.UTC),
		CompanyName:        "Hasbahce",
		Gender:             "Kadın",
		Password:           config.HashPassword("elif2005"),
	}
	config.DB.Create(&user)
}
