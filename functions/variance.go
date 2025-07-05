package functions

import (
	"fmt"
	"strconv"
)

// This function is used to calculate the variance from a slice of string numbers
func Variance(slice []string) float64 {
	var sum float64
	for _, v := range slice {
		nb, err := strconv.Atoi(string(v))
		if err != nil {
			fmt.Println("the file has invalid number:", err)
			return 0
		}
		number := (float64(nb) - Average(slice))
		sum += (number * number)
	}
	return sum / float64(len(slice))
}
