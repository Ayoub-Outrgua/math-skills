package functions

import (
	"fmt"
	"sort"
	"strconv"
)

func Median(slice []string) float64 {
	valuesSlice := []float64{}
	for _, v := range slice {
		nb, err := strconv.Atoi(string(v))
		if err != nil {
			fmt.Println("the file has non valid number:", err)
			return 0
		}
		valuesSlice = append(valuesSlice, float64(nb))
	}
	sort.Float64s(valuesSlice)
	if len(valuesSlice)%2 != 0 {
		return valuesSlice[len(valuesSlice)/2]
	} else {
		var firstNumber float64 = valuesSlice[(len(valuesSlice)/2)-1]
		var secendNumber float64 = valuesSlice[len(valuesSlice)/2]
		return (firstNumber + secendNumber) / 2
	}
}
