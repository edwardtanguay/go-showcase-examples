package main

import (
	"02370-func-as-param/tools"
	"fmt"
)

var operations = map[string]func(int, int) int{
	"add":  tools.Add,
	"sub":  tools.Substract,
	"mult": tools.Multiply,
}

func calculator(a int, b int, name string) (string, error) {
	if fn, exists := operations[name]; exists {
		switch name {
		case "add":
			return fmt.Sprintf("%d + %d = %d", a, b, fn(a, b)), nil
		}
	}
	return "", fmt.Errorf("function not defined")
}

func main() {
	first := 9
	second := 4
	answerLine, err := calculator(first, second, "add")
	if err != nil {
		fmt.Printf("Error: %v", err)
	} else {
		println(answerLine)
	}
}
