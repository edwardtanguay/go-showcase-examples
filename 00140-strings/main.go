
package main

import "fmt"

func main() {
	
	// length
	title := "LAMP"
	fmt.Printf("Length is %d\n", len(title))

	// iterate over string
	for i,ch := range title {
		fmt.Printf("%d: %c\n", i, ch)
	}
}
