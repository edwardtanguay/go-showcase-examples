package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func writeFile(dirName string, fileName string, content string) {

	filePath := filepath.Join(dirName, fileName)

	file, err := os.Create(filePath)

	if err != nil {
		fmt.Printf("Error creating file: %v\n", err)
		return
	}
	defer file.Close()

	_, err = file.WriteString(content)
	if err != nil {
		fmt.Printf("Error writing to file: %v\n", err)
		return
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Datapod: Increment WhenCreated")
		fmt.Println("Runs through a dpod.txt file and resets all whenCreated fields to current date/time, each item incremented by one second")
		fmt.Println("Usage: dpodiwc <fileName>")
		fmt.Println("Example: dpodiwc skills.dpod.txt")
		fmt.Println("Assumes file is in: C:\\edward\\datapod2023\\datapod-for-vite-react-core\\currentImport")
		return
	}

	var dirName = os.Args[1]

	err := os.MkdirAll(dirName, 0755)
	if err != nil {
		fmt.Printf("Error creating directory: %v\n", err)
		return
	}

	writeFile(dirName, "main.go", `
package main

import "fmt"

func main() {
	fmt.Println("nnn")
}
`)

	writeFile(dirName, "tools.go", `
package main

import (
	"fmt"
	"strings"
)

func separator(title ...string) {
	width := 50
	preWidth := 3
	character := "="
	if len(title) == 0 || title[0] == "" {
		fmt.Println(strings.Repeat(character, width))
	} else {
		fmt.Printf("%s %s %s\n", strings.Repeat(character, preWidth), strings.ToUpper(title[0]), strings.Repeat(character, width - len(title[0]) - preWidth - 2))
	}
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
`)

}
