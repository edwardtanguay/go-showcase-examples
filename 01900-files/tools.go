package main

import "fmt"

func separator() {
	fmt.Println("---")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}