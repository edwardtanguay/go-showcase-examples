
package main

import "fmt"

type Employee struct {
	name string
	department string
	conferences []string
}

func main() {
	emp1 := Employee{"James Ashton", "sales", []string{"Sales World", "Sales Marketplace"}}
	fmt.Printf("%#v", emp1)
}
