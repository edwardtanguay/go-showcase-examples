package main

import "testing"

func TestHumanAge(t *testing.T) {
	dogHumanAge := getHumanAge(4)
	if dogHumanAge != 28 {
		t.Fatal("No match, got", dogHumanAge)
	}

}