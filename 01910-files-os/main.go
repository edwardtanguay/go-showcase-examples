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

func main() {
	separator()

	if(fileExists("mainx.go")) {
		fmt.Println("yes")
	} else {
		fmt.Println("no")
	}
}
