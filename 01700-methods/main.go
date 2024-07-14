package main

import (
	"fmt"
	"strings"
)

func main() {
	emp := Employee{"Richard", "Wagstedt", 34}
	Separator()
	fmt.Printf("The employee %v is %v years old.\n", emp.FirstName+" "+emp.LastName, emp.Age)
	Separator()
	fmt.Printf("The employee %v is %v years old.\n", emp.FullName(), emp.Age)
	Separator()
	fmt.Printf("The employee %v is %v years old.\n", emp.StyledName(), emp.Age)
}

type Employee struct {
	FirstName string
	LastName  string
	Age       int
}

func (e Employee) FullName() string {
	return e.FirstName + " " + e.LastName
}

func (e Employee) StyledName() string {
	return strings.ToUpper(e.FullName())
}
