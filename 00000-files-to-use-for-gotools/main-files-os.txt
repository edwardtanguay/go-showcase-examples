package main

import (
	"fmt"
	"os"
)

func fileExists(pathAndFileName string) bool {
	if _, err := os.Stat(pathAndFileName); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func getSizeAndModTime(pathAndFileName string) (int64, string) {
	fileInfo, err := os.Stat(pathAndFileName)
	if err != nil {
		panic(err)
	}
	return fileInfo.Size(), fileInfo.ModTime().Format("2006-01-02 15:04:05")
}

func main() {

	if(fileExists("mainx.go")) {
		fmt.Println("yes")
	} else {
		fmt.Println("no")
	}

	size, when := getSizeAndModTime("main.go")
	fmt.Printf("Size: %.2d KB, Modification Time: %s\n", size, when)


}

