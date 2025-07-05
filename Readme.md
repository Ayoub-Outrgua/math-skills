# Statistical Calculator

## Project Overview

This Go program reads a file containing a list of numbers (one per line) and calculates key statistical measures:

- Average  
- Median  
- Variance  
- Standard Deviation  

The results are then printed as rounded integers.

## Features

- Reads numeric data from a `.txt` file passed as a command-line argument.
- Validates input for correct file extension and argument count.
- Handles invalid numbers and empty files gracefully.
- Implements core statistical calculations using Go.
- Outputs results formatted for easy reading.

## Usage

To run the program, use the following command:

```bash
go run . data.txt
