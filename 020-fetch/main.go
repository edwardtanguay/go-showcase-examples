package main

import (
	"encoding/json"
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

	json := string(bytes)
	fmt.Print(json)

	skills := getSkillsFromJson(json)

	for i, skill := range skills {
		fmt.Printf("Skill %d:\n", i+1)
		fmt.Printf("  ID Code: %s\n", skill.IdCode)
		fmt.Printf("  Name: %s\n", skill.Name)
		fmt.Printf("  URL: %s\n", skill.Url)
		fmt.Printf("  Description: %s\n\n", skill.Description)
	}
}

func getSkillsFromJson(content string) []Skill {
	var skills []Skill
	err := json.Unmarshal([]byte(content), &skills)
	checkError(err)
	return skills
}

type Skill struct {
	IdCode      string `json:"idCode"`
	Name        string `json:"name"`
	Url         string `json:"url"`
	Description string `json:"description"`
}
