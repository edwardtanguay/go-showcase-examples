
package main

import "fmt"
import "os"

func CreateFile(fileName string, content string) error {
	file, err := os.Create(fileName)
	if(err != nil) {
		return fmt.Errorf("Cannot create file: %s", fileName)
	}
	defer file.Close()
	file.WriteString(content)
	return nil
}

func main() {
	CreateFile("test.txt", "this is line 1\nand this line 2")
}
