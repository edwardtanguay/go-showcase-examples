package main

import (
	"fmt"
	"math/rand"
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