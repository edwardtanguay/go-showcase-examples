package main

import "fmt"

func main() {
	emp := Employee{"Hans", "Hansrodt", 34}
	separator()
	fmt.Println(emp)
	fmt.Printf("Employee: %v\n", emp)
	fmt.Printf("Employee: %+v\n", emp)
	separator()
	fmt.Printf("The employee %v is %v years old.\n", emp.FirstName + " " + emp.LastName, emp.Age)
	emp.Age = emp.Age + 10
	separator()
	fmt.Printf("The employee %v is %v years old.\n", emp.FirstName + " " + emp.LastName, emp.Age)
}

type Employee struct {
	FirstName string
	LastName string
	Age int
}
