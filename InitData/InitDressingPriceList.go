package InitData

import (
	"fmt"
	"github.com/mustafasaahin/Draper/config"
	"github.com/mustafasaahin/Draper/models"
)

func InitDressingPriceList() {
	var DressingPrice models.DressingPriceList
	config.DB.
		Where("id<>0").
		Delete(&models.DressingPriceList{})

	DressingPrice = models.DressingPriceList{
		Code:        "P001",
		Process:     "Polyester Tekboya",
		Price:       1.7,
		ProcessType: "Boya",
	}
	config.DB.Create(&DressingPrice)
	fmt.Println(DressingPrice.ID)

	DressingPrice = models.DressingPriceList{
		Code:        "P002",
		Process:     "Tintin",
		Price:       0,
		ProcessType: "Boya",
	}
	config.DB.Create(&DressingPrice)
	fmt.Println(DressingPrice.ID)

	DressingPrice = models.DressingPriceList{
		Code:        "P003",
		Process:     "Polyester Şifon - Yoryou",
		Price:       0,
		ProcessType: "Boya",
	}
	config.DB.Create(&DressingPrice)
	fmt.Println(DressingPrice.ID)

	DressingPrice = models.DressingPriceList{
		Code:        "P004",
		Process:     "Polyester Şifon 60 - 80 gr/mt.tül",
		Price:       0,
		ProcessType: "Boya",
	}
	config.DB.Create(&DressingPrice)
	fmt.Println(DressingPrice.ID)
	DressingPrice = models.DressingPriceList{
		Code:        "P005",
		Process:     "Katyonik Tekboya",
		Price:       1.6,
		ProcessType: "Boya",
	}
	config.DB.Create(&DressingPrice)
	fmt.Println(DressingPrice.ID)
	DressingPrice = models.DressingPriceList{
		Code:        "P006",
		Process:     "Katyonik Çift Boya",
		Price:       2.3,
		ProcessType: "Boya",
	}
	config.DB.Create(&DressingPrice)
	fmt.Println(DressingPrice.ID)

	DressingPrice = models.DressingPriceList{
		Code:        "P007",
		Process:     "Bi-Streç Tek Boya",
		Price:       0,
		ProcessType: "Boya",
	}
	config.DB.Create(&DressingPrice)
	fmt.Println(DressingPrice.ID)
	DressingPrice = models.DressingPriceList{
		Code:        "P008",
		Process:     "Bengalin Tek Boya",
		Price:       0,
		ProcessType: "Boya",
	}
	config.DB.Create(&DressingPrice)
	fmt.Println(DressingPrice.ID)
	DressingPrice = models.DressingPriceList{
		Code:        "P009",
		Process:     "Bengalin Çift Boya",
		Price:       0,
		ProcessType: "Boya",
	}
	config.DB.Create(&DressingPrice)
	fmt.Println(DressingPrice.ID)

	DressingPrice = models.DressingPriceList{
		Code:        "P010",
		Process:     "Bez Mikro",
		Price:       0,
		ProcessType: "Boya",
	}
	config.DB.Create(&DressingPrice)
	fmt.Println(DressingPrice.ID)
	DressingPrice = models.DressingPriceList{
		Code:        "P011",
		Process:     "Dimi - Soft Mikro",
		Price:       0,
		ProcessType: "Boya",
	}
	config.DB.Create(&DressingPrice)
	fmt.Println(DressingPrice.ID)
	DressingPrice = models.DressingPriceList{
		Code:        "P012",
		Process:     "Naylon Mikro-Pamuk Çift Boya",
		Price:       0,
		ProcessType: "Boya",
	}
	config.DB.Create(&DressingPrice)
	fmt.Println(DressingPrice.ID)
	DressingPrice = models.DressingPriceList{
		Code:        "P013",
		Process:     "Naylon Mikro - Pamuk Çift Boya Likra",
		Price:       0,
		ProcessType: "Boya",
	}
	config.DB.Create(&DressingPrice)
	fmt.Println(DressingPrice.ID)
	DressingPrice = models.DressingPriceList{
		Code:        "P014",
		Process:     "Pamuk Tekboya",
		Price:       1.6,
		ProcessType: "Boya",
	}
	config.DB.Create(&DressingPrice)
	fmt.Println(DressingPrice.ID)
	DressingPrice = models.DressingPriceList{
		Code:        "P015",
		Process:     "Pamuk Vual",
		Price:       0,
		ProcessType: "Boya",
	}
	config.DB.Create(&DressingPrice)
	fmt.Println(DressingPrice.ID)
	DressingPrice = models.DressingPriceList{
		Code:        "P016",
		Process:     "Poly. Pamuk Çift Boya",
		Price:       2.3,
		ProcessType: "Boya",
	}
	config.DB.Create(&DressingPrice)
	fmt.Println(DressingPrice.ID)
	DressingPrice = models.DressingPriceList{
		Code:        "P017",
		Process:     "Poly. Floş Tafta Tek Boya",
		Price:       0,
		ProcessType: "Boya",
	}
	config.DB.Create(&DressingPrice)
	fmt.Println(DressingPrice.ID)
	DressingPrice = models.DressingPriceList{
		Code:        "P018",
		Process:     "Poly. Floş Tafta Çift Boya",
		Price:       0,
		ProcessType: "Boya",
	}
	config.DB.Create(&DressingPrice)
	fmt.Println(DressingPrice.ID)
	DressingPrice = models.DressingPriceList{
		Code:        "P019",
		Process:     "Floş Viskon EN Likra",
		Price:       0,
		ProcessType: "Boya",
	}
	config.DB.Create(&DressingPrice)
	fmt.Println(DressingPrice.ID)
	DressingPrice = models.DressingPriceList{
		Code:        "P020",
		Process:     "Viskon Tek Boya",
		Price:       1.8,
		ProcessType: "Boya",
	}
	config.DB.Create(&DressingPrice)
	fmt.Println(DressingPrice.ID)
	DressingPrice = models.DressingPriceList{
		Code:        "P021",
		Process:     "Viskon Şönil Haşıl (Koyu Renk)",
		Price:       1.3,
		ProcessType: "Boya",
	}
	config.DB.Create(&DressingPrice)
	fmt.Println(DressingPrice.ID)
	DressingPrice = models.DressingPriceList{
		Code:        "P022",
		Process:     "Poly Viskon Çift Boya",
		Price:       2.5,
		ProcessType: "Boya",
	}
	config.DB.Create(&DressingPrice)
	fmt.Println(DressingPrice.ID)
	DressingPrice = models.DressingPriceList{
		Code:        "P023",
		Process:     "Poly. Viskon Çift Boya Likralı",
		Price:       1.95,
		ProcessType: "Boya",
	}
	config.DB.Create(&DressingPrice)
	fmt.Println(DressingPrice.ID)
	DressingPrice = models.DressingPriceList{
		Code:        "P024",
		Process:     "Viskon Cotton Tek Boya",
		Price:       1.7,
		ProcessType: "Boya",
	}
	config.DB.Create(&DressingPrice)
	fmt.Println(DressingPrice.ID)
	DressingPrice = models.DressingPriceList{
		Code:        "P025",
		Process:     "Viskon Cotton Lycra Tek Boya",
		Price:       2,
		ProcessType: "Boya",
	}
	config.DB.Create(&DressingPrice)
	fmt.Println(DressingPrice.ID)
	DressingPrice = models.DressingPriceList{
		Code:        "P026",
		Process:     "Işlem Yok",
		Price:       0,
		ProcessType: "Boya",
	}
	config.DB.Create(&DressingPrice)
	fmt.Println(DressingPrice.ID)
	DressingPrice = models.DressingPriceList{
		Code:        "P027",
		Process:     "Su Geçmez",
		Price:       0.6,
		ProcessType: "Ek İşlem",
	}
	config.DB.Create(&DressingPrice)
	fmt.Println(DressingPrice.ID)
	DressingPrice = models.DressingPriceList{
		Code:        "P028",
		Process:     "Yıkama Apre",
		Price:       1,
		ProcessType: "Ek İşlem",
	}
	config.DB.Create(&DressingPrice)
	fmt.Println(DressingPrice.ID)
	DressingPrice = models.DressingPriceList{
		Code:        "P029",
		Process:     "Ön Pikse",
		Price:       0.5,
		ProcessType: "Ek İşlem",
	}
	config.DB.Create(&DressingPrice)
	fmt.Println(DressingPrice.ID)
	DressingPrice = models.DressingPriceList{
		Code:        "P030",
		Process:     "Kuru Apre",
		Price:       0.7,
		ProcessType: "Ek İşlem",
	}
	config.DB.Create(&DressingPrice)
	fmt.Println(DressingPrice.ID)
	DressingPrice = models.DressingPriceList{
		Code:        "P031",
		Process:     "İnceltme",
		Price:       0.4,
		ProcessType: "Ek İşlem",
	}
	config.DB.Create(&DressingPrice)
	fmt.Println(DressingPrice.ID)
	DressingPrice = models.DressingPriceList{
		Code:        "P032",
		Process:     "Tüy Dökme",
		Price:       0.4,
		ProcessType: "Ek İşlem",
	}
	config.DB.Create(&DressingPrice)
	fmt.Println(DressingPrice.ID)
	DressingPrice = models.DressingPriceList{
		Code:        "P033",
		Process:     "Gaze",
		Price:       0.2,
		ProcessType: "Ek İşlem",
	}
	config.DB.Create(&DressingPrice)
	fmt.Println(DressingPrice.ID)
	DressingPrice = models.DressingPriceList{
		Code:        "P034",
		Process:     "Kalender",
		Price:       0.1,
		ProcessType: "Ek İşlem",
	}
	config.DB.Create(&DressingPrice)
	fmt.Println(DressingPrice.ID)
	DressingPrice = models.DressingPriceList{
		Code:        "P035",
		Process:     "Samfor",
		Price:       0.15,
		ProcessType: "Ek İşlem",
	}
	config.DB.Create(&DressingPrice)
	fmt.Println(DressingPrice.ID)
}
