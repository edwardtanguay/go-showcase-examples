package main

import (
	"fmt"
	"sort"
)

func separator() {
	fmt.Println("---")
}

func displayColors(colors []string) {
	separator()
	for index, color := range colors {
		fmt.Printf("%v. %v\n", index+1, color)
	}
}

func main() {

	colors := []string{"yellow", "orange"}

	fmt.Printf("There are %v colors:\n", len(colors))

	displayColors(colors)

	// append to end
	colors = append(colors, "black")

	displayColors(colors)

	// prepend to beginning
	colors = append([]string{"red"}, colors...)

	displayColors(colors)

	sort.Strings(colors)

	displayColors(colors)

	// remove last item
	colors = append(colors[:len(colors)-1])

	displayColors(colors)

	// remove first item
	colors = append(colors[1:len(colors)])

	displayColors(colors)

}
