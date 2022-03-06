package main

import (
	"fmt"
)

func main() {
	colors := []string{"Red", "Green", "Blue"}
	fmt.Println(colors)

	for i := 0; i < len(colors); i++ {
		fmt.Println(colors[i])
	}

	for i := range colors {
		fmt.Println(colors[i])
	}

	for _, color := range colors {
		fmt.Println(color)
	}

	value := 1

	for value < 10 {
		fmt.Println("Value:", value)
		value++
	}

	sum := 1

	for sum < 1000 {
		sum += sum
		fmt.Println("Sum: ", sum)
		// go to ends for loop and prints label
		if sum > 200 {
			goto theEnd
		}
	}
theEnd:
	fmt.Println("End of program")

	var out int
	for j := 0; j < 20; j++ {
		out = j*j + out
		if out > 12 {
			goto finish
		}
	}
finish:
	fmt.Println(out)

	doSomething()

	addedValues := addValues(1, 2)
	fmt.Println(addedValues)

	multiSum := addAllValues(1, 2, 3, 4)
	fmt.Println(multiSum)

	firstVal, secondVal := addAllValuesAgain(1, 2, 3, 4, 5)
	fmt.Println("Sum of Firstval:", firstVal)
	fmt.Println("Count of Array:", secondVal)
}

func doSomething() {
	fmt.Println("Doing Something")
}

func addValues(value1 int, value2 int) int {
	return value1 + value2
}

func addAllValues(values ...int) int {
	total := 0
	for _, v := range values {
		total += v
	}
	return total
}

func addAllValuesAgain(values ...int) (int, int) {
	total := 0
	for _, v := range values {
		total += v
	}
	return total, len(values)
}
