package main

import (
	"fmt"
	"os"
	"strings"
)

func getLinesFromFile(fileName string) []string {
	byteContents, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	contents := string(byteContents)
	lines := strings.Split(contents, "\n")
	return lines
}

func main() {
	lines := getLinesFromFile("main.go")
	for i,line := range lines {
		fmt.Printf("%d. %s\n", i+1, line)
	}
}
