package models

type Setup struct {
	ID                  uint64  `json:"id" gorm:"primaryKey"`
	WeavingWarpPrice    float64 `json:"weaving_warp_price" gorm:"Comment:Dokuma atkı ücreti"`
	ExchangeUSD         float64 `json:"exchange_usd" gorm:"Comment:Usd kuru"`
	ExchangeEUR         float64 `json:"exchange_eur" gorm:"Comment:Eur kuru"`
	DressingExchaneType string  `json:"dressing_exchane_type" gorm:"Comment: Apre döviz tipi"`
}
