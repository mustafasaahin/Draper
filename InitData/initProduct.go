package InitData

import (
	"fmt"
	"github.com/mustafasaahin/Draper/config"
	"github.com/mustafasaahin/Draper/models"
)

func InitProduct() {
	var product models.Product
	config.DB.Where("id<>0").
		Delete(&models.Product{})
	product = models.Product{
		Code: "ARYA",
		Name: "BLACK",

		WarpYarn1:  "20/1 PENYE PAMUK- ERDEM TEKSTİL USA 100%",
		Gram1:      0.19,
		Wastage1:   1.05,
		TotalGram1: 0.2,
		UnitPrice1: 4.3,
		Cost1:      0.86,

		WarpYarn2:  "28/1 renkli +40/1 polyvison likra",
		Gram2:      0.35,
		Wastage2:   1.05,
		TotalGram2: 0.351,
		UnitPrice2: 0,
		Cost2:      0,

		WarpYarn3:  "",
		Gram3:      0,
		Wastage3:   1.05,
		TotalGram3: 0,
		UnitPrice3: 0,
		Cost3:      0,

		WarpYarn4:  "",
		Gram4:      0,
		Wastage4:   1.05,
		TotalGram4: 0,
		UnitPrice4: 0,
		Cost4:      0,

		WarpYarn5:  "",
		Gram5:      0,
		Wastage5:   1.05,
		TotalGram5: 0,
		UnitPrice5: 0,
		Cost5:      0,
	}
	config.DB.Create(&product)
	fmt.Println(product.ID)
}
