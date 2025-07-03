package main

import (
	"fmt"
	"mathskills/functions"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Error: Args not exist!")
		return
	}
	fileName := os.Args[1]
	fileData, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	// slice := strings.Split(string(fileData), "\n")
	slice := strings.Fields(string(fileData))
	if len(slice) == 0 {
		fmt.Println("the file is empty!")
		return
	}
	average := functions.Average(slice)
	fmt.Println(slice)
	fmt.Println(average)
}
