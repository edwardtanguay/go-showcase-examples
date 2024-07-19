
package main

import "fmt"

type Framework struct {
	name string
	yearCreated int
}

func (f Framework) display() string {
	return fmt.Sprintf("%s was created in %d.", f.name, f.yearCreated)
}

func main() {
	f1 := Framework{"React", 2013}
	f2 := Framework{"Vue", 2014}
	println(f1.display())	
	println(f2.display())	
}
