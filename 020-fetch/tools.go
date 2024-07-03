package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

const outputDirName = "output"

func separator() {
	fmt.Println("---")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func createTextFile(fileName, content string) {
	err := os.MkdirAll(outputDirName, 0755)
	pathAndFileName := filepath.Join(outputDirName, fileName)
	file, err := os.Create(pathAndFileName)
	checkError(err)
	_, err = io.WriteString(file, content)
	checkError(err)
	defer file.Close()
}
