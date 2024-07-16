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
		fmt.Println("Usage: ce <directory-name>")
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

}
