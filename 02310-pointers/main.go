
package main

import "fmt"

type Person struct {
	firstName string
	lastName string
	fullName string
	age int
}

func Display(p Person) {
	p.fullName = p.firstName + " " + p.lastName
	fmt.Printf("person.fullName = %s\n", p.fullName)
	// person is internally but not externally changed
}

func Update(p *Person) {
	p.fullName = p.firstName + " " + p.lastName
	fmt.Printf("person.fullName = %s\n", p.fullName)
	// person IS internally AS WELL AS externally changed
}

func main() {
	person := Person{"Henri", "Lefèvre", "", 44}
	fmt.Printf("%#v\n", person)
	Display(person)
	fmt.Printf("%#v\n", person)
	Update(&person)
	fmt.Printf("%#v\n", person)
}
