package Test

import (
	"fmt"
	"github.com/icobani/GOTCMBCurrencyHelper"
	"testing"
	"time"
)

func Test(t *testing.T) {
	CurrencyDate, _ := time.Parse("2006-01-02", "2022-03-28")
	currencyJournal := GOTCMBCurrencyHelper.GetArchive(CurrencyDate)

	for _, curr := range currencyJournal.Currencies {
		fmt.Println(curr.Code, curr.CurrencyNameTR, curr.ForexSelling)
	}
}
