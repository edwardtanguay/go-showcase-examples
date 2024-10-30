package main

import "os"

func writeFile(fileName, content string) {
	file, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}
	_, err = file.WriteString(content)
	if err != nil {
		panic(err)
	}

}

func main() {
	writeFile("test.txt", "first line\nsecond line")
}




