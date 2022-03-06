package main

import "math"

//Multiply values
func MultiplyValues(value1 float64, value2 float64) float64 {
	multi := value1 * value2
	return math.Round(multi*100) / 100
}
