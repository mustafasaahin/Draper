package InitData

import (
	"fmt"
	"github.com/mustafasaahin/Draper/config"
	"github.com/mustafasaahin/Draper/models"
)

func InitYarns() {
	var yarns models.Yarns
	config.DB.
		Where("id<>0").
		Delete(&models.Yarns{})
	yarns = models.Yarns{
		Code:      "001",
		Name:      "13/1 KAPLANLAR 30 WOOL 70 PES",
		UnitPrice: 7.5,
	}
	config.DB.Create(&models.Yarns{})
	fmt.Println(yarns.ID)
	yarns = models.Yarns{
		Code:      "002",
		Name:      "70dn lycra(75f36+40f1)  115f37   COMBAT 18.10.2021",
		UnitPrice: 4.29,
	}
	config.DB.Create(&models.Yarns{})
	fmt.Println(yarns.ID)
	yarns = models.Yarns{
		Code:      "003",
		Name:      "150 Denye Koyu Tintin",
		UnitPrice: 3.5,
	}
	config.DB.Create(&models.Yarns{})
	fmt.Println(yarns.ID)
	yarns = models.Yarns{
		Code:      "004",
		Name:      "150 F 72 Tekstüre  Melanj     18.10.2021",
		UnitPrice: 3.33,
	}
	config.DB.Create(&models.Yarns{})
	fmt.Println(yarns.ID)
	yarns = models.Yarns{
		Code:      "005",
		Name:      "450 F 144 Tekstüre Melanj     18.10.2021",
		UnitPrice: 2.76,
	}
	config.DB.Create(&models.Yarns{})
	fmt.Println(yarns.ID)
	yarns = models.Yarns{
		Code:      "006",
		Name:      "28/1boyalıK.ELYAF flam + 40/1 renkli likra",
		UnitPrice: 6.3,
	}
	config.DB.Create(&models.Yarns{})
	fmt.Println(yarns.ID)
	yarns = models.Yarns{
		Code:      "007",
		Name:      "150 F 48T Teksture Img - SPR       18.10.2021",
		UnitPrice: 2.44,
	}
	config.DB.Create(&models.Yarns{})
	fmt.Println(yarns.ID)
	yarns = models.Yarns{
		Code:      "008",
		Name:      "300 F 96 S Texture Img Siyah   18.10.2021",
		UnitPrice: 2.58,
	}
	config.DB.Create(&models.Yarns{})
	fmt.Println(yarns.ID)
	yarns = models.Yarns{
		Code:      "009",
		Name:      "300 F 120 Teks Melanj       18.10.2021",
		UnitPrice: 2.78,
	}
	config.DB.Create(&models.Yarns{})
	fmt.Println(yarns.ID)

}
