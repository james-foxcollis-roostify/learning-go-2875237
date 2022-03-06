package main

import "math"

//Subtract values
func SubtractValues(value1 float64, value2 float64) float64 {
	sub := value1 - value2
	return math.Round(sub*100) / 100
}
