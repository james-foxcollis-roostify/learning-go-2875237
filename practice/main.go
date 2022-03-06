package main

import (
	"fmt"
	"sort"
)

func main() {
	var colors = []string{"Red", "Green", "Blue"}
	fmt.Println(colors)
	colors = append(colors, "Black")
	fmt.Println(colors)

	//Remove first item in slice
	colors = append(colors[1:len(colors)])
	fmt.Println(colors)
	//Remove list item in slice
	colors = append(colors[:len(colors)-1])
	fmt.Println(colors)

	numbers := make([]int, 5)
	numbers[0] = 134
	numbers[1] = 5
	numbers[2] = 33
	numbers[3] = 555
	numbers[4] = 1
	numbers = append(numbers, 35)
	fmt.Println(numbers)

	sort.Ints(numbers)
	fmt.Println(numbers)
}
