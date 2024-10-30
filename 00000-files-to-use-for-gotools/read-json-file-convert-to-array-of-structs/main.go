package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Person struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Age       int    `json:"age"`
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

	persons := []Person{{FirstName: "Hans", LastName: "Hackert", Age: 43}, {FirstName: "Petra", LastName: "Reinecke", Age: 52}}

	fmt.Println("Saving to file as JSON:")
	for i, person := range persons {
		fmt.Printf("%d. %s\n", i+1, person.FirstName+" "+person.LastName)
	}

	json, err := json.MarshalIndent(persons, "", "	")
	if err != nil {
		println("could not convert struct to JSON text")
	}
	writeFile(fileName, string(json))

}
