package main

import "fmt"

type Employee struct {
	name        string
	department  string
	conferences []string
}

func (emp *Employee) addSalesConference(name string) error {
	if emp.department == "sales" {
		emp.conferences = append(emp.conferences, name)
	} else {
		return fmt.Errorf("%s is not in Sales", emp.name)
	}

	return nil
}

func main() {
	emp1 := Employee{"James Ashton", "sales", []string{"Sales World", "Sales Marketplace"}}
	emp2 := Employee{"Sven Techton", "it", []string{"JavaScript World", "IT trends"}}

	err := emp1.addSalesConference("Sales Experts Conference")
	if(err != nil) {
		fmt.Printf("ERROR: %v", err)
	}

	err2 := emp2.addSalesConference("Sales Experts Conference")
	if(err2 != nil) {
		fmt.Printf("ERROR: %v\n", err2)
	}

	fmt.Printf("%#v\n", emp1)
	fmt.Printf("%#v\n", emp2)
}
