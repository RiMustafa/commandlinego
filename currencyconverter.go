package main

var ConversionRates = map[string]float64{
	"SEK": 11.10,
	"AUD": 1.50,
	"BRL": 4.90,
	"CAD": 1.36,
	"CHF": 0.90,
	"CNY": 7.12,
	"CZK": 22.35,
	"DKK": 6.89,
	"EUR": 0.94,
	"GBP": 0.79,
	"HKD": 7.83,
	"HUF": 361.95,
	"IDR": 15000.00,
	"INR": 82.00,
	"ISK": 136.00,
	"JPY": 139.60,
	"KRW": 1303.50,
	"KWD": 0.31,
	"MAD": 10.50,
	"MXN": 17.66,
	"MYR": 4.60,
	"NOK": 10.51,
	"NZD": 1.62,
	"PLN": 4.31,
	"RUB": 77.21,
	"SAR": 3.75,
	"SGD": 1.34,
	"THB": 35.11,
	"TRY": 23.60,
	"USD": 1.0,
	"ZAR": 18.58,
}

func ConvertCurrency(amount float64, fromCurrency string, toCurrency string) float64 {
	fromRate, fromFound := ConversionRates[fromCurrency]
	toRate, toFound := ConversionRates[toCurrency]

	if fromFound && toFound {
		usdAmount := amount / fromRate
		return usdAmount * toRate
	}
	return -1
}
