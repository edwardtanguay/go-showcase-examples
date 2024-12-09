package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

// build command: go build -o crex.exe main.go

func writeFile(dirName string, fileName string, content string) error {

	filePath := filepath.Join(dirName, fileName)

	file, err := os.Create(filePath)

	if err != nil {
		return fmt.Errorf("error creating file: %v", err)
	}
	defer file.Close()

	_, err = file.WriteString(content)
	if err != nil {
		return fmt.Errorf("error writing to file: %v", err)
	}

	return nil
}

func createGoModFile(directory string) error {
	err := os.Chdir(directory)
	if err != nil {
		return fmt.Errorf("could not change to directory \"%s\", error: %v", directory, err)
	}

	cmd := exec.Command("go", "mod", "init", directory)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()

	if err != nil {
		return fmt.Errorf("could not initialize go.mod file, error: %v", err)
	}

	return nil
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("This utility will create a main.go and go.mod, run it with: go run .")
		fmt.Println("Usage: crex <directory-name>")
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

	err = createGoModFile(dirName)
	if err != nil {
		fmt.Printf("ERROR: %v", err)
	}

}
