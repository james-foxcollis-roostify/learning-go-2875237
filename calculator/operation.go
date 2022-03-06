package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

//Refactored to separate operator into function
func Operation() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Select an Operation (+ - * /)")
	input, _ := reader.ReadString('\n')
	operator := strings.TrimSpace(input)

	return operator
}
