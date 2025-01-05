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
	doneChanls := make([]chan bool, len(taxRates))
	errChanls := make([]chan error, len(taxRates))

	for i, taxRate := range taxRates {
		doneChanls[i] = make(chan bool)
		errChanls[i] = make(chan error)
		fm := filemanager.NewFileManager(inputFilePath, fmt.Sprintf("tax_included_prices_%v.json", taxRate*100))
		go prices.NewTaxIncludedPricesJob(fm, taxRate).Process(doneChanls[i], errChanls[i])
	}

	for i := range taxRates {
		select {
		case err := <-errChanls[i]:
			if err != nil {
				fmt.Println(err)
			}
		case <-doneChanls[i]:
			fmt.Printf("Tax included prices for tax rate %.1f%% has been successfully written to file\n", taxRates[i]*100)
		}
	}
}
