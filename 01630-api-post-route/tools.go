package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
)

func getSuuid() string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	byteSuuid := make([]byte, 6)
	for i := range byteSuuid {
		byteSuuid[i] = charset[rand.Intn(len(charset))]
	}
	fmt.Printf("%v", byteSuuid)
	return string(byteSuuid)
}

func getFlashcardsFromJsonFile(fileName string) ([]Flashcard, error) {
	byteData, err := os.ReadFile(fileName)
	if err != nil {
		return nil, fmt.Errorf("error reading file %s: %v", fileName, err)
	}

	var flashcards []Flashcard
	err = json.Unmarshal(byteData, &flashcards)
	if err != nil {
		return nil, fmt.Errorf("invalid JSON : %v", err)
	}

	return flashcards, nil
}

func writeFlashcardsToJsonFile(fileName string, flashcards []Flashcard) error {
	byteData, err := json.MarshalIndent(flashcards, "", "\t")
	if err != nil {
		return fmt.Errorf("error processing flashcards: %v", err)
	}

	err = os.WriteFile(fileName, byteData, 0644)
	if err != nil {
		return fmt.Errorf("error writing file: %v", err)
	}

	return nil
}
