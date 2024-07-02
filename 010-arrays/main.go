package main

import "fmt"

func main() {

	// multiline defintion
	var colors [3]string
	colors[0] = "red"
	colors[1] = "blue"
	colors[2] = "green"
	fmt.Printf("The second color is %s.\n", colors[1])

	// single line defintion
	scores := [4]int{89, 75, 99, 66}
	fmt.Println("scores", scores)
	fmt.Printf("The last score is %d.\n", scores[len(scores)-1])

}
