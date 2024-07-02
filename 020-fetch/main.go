package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "https://edwardtanguay.vercel.app/share/skills.json"

func main() {
	resp, err := http.Get(url)
	checkError(err)

	fmt.Printf("Response type: %T\n", resp)
	defer resp.Body.Close()

	bytes, err := io.ReadAll(resp.Body)
	checkError(err)

	content := string(bytes)
	fmt.Print(content)
}
