package main

import (
	"fmt"
)

func main() {
	float1 := inputValue()

	float2 := inputValue()

	operator := Operation()

	var result float64

	switch operator {
	case "+":
		result = AddValues(float1, float2)
	case "-":
		result = SubtractValues(float1, float2)
	case "*":
		result = MultiplyValues(float1, float2)
	case "/":
		result = DivideValues(float1, float2)
	default:
		panic("Invalid Operation")
	}

	fmt.Printf("The result is %v\n\n", result)

}
