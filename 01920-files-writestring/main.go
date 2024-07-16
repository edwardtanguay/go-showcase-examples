
package main

import "os"

func writeFile(fileName, content string) {
	file, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}
	println(999, file.Name())
}

func main() {
	writeFile("test.txt", "first line")
}
