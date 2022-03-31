package InitData

import (
	"fmt"
	"github.com/mustafasaahin/Draper/config"
	"github.com/mustafasaahin/Draper/models"
	"github.com/xuri/excelize/v2"
	"strconv"
)

func InitProduct() {
	f, err := excelize.OpenFile("Draper.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func() {
		// Close the spreadsheet.
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	// Get all the rows in the Sheet1.
	rows, err := f.GetRows("Sayfa1")
	if err != nil {
		fmt.Println(err)
		return
	}
	config.DB.Where("id<>0").
		Delete(&models.Product{})
	for _, row := range rows {
		var product = models.Product{}
		for colNumber, cell := range row {
			switch colNumber {
			case 0: // Product Code
				{
					product.Code = cell
				}
			case 1: // Product Name
				{
					product.Name = cell
				}
			case 2: // Çözgü İplik
				{
					product.WarpYarn1 = cell
				}
			case 3: // Gram1
				{
					product.Gram1, _ = strconv.ParseFloat(cell, 64)
				}
			case 4:
				{
					product.Wastage1, _ = strconv.ParseFloat(cell, 64)
				}
			case 5:
				{
					product.UnitPrice1, _ = strconv.ParseFloat(cell, 64)
				}
			case 6:
				{
					product.TotalGram1, _ = strconv.ParseFloat(cell, 64)
				}
			case 7:
				{
					product.Cost1, _ = strconv.ParseFloat(cell, 64)
				}
			case 8:
				{
					product.WarpYarn2 = cell
				}
			case 9:
				{
					product.Gram2, _ = strconv.ParseFloat(cell, 64)
				}
			case 10:
				{
					product.Wastage2, _ = strconv.ParseFloat(cell, 64)
				}
			case 11:
				{
					product.UnitPrice2, _ = strconv.ParseFloat(cell, 64)
				}
			case 12:
				{
					product.TotalGram2, _ = strconv.ParseFloat(cell, 64)
				}
			case 13:
				{
					product.Cost2, _ = strconv.ParseFloat(cell, 64)
				}
			case 14:
				{
					product.WarpYarn3 = cell
				}
			case 15:
				{
					product.Gram3, _ = strconv.ParseFloat(cell, 64)
				}
			case 16:
				{
					product.Wastage3, _ = strconv.ParseFloat(cell, 64)
				}
			case 17:
				{
					product.UnitPrice3, _ = strconv.ParseFloat(cell, 64)
				}
			case 18:
				{
					product.TotalGram3, _ = strconv.ParseFloat(cell, 64)
				}
			case 19:
				{
					product.Cost3, _ = strconv.ParseFloat(cell, 64)
				}
			case 20:
				{
					product.WarpYarn4 = cell
				}
			case 21:
				{
					product.Gram4, _ = strconv.ParseFloat(cell, 64)
				}
			case 22:
				{
					product.Wastage4, _ = strconv.ParseFloat(cell, 64)
				}
			case 23:
				{
					product.TotalGram4, _ = strconv.ParseFloat(cell, 64)
				}
			case 24:
				{
					product.UnitPrice4, _ = strconv.ParseFloat(cell, 64)
				}
			case 25:
				{
					product.Cost4, _ = strconv.ParseFloat(cell, 64)
				}
			case 26:
				{
					product.Gram5, _ = strconv.ParseFloat(cell, 64)
				}
			case 27:
				{
					product.Wastage5, _ = strconv.ParseFloat(cell, 64)
				}
			case 28:
				{
					product.TotalGram5, _ = strconv.ParseFloat(cell, 64)
				}
			case 29:
				{
					product.UnitPrice5, _ = strconv.ParseFloat(cell, 64)
				}
			case 30:
				{
					product.Cost05, _ = strconv.ParseFloat(cell, 64)
				}
			case 31:
				{
					product.WeftYarn1 = cell
				}
			case 32:
				{
					product.Gram01, _ = strconv.ParseFloat(cell, 64)
				}
			case 33:
				{
					product.Wastage01, _ = strconv.ParseFloat(cell, 64)
				}
			case 34:
				{
					product.TotalGram01, _ = strconv.ParseFloat(cell, 64)
				}
			case 35:
				{
					product.UnitPrice01, _ = strconv.ParseFloat(cell, 64)
				}
			case 37:
				{
					product.Cost01, _ = strconv.ParseFloat(cell, 64)
				}
			case 38:
				{
					product.WeftYarn2 = cell
				}
			case 39:
				{
					product.Gram02, _ = strconv.ParseFloat(cell, 64)
				}
			case 40:
				{
					product.Wastage02, _ = strconv.ParseFloat(cell, 64)
				}
			case 41:
				{
					product.TotalGram02, _ = strconv.ParseFloat(cell, 64)
				}
			case 42:
				{
					product.UnitPrice02, _ = strconv.ParseFloat(cell, 64)
				}
			case 43:
				{
					product.Cost02, _ = strconv.ParseFloat(cell, 64)
				}
			case 44:
				{
					product.WeftYarn3 = cell
				}
			case 45:
				{
					product.Gram03, _ = strconv.ParseFloat(cell, 64)
				}
			case 46:
				{
					product.Wastage03, _ = strconv.ParseFloat(cell, 64)
				}
			case 47:
				{
					product.TotalGram03, _ = strconv.ParseFloat(cell, 64)
				}
			case 48:
				{
					product.UnitPrice03, _ = strconv.ParseFloat(cell, 64)
				}
			case 49:
				{
					product.Cost03, _ = strconv.ParseFloat(cell, 64)
				}
			case 50:
				{
					product.WeftYarn4 = cell
				}
			case 51:
				{
					product.Gram04, _ = strconv.ParseFloat(cell, 64)
				}
			case 52:
				{
					product.Wastage04, _ = strconv.ParseFloat(cell, 64)
				}
			case 53:
				{
					product.TotalGram04, _ = strconv.ParseFloat(cell, 64)
				}
			case 54:
				{
					product.UnitPrice04, _ = strconv.ParseFloat(cell, 64)
				}
			case 55:
				{
					product.Cost04, _ = strconv.ParseFloat(cell, 64)
				}
			case 56:
				{
					product.WeftYarn5 = cell
				}
			case 57:
				{
					product.Gram05, _ = strconv.ParseFloat(cell, 64)
				}
			case 58:
				{
					product.Wastage05, _ = strconv.ParseFloat(cell, 64)
				}
			case 59:
				{
					product.TotalGram05, _ = strconv.ParseFloat(cell, 64)
				}
			case 60:
				{
					product.UnitPrice05, _ = strconv.ParseFloat(cell, 64)
				}
			case 61:
				{
					product.Cost05, _ = strconv.ParseFloat(cell, 64)
				}
			case 62:
				{
					product.WeavingWeftOfQuantity, _ = strconv.ParseFloat(cell, 64)
				}
			case 63:
				{
					product.UnitPriceOfWeft, _ = strconv.ParseFloat(cell, 64)
				}
			case 64:
				{
					product.DeverePrice, _ = strconv.ParseFloat(cell, 64)
				}
			case 65:
				{
					product.TotalWeftPrice, _ = strconv.ParseFloat(cell, 64)
				}
			case 66:
				{
					product.TotalWeftPriceUSD, _ = strconv.ParseFloat(cell, 64)
				}
			case 67:
				{
					product.PaintDressingType = cell
				}
			case 68:
				{
					product.Waterproof, _ = strconv.ParseBool(cell)
				}
			case 69:
				{
					product.YikamaApre, _ = strconv.ParseBool(cell)
				}
			case 70:
				{
					product.Onfikse, _ = strconv.ParseBool(cell)
				}
			case 71:
				{
					product.KuruApre, _ = strconv.ParseBool(cell)
				}
			case 72:
				{
					product.Inceltme, _ = strconv.ParseBool(cell)
				}
			case 73:
				{
					product.TuyDokme, _ = strconv.ParseBool(cell)
				}
			case 74:
				{
					product.Gaze, _ = strconv.ParseBool(cell)
				}
			case 75:
				{
					product.Kalender, _ = strconv.ParseBool(cell)
				}
			case 76:
				{
					product.Samfor, _ = strconv.ParseBool(cell)
				}
			case 77:
				{
					product.ApreGram, _ = strconv.ParseFloat(cell, 64)
				}
			case 78:
				{
					product.ApreUnitPrice, _ = strconv.ParseFloat(cell, 64)
				}
			case 79:
				{
					product.
						ApreTotalPrice, _ = strconv.ParseFloat(cell, 64)
				}
			case 80:
				{
					product.IntermediateCost, _ = strconv.ParseFloat(cell, 64)
				}
			case 81:
				{
					product.CekmeYuzde, _ = strconv.ParseFloat(cell, 64)
				}
			case 82:
				{
					product.CekmePrice, _ = strconv.ParseFloat(cell, 64)
				}
			case 83:
				{
					product.TotalCost, _ = strconv.ParseFloat(cell, 64)
				}
			case 84:
				{
					product.Photo = cell
				}
			case 85:
				{
					product.SalesPrices, _ = strconv.ParseFloat(cell, 64)
				}
			case 86:
				{
					product.AdditionalProcess = cell
				}
			}
		}
		config.DB.Create(&product)
	}
}
