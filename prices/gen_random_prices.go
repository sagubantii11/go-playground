// Get 5 random float64 numbers in a range (0, 213) using Go Language and store in a file named `prices.txt`
package prices

import (
	"fmt"
	"math/rand"
	"os"
	// "time"
)

func GenerateRandomPrices(filePath string) {
	// rand.Seed(time.Now().UnixNano())
	file, err := os.Create(filePath)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	for i := 0; i < 5; i++ {
		price := rand.Float64() * 213
		_, err := file.WriteString(fmt.Sprintf("%.2f\n", price))
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}
