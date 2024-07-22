package main

import (
	"fmt"
	"os"
)

// build command: go build -o dpodcsg.exe main.go

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Datapod: Copy Skill Graphics")
		fmt.Println("1. runs through the file C:\\edward\\datapod2023\\datapod-for-vite-react-core\\currentImport\\skills.dpod.txt")
		fmt.Println("2. finds all lines ending with ##<idCode>")
		fmt.Println("3. copies <idCode>.png from C:\\WORK\\import to C:\\edward\\datapod2023\\tanguay-eu\\public\\images\\outline")
		fmt.Println("Usage: dpodcsg = get this instructions")
		fmt.Println("Usage: dpodcsg copy = copy the files")
		return
	} else {

		command := os.Args[1]

		if command != "copy" {
			fmt.Printf("ERROR: the only valid command is \"copy\"")
		} else {
			fmt.Printf("copying...")	
		}

	}
}
