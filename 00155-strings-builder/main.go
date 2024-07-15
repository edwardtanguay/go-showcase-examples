package main

import (
	"strings"
)

func main() {
	var sb strings.Builder

	println("capacity", sb.Cap())

	sb.WriteString("Hello")
	sb.WriteString(" ")
	sb.WriteString("World")
	println("Size of string will be ", sb.Len())
	println("string: ", sb.String())

	// expand space for string content
	println("capacity", sb.Cap())
	sb.Grow(1024) // can improve performance
	println("capacity", sb.Cap())

	sb.Reset()
	println("capacity", sb.Cap())
	println("string: ", sb.String())
}
