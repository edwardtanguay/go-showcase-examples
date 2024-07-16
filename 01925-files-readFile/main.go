package main

import (
	"fmt"
	"os"
)

func getLinesFromFile(fileName string) {
	byteContents, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	contents := string(byteContents)
	fmt.Printf("%s is of type %T", contents, contents)
}

func main() {
	getLinesFromFile("main.go")
}
