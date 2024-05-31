package main

import (
	"fmt"
)

func main() {
	var amount float64
	var toCurrency string

	fmt.Print("Enter the amount in USD: ")
	fmt.Scan(&amount)

	fmt.Println("List of available currencies:")
	for currencyCode, currencyName := range Names {
		fmt.Printf("%s - %s\n", currencyCode, currencyName)
	}

	fmt.Print("Enter the currency code to convert to: ")
	fmt.Scan(&toCurrency)

	result := ConvertUSDToOther(amount, toCurrency)
	if result != -1 {
		fmt.Printf("USD %.2f --> %s %.2f\n", amount, toCurrency, result)
	} else {
		fmt.Println("Invalid currency provided.")
	}
}
