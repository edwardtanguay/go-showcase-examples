
package main

import "fmt"
import "github.com/edwardtanguay/gotools"

func main() {
	lines, _ := files.GetLinesFromFile("main.go")
	for index,line := range lines {
		fmt.Printf("%02d | %s\n", index+1, line)
	}
}
