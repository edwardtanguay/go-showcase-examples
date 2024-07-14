package main

import "fmt"

func main() {

	languages := []string{"C#", "Python", "Ruby", "Java"}

	// for classic
	separator()
	for i := 0; i < len(languages); i++ {
		fmt.Printf("%v. %v\n", i+1, languages[i])
	}

	// for index/range
	separator()
	for i := range languages {
		fmt.Printf("%v. %v\n", i+1, languages[i])
	}

	// for of
	separator()
	for i, language := range languages {
		fmt.Printf("%v. %v\n", i+1, language)
	}

	// for exterior index
	separator()
	index := 0
	for index < len(languages) {
		fmt.Printf("%v. %v\n", index+1, languages[index])
		index++
	}
}
