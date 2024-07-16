
package main

import "fmt"
import "os"
import "encoding/json"

type HtmlColor struct {
	Hex string 			`json:"hex"`
	HtmlName string 	`json:"name"`
	Color string 		`json:"textColor"`
}

func main() {
	content, _ := os.ReadFile("htmlColors.json")
	fmt.Printf("Content of file: %s\n", content)

	var htmlColors []HtmlColor
	json.Unmarshal(content, &htmlColors)

	fmt.Printf("%#v\n", htmlColors)
}
