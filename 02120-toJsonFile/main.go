package main

import (
	"fmt"
	"os"
)

type Person struct {
	firstName string
	lastName  string
	age       int
}

func writeFile(fileName string, content string) {
	file, err := os.Create(fileName)
	if err != nil {
		println("file could not be created")
	}
	_, err2 := file.WriteString(content)
	if err2 != nil {
		println("could not write to file")
	}
}

func main() {
	fileName := "persons.json"

	persons := []Person{{firstName: "Hans", lastName: "Hackert", age: 43}, {firstName: "Petra", lastName: "Reinecke", age: 52}}

	fmt.Println("Saving to file as JSON:")
	for i, person := range persons {
		fmt.Printf("%d. %s\n", i+1, person.firstName+" "+person.lastName)
	}

	writeFile(fileName, "nnn\nooo")

}
