package main

import (
	// "fmt"
	"io"
	"os"
)

func createTextFile(idCode, content string) {
	fileName := idCode + ".txt"
	file, err := os.Create(fileName)
	checkError(err)
	_, err = io.WriteString(file, content)
	checkError(err)
	defer file.Close()

}

func main() {
	colors := []string{"red", "blue", "yellow"}
	for _, color := range colors {
		createTextFile(color, "the content")
	}
}
