package main

import (
	"fmt"
	"math"
	"math/rand"
	"strings"
)

func main() {

	separator()

	fmt.Println(math.Pow(2, 8))

	separator("random")

	fmt.Println(rand.Int())
	fmt.Println(rand.Intn(3)) // 0-2
	
	fmt.Println(rand.Float32())
	fmt.Println(rand.Float64())

	separator("randomizing slices")

	colors := []string{"red", "green", "blue", "yellow", "purple"}
	randomIndexes := rand.Perm(len(colors))
	fmt.Println(colors)
	fmt.Println(randomIndexes)
	for _,randomIndex := range randomIndexes {
		fmt.Printf("random index %d is %s\n", randomIndex, colors[randomIndex])
	}

	separator("shuffle")

	words := strings.Fields("one two three four five")
	rand.Shuffle(len(words), func(i, j int) {
		words[i], words[j] = words[j], words[i]
	})
	fmt.Println(words)
}
