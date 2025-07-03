package functions

import (
	"fmt"
	"sort"
	"strconv"
)

func Median(slice []string) int {
	intSlice := []int{}
	for _, v := range slice {
		nb, err := strconv.Atoi(string(v))
		if err != nil {
			fmt.Println("the file has non valid number")
			return 0
		}
		intSlice = append(intSlice, nb)
	}
	sort.Ints(intSlice)
	if len(intSlice)%2 != 0 {
		return intSlice[len(intSlice)/2]
	} else {
		
	}
}
