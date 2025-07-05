package functions

import "math"

// This function calculates the standard deviation by taking the square root of the variance
func StandardDeviation(variance float64) float64 {
	return math.Sqrt(variance)
}
