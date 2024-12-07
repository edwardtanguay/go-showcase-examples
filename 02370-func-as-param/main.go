package main

import (
	"02370-func-as-param/tools"
	"fmt"
)

var operations = map[string]func(int, int) int{
	"add":  tools.Add,
	"sub":  tools.Subtract,
	"mult": tools.Multiply,
}

func calculator(a int, b int, name string) (string, error) {
	if fn, exists := operations[name]; exists {
		switch name {
		case "add":
			return fmt.Sprintf("%d + %d = %d", a, b, fn(a, b)), nil
		case "sub":
			return fmt.Sprintf("%d - %d = %d", a, b, fn(a, b)), nil
		case "mult":
			return fmt.Sprintf("%d * %d = %d", a, b, fn(a, b)), nil
		}
	}
	return "", fmt.Errorf("function \"%s\" not defined", name)
}

func main() {
	first := 9
	second := 4
	ops := []string{"sub", "mult", "whatever", "add"}
	for _, op := range ops {
		answerLine, err := calculator(first, second, op)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
		} else {
			println(answerLine)
		}
	}
}
