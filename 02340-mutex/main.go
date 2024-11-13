package main

import (
	"fmt"
	"sync"
)

func main() {
	var counter int
	var mutex sync.Mutex
	var wg sync.WaitGroup

	increment := func() {
		defer wg.Done()
		mutex.Lock()
		counter++
		mutex.Unlock()
	}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go increment()
	}

	wg.Wait()
	fmt.Println("Final counter value:", counter)
}
