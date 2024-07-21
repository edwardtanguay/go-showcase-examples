package main

import "fmt"
import "github.com/fatih/color"

func PrintAnything(v interface{}) {
	red := color.New(color.FgRed).SprintFunc()
	green := color.New(color.FgGreen).SprintFunc()
	fmt.Printf("The variable of type %s has a value of: %v\n", green(fmt.Sprintf("%T", v)), red(v))
}

func main() {
	PrintAnything("hello")
	PrintAnything(123)
	PrintAnything(1.34)
	PrintAnything(true)
}