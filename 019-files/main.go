package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

const dirName = "output"

func createTextFile(idCode, content string) {
	err := os.MkdirAll(dirName, 0755)
	pathAndFileName := filepath.Join(dirName, idCode+".txt")
	file, err := os.Create(pathAndFileName)
	checkError(err)
	_, err = io.WriteString(file, content)
	checkError(err)
	defer file.Close()
}

func readTextFile(idCode string) {
	pathAndFileName := filepath.Join(dirName, idCode+".txt")
	data, err := os.ReadFile(pathAndFileName)
	checkError(err)
	fmt.Printf(`
---%s---------
%s
--------------
`, pathAndFileName, data)
}

func main() {
	colors := []string{"red", "blue", "yellow"}
	for _, color := range colors {
		createTextFile(color, fmt.Sprintf("This is content about the color %s.", color))
	}
}
