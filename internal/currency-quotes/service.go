package currency_quotes

import (
	"errors"
	"strings"
)

// Интерфейс провайдера с данными курса валют, необходимыми для конвертации
type CurrencyQuotesProvider interface {
	Get(incoming string, outgoing string) (float64, error)
}

// Структура провайдера курса валют, необходимыми для конвертации
type currencyQuotesRepository struct {
	quotes map[string]float64
}

// Возвращает данные курса валют для валют, переданных на вход
func (r *currencyQuotesRepository) Get (incoming string, outgoing string) (float64, error) {
	if incoming == outgoing {
		return 1., nil
	}
	currency := strings.Join([]string{incoming, outgoing}, "/")
	if value, ok := r.quotes[currency]; ok {
		return value, nil
	} else {
		return 0., errors.New("Unknown currency quotes")
	}
}

// Создает конкретную реализацию провайдера курса валют, с заранее подготовленными для конвертации данными
func NewCurrencyQuotesRepository() CurrencyQuotesProvider {
	return &currencyQuotesRepository{
		quotes: map[string]float64{
			"USD/RUB": 73.,
			"RUB/USD": 0.014,
		},
	}
}
