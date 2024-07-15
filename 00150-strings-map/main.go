package main

import "strings"

func main() {

	sentence := "You want to inspect strings that contain special characters or non-visible characters"
	shift := 10

	transform := func(r rune) rune {
		switch {
		case r >= 'A' && r <= 'Z':
			value := int('A') + (int(r) - int('A') + shift)%26
			if value > 91 {
				value -= 26
			} else if value < 65 {
				value += 26
			}
			return rune(value)
		case r >= 'a' && r <= 'z':
			value := int('a') + (int(r) - int('a') + shift)%26
			if value > 122 {
				value -= 26
			} else if value < 97 {
				value += 26
			}
			return rune(value)
		}
		return r
	}
	encodedString := strings.Map(transform, sentence)
	println(sentence)
	println(encodedString)
}
