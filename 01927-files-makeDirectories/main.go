
package main

import "os"

func createNestedDirectories(directoryPath string) {
	err := os.MkdirAll(directoryPath, os.ModePerm);
	 if err != nil {
		panic(err)
	 }
}

func deleteNestedDirectories(directoryPath string) {
	 err := os.RemoveAll(directoryPath)
	 if err != nil {
		panic(err)
	 }
}

func main() {
	createNestedDirectories("output/images")	
	createNestedDirectories("output/data")	
	// deleteNestedDirectories("output")
}
