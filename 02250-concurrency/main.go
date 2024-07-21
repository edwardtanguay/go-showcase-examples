package main

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

func fetchUrl(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprintf("Error fetching %s: %v\n", url, err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		ch <- fmt.Sprintf("Error reading response from %s: %v\n", url, err)
	}

	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("Fetched %s in %.2f seconds with response length %d bytes\n", url, secs, len(body))
}

func main() {
	urls := []string{
		"https://google.com",
		"https://amazon.com",
		"https://stackoverflow.com",
		"https://www.heise.de",
	}

	ch := make(chan string)

	for _, url := range urls {
		go fetchUrl(url, ch)
	}

	for range urls {
		fmt.Printf(<-ch)
	}
}
