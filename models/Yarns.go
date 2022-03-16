package models

type Yarns struct {
	ID        uint64  `json:"id" gorm:"primaryKey"`
	Code      string  `json:"code" gorm:"Comment:İplik kodu"`
	Name      string  `json:"name" gorm:"Comment:İplik ismi"`
	UnitPrice float64 `json:"unit_price" gorm:"Comment:Birim fiyat"`
}
