package main

import (
	"fmt"
)

func main() {
	anInt := 42
	var p *int = &anInt
	fmt.Println("Value of p:", *p)

	value1 := 42.3
	pointer1 := &value1
	fmt.Println("Value of pointer1:", *pointer1)
	fmt.Printf("type of pointer1 %T\n", pointer1)

	*pointer1 = *pointer1 / 31
	fmt.Println("Value of pointer1:", *pointer1)
	fmt.Println("Value of pointer1:", value1)

}
