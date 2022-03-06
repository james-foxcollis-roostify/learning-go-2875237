package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	// content := "Hello From Go!"
	// file, err := os.Create("./fromString.txt")
	// checkError(err)
	// length, err := io.WriteString(file, content)
	// checkError(err)
	// fmt.Printf("Wrote a file with %v characters\n", length)
	// defer file.Close()
	// defer readFile("./fromString.txt")
	defer os.Remove("./fromString.txt")
}

func readFile(fileName string) {
	data, err := ioutil.ReadFile(fileName)
	checkError(err)
	fmt.Println("Text read from file:", string(data))
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
