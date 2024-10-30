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

	defer resp.Body.Close()

	bytes, err := io.ReadAll(resp.Body)
	checkError(err)

	json := string(bytes)

	skills := convertJsonToSkills(json)
	htmlContent := convertSkillsToHtml(skills)
	createTextFile("skills.html", htmlContent)
}

func convertSkillsToHtml(skills []Skill) string {
	html := ""
	title := "Skills"
	html += getHtmlBoilerplateBegin(title)
	html += fmt.Sprintf("<h1>%s</h1>", title)
	for _, skill := range skills {
		html += fmt.Sprintf("<h2><a target=\"_blank\" href=\"%s\">%s</a></h2>\n", skill.Url, skill.Name)
		html += fmt.Sprintf("<p>%s</p>\n", skill.Description)
		html += fmt.Sprintf("\n")
	}
	html += getHtmlBoilerplateEnd()
	return html
}

func convertJsonToSkills(content string) []Skill {
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
