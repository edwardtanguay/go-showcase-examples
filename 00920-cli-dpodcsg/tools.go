package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
)

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

func GetGraphicCodeFromOutlineLine(line string) string {
	line = strings.TrimSpace(line)
	pattern := regexp.MustCompile(`##[a-zA-Z0-9]+$`)
	matches := pattern.FindStringSubmatch(line)
	if len(matches) > 0 {
		markerAndIdCode := matches[0]
		return strings.TrimPrefix(markerAndIdCode, "##")
	}
	return ""
}

func CopyFile(srcPathAndFileName, dstPathAndFileName string) error {
	sourceFile, err := os.Open(srcPathAndFileName)
	if err != nil {
		return fmt.Errorf("could not open source file: %v", err)
	}
	defer sourceFile.Close()

	destFile, err := os.Create(dstPathAndFileName)
	if err != nil {
		return fmt.Errorf("could not create or truncate destination file: %v", err)
	}
	defer destFile.Close()

	_, err = io.Copy(destFile, sourceFile)
	if err != nil {
		return fmt.Errorf("could not copy content: %v", err)
	}

	err = destFile.Sync()
	if err != nil {
		return fmt.Errorf("could not flush content to disk: %v", err)
	}

	return nil
}
