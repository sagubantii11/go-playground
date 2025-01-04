package conversion

import (
	"fmt"
	"strconv"
)

func StringsToFloat64s(strings []string) ([]float64, error) {
	float64s := make([]float64, len(strings))

	for index, str := range strings {
		float64Val, err := strconv.ParseFloat(str, 64)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		float64s[index] = float64Val
	}

	return float64s, nil
}
