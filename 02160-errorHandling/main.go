
package main

import "fmt"

type Employee struct {
	name string
	department string
	conferences []string
}

func (emp *Employee) addSalesConference(name string) {
	emp.conferences = append(emp.conferences, name)
}

func main() {
	emp1 := Employee{"James Ashton", "sales", []string{"Sales World", "Sales Marketplace"}}
	emp2 := Employee{"Sven Techton", "it", []string{"JavaScript World", "IT trends"}}

	emp1.addSalesConference("Sales Experts Conference")

	fmt.Printf("%#v\n", emp1)
	fmt.Printf("%#v\n", emp2)
}
