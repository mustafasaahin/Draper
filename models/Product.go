package models

type Product struct {
	ID   uint64 `json:"id" gorm:"primaryKey"`
	Code string `json:"code" gorm:"Comment:Ürün kodu"`
	Name string `json:"name" gorm:"Comment:Ürün adi"`

	WarpYarn1  string  `json:"warp_yarn_1" gorm:"Comment:Çözgü iplik 1"`
	Gram1      float64 `json:"gram_1" gorm:"Comment:Çözgü iplik gram 1"`
	Wastage1   float64 `json:"wastage_1" gorm:"Comment:Çözgü iplik % olarak fire"`
	UnitPrice1 float64 `json:"unit_price" gorm:"Comment:Çözgü iplik Birim fiyat1"`
	TotalGram1 float64 `json:"total_gram_1" gorm:"Comment:Çözgü iplik gram 1"`
	Cost1      float64 `json:"cost_1" gorm:"Comment:Çözgü iplik Maliyet1"`

	WarpYarn2  string  `json:"warp_yarn_2" gorm:"Comment:Çözgü iplik adı 2"`
	Gram2      float64 `json:"gram_2" gorm:"Comment:Çözgü iplik gram 2"`
	Wastage2   float64 `json:"wastage_2" gorm:"Comment: % olarak fire 2"`
	UnitPrice2 float64 `json:"unit_price2" gorm:"Comment:gorm:Çözgü iplik Birim fiyat 2"`
	TotalGram2 float64 `json:"total_gram_2" gorm:"Comment:Çözgü iplik gram 2"`
	Cost2      float64 `json:"cost_2" gorm:"Comment:Çözgü iplik Maliyet 2"`

	WarpYarn3  string  `json:"warp_yarn_3" gorm:"Comment:Çözgü iplik adı 3"`
	Gram3      float64 `json:"gram_3"`
	Wastage3   float64 `json:"wastage_3" gorm:"Comment:Çözgü iplik % olarak fire 3"`
	UnitPrice3 float64 `json:"unit_price_3" gorm:"Comment:Çözgü iplik Birim fiyat 3"`
	TotalGram3 float64 `json:"total_gram_3"`
	Cost3      float64 `json:"cost_3" gorm:"Comment:Çözgü iplik Maliyet 3"`

	WarpYarn4  string  `json:"warp_yarn_4" gorm:"Comment:Çözgü iplik adı 4"`
	Gram4      float64 `json:"gram_4"`
	Wastage4   float64 `json:"wastage_4" gorm:"Comment:Çözgü iplik % olarak fire 4"`
	UnitPrice4 float64 `json:"unit_price_4" gorm:"Comment:Çözgü iplik Birim fiyat 4"`
	TotalGram4 float64 `json:"total_gram_4"`
	Cost4      float64 `json:"cost_4" gorm:"Comment:Çözgü iplik Maliyet 4"`

	WarpYarn5  string  `json:"warp_yarn_5" gorm:"Comment:Çözgü iplik adı 5"`
	Gram5      float64 `json:"gram_5"`
	Wastage5   float64 `json:"wastage_5" gorm:"Comment:Çözgü iplik % olarak fire 5"`
	UnitPrice5 float64 `json:"unit_price_5" gorm:"Comment:Çözgü iplik Birim fiyat5"`
	TotalGram5 float64 `json:"total_gram_5"`
	Cost5      float64 `json:"cost_5" gorm:"Comment:Çözgü iplik Maliyet5"`

	WeftYarn1   string  `json:"weft_yarn_1" gorm:"Comment:Atkı iplik adı 1"`
	Gram01      float64 `json:"gram_01" gorm:"Comment:Atkı iplik gram 1"`
	Wastage01   float64 `json:"wastage_01" gorm:"Comment:Atkı iplik yüzde olarak fire 1"`
	UnitPrice01 float64 `json:"unit_price_01" gorm:"Comment:Atkı iplik birim fiyatı 1"`
	TotalGram01 float64 `json:"total_gram_01" gorm:"Comment:Atkı iplik toplam gram 1"`
	Cost01      float64 `json:"cost_01" gorm:"Comment:Atkı iplik maliyet 1"`

	WeftYarn2   string  `json:"weft_yarn_2" gorm:"Comment:Atkı iplik adı 2"`
	Gram02      float64 `json:"gram_02" gorm:"Comment:Atkı iplik gram 2"`
	Wastage02   float64 `json:"wastage_02" gorm:"Comment:Atkı iplik yüzde olarak fire 2"`
	UnitPrice02 float64 `json:"unit_price_02" gorm:"Comment:Atkı iplik birim fiyatı 2"`
	TotalGram02 float64 `json:"total_gram_02" gorm:"Comment:Atkı iplik toplam gram 2"`
	Cost02      float64 `json:"cost_02" gorm:"Comment:Atkı iplik maliyet 2"`

	WeftYarn3   string  `json:"weft_yarn_3" gorm:"Comment:Atkı iplik adı 3"`
	Gram03      float64 `json:"gram_03" gorm:"Comment:Atkı iplik gram 3"`
	Wastage03   float64 `json:"wastage_03" gorm:"Comment:Atkı iplik yüzde olarak fire 3"`
	UnitPrice03 float64 `json:"unit_price_03" gorm:"Comment:Atkı iplik birim fiyatı 3"`
	TotalGram03 float64 `json:"total_gram_03" gorm:"Comment:Atkı iplik toplam gram 3"`
	Cost03      float64 `json:"cost_03" gorm:"Comment:Atkı iplik maliyet 3"`

	WeftYarn4   string  `json:"weft_yarn_4" gorm:"Comment:Atkı iplik adı 4"`
	Gram04      float64 `json:"gram_04" gorm:"Comment:Atkı iplik gram 4"`
	Wastage04   float64 `json:"wastage_04" gorm:"Comment:Atkı iplik yüzde olarak fire 4"`
	UnitPrice04 float64 `json:"unit_price_04" gorm:"Comment:Atkı iplik birim fiyatı 4"`
	TotalGram04 float64 `json:"total_gram_04" gorm:"Comment:Atkı iplik toplam gram 4"`
	Cost04      float64 `json:"cost_04" gorm:"Comment:Atkı iplik maliyet 4"`

	WeftYarn5   string  `json:"weft_yarn_5" gorm:"Comment:Atkı iplik adı 5"`
	Gram05      float64 `json:"gram_05" gorm:"Comment:Atkı iplik gram 5"`
	Wastage05   float64 `json:"wastage_05" gorm:"Comment:Atkı iplik yüzde olarak fire 5"`
	UnitPrice05 float64 `json:"unit_price_05" gorm:"Comment:Atkı iplik birim fiyatı 5"`
	TotalGram05 float64 `json:"total_gram_05" gorm:"Comment:Atkı iplik toplam gram 5"`
	Cost05      float64 `json:"cost_05" gorm:"Comment:Atkı iplik maliyet 5"`

	WeavingWeftOfQuantity uint32  `json:"weaving_weft_of_quantity" gorm:"Comment:Dokuma atkı sayısı"`
	UnitPriceOfWeft       float64 `json:"unit_price_of_weft" gorm:"Comment:Atkı birim fiyatı"`
	TotalWeftPrice        float64 `json:"total_weft_price" gorm:"Comment:Atkı toplam fiyat"`
	TotalWeftPriceUSD     float64 `json:"total_weft_price_usd" gorm:"Comment:Atkı toplam fiyat USD"`
	PaintDressingType     string  `json:"paint_dressing_type" gorm:"Comment:Boya apre tipi"`
}
