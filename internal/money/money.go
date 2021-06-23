package money

const CurrencyRuble string = "RUB"
const CurrencyDollar string = "USD"

type Money struct {
	Amount   float64
	Currency string
}
