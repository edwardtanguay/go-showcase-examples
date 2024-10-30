package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

const outputDirName = "output"

func separator() {
	fmt.Println("---")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func createTextFile(fileName, content string) {
	err := os.MkdirAll(outputDirName, 0755)
	pathAndFileName := filepath.Join(outputDirName, fileName)
	file, err := os.Create(pathAndFileName)
	checkError(err)
	_, err = io.WriteString(file, content)
	checkError(err)
	defer file.Close()
}

func getHtmlBoilerplateBegin(title string) string {
	return `
<!DOCTYPE html>
<html lang="en">
<head>
	<meta charset="UTF-8">
	<meta name="viewport" content="width=device-width, initial-scale=1.0">
	<title>` + title + `</title>
	<style>
		* {
			margin: 0;
			padding: 0;
			box-sizing: border-box;
		}
		body {
			font-family: monospace;
  			line-height: 1.5;
			padding: 1rem;
		}
		h2 {
			font-size: 1.2rem;
			margin-top: .3rem;
		}
		p {
			font-size: 1rem;
			font-style: italic;
		}
		a {
			color: #333;
		}
	</style>
</head>
<body>
`
}

func getHtmlBoilerplateEnd() string {
	return `
</body>
</html>`
}
