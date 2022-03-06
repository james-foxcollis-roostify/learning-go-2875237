package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	// dow := rand.Intn(7) + 1
	// fmt.Println("Day", dow)

	var result string
	switch dow := rand.Intn(7) + 1; dow {
	case 1:
		result = "It's Sunday"
	case 2:
		result = "It's Monday"
	default:
		result = "It's some other day"
	}

	fmt.Println(result)

	//In go, once a case is met it breaks, to match js/java/c case, add fallthrough to execute the next case
	switch down := 1; down {
	case 1:
		result = "It's Sunday"
		fallthrough
	case 2:
		result = "It's Monday"
		fallthrough
	default:
		result = "It's some other day"
	}

	fmt.Println(result)
}
