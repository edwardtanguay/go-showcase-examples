package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func getJsonFiles(dirPath string) ([]string, error) {
	var fileList []string

	err := filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && strings.HasSuffix(info.Name(), ".json") {
			fileList = append(fileList, path)
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return fileList, nil
}

func main() {
	dirPath := "." 
	files, err := getJsonFiles(dirPath)
	if err != nil {
		panic(err)
	}

	fmt.Println("JSON Files:")
	for _, file := range files {
		fmt.Println(file)
	}
}
