package main

import "net/http"

func getContentFromUrl(url string) {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	println(resp.StatusCode)
}

func main() {
	getContentFromUrl("https://edwardtanguay.vercel.app/share/htmlcolors.json")
}
