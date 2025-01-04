package prices

import (
	"fmt"

	"github.com/sagubantii11/go-playground/conversion"
	"github.com/sagubantii11/go-playground/filemanager"
)

type TaxIncludedPricesJob struct {
	IOFManager        filemanager.FileManager
	TaxRate           float64
	InputPrices       []float64
	TaxIncludedPrices map[string]float64
}

func NewTaxIncludedPricesJob(fm filemanager.FileManager, taxRate float64) *TaxIncludedPricesJob {
	return &TaxIncludedPricesJob{
		IOFManager: fm,
		TaxRate:    taxRate,
	}
}

func (job *TaxIncludedPricesJob) LoadPrices() {
	lines, err := job.IOFManager.ReadLinesFromFile()

	if err != nil {
		fmt.Println(err)
		return
	}

	prices, err := conversion.StringsToFloat64s(lines)

	if err != nil {
		fmt.Println(err)
		return
	}

	job.InputPrices = prices
}

func (job *TaxIncludedPricesJob) Process() {
	job.LoadPrices() // Load prices from file

	result := make(map[string]float64)

	for _, price := range job.InputPrices {
		result[fmt.Sprintf("%.2f", price)] = price * (1 + job.TaxRate)
	}

	job.TaxIncludedPrices = result

	err := job.IOFManager.WriteJsonToFile(job)
	if err != nil {
		fmt.Println(err)
		return
	}
}
