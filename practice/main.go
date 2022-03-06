package main

import (
	"fmt"
	"math"
)

func main() {

	i1, i2, i3 := 12, 45, 68
	intSum := i1 + i2 + i3
	fmt.Println("Integer Sum is: ", intSum)

	fmt.Println("Math")

	f1, f2, f3 := 23.5, 65.1, 76.3
	floatSum := f1 + f2 + f3
	fmt.Println("Float Sum is:", floatSum)
	floatSum = math.Round(floatSum*100) / 100
	fmt.Println("The sum is now:", floatSum)

	circleRadius := 15.5
	circum := circleRadius * 2 * math.Pi
	fmt.Printf("Circumference: %.2f\n", circum)

}
