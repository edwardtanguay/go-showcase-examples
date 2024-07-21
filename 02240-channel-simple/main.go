
package main

import "fmt"
import "sync"

func main() {
	score := make(chan int)

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		_score := <-score
		fmt.Printf("Score received: %d\n", _score)
	}()

	score <- 98
	wg.Wait()
}
