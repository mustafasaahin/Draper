package models

type DressingPriceList struct {
	ID          uint64  `json:"id" gorm:"primaryKey"`
	Code        string  `json:"code" gorm:"Comment:İşlem kodu"`
	Process     string  `json:"process" gorm:"Comment:işlem"`
	ProcessType string  `json:"process_type" gorm:"Comment:İşlem tipi"`
	Price       float64 `json:"price" gorm:"Comment:İşlem fiyatı"`
}
