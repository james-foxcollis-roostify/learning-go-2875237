package main

import "math"

//Divide values
func DivideValues(value1 float64, value2 float64) float64 {
	divided := value1 / value2
	return math.Round(divided*100) / 100
}
