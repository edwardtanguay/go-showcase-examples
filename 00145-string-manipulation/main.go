package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func main() {

	// split
	scoreList := "92;77;sdf;82;95"
	scores := strings.Split(scoreList, ";")
	cleanScores := []int{}
	for _, strScore := range scores {
		score, err := strconv.Atoi(strScore)
		if err != nil {
			fmt.Printf("Error converting %s: %v\n", strScore, err)
		} else {
			fmt.Printf("score = %d, with curve = %d\n", score, score+2)
			cleanScores = append(cleanScores, score)
		}
	}
	fmt.Printf("scores = %v is of type %T\n", scores, scores)
	fmt.Printf("cleanScores = %v is of type %T\n", cleanScores, cleanScores)

	// join
	names := []string{"Harald", "Michael", "Robert"}
	nameList := fmt.Sprintf(">>> %s <<<", strings.Join(names, "|"))
	println(nameList)

	// fields
	sentence := "He takes a look at several useful packages."
	wordItems := strings.Fields(sentence)
	fmt.Printf("words is %q and of type %T\n", wordItems, wordItems)

	sentence2 := "The conference, scheduled for next week, will cover various topics: artificial intelligence, machine learning, data analysis; cybersecurity; ethical considerations - are you attending?"
	findPunct := func(c rune) bool {
		return unicode.IsPunct(c)
	}
	sentenceParts := strings.FieldsFunc(sentence2, findPunct)
	fmt.Printf("parts: %q\n", sentenceParts)

	// replacing
	rep := strings.NewReplacer("a", "_","e", "_","i", "_","o", "_","u", "_")
	blankedSentence := rep.Replace(sentence2)
	println(sentence2)
	println(blankedSentence)
}
