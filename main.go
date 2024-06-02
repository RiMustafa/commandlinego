package main

import (
	"fmt"
)

func main() {
	var amount float64
	var fromCurrency, toCurrency string

	fmt.Print("Enter the amount: ")
	fmt.Scan(&amount)

	fmt.Println("List of available currencies:")
	for currencyCode, currencyName := range Names {
		fmt.Printf("%s - %s\n", currencyCode, currencyName)
	}

	fmt.Print("Enter the currency code to convert from: ")
	fmt.Scan(&fromCurrency)

	fmt.Print("Enter the currency code to convert to: ")
	fmt.Scan(&toCurrency)

	result := ConvertCurrency(amount, fromCurrency, toCurrency)
	if result != -1 {
		fmt.Printf("%s %.2f --> %s %.2f\n", fromCurrency, amount, toCurrency, result)
	} else {
		fmt.Println("Invalid currency provided.")
	}
}
