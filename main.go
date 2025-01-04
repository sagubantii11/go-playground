package main

import (
	"fmt"

	"github.com/sagubantii11/go-playground/filemanager"
	"github.com/sagubantii11/go-playground/prices"
)

func main() {
	var inputFilePath string
	fmt.Println("Enter the input file path: ")
	fmt.Scan(&inputFilePath)
	prices.GenerateRandomPrices(inputFilePath) // Generate random prices and store in a file named `prices.txt`
	taxRates := []float64{0, 0.05, 0.1, 0.12, 0.15, 0.2, 0.3}

	for _, taxRate := range taxRates {
		fm := filemanager.NewFileManager(inputFilePath, fmt.Sprintf("tax_included_prices_%v.json", taxRate*100))
		prices.NewTaxIncludedPricesJob(fm, taxRate).Process()
	}

}
