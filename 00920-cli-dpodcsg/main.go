package main

import (
	"fmt"
	"os"

	"github.com/fatih/color"
)

// build command: go build -o dpodcsg.exe main.go

const toScanPathAndFileName = "C:\\edward\\datapod2023\\datapod-for-vite-react-core\\currentImport\\skills.dpod.txt"
const graphicsFromPath = "C:\\WORK\\import"
const graphicsToPath = "C:\\edward\\datapod2023\\tanguay-eu\\public\\images\\outline"

func main() {
	blue := color.New(color.FgBlue).SprintFunc()
	green := color.New(color.FgGreen).SprintFunc()
	yellow := color.New(color.FgYellow).SprintFunc()
	red := color.New(color.FgRed).SprintFunc()

	if len(os.Args) < 2 {
		fmt.Println("Datapod: Copy Skill Graphics")
		fmt.Printf("1. runs through the file: %s\n", yellow(toScanPathAndFileName))
		fmt.Println("2. finds all lines ending with ##<idCode>")
		fmt.Printf("3. copies <idCode>.png from %s to %s\n", blue(graphicsFromPath), green(graphicsToPath))
		fmt.Printf("Usage: %s = get this instructions\n", red("dpodcsg"))
		fmt.Printf("Usage: %s = copy the files\n", red("dpodcsg copy"))
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
