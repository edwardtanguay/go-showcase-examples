package main

import "fmt"

type Report struct {
	Title string
	Pages int
}

func main() {

	r1 := Report{"Quarterly Sales Report", 23}
	fmt.Printf("Report: %v", r1)
}
