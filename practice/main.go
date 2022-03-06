package main

import (
	"fmt"
)

func main() {
	fmt.Println("Conditional logic")

	theAnswer := 46
	var result string

	if theAnswer < 0 {
		result = "Less than zero"
	} else if theAnswer == 1 && (theAnswer == 0 || theAnswer > 0) {
		result = "Equal to zero"
	} else {
		result = "Greater than zero"
	}
	fmt.Println(result)

	// initialize variable in if statement
	if x := 42; x < 0 {
		result = "Less than zero"
	} else if x == 0 {
		result = "Equal to zero"
	} else {
		result = "Greater than zero"
	}
	fmt.Println(result)

}
