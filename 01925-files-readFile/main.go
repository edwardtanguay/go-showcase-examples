package main

import (
	"fmt"
	"os"
	"strings"
)

func getLinesFromFile(fileName string) {
	byteContents, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	contents := string(byteContents)
	lines := strings.Split(contents, "\n")
	fmt.Println(strings.Join(lines, "|"))
}

func main() {
	getLinesFromFile("main.go")
}
