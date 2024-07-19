
package main

import "fmt"

type Framework struct {
	name string
	yearCreated int
}

func (f Framework) display() string {
	return fmt.Sprintf("nnn")
}

func main() {
	react := Framework{"React", 2013}
	println(react.display())	
}
