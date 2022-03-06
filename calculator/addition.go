package main

import "math"

//Add values
func AddValues(value1 float64, value2 float64) float64 {
	sum := value1 + value2
	return math.Round(sum*100) / 100
}
