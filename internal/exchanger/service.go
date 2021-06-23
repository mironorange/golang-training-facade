package exchanger

import (
	"errors"
	currency_quotes "github.com/mironorange/golang-training-facade/internal/currency-quotes"
	"github.com/mironorange/golang-training-facade/internal/money"
)

// Интерфейс конвертера денег
type CurrencyExchanger interface {
	Exchange(m money.Money, outgoingCurrency string) (money.Money, error)
}

type currencyExchanger struct {
	quotes currency_quotes.CurrencyQuotesProvider
}

// Преобразует деньги, переданные на вход в формат денег ожидаемый на выход
func (e *currencyExchanger) Exchange(m money.Money, outgoingCurrency string) (money.Money, error) {
	r := currency_quotes.NewCurrencyQuotesRepository()
	q, error := r.Get(m.Currency, outgoingCurrency)
	if error != nil {
		return money.Money{}, errors.New("Unknown operation with currency")
	}
	return money.Money{Amount: q * m.Amount, Currency: outgoingCurrency}, nil
}

// Создает новый конвертор валют
func NewCurrencyExchanger(quotes currency_quotes.CurrencyQuotesProvider) CurrencyExchanger {
	return &currencyExchanger{
		quotes: quotes,
	}
}
