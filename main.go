package main

import (
	"fmt"
	"math"
	"os"
	"strings"

	"mathskills/functions"
)

func main() {
	// Ensure exactly one argument is passed
	if len(os.Args) != 2 {
		fmt.Println("Error: Args not exist!")
		return
	}
	fileName := os.Args[1]

	// Protect the main function by ensuring the argument has a ".txt"
	if !strings.HasSuffix(fileName, ".txt") {
		fmt.Println("Error: the file is not exist!")
		return
	}

	// Read the content of the file
	fileData, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	// Split the file content into fields separated by whitespace
	slice := strings.Fields(string(fileData))

	// Check if the slice is empty
	if len(slice) == 0 {
		fmt.Println("the file is empty!")
		return
	}

	// Call my functions and round the results
	average := math.Round(functions.Average(slice))
	median := math.Round(functions.Median(slice))
	variance := math.Round(functions.Variance(slice))
	standardDeviation := math.Round(functions.StandardDeviation(variance))

	// Print The results
	fmt.Printf("Average : %0.f\n", average)
	fmt.Printf("Median : %0.f\n", median)
	fmt.Printf("Variance : %0.f\n", variance)
	fmt.Printf("Standard Deviation : %0.f\n", standardDeviation)
}
