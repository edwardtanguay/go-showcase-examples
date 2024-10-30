
package main

import (
	"fmt"
	"strings"
)

func separator(title ...string) {
	width := 50
	preWidth := 3
	character := "="
	if len(title) == 0 || title[0] == "" {
		fmt.Println(strings.Repeat(character, width))
	} else {
		fmt.Printf("%s %s %s\n", strings.Repeat(character, preWidth), strings.ToUpper(title[0]), strings.Repeat(character, width - len(title[0]) - preWidth - 2))
	}
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
