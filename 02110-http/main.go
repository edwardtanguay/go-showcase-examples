package main

import (
	"io"
	"net/http"
)

type HtmlColor struct {
	hex string
	name string
	dark string
}

func getContentFromUrl(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	statusCode := resp.StatusCode
	if statusCode != 200 {
		return ""
	} else {
		byteContent, _ := io.ReadAll(resp.Body)
		return string(byteContent)
	}

}

func main() {
	content := getContentFromUrl("https://edwardtanguay.vercel.app/share/htmlcolors.json")
	println(content)
}
