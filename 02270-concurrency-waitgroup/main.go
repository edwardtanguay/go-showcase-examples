package main

import (
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"
)

func fetchUrl(url string, wg *sync.WaitGroup, results chan<- string) {
	defer wg.Done()

	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		results <- fmt.Sprintf("Error fetching %s: %v\n", url, err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		results <- fmt.Sprintf("Error reading response from %s: %v\n", url, err)
	}

	secs := time.Since(start).Seconds()
	results <- fmt.Sprintf("Fetched %s in %.2f seconds with response length %d bytes\n", url, secs, len(body))
}

func main() {
	urls := []string{
		"https://google.com",
		"https://amazon.com",
		"https://stackoverflow.com",
		"https://www.heise.de",
	}

	var wg sync.WaitGroup
	results := make(chan string, len(urls))

	for _, url := range urls {
		wg.Add(1)
		go fetchUrl(url, &wg, results)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	for result := range results {
		fmt.Print(result)
	}
}
