package functions

import (
	"fmt"
	"strconv"
)

func Average(slice []string) float64 {
	var sum float64 = 0
	for _, v := range slice {
		nb, err := strconv.Atoi(string(v))
		if err != nil {
			fmt.Println("the file has non valid number")
			return 0
		}
		sum += float64(nb)
	}
	return sum / float64(len(slice))
}
