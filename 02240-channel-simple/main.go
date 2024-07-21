
package main

import "fmt"
import "sync"

func main() {
	score := make(chan int)
	var wg sync.WaitGroup

	fmt.Println("test")

	go func() {
		_score := <-score
		fmt.Printf("Score received: %d\n", _score)
	}()

	score <- 98
	wg.Wait()
}
