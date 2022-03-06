package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//Refactored function for input
func inputValue() float64 {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Value 1: ")
	input1, _ := reader.ReadString('\n')
	float1, err := strconv.ParseFloat(strings.TrimSpace(input1), 64)
	CheckError(err)

	return float1
}
