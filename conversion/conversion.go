package conversion

import (
	"fmt"
	"strconv"
)

func StringsToFloat(strings []string) ([]float64, error) {
	var floats []float64

	for _, stringVal := range strings {

		floatVal, err := strconv.ParseFloat(stringVal, 64)

		if err != nil {
			fmt.Println("Error parsing price:", err)
			return nil, err
		}

		floats = append(floats, floatVal)
	}
	return floats, nil
}
