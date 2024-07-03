package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
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

func getSkillsFromJson(content string) []Skill {
	skills := make([]Skill, 0, 20)
	decoder := json.NewDecoder(strings.NewReader(content))
	_, err := decoder.Token()
	checkError(err)
	var skill Skill
	for decoder.More() {

	}
}

type Skill struct {
	idCode      string
	name        string
	url         string
	description string
}
