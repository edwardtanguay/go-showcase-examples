package main

import (
	"fmt"
)

type Person struct {
	FirstName string
	LastName  string
}

func main() {
	person := Person{"Joseph", "Maltable"}
	var p *Person = &person
	fmt.Printf("%#v\n", person)
	p.FirstName = "Ralph"
	(*p).LastName = "Newton"
	fmt.Printf("%#v\n", person)
}
