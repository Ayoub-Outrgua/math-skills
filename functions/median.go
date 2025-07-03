package functions

import (
	"fmt"
	"sort"
	"strconv"
)

func Median(slice []string) float64 {
	intSlice := []int{}
	for _, v := range slice {
		nb, err := strconv.Atoi(string(v))
		if err != nil {
			fmt.Println("the file has non valid number:", err)
			return 0
		}
		intSlice = append(intSlice, nb)
	}
	sort.Ints(intSlice)
	if len(intSlice)%2 != 0 {
		return float64(intSlice[len(intSlice)/2])
	} else {
		var firstNumber float64 = float64(intSlice[(len(intSlice)/2)-1])
		var secendNumber float64 = float64(intSlice[len(intSlice)/2])
		return (firstNumber + secendNumber) / 2
	}
}
