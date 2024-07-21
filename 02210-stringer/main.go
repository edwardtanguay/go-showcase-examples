package main

import "fmt"

type Report struct {
	Title string
	Pages int
}

func (r Report) String() string {
	return fmt.Sprintf("The report named \"%s\" has %d pages.", r.Title, r.Pages)
}

func main() {

	r1 := Report{"Quarterly Sales Report", 23}
	fmt.Printf("Report: %v", r1)
}
