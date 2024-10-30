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
	emp1 := Employee{"James Adverty", "sales", []string{"Sales World", "Sales Marketplace"}}
	emp2 := Employee{"Sven Techton", "it", []string{"JavaScript World", "IT trends"}}
	emp3 := Employee{"Angel Marketon", "sales", []string{"Sales World", "Sales Marketplace"}}

	employees := []*Employee{&emp1, &emp2, &emp3}

	for _, employee := range employees {
		fmt.Printf("%#v\n", employee)
	}

	for _, employee := range employees {
		err := employee.addSalesConference("Sales Experts Conference")
		if err != nil {
			fmt.Printf("ERROR: %v\n", err)
		}
	}

	for _, employee := range employees {
		fmt.Printf("%#v\n", employee)
	}

}
