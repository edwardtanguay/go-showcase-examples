package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type HtmlColor struct {
	Hex      string `json:"hex"`
	HtmlName string `json:"name"`
	Color    string `json:"textColor"`
}

func main() {
	content, _ := os.ReadFile("htmlColors.json")
	fmt.Printf("Content of file: %s\n", content)

	var htmlColors []HtmlColor
	json.Unmarshal(content, &htmlColors)

	fmt.Printf("%#v\n", htmlColors)
	for i, htmlColor := range htmlColors {
		fmt.Printf("%d. %s\n", i+1, htmlColor.HtmlName)
	}
}
