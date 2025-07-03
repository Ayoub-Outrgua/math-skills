package main

import (
	"fmt"
	"math"
	"os"
	"strings"

	"mathskills/functions"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Error: Args not exist!")
		return
	}
	fileName := os.Args[1]
	if !strings.HasSuffix(fileName, ".txt") || len(fileName) < 5 {
		fmt.Println("Error: the file is not exist!")
		return
	}
	fileData, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	slice := strings.Split(string(fileData), "\n")
	// slice := strings.Fields(string(fileData))
	if len(slice) == 0 {
		fmt.Println("the file is empty!")
		return
	}
	average := math.Round(functions.Average(slice))
	median := math.Round(functions.Median(slice))
	variance := math.Round(functions.Variance(slice))
	standardDeviation := math.Round(functions.StandardDeviation(variance))

	fmt.Println("Average : ", average)
	fmt.Println("Median : ", median)
	fmt.Println("Variance : ", variance)
	fmt.Println("Standard Deviation : ", standardDeviation)
}
