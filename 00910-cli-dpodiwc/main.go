package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"
)

// build command: go build -o dpodiwc.exe main.go

const validFileEnding = ".dpod.txt"
const basePath = "C:\\edward\\datapod2023\\datapod-for-vite-react-core\\currentImport"

func GetLinesFromFile(pathAndFileName string) ([]string, error) {
	file, err := os.Open(pathAndFileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}

func WriteLinesToFile(pathAndFileName string, lines []string) error {
	file, err := os.OpenFile(pathAndFileName, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if(err != nil) {
		return fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for _, line := range lines {
		_, err := writer.WriteString(line + "\n")
		if err != nil {
			return fmt.Errorf("failed to write line to file: %w", err)
		}
	}
	if err := writer.Flush(); err != nil {
		return fmt.Errorf("failed to flush writer: %w", err)
	}

	return nil
}

func addSecondsToDateTime(dateTime string, seconds int) (string, error) {
	const layout = "2006-01-02 15:04:05"

	t, err := time.Parse(layout, dateTime)
	if err != nil {
		return "", fmt.Errorf("failed to parse dateTime: %w", err)
	}

	t = t.Add(time.Duration(seconds) * time.Second)

	return t.Format(layout), nil
}


func replaceDateTime(line string, newDateTime string) string {
	pattern := `\d{4}-\d{2}-\d{2} \d{2}:\d{2}:\d{2}`
	re := regexp.MustCompile(pattern)
	return re.ReplaceAllString(line, newDateTime)
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Datapod: Increment WhenCreated")
		fmt.Println("Runs through a dpod.txt file and resets all whenCreated fields to current date/time, each item incremented by one second")
		fmt.Println("Usage: dpodiwc <fileName>")
		fmt.Println("Example: dpodiwc skills.dpod.txt")
		fmt.Println("Assumes file is in: C:\\edward\\datapod2023\\datapod-for-vite-react-core\\currentImport")
		return
	} else {

		fileName := os.Args[1]

		if !strings.HasSuffix(fileName, validFileEnding) {
			fmt.Printf("ERROR: File name must end with \"%s\"\n", validFileEnding)
		} else {
			pathAndFileName := filepath.Join(basePath, fileName)
			lines, err := GetLinesFromFile(pathAndFileName)
			if(err != nil) {
				fmt.Printf("ERROR: %v\n", err)
			} else {
				var newLines []string
				for _,line := range lines {
					currentDateTime := time.Now().Format("2006-01-02 15:04:05")
					newDateTime, err := addSecondsToDateTime(currentDateTime, 1)
					if(err != nil) {
						fmt.Printf("ERROR: %v", err)
					}
					newLines = append(newLines, replaceDateTime(line, newDateTime))
				}

				err := WriteLinesToFile(pathAndFileName, newLines)
				if(err != nil) {
					fmt.Printf("ERROR: %v\n", err)
				}
			}
		}

	}
}
