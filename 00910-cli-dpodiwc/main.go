package main

import (
	"fmt"
	"os"
	"strings"
)

// build command: go build -o dpodiwc.exe main.go

const validFileEnding = ".dpod.txt"

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Datapod: Increment WhenCreated")
		fmt.Println("Runs through a dpod.txt file and resets all whenCreated fields to current date/time, each item incremented by one second")
		fmt.Println("Usage: dpodiwc <fileName>")
		fmt.Println("Example: dpodiwc skills.dpod.txt")
		fmt.Println("Assumes file is in: C:\\edward\\datapod2023\\datapod-for-vite-react-core\\currentImport")
		return
	} else {

		fileName := os.Args[1]

		if !strings.HasSuffix(fileName, validFileEnding) {
			fmt.Printf("ERROR: File name must end with \"%s\"\n", validFileEnding)
		} else {
			fmt.Printf("processing %s\n", fileName)
		}

	}
}
