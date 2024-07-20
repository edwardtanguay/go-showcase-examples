
package main

import "fmt"

type StringPredicate = func(str string) bool

func StringFilter (stringItems []string, predicate StringPredicate) []string {
	var results []string

	for _,stringItem := range stringItems {
		if(predicate(stringItem)) {
			results = append(results, stringItem)
		}
	}

	return results
}

func main() {
	names := []string{"Jean", "Marie", "Pierre", "Claire", "Luc"}

	shortNames := StringFilter(names, func(name string) bool {
		return len(name) <= 4
	})

	longNames := StringFilter(names, func(name string) bool {
		return len(name) > 4
	})

	fmt.Printf("names = %#v\n", names)
	fmt.Printf("shortNames = %#v\n", shortNames)
	fmt.Printf("longNames = %#v\n", longNames)
}
