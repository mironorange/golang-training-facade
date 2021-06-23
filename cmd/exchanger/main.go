package main

import (
	"fmt"
	quotes "github.com/mironorange/golang-training-facade/internal/currency-quotes"
	"github.com/mironorange/golang-training-facade/internal/exchanger"
	"github.com/mironorange/golang-training-facade/internal/money"
)

func main() {
	m := money.Money{Amount: 3600, Currency: money.CurrencyRuble}
	fmt.Printf("Give me %v rubles in %v\n", m.Amount, m.Currency)

	// Преобразуем валюты, переданные на вход, в ожидаемый для клиента формат
	e := exchanger.NewCurrencyExchanger(quotes.NewCurrencyQuotesRepository())
	em, error := e.Exchange(m, money.CurrencyDollar)
	if error != nil {
		fmt.Println(error.Error())
	} else {
		fmt.Printf("Take %v %v\n", em.Amount, em.Currency)
	}
}
