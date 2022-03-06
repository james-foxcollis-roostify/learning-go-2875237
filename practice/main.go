package main

import (
	"fmt"
)

const aConst = 64

func main() {

	//explicit type
	var aString string = "This is Go!"
	fmt.Println(aString)
	fmt.Printf("The variables type is %T\n", aString)

	var anInt int = 42
	fmt.Println(anInt)

	var defaultInt int
	fmt.Println(defaultInt)

	// Implicit type
	var anotherString = "Another String"
	fmt.Println(anotherString)
	fmt.Printf("Variable type is %T\n", anotherString)

	// Alternate var assignment, only works inside functions

	thirdString := "Third String"
	fmt.Println(thirdString)
	fmt.Printf("Type %T\n", thirdString)

	fmt.Println(aConst)

}
