package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
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
	}

	fmt.Printf("The result is %v\n\n", result)

}

//Add values
func AddValues(value1 float64, value2 float64) float64 {
	sum := value1 + value2
	return math.Round(sum*100) / 100
}

//Subtract values
func SubtractValues(value1 float64, value2 float64) float64 {
	sub := value1 - value2
	return math.Round(sub*100) / 100
}

//Multiply values
func MultiplyValues(value1 float64, value2 float64) float64 {
	multi := value1 * value2
	return math.Round(multi*100) / 100
}

//Divide values
func DivideValues(value1 float64, value2 float64) float64 {
	divided := value1 / value2
	return math.Round(divided*100) / 100
}

//Function for input
// func inputValue() (float64, float64, string) {
// 	reader := bufio.NewReader(os.Stdin)
// 	fmt.Print("Value 1: ")
// 	input1, _ := reader.ReadString('\n')
// 	float1, err := strconv.ParseFloat(strings.TrimSpace(input1), 64)
// 	if err != nil {
// 		fmt.Println(err)
// 		panic("Value 1 must be a number")
// 	}

// 	fmt.Print("Value 2: ")
// 	input2, _ := reader.ReadString('\n')
// 	float2, err := strconv.ParseFloat(strings.TrimSpace(input2), 64)
// 	if err != nil {
// 		fmt.Println(err)
// 		panic("Value 2 must be a number")
// 	}

// 	fmt.Print("Select an Operation (+ - * /)")
// 	input3, _ := reader.ReadString('\n')
// 	operator := strings.TrimSpace(input3)

// 	return float1, float2, operator
// }

//Refactored function for input
func inputValue() float64 {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Value 1: ")
	input1, _ := reader.ReadString('\n')
	float1, err := strconv.ParseFloat(strings.TrimSpace(input1), 64)
	if err != nil {
		fmt.Println(err)
		panic("Value 1 must be a number")
	}

	return float1
}

//Refactored to separate operator into function
func Operation() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Select an Operation (+ - * /)")
	input, _ := reader.ReadString('\n')
	operator := strings.TrimSpace(input)

	return operator
}
